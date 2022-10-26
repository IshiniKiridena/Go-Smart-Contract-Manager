package main

import (
	"fmt"

	"github.com/IshiniKiridena/SmartC/deploy"
)

func main() {
	//keys.GenerateAccounts()
	errInDepl := deploy.DeployContract()
	if errInDepl != nil {
		fmt.Println(errInDepl)
	} else {
		fmt.Println("Contract deployed")
	}
}
