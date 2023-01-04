package deploy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetGasPrice() int {
	url := "https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=AER6M2C3436231IGT7SV7JZ2URFYFX7MZ1"
	result, errWwhenCallingUrl := http.Get(url)
	if errWwhenCallingUrl != nil {
		fmt.Println("Unable to reach etherscan API: ", errWwhenCallingUrl)
		return 0
	}

	if result.StatusCode != 200 {
		fmt.Println("Status not 200 : ", result)
		return 0
	}

	data, errWhenReading := ioutil.ReadAll(result.Body)
	if errWhenReading != nil {
		fmt.Println("Unable to read response from etherscan API: ", errWhenReading)
		return 0
	}
	defer result.Body.Close()

	value := gjson.GetBytes(data, "result.FastGasPrice")
	fmt.Println(value.String())

	intValue, _ := strconv.Atoi(value.String())
	finalWeiAmount := intValue * 1000000000

	fmt.Println(finalWeiAmount)
	return finalWeiAmount
}
