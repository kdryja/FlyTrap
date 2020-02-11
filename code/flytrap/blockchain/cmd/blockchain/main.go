package main

import (
	"flag"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
)

const CLIENT = "0x641c46D43A3c552a76318c12D8Fc2839b913F32F"

var (
	contract = flag.String("contract", blockchain.CONTRACT, "Specify address of your contract")
	newTopic = flag.Bool("new_topic", false, "Whether new topic should be created")
	addPub   = flag.Bool("pub", false, "Client to add as publisher")
)

func main() {
	flag.Parse()
	b, err := blockchain.New()
	if err != nil {
		log.Fatal(err)
	}
	var addr common.Address
	if *contract == "" {
		if addr, err = b.GenerateContract(100); err != nil {
			log.Fatal(err)
		}
	} else {
		addr = common.HexToAddress(*contract)
	}
	if err := b.SetInstance(addr); err != nil {
		log.Fatal(err)
	}
	t := [32]byte{}
	copy(t[:], "MyTopic3")
	if *newTopic {
		b.Opts.Value = big.NewInt(1000)
		if _, err := b.Instance.AddTopic(b.Opts, t, big.NewInt(0), big.NewInt(0)); err != nil {
			log.Fatal(err)
		}
	}
	if *addPub {
		if err := b.AddPub(CLIENT, "MyTopic3"); err != nil {
			log.Fatal(err)
		}
	}
	ok, err := blockchain.Verify("MyTopic3", CLIENT, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(ok)
}
