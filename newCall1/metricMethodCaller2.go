package call

import (
	"context"
	"crypto/ecdsa"
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
//
//	transactionAddDetails, errWhenAddingDetails := contractBuilded.BuildTransactor.contract.Transact(auth1, "addDetails")
//	if errWhenAddingDetails != nil {
//		fmt.Println("8 ", fmt.Errorf(errWhenAddingDetails.Error()))
//		return
//	}
//	fmt.Println("Data added : TRANSACTION : ", transactionAddDetails.Hash().Hex())
//
//	time.Sleep(30 * time.Second)

	//set the arrays
	methodArrayForFormulas := [...]string{"getFertilizer_Emission", "getWater_Emission"}
	methodArrayForValues := [...]string{"getFERTILIZER_EMISSION", "getCOMPOST_QUANTITY", "getWATER"}

	for i := 0; i < len(methodArrayForFormulas); i++ {
		fmt.Println("--------------------------Executing Formula Getter : ", methodArrayForFormulas[i])

		var formula []interface{}
		errWhenGettingFormulaDetails := contractBuilded.BuildCaller.contract.Call(&bind.CallOpts{}, &formula, methodArrayForFormulas[i])
		if errWhenGettingFormulaDetails != nil {
			fmt.Println("9 ", fmt.Errorf(errWhenGettingFormulaDetails.Error()))
			return
		}
		fmt.Println("Formula : ", formula[0])
		formulaString1, errInMarshalling1 := json.Marshal(formula[0])
		if errInMarshalling1 != nil {
			fmt.Println(errInMarshalling1)
			return
		}
		fmt.Println(string(formulaString1))
	
		var formula1 Formula
		errWhenUnMarshalling1 := json.Unmarshal([]byte(formulaString1), &formula1)
		if errWhenUnMarshalling1 != nil {
			fmt.Println("10 ", fmt.Errorf(errWhenUnMarshalling1.Error()))
			return
		}
		fmt.Println("Formula details for formula ID: " + formula1.FormulaID + " -> ")
		fmt.Printf("%v\n", formula1)
		fmt.Println()
	}

	for i := 0; i < len(methodArrayForValues); i++ {
		fmt.Println("--------------------------Executing Value Getter : ", methodArrayForValues[i])

		var value []interface{}
		errWhenGettingValueDetails := contractBuilded.BuildCaller.contract.Call(&bind.CallOpts{}, &value, methodArrayForValues[i])
		if errWhenGettingValueDetails != nil {
			fmt.Println("11 ", fmt.Errorf(errWhenGettingValueDetails.Error()))
			return
		}
		fmt.Println("Value : ", value[0])
		valueString1, errInMarshalling2 := json.Marshal(value[0])
		if errInMarshalling2 != nil {
			fmt.Println(errInMarshalling2)
			return
		}
		fmt.Println(string(valueString1))

		var value1 Value

		errWhenUnMarshalling2 := json.Unmarshal([]byte(valueString1), &value1)
		if errWhenUnMarshalling2 != nil {
			fmt.Println("12 ", fmt.Errorf(errWhenUnMarshalling2.Error()))
			return
		}
		fmt.Println("Value details for value ID: " + value1.ValueID + " -> ")
		fmt.Printf("%v\n", value1)
		fmt.Println()
	}
}

func bindBuild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuildABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}
