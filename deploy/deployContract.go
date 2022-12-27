package deploy

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	metric1 "github.com/IshiniKiridena/SmartC/build"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var pubKey = "0x7fD5918Bc9fB9c0135865F3A36cE5E1295Ec54Ce"
var privKey = "4752159e7dd9f582f651c0feb8a53eb93c2737a8491341ad5869dd6ca067acfd"

func DeployContract() error {
	//this dial should be the infura project id - api key
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/f7ad4e6f2bd54303b26fb0e0679752f8")
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New(err.Error())
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	fmt.Println("Nonce : ", nonce)

	var gasP = 47045172800000

	gasPrice, err := client.EstimateGas(context.Background(), ethereum.CallMsg{})
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	fmt.Println("Gas price pure : ", gasPrice)
	fmt.Println("Generated gas price : ", gasP)

	//creating new keyed transactor
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(int64(gasPrice))
	fmt.Println("Auth  : ", auth)

	//input := "1.0"
	address, tx, instance, err := metric1.DeployStore(auth, client)
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}

	fmt.Println("Contract address : https://sepolia.etherscan.io/address/" + address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println("Transaction Hash ; ", "https://sepolia.etherscan.io/tx/"+tx.Hash().Hex())  // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance

	return nil
}
