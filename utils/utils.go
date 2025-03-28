package utils

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
)

type ContextKey string

func LoadKey(keyPath, password string) (*keystore.Key, error) {
	keyjson, err := os.ReadFile(keyPath)
	if err != nil {
		log.Error("Failed to read key file: ", err)
		return nil, errors.Wrap(err, "Failed to read key file")
	}

	return keystore.DecryptKey(keyjson, password)
}

func LoadETHClient(endpoint string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Error("Failed to connect to the Ethereum client: ", err)
		return nil, errors.Wrap(err, "Failed to connect to the Ethereum client")
	}

	log.Info("Successfully connected to the Ethereum client")
	return client, nil
}

func LoadAuth(key *keystore.Key, network int64) (*bind.TransactOpts, error) {
	auth,err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, big.NewInt(network)) // 1 是主网链 ID
	if err != nil{
		log.Error("Failed to create authorized transactor: ", err)
		return nil,errors.Wrap(err,"Failed to create authorized transactor")
	}
	return auth,nil
}