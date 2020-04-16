package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	keystoreFile = flag.String("key", "", "location of your UTC file")
	public       = flag.Bool("pub", false, "whether to output public key")
	private      = flag.Bool("priv", false, "whether to output private key")
	password     = flag.String("pass", "password", "password to unlock your keystore")
)

func main() {
	flag.Parse()
	if *keystoreFile == "" {
		log.Fatal("Please specify keystore file")
	}
	jsonBytes, err := ioutil.ReadFile(*keystoreFile)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(jsonBytes, *password)
	if err != nil {
		log.Fatal(err)
	}
	if *public {
		fmt.Println(key.Address.Hex())
	}
	if *private {
		fmt.Printf("%x\n", crypto.FromECDSA(key.PrivateKey))
	}

}
