// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethr_did_registry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EthrDidRegistryMetaData contains all meta data concerning the EthrDidRegistry contract.
var EthrDidRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDDelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDOwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegateSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"changed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegateSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"validDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506123ad8061001c5f395ff3fe608060405234801561000f575f80fd5b50600436106100fd575f3560e01c806380b29f7c11610095578063a7068d6611610064578063a7068d66146102b5578063e476af5c146102d1578063f00d4b5d146102ed578063f96d0f9f14610309576100fd565b806380b29f7c146102315780638733d4e81461024d578063930726841461027d5780639c2c1b2b14610299576100fd565b8063240cf1fa116100d1578063240cf1fa14610199578063622b2a3c146101b557806370ae92d2146101e55780637ad4b0a414610215576100fd565b8062c023da14610101578063022914a71461011d5780630d44625b1461014d578063123b5e981461017d575b5f80fd5b61011b60048036038101906101169190611446565b610339565b005b610137600480360381019061013291906114b2565b61034a565b60405161014491906114ec565b60405180910390f35b61016760048036038101906101629190611505565b610379565b604051610174919061156d565b60405180910390f35b610197600480360381019061019291906115e6565b6103a4565b005b6101b360048036038101906101ae919061169f565b61044b565b005b6101cf60048036038101906101ca9190611505565b6104ea565b6040516101dc9190611730565b60405180910390f35b6101ff60048036038101906101fa91906114b2565b6105a9565b60405161020c919061156d565b60405180910390f35b61022f600480360381019061022a9190611749565b6105be565b005b61024b60048036038101906102469190611505565b6105d1565b005b610267600480360381019061026291906114b2565b6105e2565b60405161027491906114ec565b60405180910390f35b610297600480360381019061029291906117c9565b610688565b005b6102b360048036038101906102ae9190611852565b61072b565b005b6102cf60048036038101906102ca91906118ef565b6107d2565b005b6102eb60048036038101906102e69190611953565b6107e5565b005b610307600480360381019061030291906119f8565b610888565b005b610323600480360381019061031e91906114b2565b610897565b604051610330919061156d565b60405180910390f35b610345833384846108ac565b505050565b5f602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052825f5260405f20602052815f5260405f20602052805f5260405f205f9250925050505481565b5f601960f81b5f60f81b3060035f6103bb8d6105e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548b888888604051602001610410989796959493929190611c0f565b604051602081830303815290604052805190602001209050610441886104398a8a8a8a876109ff565b868686610b24565b5050505050505050565b5f601960f81b5f60f81b3060035f6104628b6105e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205489866040516020016104b396959493929190611cf1565b6040516020818303038152906040528051906020012090506104e2866104dc88888888876109ff565b84610c83565b505050505050565b5f8060015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8560405160200161053b9190611d7a565b6040516020818303038152906040528051906020012081526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490504281119150509392505050565b6003602052805f5260405f205f915090505481565b6105cb8433858585610b24565b50505050565b6105dd83338484610e4b565b505050565b5f805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461067e5780915050610683565b829150505b919050565b5f601960f81b5f60f81b3060035f61069f8c6105e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548a87876040516020016106f29796959493929190611ddd565b6040516020818303038152906040528051906020012090506107228761071b89898989876109ff565b8585610e4b565b50505050505050565b5f601960f81b5f60f81b3060035f6107428d6105e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548b888888604051602001610797989796959493929190611eb2565b6040516020818303038152906040528051906020012090506107c8886107c08a8a8a8a876109ff565b868686611050565b5050505050505050565b6107df8433858585611050565b50505050565b5f601960f81b5f60f81b3060035f6107fc8c6105e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548a878760405160200161084f9796959493929190611f98565b60405160208183030381529060405280519060200120905061087f8761087889898989876109ff565b85856108ac565b50505050505050565b610893823383610c83565b5050565b6002602052805f5260405f205f915090505481565b83836108b7826105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610924576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091b90612079565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff167f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e485855f60025f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546040516109ad9493929190612118565b60405180910390a24360025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505050505050565b5f806001838787876040515f8152602001604052604051610a239493929190612171565b6020604051602081039080840390855afa158015610a43573d5f803e3d5ffd5b505050602060405103519050610a58876105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abc906121fe565b60405180910390fd5b60035f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610b1290612249565b91905055508091505095945050505050565b8484610b2f826105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9390612079565b60405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff167f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e486868642610be29190612290565b60025f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054604051610c3094939291906122c3565b60405180910390a24360025f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050505050505050565b8282610c8e826105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610cfb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf290612079565b60405180910390fd5b825f808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff167f38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a38460025f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054604051610dfa92919061230d565b60405180910390a24360025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050505050565b8383610e56826105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ec3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eba90612079565b60405180910390fd5b4260015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f86604051602001610f139190611d7a565b6040516020818303038152906040528051906020012081526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508573ffffffffffffffffffffffffffffffffffffffff167f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f785854260025f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054604051610ffe9493929190612334565b60405180910390a24360025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505050505050565b848461105b826105e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110bf90612079565b60405180910390fd5b82426110d49190612290565b60015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f876040516020016111239190611d7a565b6040516020818303038152906040528051906020012081526020019081526020015f205f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508673ffffffffffffffffffffffffffffffffffffffff167f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7868686426111cb9190612290565b60025f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546040516112199493929190612334565b60405180910390a24360025f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050505050505050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112a68261127d565b9050919050565b6112b68161129c565b81146112c0575f80fd5b50565b5f813590506112d1816112ad565b92915050565b5f819050919050565b6112e9816112d7565b81146112f3575f80fd5b50565b5f81359050611304816112e0565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61135882611312565b810181811067ffffffffffffffff8211171561137757611376611322565b5b80604052505050565b5f61138961126c565b9050611395828261134f565b919050565b5f67ffffffffffffffff8211156113b4576113b3611322565b5b6113bd82611312565b9050602081019050919050565b828183375f83830152505050565b5f6113ea6113e58461139a565b611380565b9050828152602081018484840111156114065761140561130e565b5b6114118482856113ca565b509392505050565b5f82601f83011261142d5761142c61130a565b5b813561143d8482602086016113d8565b91505092915050565b5f805f6060848603121561145d5761145c611275565b5b5f61146a868287016112c3565b935050602061147b868287016112f6565b925050604084013567ffffffffffffffff81111561149c5761149b611279565b5b6114a886828701611419565b9150509250925092565b5f602082840312156114c7576114c6611275565b5b5f6114d4848285016112c3565b91505092915050565b6114e68161129c565b82525050565b5f6020820190506114ff5f8301846114dd565b92915050565b5f805f6060848603121561151c5761151b611275565b5b5f611529868287016112c3565b935050602061153a868287016112f6565b925050604061154b868287016112c3565b9150509250925092565b5f819050919050565b61156781611555565b82525050565b5f6020820190506115805f83018461155e565b92915050565b5f60ff82169050919050565b61159b81611586565b81146115a5575f80fd5b50565b5f813590506115b681611592565b92915050565b6115c581611555565b81146115cf575f80fd5b50565b5f813590506115e0816115bc565b92915050565b5f805f805f805f60e0888a03121561160157611600611275565b5b5f61160e8a828b016112c3565b975050602061161f8a828b016115a8565b96505060406116308a828b016112f6565b95505060606116418a828b016112f6565b94505060806116528a828b016112f6565b93505060a088013567ffffffffffffffff81111561167357611672611279565b5b61167f8a828b01611419565b92505060c06116908a828b016115d2565b91505092959891949750929550565b5f805f805f60a086880312156116b8576116b7611275565b5b5f6116c5888289016112c3565b95505060206116d6888289016115a8565b94505060406116e7888289016112f6565b93505060606116f8888289016112f6565b9250506080611709888289016112c3565b9150509295509295909350565b5f8115159050919050565b61172a81611716565b82525050565b5f6020820190506117435f830184611721565b92915050565b5f805f806080858703121561176157611760611275565b5b5f61176e878288016112c3565b945050602061177f878288016112f6565b935050604085013567ffffffffffffffff8111156117a05761179f611279565b5b6117ac87828801611419565b92505060606117bd878288016115d2565b91505092959194509250565b5f805f805f8060c087890312156117e3576117e2611275565b5b5f6117f089828a016112c3565b965050602061180189828a016115a8565b955050604061181289828a016112f6565b945050606061182389828a016112f6565b935050608061183489828a016112f6565b92505060a061184589828a016112c3565b9150509295509295509295565b5f805f805f805f60e0888a03121561186d5761186c611275565b5b5f61187a8a828b016112c3565b975050602061188b8a828b016115a8565b965050604061189c8a828b016112f6565b95505060606118ad8a828b016112f6565b94505060806118be8a828b016112f6565b93505060a06118cf8a828b016112c3565b92505060c06118e08a828b016115d2565b91505092959891949750929550565b5f805f806080858703121561190757611906611275565b5b5f611914878288016112c3565b9450506020611925878288016112f6565b9350506040611936878288016112c3565b9250506060611947878288016115d2565b91505092959194509250565b5f805f805f8060c0878903121561196d5761196c611275565b5b5f61197a89828a016112c3565b965050602061198b89828a016115a8565b955050604061199c89828a016112f6565b94505060606119ad89828a016112f6565b93505060806119be89828a016112f6565b92505060a087013567ffffffffffffffff8111156119df576119de611279565b5b6119eb89828a01611419565b9150509295509295509295565b5f8060408385031215611a0e57611a0d611275565b5b5f611a1b858286016112c3565b9250506020611a2c858286016112c3565b9150509250929050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b611a7b611a7682611a36565b611a61565b82525050565b5f819050919050565b5f611aa4611a9f611a9a8461127d565b611a81565b61127d565b9050919050565b5f611ab582611a8a565b9050919050565b5f611ac682611aab565b9050919050565b5f8160601b9050919050565b5f611ae382611acd565b9050919050565b5f611af482611ad9565b9050919050565b611b0c611b0782611abc565b611aea565b82525050565b5f819050919050565b611b2c611b2782611555565b611b12565b82525050565b611b43611b3e8261129c565b611aea565b82525050565b5f81905092915050565b7f73657441747472696275746500000000000000000000000000000000000000005f82015250565b5f611b87600c83611b49565b9150611b9282611b53565b600c82019050919050565b5f819050919050565b611bb7611bb2826112d7565b611b9d565b82525050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f611be982611bbd565b611bf38185611bc7565b9350611c03818560208601611bd1565b80840191505092915050565b5f611c1a828b611a6a565b600182019150611c2a828a611a6a565b600182019150611c3a8289611afb565b601482019150611c4a8288611b1b565b602082019150611c5a8287611b32565b601482019150611c6982611b7b565b9150611c758286611ba6565b602082019150611c858285611bdf565b9150611c918284611b1b565b6020820191508190509998505050505050505050565b7f6368616e67654f776e65720000000000000000000000000000000000000000005f82015250565b5f611cdb600b83611b49565b9150611ce682611ca7565b600b82019050919050565b5f611cfc8289611a6a565b600182019150611d0c8288611a6a565b600182019150611d1c8287611afb565b601482019150611d2c8286611b1b565b602082019150611d3c8285611b32565b601482019150611d4b82611ccf565b9150611d578284611b32565b601482019150819050979650505050505050565b611d74816112d7565b82525050565b5f602082019050611d8d5f830184611d6b565b92915050565b7f7265766f6b6544656c65676174650000000000000000000000000000000000005f82015250565b5f611dc7600e83611b49565b9150611dd282611d93565b600e82019050919050565b5f611de8828a611a6a565b600182019150611df88289611a6a565b600182019150611e088288611afb565b601482019150611e188287611b1b565b602082019150611e288286611b32565b601482019150611e3782611dbb565b9150611e438285611ba6565b602082019150611e538284611b32565b60148201915081905098975050505050505050565b7f61646444656c65676174650000000000000000000000000000000000000000005f82015250565b5f611e9c600b83611b49565b9150611ea782611e68565b600b82019050919050565b5f611ebd828b611a6a565b600182019150611ecd828a611a6a565b600182019150611edd8289611afb565b601482019150611eed8288611b1b565b602082019150611efd8287611b32565b601482019150611f0c82611e90565b9150611f188286611ba6565b602082019150611f288285611b32565b601482019150611f388284611b1b565b6020820191508190509998505050505050505050565b7f7265766f6b6541747472696275746500000000000000000000000000000000005f82015250565b5f611f82600f83611b49565b9150611f8d82611f4e565b600f82019050919050565b5f611fa3828a611a6a565b600182019150611fb38289611a6a565b600182019150611fc38288611afb565b601482019150611fd38287611b1b565b602082019150611fe38286611b32565b601482019150611ff282611f76565b9150611ffe8285611ba6565b60208201915061200e8284611bdf565b915081905098975050505050505050565b5f82825260208201905092915050565b7f6261645f6163746f7200000000000000000000000000000000000000000000005f82015250565b5f61206360098361201f565b915061206e8261202f565b602082019050919050565b5f6020820190508181035f83015261209081612057565b9050919050565b5f82825260208201905092915050565b5f6120b182611bbd565b6120bb8185612097565b93506120cb818560208601611bd1565b6120d481611312565b840191505092915050565b5f819050919050565b5f6121026120fd6120f8846120df565b611a81565b611555565b9050919050565b612112816120e8565b82525050565b5f60808201905061212b5f830187611d6b565b818103602083015261213d81866120a7565b905061214c6040830185612109565b612159606083018461155e565b95945050505050565b61216b81611586565b82525050565b5f6080820190506121845f830187611d6b565b6121916020830186612162565b61219e6040830185611d6b565b6121ab6060830184611d6b565b95945050505050565b7f6261645f7369676e6174757265000000000000000000000000000000000000005f82015250565b5f6121e8600d8361201f565b91506121f3826121b4565b602082019050919050565b5f6020820190508181035f830152612215816121dc565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61225382611555565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036122855761228461221c565b5b600182019050919050565b5f61229a82611555565b91506122a583611555565b92508282019050808211156122bd576122bc61221c565b5b92915050565b5f6080820190506122d65f830187611d6b565b81810360208301526122e881866120a7565b90506122f7604083018561155e565b612304606083018461155e565b95945050505050565b5f6040820190506123205f8301856114dd565b61232d602083018461155e565b9392505050565b5f6080820190506123475f830187611d6b565b61235460208301866114dd565b612361604083018561155e565b61236e606083018461155e565b9594505050505056fea26469706673582212207bb24bcbe62e66b70241ecc87fcf753f82515d93b5869cea10813394cfba3dc964736f6c634300081a0033",
}

// EthrDidRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EthrDidRegistryMetaData.ABI instead.
var EthrDidRegistryABI = EthrDidRegistryMetaData.ABI

// EthrDidRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthrDidRegistryMetaData.Bin instead.
var EthrDidRegistryBin = EthrDidRegistryMetaData.Bin

// DeployEthrDidRegistry deploys a new Ethereum contract, binding an instance of EthrDidRegistry to it.
func DeployEthrDidRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthrDidRegistry, error) {
	parsed, err := EthrDidRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthrDidRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthrDidRegistry{EthrDidRegistryCaller: EthrDidRegistryCaller{contract: contract}, EthrDidRegistryTransactor: EthrDidRegistryTransactor{contract: contract}, EthrDidRegistryFilterer: EthrDidRegistryFilterer{contract: contract}}, nil
}

// EthrDidRegistry is an auto generated Go binding around an Ethereum contract.
type EthrDidRegistry struct {
	EthrDidRegistryCaller     // Read-only binding to the contract
	EthrDidRegistryTransactor // Write-only binding to the contract
	EthrDidRegistryFilterer   // Log filterer for contract events
}

// EthrDidRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthrDidRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthrDidRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthrDidRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthrDidRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthrDidRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthrDidRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthrDidRegistrySession struct {
	Contract     *EthrDidRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthrDidRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthrDidRegistryCallerSession struct {
	Contract *EthrDidRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EthrDidRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthrDidRegistryTransactorSession struct {
	Contract     *EthrDidRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EthrDidRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthrDidRegistryRaw struct {
	Contract *EthrDidRegistry // Generic contract binding to access the raw methods on
}

// EthrDidRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthrDidRegistryCallerRaw struct {
	Contract *EthrDidRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EthrDidRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthrDidRegistryTransactorRaw struct {
	Contract *EthrDidRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthrDidRegistry creates a new instance of EthrDidRegistry, bound to a specific deployed contract.
func NewEthrDidRegistry(address common.Address, backend bind.ContractBackend) (*EthrDidRegistry, error) {
	contract, err := bindEthrDidRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistry{EthrDidRegistryCaller: EthrDidRegistryCaller{contract: contract}, EthrDidRegistryTransactor: EthrDidRegistryTransactor{contract: contract}, EthrDidRegistryFilterer: EthrDidRegistryFilterer{contract: contract}}, nil
}

// NewEthrDidRegistryCaller creates a new read-only instance of EthrDidRegistry, bound to a specific deployed contract.
func NewEthrDidRegistryCaller(address common.Address, caller bind.ContractCaller) (*EthrDidRegistryCaller, error) {
	contract, err := bindEthrDidRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryCaller{contract: contract}, nil
}

// NewEthrDidRegistryTransactor creates a new write-only instance of EthrDidRegistry, bound to a specific deployed contract.
func NewEthrDidRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EthrDidRegistryTransactor, error) {
	contract, err := bindEthrDidRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryTransactor{contract: contract}, nil
}

// NewEthrDidRegistryFilterer creates a new log filterer instance of EthrDidRegistry, bound to a specific deployed contract.
func NewEthrDidRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EthrDidRegistryFilterer, error) {
	contract, err := bindEthrDidRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryFilterer{contract: contract}, nil
}

// bindEthrDidRegistry binds a generic wrapper to an already deployed contract.
func bindEthrDidRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthrDidRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthrDidRegistry *EthrDidRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthrDidRegistry.Contract.EthrDidRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthrDidRegistry *EthrDidRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.EthrDidRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthrDidRegistry *EthrDidRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.EthrDidRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthrDidRegistry *EthrDidRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthrDidRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthrDidRegistry *EthrDidRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthrDidRegistry *EthrDidRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.contract.Transact(opts, method, params...)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCaller) Changed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "changed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistrySession) Changed(arg0 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Changed(&_EthrDidRegistry.CallOpts, arg0)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Changed(&_EthrDidRegistry.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCaller) Delegates(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "delegates", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistrySession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Delegates(&_EthrDidRegistry.CallOpts, arg0, arg1, arg2)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Delegates(&_EthrDidRegistry.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthrDidRegistry *EthrDidRegistryCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "identityOwner", identity)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthrDidRegistry *EthrDidRegistrySession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _EthrDidRegistry.Contract.IdentityOwner(&_EthrDidRegistry.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _EthrDidRegistry.Contract.IdentityOwner(&_EthrDidRegistry.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistrySession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Nonce(&_EthrDidRegistry.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _EthrDidRegistry.Contract.Nonce(&_EthrDidRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthrDidRegistry *EthrDidRegistryCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthrDidRegistry *EthrDidRegistrySession) Owners(arg0 common.Address) (common.Address, error) {
	return _EthrDidRegistry.Contract.Owners(&_EthrDidRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _EthrDidRegistry.Contract.Owners(&_EthrDidRegistry.CallOpts, arg0)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthrDidRegistry *EthrDidRegistryCaller) ValidDelegate(opts *bind.CallOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _EthrDidRegistry.contract.Call(opts, &out, "validDelegate", identity, delegateType, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthrDidRegistry *EthrDidRegistrySession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _EthrDidRegistry.Contract.ValidDelegate(&_EthrDidRegistry.CallOpts, identity, delegateType, delegate)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthrDidRegistry *EthrDidRegistryCallerSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _EthrDidRegistry.Contract.ValidDelegate(&_EthrDidRegistry.CallOpts, identity, delegateType, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) AddDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "addDelegate", identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.AddDelegate(&_EthrDidRegistry.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.AddDelegate(&_EthrDidRegistry.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) AddDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "addDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.AddDelegateSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.AddDelegateSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.ChangeOwner(&_EthrDidRegistry.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.ChangeOwner(&_EthrDidRegistry.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.ChangeOwnerSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.ChangeOwnerSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeAttribute(&_EthrDidRegistry.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeAttribute(&_EthrDidRegistry.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeAttributeSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeAttributeSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) RevokeDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "revokeDelegate", identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeDelegate(&_EthrDidRegistry.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeDelegate(&_EthrDidRegistry.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) RevokeDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "revokeDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeDelegateSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.RevokeDelegateSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.SetAttribute(&_EthrDidRegistry.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.SetAttribute(&_EthrDidRegistry.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistrySession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.SetAttributeSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthrDidRegistry *EthrDidRegistryTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthrDidRegistry.Contract.SetAttributeSigned(&_EthrDidRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// EthrDidRegistryDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDAttributeChangedIterator struct {
	Event *EthrDidRegistryDIDAttributeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthrDidRegistryDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthrDidRegistryDIDAttributeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthrDidRegistryDIDAttributeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthrDidRegistryDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthrDidRegistryDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthrDidRegistryDIDAttributeChanged represents a DIDAttributeChanged event raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*EthrDidRegistryDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryDIDAttributeChangedIterator{contract: _EthrDidRegistry.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *EthrDidRegistryDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthrDidRegistryDIDAttributeChanged)
				if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDAttributeChanged is a log parse operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) ParseDIDAttributeChanged(log types.Log) (*EthrDidRegistryDIDAttributeChanged, error) {
	event := new(EthrDidRegistryDIDAttributeChanged)
	if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthrDidRegistryDIDDelegateChangedIterator is returned from FilterDIDDelegateChanged and is used to iterate over the raw logs and unpacked data for DIDDelegateChanged events raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDDelegateChangedIterator struct {
	Event *EthrDidRegistryDIDDelegateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthrDidRegistryDIDDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthrDidRegistryDIDDelegateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthrDidRegistryDIDDelegateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthrDidRegistryDIDDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthrDidRegistryDIDDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthrDidRegistryDIDDelegateChanged represents a DIDDelegateChanged event raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDDelegateChanged struct {
	Identity       common.Address
	DelegateType   [32]byte
	Delegate       common.Address
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDDelegateChanged is a free log retrieval operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) FilterDIDDelegateChanged(opts *bind.FilterOpts, identity []common.Address) (*EthrDidRegistryDIDDelegateChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.FilterLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryDIDDelegateChangedIterator{contract: _EthrDidRegistry.contract, event: "DIDDelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDIDDelegateChanged is a free log subscription operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) WatchDIDDelegateChanged(opts *bind.WatchOpts, sink chan<- *EthrDidRegistryDIDDelegateChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.WatchLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthrDidRegistryDIDDelegateChanged)
				if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDDelegateChanged is a log parse operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) ParseDIDDelegateChanged(log types.Log) (*EthrDidRegistryDIDDelegateChanged, error) {
	event := new(EthrDidRegistryDIDDelegateChanged)
	if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthrDidRegistryDIDOwnerChangedIterator is returned from FilterDIDOwnerChanged and is used to iterate over the raw logs and unpacked data for DIDOwnerChanged events raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDOwnerChangedIterator struct {
	Event *EthrDidRegistryDIDOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthrDidRegistryDIDOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthrDidRegistryDIDOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthrDidRegistryDIDOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthrDidRegistryDIDOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthrDidRegistryDIDOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthrDidRegistryDIDOwnerChanged represents a DIDOwnerChanged event raised by the EthrDidRegistry contract.
type EthrDidRegistryDIDOwnerChanged struct {
	Identity       common.Address
	Owner          common.Address
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDOwnerChanged is a free log retrieval operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) FilterDIDOwnerChanged(opts *bind.FilterOpts, identity []common.Address) (*EthrDidRegistryDIDOwnerChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.FilterLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthrDidRegistryDIDOwnerChangedIterator{contract: _EthrDidRegistry.contract, event: "DIDOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchDIDOwnerChanged is a free log subscription operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) WatchDIDOwnerChanged(opts *bind.WatchOpts, sink chan<- *EthrDidRegistryDIDOwnerChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthrDidRegistry.contract.WatchLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthrDidRegistryDIDOwnerChanged)
				if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDOwnerChanged is a log parse operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthrDidRegistry *EthrDidRegistryFilterer) ParseDIDOwnerChanged(log types.Log) (*EthrDidRegistryDIDOwnerChanged, error) {
	event := new(EthrDidRegistryDIDOwnerChanged)
	if err := _EthrDidRegistry.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
