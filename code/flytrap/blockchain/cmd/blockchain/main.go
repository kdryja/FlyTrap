package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
)

const CLIENT = "0x8854e627347fD04b4a77DB7dcc6aaa43b7e7135c"

var (
	contract    = flag.String("contract", blockchain.FLYTRAP_CONTRACT, "Specify address of your contract")
	newTopic    = flag.Bool("new_topic", false, "Whether new topic should be created")
	addPub      = flag.Bool("pub", false, "Client to add as publisher")
	addSub      = flag.Bool("sub", false, "Client to add as subscriber")
	addContract = flag.Bool("new", false, "Whether to create new contract for flytrap")
	topicName   = flag.String("topic", "MyTopic", "name of the topic to modify")
	logsRead    = flag.Bool("logs", false, "receive event logs")
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
		if _, err := b.Instance.AddTopic(b.Opts, t, [2]byte{'U', 'K'}, big.NewInt(0), big.NewInt(0)); err != nil {
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
	if *logsRead {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(0),
			ToBlock:   big.NewInt(1000000),
			Addresses: []common.Address{
				common.HexToAddress(blockchain.FLYTRAP_CONTRACT),
			},
		}
		logs, err := b.Client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		contractAbi, err := abi.JSON(strings.NewReader(string(blockchain.FlytrapABI)))
		if err != nil {
			log.Fatal(err)
		}
		events := []blockchain.Event{}
		for _, vLog := range logs {
			event := blockchain.Event{}
			err = contractAbi.Unpack(&event, "ACLChange", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			event.From = common.BytesToAddress(vLog.Topics[1].Bytes())
			event.To = common.BytesToAddress(vLog.Topics[2].Bytes())
			event.Action = blockchain.Action(vLog.Topics[3].Big().Int64())
			events = append(events, event)
		}
		for _, event := range events {
			log.Printf("Timestamp: %s", time.Unix(event.Timestamp.Int64(), 0))
			log.Printf("Topic affected: %s", event.Name)
			log.Printf("Person initiating request: %s", event.From.Hex())
			log.Printf("Person affected: %s", event.To.Hex())
			log.Printf("Action performed: %s\n\n", event.Action)
		}
	}
}
