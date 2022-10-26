package load

import (
	"fmt"
	"log"

	store "github.com/IshiniKiridena/SmartC/build"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadContract() {
	//contrct loading
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/b058f78814274e9eb6cb3796fa0527e4")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xd07a669b8046a640cb468b4ce038569d72695868")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded, instance : ", instance)

	//contrat reading
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version : ", version)

	_ = instance
}
