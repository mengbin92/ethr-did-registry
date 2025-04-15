```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract EthereumDIDRegistryV2 is Initializable, UUPSUpgradeable, AccessControlUpgradeable {
    using ECDSA for bytes32;

    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant MULTISIG_ROLE = keccak256("MULTISIG_ROLE");

    mapping(address => address) private _owners;
    mapping(address => mapping(bytes32 => mapping(address => uint))) private _delegates;
    mapping(address => uint) public changed;
    mapping(address => uint) public nonce;

    // 多签 pending 操作
    enum ActionType { ChangeOwner, AddDelegate, RevokeDelegate, SetAttribute, RevokeAttribute }
    struct Pending {
        ActionType action;
        address initiator;
        bytes   data;         // abi.encode(...) of parameters
        uint    approvals;    // 已确认数
        bool    executed;
    }
    uint public pendingCount;
    mapping(uint => Pending) public pendings;
    mapping(uint => mapping(address => bool)) public hasApproved;

    event DIDOwnerChanged(address indexed identity, address owner, uint previousChange);
    event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint validTo, uint previousChange);
    event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint validTo, uint previousChange);
    event PendingCreated(uint indexed id, ActionType action, address indexed initiator);
    event PendingApproved(uint indexed id, address indexed approver, uint approvals);
    event PendingExecuted(uint indexed id);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() initializer {}

    function initialize(address admin, address upgrader, address[] calldata multisigs) public initializer {
        __AccessControl_init();
        __UUPSUpgradeable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(UPGRADER_ROLE, upgrader);

        for (uint i = 0; i < multisigs.length; i++) {
            _grantRole(MULTISIG_ROLE, multisigs[i]);
        }
    }

    // UUPS 升级授权
    function _authorizeUpgrade(address newImpl) internal override onlyRole(UPGRADER_ROLE) {}

    modifier onlyOwner(address identity) {
        require(_msgSender() == identityOwner(identity), "bad_actor");
        _;
    }

    function identityOwner(address identity) public view returns (address) {
        address o = _owners[identity];
        return o == address(0) ? identity : o;
    }

    function validDelegate(address identity, bytes32 delegateType, address delegate) public view returns (bool) {
        return _delegates[identity][keccak256(abi.encode(delegateType))][delegate] > block.timestamp;
    }

    // —— 多签流程 —— //

    function _newPending(ActionType action, bytes memory data) internal returns (uint) {
        uint id = ++pendingCount;
        pendings[id] = Pending({
            action:    action,
            initiator: _msgSender(),
            data:      data,
            approvals: 0,
            executed:  false
        });
        emit PendingCreated(id, action, _msgSender());
        return id;
    }

    function approvePending(uint id) external onlyRole(MULTISIG_ROLE) {
        Pending storage p = pendings[id];
        require(!p.executed, "already_executed");
        require(!hasApproved[id][_msgSender()], "already_approved");
        hasApproved[id][_msgSender()] = true;
        p.approvals++;
        emit PendingApproved(id, _msgSender(), p.approvals);

        // 简单规则：至少 2 个 multisig 批准即可执行
        if (p.approvals >= 2) {
            _executePending(id);
        }
    }

    function _executePending(uint id) internal {
        Pending storage p = pendings[id];
        require(!p.executed, "already_executed");
        p.executed = true;

        // 根据 action 解码 data 并调用相应内部函数
        if (p.action == ActionType.ChangeOwner) {
            (address identity, address newOwner) = abi.decode(p.data, (address,address));
            _changeOwner(identity, newOwner);
        } else if (p.action == ActionType.AddDelegate) {
            (address identity, bytes32 dt, address delegate, uint validity) = abi.decode(p.data, (address,bytes32,address,uint));
            _addDelegate(identity, dt, delegate, validity);
        } else if (p.action == ActionType.RevokeDelegate) {
            (address identity, bytes32 dt, address delegate) = abi.decode(p.data, (address,bytes32,address));
            _revokeDelegate(identity, dt, delegate);
        } else if (p.action == ActionType.SetAttribute) {
            (address identity, bytes32 name, bytes memory value, uint validity) = abi.decode(p.data, (address,bytes32,bytes,uint));
            _setAttribute(identity, name, value, validity);
        } else if (p.action == ActionType.RevokeAttribute) {
            (address identity, bytes32 name, bytes memory value) = abi.decode(p.data, (address,bytes32,bytes));
            _revokeAttribute(identity, name, value);
        }

        emit PendingExecuted(id);
    }

    // —— 内部核心逻辑 —— //

    function _changeOwner(address identity, address newOwner) internal {
        emit DIDOwnerChanged(identity, newOwner, changed[identity]);
        _owners[identity] = newOwner;
        changed[identity] = block.number;
    }

    function _addDelegate(address identity, bytes32 dt, address delegate, uint validity) internal {
        uint validTo = block.timestamp + validity;
        _delegates[identity][keccak256(abi.encode(dt))][delegate] = validTo;
        emit DIDDelegateChanged(identity, dt, delegate, validTo, changed[identity]);
        changed[identity] = block.number;
    }

    function _revokeDelegate(address identity, bytes32 dt, address delegate) internal {
        _delegates[identity][keccak256(abi.encode(dt))][delegate] = block.timestamp;
        emit DIDDelegateChanged(identity, dt, delegate, block.timestamp, changed[identity]);
        changed[identity] = block.number;
    }

    function _setAttribute(address identity, bytes32 name, bytes memory value, uint validity) internal {
        emit DIDAttributeChanged(identity, name, value, block.timestamp + validity, changed[identity]);
        changed[identity] = block.number;
    }

    function _revokeAttribute(address identity, bytes32 name, bytes memory value) internal {
        emit DIDAttributeChanged(identity, name, value, 0, changed[identity]);
        changed[identity] = block.number;
    }

    // —— 对外接口：仅创建 Pending —— //

    function changeOwner(address identity, address newOwner) external onlyOwner(identity) {
        _newPending(ActionType.ChangeOwner, abi.encode(identity, newOwner));
    }

    function addDelegate(address identity, bytes32 dt, address delegate, uint validity) external onlyOwner(identity) {
        _newPending(ActionType.AddDelegate, abi.encode(identity, dt, delegate, validity));
    }

    function revokeDelegate(address identity, bytes32 dt, address delegate) external onlyOwner(identity) {
        _newPending(ActionType.RevokeDelegate, abi.encode(identity, dt, delegate));
    }

    function setAttribute(address identity, bytes32 name, bytes memory value, uint validity) external onlyOwner(identity) {
        _newPending(ActionType.SetAttribute, abi.encode(identity, name, value, validity));
    }

    function revokeAttribute(address identity, bytes32 name, bytes memory value) external onlyOwner(identity) {
        _newPending(ActionType.RevokeAttribute, abi.encode(identity, name, value));
    }
}
```

**solc version**: 0.8.26
**evm version**: default