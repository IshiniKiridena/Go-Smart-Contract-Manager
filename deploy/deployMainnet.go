package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployToMainnet(abi string, bin string) {
	seed := ""

	//Dial infura client
	client, errWhenDialingEthClinet := ethclient.Dial("https://mainnet.infura.io/v3/f7ad4e6f2bd54303b26fb0e0679752f8")
	if errWhenDialingEthClinet != nil {
		fmt.Println("Error when dialing the eth client " + errWhenDialingEthClinet.Error())
		return
	}

	//load ECDSA private key
	privateKey, errWhenGettingECDSAKey := crypto.HexToECDSA(seed)
	if errWhenGettingECDSAKey != nil {
		fmt.Println("Error when getting ECDSA key " + errWhenGettingECDSAKey.Error())
		return
	}

	//get the public key
	publicKey := privateKey.Public()

	//get public key ECDSA
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	//get nonce
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, errWhenGettingNonce := client.PendingNonceAt(context.Background(), fromAddress)
	if errWhenGettingNonce != nil {
		fmt.Println("Error when getting nonce " + errWhenGettingNonce.Error())
		return
	}

	// //get the gas price
	// gasPrice, errWhenGettingGasPrice := client.EstimateGas(context.Background(), ethereum.CallMsg{})
	// if errWhenGettingGasPrice != nil {
	// 	fmt.Println("Error when getting gas price " + errWhenGettingGasPrice.Error())
	// 	return
	// }
	// fmt.Println("Nonce : ", nonce)
	// fmt.Println("Gas price: ", gasPrice)
	gasPrice := GetGasPrice()
	fmt.Println(gasPrice)

	//create the keyed transactor
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(100000) // in units
	auth.GasPrice = big.NewInt(int64(gasPrice))

	//assign metadata for the contract
	var BuildData = &bind.MetaData{
		ABI: abi,
		Bin: bin,
	}

	//var ContractABI = BuildData.ABI
	var ContractBIN = BuildData.Bin

	parsed, errWhenGettingABI := BuildData.GetAbi()
	if errWhenGettingABI != nil {
		fmt.Println("Error when getting abi from passed ABI string " + errWhenGettingABI.Error())
		return
	}

	if parsed == nil {
		fmt.Println("Error when getting ABI string , ERROR : GetAbi() returned nil")
		return
	}

	address, tx, contract, errWhenDeployingContract := bind.DeployContract(auth, *parsed, common.FromHex(ContractBIN), client)
	if errWhenDeployingContract != nil {
		fmt.Println("Error when deploying contract " + errWhenDeployingContract.Error())
		return
	}

	fmt.Println("1 ", address)
	fmt.Println("2 ", tx)
	fmt.Println("3 ", contract)

	contractAddress := address.Hex()
	transactionHash := tx.Hash().Hex()
	_ = contract

	fmt.Println("View contract at : https://etherscan.io/address/", address.Hex())
	fmt.Println("View transaction at : https://etherscan.io/tx/", tx.Hash().Hex())

	fmt.Println("Contract address : ", contractAddress)
	fmt.Println("Transaction hash : ", transactionHash)

	// Wait for the transaction to be mined and calculate the cost
	receipt, errInGettingReceipt := bind.WaitMined(context.Background(), client, tx)
	fmt.Println("Receipt", receipt)
	if errInGettingReceipt != nil {
		fmt.Println("Error in getting receipt: Error: " + errInGettingReceipt.Error())
	} else {
		if receipt.Status == 0 {
			fmt.Println("Transaction failed.")
			return
		} else {
			fmt.Println("Contract address : ", contractAddress)
			fmt.Println("Transaction hash : ", transactionHash)
		}
	}

}
