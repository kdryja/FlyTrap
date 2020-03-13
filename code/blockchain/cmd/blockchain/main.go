package main

import (
	"flag"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kdryja/thesis/code/blockchain"
)

var (
	contract    = flag.String("contract", blockchain.FLYTRAP_CONTRACT, "Specify address of your contract")
	newTopic    = flag.String("new_topic", "", "Reason to create a new topic")
	addPub      = flag.String("pub", "", "Reason to add as publisher")
	addSub      = flag.String("sub", "", "Reason to add as subscriber")
	revokePub   = flag.String("rpub", "", "Reason to revoke as publisher")
	revokeSub   = flag.String("rsub", "", "Reason to revoke as subscriber")
	addContract = flag.Bool("new", false, "Whether to create new contract for flytrap")
	topicName   = flag.String("topic", "MyTopic1", "name of the topic to modify")
	client      = flag.String("client", os.Getenv("BLOCKCHAIN_CLIENT"), "Public key of client to manipulate")
	logsRead    = flag.Bool("logs", false, "Receive relevant event logs. Providing -client and -topic flags will restrict the search only to relevant fields.")
	timestamp   = flag.String("date", "", "Provide date that should be searched for. E.g. 2020-02-29")
	verbose     = flag.Bool("v", false, "Enable verbose mode")
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
	if *newTopic != "" {
		b.Opts.Value = big.NewInt(1000)
		if _, err := b.Instance.AddTopic(b.Opts, *topicName, [2]byte{'G', 'B'}, big.NewInt(0), big.NewInt(0), *newTopic, true); err != nil {
			log.Fatal(err)
		}
	}
	if *addPub != "" {
		if err := b.AddPub(*client, *topicName, *addPub); err != nil {
			log.Fatal(err)
		}
	}
	if *addSub != "" {
		if err := b.AddSub(*client, *topicName, *addSub); err != nil {
			log.Fatal(err)
		}
	}
	if *revokeSub != "" {
		if _, err := b.Instance.RevokeSub(b.Opts, common.HexToAddress(*client), *topicName, *revokeSub); err != nil {
			log.Fatal(err)
		}
	}
	if *revokePub != "" {
		if _, err := b.Instance.RevokePub(b.Opts, common.HexToAddress(*client), *topicName, *revokePub); err != nil {
			log.Fatal(err)
		}
	}
	if *logsRead {
		events := b.AllLogs(nil, nil, *contract)
		var parsedTime *time.Time
		if *timestamp != "" {
			tmpTime, err := time.Parse("2006-01-02", *timestamp)
			if err != nil {
				log.Fatal(err)
			}
			parsedTime = &tmpTime
			pubs, subs := blockchain.TopicLogsAt(*parsedTime, *topicName, events)
			log.Printf("At %q and at topic %q following clients had subscribe access:\n%v\nand following topics had publish access:\n%v\n\n", *timestamp, *topicName, subs, pubs)
		}
		if *client != "" {
			parsedTime, err := time.Parse("2006-01-02", *timestamp)
			if err != nil {
				log.Fatal(err)
			}
			pubs, subs := blockchain.PersonLogsAt(parsedTime, *client, events)
			log.Printf("At %q and for %q had subscribe access to:\n%v\nand publish access to:\n%v\n\n", *timestamp, *client, subs, pubs)
		}
		if *verbose {
			for _, event := range events {
				if parsedTime != nil && parsedTime.Before(time.Unix(event.Timestamp.Int64(), 0)) {
					break
				}
				log.Printf("Timestamp: %s", time.Unix(event.Timestamp.Int64(), 0))
				log.Printf("Topic affected: %s", event.Name)
				log.Printf("Person initiating request: %s", event.From.Hex())
				log.Printf("Person affected: %s", event.To.Hex())
				log.Printf("Action performed: %s", event.Action)
				log.Printf("Provided reason: %q\n\n", event.Reason)
			}
		}
	}
}
