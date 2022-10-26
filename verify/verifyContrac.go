package verify

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func VerifyContract() {
	//api key is the ethereum api key
	apiUrl := "https://api-sepolia.etherscan.io/api?module=contract&action=verifyproxycontract&apikey=X9YDUDC3IBDISW6HFU7NU2S7N7BI6MQ1RB"
	data := url.Values{}
	data.Set("address", "0x364cca98C6886c96e8Ff1FF9BE0778ba84aAE07F")

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String() // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("Error, Not verified  ", resp.Status)
	}

	fmt.Println("Contract verified")
}
