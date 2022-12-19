package call

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var BuildABI string

func CreateInstance(abi string, bin string, contractAdd string) {
	type BuildCaller struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	type BuildTransactor struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	type BuildFilterer struct {
		contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	type Build struct {
		BuildCaller
		BuildTransactor
		BuildFilterer // Log filterer for contract events
	}

	//Dial infura client
	client, errWhenDialingEthClinet := ethclient.Dial("https://sepolia.infura.io/v3/f7ad4e6f2bd54303b26fb0e0679752f8")
	if errWhenDialingEthClinet != nil {
		fmt.Println("1 ",
			fmt.Errorf(errWhenDialingEthClinet.Error()))
		return
	}

	//load ECDSA private key
	privateKey, errWhenGettingECDSAKey := crypto.HexToECDSA("4752159e7dd9f582f651c0feb8a53eb93c2737a8491341ad5869dd6ca067acfd")
	if errWhenGettingECDSAKey != nil {
		fmt.Println("2 ",
			fmt.Errorf(errWhenGettingECDSAKey.Error()))
		return
	}

	//get the public key
	publicKey := privateKey.Public()
	//get public key ECDSA
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("3 Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	//get nonce
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce1, errWhenGettingNonce := client.PendingNonceAt(context.Background(), fromAddress)
	if errWhenGettingNonce != nil {
		fmt.Println("4 ", fmt.Errorf(errWhenGettingNonce.Error()))
		return
	}

	//get the gas price
	gasPrice, errWhenGettingGasPrice := client.EstimateGas(context.Background(), ethereum.CallMsg{})
	if errWhenGettingGasPrice != nil {
		fmt.Println("5 ", fmt.Errorf(errWhenGettingGasPrice.Error()))
		return
	}

	//create the keyed transactor
	auth1 := bind.NewKeyedTransactor(privateKey)
	auth1.Nonce = big.NewInt(int64(nonce1))
	auth1.Value = big.NewInt(0)      // in wei
	auth1.GasLimit = uint64(3000000) // in units
	auth1.GasPrice = big.NewInt(int64(gasPrice))

	//assign metadata for the contract
	var BuildData = &bind.MetaData{
		ABI: abi,
		Bin: bin,
	}
	BuildABI = BuildData.ABI
	//var ContractABI = BuildData.ABI
	//var ContractBIN = BuildData.Bin

	parsed, errWhenGettingABI := BuildData.GetAbi()
	if errWhenGettingABI != nil {
		fmt.Println("6 ", fmt.Errorf(errWhenGettingABI.Error()))
		return
	}

	if parsed == nil {
		fmt.Println("ABI returned nil")
		return
	}
	address := common.HexToAddress(contractAdd)
	contract, errWhenBindingContract := bindBuild(address, client, client, client)
	if errWhenBindingContract != nil {
		fmt.Println("7 ", fmt.Errorf(errWhenBindingContract.Error()))
		return
	}

	contractBuilded := &Build{
		BuildCaller:     BuildCaller{contract: contract},
		BuildTransactor: BuildTransactor{contract: contract},
		BuildFilterer:   BuildFilterer{contract: contract},
	}

	//set water method calling
	waterValue := big.NewInt(2)
	exponent := big.NewInt(0)

	transactionWater, errWhenSettingVariableValue := contractBuilded.BuildTransactor.contract.Transact(auth1, "setWATER", waterValue, exponent)
	if errWhenSettingVariableValue != nil {
		fmt.Println("8 ", fmt.Errorf(errWhenSettingVariableValue.Error()))
		return
	}
	fmt.Println("Set water data added : TRANSACTION : ", transactionWater.Hash().Hex())

	time.Sleep(30 * time.Second)

	//get nonce
	nonce2, errWhenGettingNonce := client.PendingNonceAt(context.Background(), fromAddress)
	if errWhenGettingNonce != nil {
		fmt.Println("4 ", fmt.Errorf(errWhenGettingNonce.Error()))
		return
	}

	//create the keyed transactor
	auth2 := bind.NewKeyedTransactor(privateKey)
	auth2.Nonce = big.NewInt(int64(nonce2))
	auth2.Value = big.NewInt(0)      // in wei
	auth2.GasLimit = uint64(3000000) // in units
	auth2.GasPrice = big.NewInt(int64(gasPrice))

	//calculation execution call
	transactionCalculation, errWhenCalculating := contractBuilded.BuildTransactor.contract.Transact(auth2, "executeCalculation")
	if errWhenCalculating != nil {
		fmt.Println("9 ", fmt.Errorf(errWhenCalculating.Error()))
		return
	}

	fmt.Println("Calculation : TRANSACTION : ", transactionCalculation.Hash().Hex())

}

func bindBuild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuildABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}
