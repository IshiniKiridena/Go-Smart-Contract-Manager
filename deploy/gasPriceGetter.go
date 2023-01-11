package deploy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GasPrice struct {
	Status  string
	Message string
	Result  struct {
		LastBlock       string
		SafeGasPrice    string
		ProposeGasPrice string
		FastGasPrice    string
		SuggestBaseFee  string
		GasUsedRatio    string
	}
}

func GetGasPrice() (string, error) {

	url := "https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=AER6M2C3436231IGT7SV7JZ2URFYFX7MZ1"
	method := "GET"

	client := &http.Client{}
	req, err1 := http.NewRequest(method, url, nil)

	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return "", err2
	}
	defer res.Body.Close()

	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println(err3)
		return "", err3
	}

	var gasPrice GasPrice 
	err4 := json.Unmarshal(body, &gasPrice)
	if err4 != nil {
		fmt.Println(err4)
		return "", err4
	}

	return gasPrice.Result.SafeGasPrice, nil
}
