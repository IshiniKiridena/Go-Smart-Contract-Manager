package call

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PivotField struct {
	Name string
	Key string 
	Field string
	Condition string
	Value string
	ArtifactTemplateId string
	ArtifactDataId string
	FormulaId string
}

func CreateInstance2(abi string, bin string, contractAdd string) {
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

//	transactionAddDetails, errWhenAddingDetails := contractBuilded.BuildTransactor.contract.Transact(auth1, "addDetails")
//	if errWhenAddingDetails != nil {
//		fmt.Println("8 ", fmt.Errorf(errWhenAddingDetails.Error()))
//		return
//	}
//	fmt.Println("Data added : TRANSACTION : ", transactionAddDetails.Hash().Hex())
//
//	time.Sleep(30 * time.Second)

	fmt.Println("\n-------------------Calling getter for formulas-------------------")
	// Calling getter for formula
	var formula []interface{}
	errWhenGettingFormulaDetails := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &formula, "getFormulaDetails")
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

	var formula1 []Formula
	errWhenUnMarshalling1 := json.Unmarshal([]byte(formulaString1), &formula1)
	if errWhenUnMarshalling1 != nil {
		fmt.Println("10 ", fmt.Errorf(errWhenUnMarshalling1.Error()))
		return
	}

	for i:=0; i<len(formula1); i++ {
		fmt.Println("Formula ID: ", formula1[i].FormulaID)
		fmt.Println("Contract Address: ", formula1[i].ContractAddress)
		fmt.Println("No of values: ", formula1[i].NoOfValues)
		fmt.Println("Activity ID: ", formula1[i].ActivityID)
		fmt.Println("Activity Name: ", HexStringToString(formula1[i].ActivityName))
		fmt.Println("Value IDs: ", formula1[i].ValueIDs)
		fmt.Println()
	}

	fmt.Println("\n-------------------Calling getter for values-------------------")

	// Calling getter for value
	var value []interface{}
	errWhenGettingFormulaDetails1 := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &value, "getValueDetails")
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

	var value1 []Value

	errWhenUnMarshalling2 := json.Unmarshal([]byte(valueString1), &value1)
	if errWhenUnMarshalling2 != nil {
		fmt.Println("12 ", fmt.Errorf(errWhenUnMarshalling2.Error()))
		return
	}

	for i:=0; i<len(value1); i++ {
		fmt.Println("Value ID: ", value1[i].ValueID)
		fmt.Println("Value Name: ", value1[i].ValueName)
		fmt.Println("Workflow ID: ", value1[i].WorkflowID)
		fmt.Println("Stage ID: ", value1[i].StageID)
		fmt.Println("Stage Name: ", HexStringToString(value1[i].StageName))
		fmt.Println("Key Name: ", HexStringToString(value1[i].KeyName))
		fmt.Println("TDP Type: ", value1[i].TdpType)
		fmt.Println("Binding Type: ", value1[i].BindingType)
		fmt.Println("Artifact ID: ", value1[i].ArtifactID)
		fmt.Println("Primary Key Row ID: ", value1[i].PrimaryKeyRowID)
		fmt.Println("Artifact Template Name: ", HexStringToString(value1[i].ArtifactTemplateName))
		fmt.Println("Field Key: ", HexStringToString(value1[i].FieldKey))
		fmt.Println("Field Name: ", HexStringToString(value1[i].FieldName))
		fmt.Println()
	}

	fmt.Println("\n-------------------Calling getter for pivot fields-------------------")
	// Calling getter for pivot field
	var pivot []interface{}
	errWhenGettingPivotFields := contractBuilded.BuildTransactor.contract.Call(&bind.CallOpts{}, &pivot, "getPivotFields")
	if errWhenGettingPivotFields != nil {
		fmt.Println("13 ", fmt.Errorf(errWhenGettingPivotFields.Error()))
		return
	}
	//fmt.Println("Pivots : ", pivot[0])
	PivotsString1, errInMarshalling1 := json.Marshal(pivot[0])
	if errInMarshalling1 != nil {
		fmt.Println(errInMarshalling1)
		return
	}
	//fmt.Println(string(PivotsString1))

	var pivot1 []PivotField
	errWhenUnMarshalling3 := json.Unmarshal([]byte(PivotsString1), &pivot1)
	if errWhenUnMarshalling3 != nil {
		fmt.Println("14 ", fmt.Errorf(errWhenUnMarshalling3.Error()))
		return
	}

	for i:=0; i<len(pivot1); i++ {
		fmt.Println("Name: ", pivot1[i].Name)
		fmt.Println("Key: ", HexStringToString(pivot1[i].Key))
		fmt.Println("Field: ", pivot1[i].Field)
		fmt.Println("Condition: ", pivot1[i].Condition)
		fmt.Println("Value: ", pivot1[i].Value)
		fmt.Println("Artifact Template ID: ", pivot1[i].ArtifactTemplateId)
		fmt.Println("Artifact Data ID: ", pivot1[i].ArtifactDataId)
		fmt.Println("Formula ID: ", pivot1[i].FormulaId)
		fmt.Println()
	}
}
