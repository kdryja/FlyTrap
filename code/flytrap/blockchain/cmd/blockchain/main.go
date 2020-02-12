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
	contract    = flag.String("contract", blockchain.FLYTRAP_CONTRACT, "Specify address of your contract")
	newTopic    = flag.Bool("new_topic", false, "Whether new topic should be created")
	addPub      = flag.Bool("pub", false, "Client to add as publisher")
	addSub      = flag.Bool("sub", false, "Client to add as subscriber")
	addContract = flag.Bool("new", false, "Whether to create new contracts for authorizer and flytrap")
	topicName   = flag.String("topic", "MyTopic", "name of the topic to modify")
)

func main() {
	flag.Parse()
	b, err := blockchain.New()
	if err != nil {
		log.Fatal(err)
	}
	var addr common.Address
	if *addContract {
		addr, _, _, err := blockchain.DeployFlytrap(b.Opts, b.Client, big.NewInt(0))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Generated new contract, address is: %s", addr.Hex())
	}
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
	if *newTopic {
		t := [32]byte{}
		copy(t[:], *topicName)
		b.Opts.Value = big.NewInt(1000)
		if _, err := b.Instance.AddTopic(b.Opts, t, big.NewInt(0), big.NewInt(0)); err != nil {
			log.Fatal(err)
		}
	}
	if *addPub {
		if err := b.AddPub(CLIENT, *topicName); err != nil {
			log.Fatal(err)
		}
	}
	if *addSub {
		if err := b.AddSub(CLIENT, *topicName); err != nil {
			log.Fatal(err)
		}
	}
}
