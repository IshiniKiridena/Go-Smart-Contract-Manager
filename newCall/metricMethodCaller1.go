package call

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Value struct {
	ValueID              string
	ValueName            string
	WorkflowID           string
	StageID              string
	StageName            string
	KeyName              string
	TdpType              string
	BindingType          int
	ArtifactID           string
	PrimaryKeyRowID      string
	ArtifactTemplateName string
	FieldKey             string
	FieldName            string
}

type Formula struct {
	FormulaID       string
	ContractAddress string
	NoOfValues      int
	ActivityID      string
	ActivityName    string
	ValueIDs        string
}

var BuildABI string

func CreateInstance1(abi string, bin string, contractAdd string) {
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

	//transactionAddDetails, errWhenAddingDetails := contractBuilded.BuildTransactor.contract.Transact(auth1, "addDetails")
	//if errWhenAddingDetails != nil {
	//	fmt.Println("8 ", fmt.Errorf(errWhenAddingDetails.Error()))
	//	return
	//}
	//fmt.Println("Data added : TRANSACTION : ", transactionAddDetails.Hash().Hex())

	//time.Sleep(30 * time.Second)

	//set the arrays
	valueArrayForFormulas := [...]string{"631a0b4ad9241a9374fConfig00", "631a0b4ad9241a9374fConfig02"}
	valueArrayForValues := [...]string{"123123123123aaaaa00", "e1223109481cf739d19e6735d0236577", "e1223109481cf739d19e6735d0236577"}

	fmt.Println("\n-------------------Calling getter for formulas-------------------")
	for i := 0; i < len(valueArrayForFormulas); i++ {
		// Calling getter for formula
		var formula []interface{}
		errWhenGettingFormulaDetails := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &formula, "getFormulaDetails", valueArrayForFormulas[i])
		if errWhenGettingFormulaDetails != nil {
			fmt.Println("9 ", fmt.Errorf(errWhenGettingFormulaDetails.Error()))
			return
		}
		//fmt.Println("Formula : ", formula[0])
		formulaString1, errInMarshalling1 := json.Marshal(formula[0])
		if errInMarshalling1 != nil {
			fmt.Println(errInMarshalling1)
			return
		}
		//fmt.Println(string(formulaString1))

		var formula1 Formula
		errWhenUnMarshalling1 := json.Unmarshal([]byte(formulaString1), &formula1)
		if errWhenUnMarshalling1 != nil {
			fmt.Println("10 ", fmt.Errorf(errWhenUnMarshalling1.Error()))
			return
		}
		fmt.Println("Formula details for formula ID: " + valueArrayForFormulas[i] + " -> ")
		fmt.Println("Formula ID: ", formula1.FormulaID)
		fmt.Println("Contract Address: ", formula1.ContractAddress)
		fmt.Println("No of values: ", formula1.NoOfValues)
		fmt.Println("Activity ID: ", formula1.ActivityID)
		fmt.Println("Activity Name: ", HexStringToString(formula1.ActivityName))
		fmt.Println("Value IDs: ", formula1.ValueIDs)
		fmt.Println()
	}

	fmt.Println("\n-------------------Calling getter for values-------------------")
	for i := 0; i < len(valueArrayForValues); i++ {
		// Calling getter for value
		var value []interface{}
		errWhenGettingFormulaDetails1 := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &value, "getValueDetails", valueArrayForValues[i])
		if errWhenGettingFormulaDetails1 != nil {
			fmt.Println("11 ", fmt.Errorf(errWhenGettingFormulaDetails1.Error()))
			return
		}
		//fmt.Println("Value : ", value[0])
		valueString1, errInMarshalling2 := json.Marshal(value[0])
		if errInMarshalling2 != nil {
			fmt.Println(errInMarshalling2)
			return
		}
		//fmt.Println(string(valueString1))

		var value1 Value

		errWhenUnMarshalling2 := json.Unmarshal([]byte(valueString1), &value1)
		if errWhenUnMarshalling2 != nil {
			fmt.Println("12 ", fmt.Errorf(errWhenUnMarshalling2.Error()))
			return
		}
		fmt.Println("Value details for value ID: " + valueArrayForValues[i] + " -> ")
		fmt.Println("Value ID: ", value1.ValueID)
		fmt.Println("Value Name: ", value1.ValueName)
		fmt.Println("Workflow ID: ", value1.WorkflowID)
		fmt.Println("Stage ID: ", value1.StageID)
		fmt.Println("Stage Name: ", HexStringToString(value1.StageName))
		fmt.Println("Key Name: ", HexStringToString(value1.KeyName))
		fmt.Println("TDP Type: ", value1.TdpType)
		fmt.Println("Binding Type: ", value1.BindingType)
		fmt.Println("Artifact ID: ", value1.ArtifactID)
		fmt.Println("Primary Key Row ID: ", value1.PrimaryKeyRowID)
		fmt.Println("Artifact Template Name: ", HexStringToString(value1.ArtifactTemplateName))
		fmt.Println("Field Key: ", HexStringToString(value1.FieldKey))
		fmt.Println("Field Name: ", HexStringToString(value1.FieldName))
		fmt.Println()
	}
	//	// get the details and add if not present
	//	var out []interface{}
	//	errWhenGettingFormulaDetails := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &out, "getFormulaDetails")
	//	if errWhenGettingFormulaDetails != nil {
	//		fmt.Println("9 ", fmt.Errorf(errWhenGettingFormulaDetails.Error()))
	//		return
	//	}
	//
	//	// if the details are not present then add them
	//	if len(out[0].([]string)) == 0 {
	//		transactionAddDetails, errWhenAddingDetails := contractBuilded.BuildTransactor.contract.Transact(auth1, "addDetails")
	//		if errWhenAddingDetails != nil {
	//			fmt.Println("8 ", fmt.Errorf(errWhenAddingDetails.Error()))
	//			return
	//		}
	//		fmt.Println("Data added : TRANSACTION : ", transactionAddDetails.Hash().Hex())
	//
	//		time.Sleep(30 * time.Second)
	//
	//		out = nil
	//		// get the details after adding
	//		errWhenGettingFormulaDetails1 := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &out, "getFormulaDetails")
	//		if errWhenGettingFormulaDetails1 != nil {
	//			fmt.Println("10 ", fmt.Errorf(errWhenGettingFormulaDetails1.Error()))
	//			return
	//		}
	//	}

	// print the details
	//for _, v1 := range out[0].([]string) {
	//	fmt.Println("v1 : ", v1)
	//}

}

func bindBuild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuildABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func HexStringToString(value string) string {
	hx, _ := hex.DecodeString(value)
	return string(hx)
}