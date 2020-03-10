package blockchain

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (b *Blockchain) AllLogs(start, end *time.Time, contract string) []Event {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(1000000),
		Addresses: []common.Address{
			common.HexToAddress(contract),
		},
	}
	logs, err := b.Client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(FlytrapABI)))
	if err != nil {
		log.Fatal(err)
	}
	events := []Event{}
	for _, vLog := range logs {
		event := Event{}
		err = contractAbi.Unpack(&event, "ACLChange", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		if start != nil && start.After(time.Unix(event.Timestamp.Int64(), 0)) {
			continue
		}
		if end != nil && end.Before(time.Unix(event.Timestamp.Int64(), 0)) {
			break
		}
		event.From = common.BytesToAddress(vLog.Topics[1].Bytes())
		event.To = common.BytesToAddress(vLog.Topics[2].Bytes())
		event.Action = Action(vLog.Topics[3].Big().Int64())
		events = append(events, event)
	}
	return events
}

func TopicLogsAt(t time.Time, topic string, events []Event) ([]string, []string) {
	pubs := make(map[string]bool)
	subs := make(map[string]bool)
	for _, event := range events {
		if time.Unix(event.Timestamp.Int64(), 0).After(t) {
			break
		}
		if event.Name != topic {
			continue
		}
		switch event.Action {
		case AddPub:
			pubs[event.To.String()] = true
		case AddSub:
			subs[event.To.String()] = true
		case RevokePub:
			delete(pubs, event.To.String())
		case RevokeSub:
			delete(subs, event.To.String())
		}
	}
	pubSlice := []string{}
	subSlice := []string{}
	for k, v := range pubs {
		if v {
			pubSlice = append(pubSlice, k)
		}
	}
	for k, v := range subs {
		if v {
			subSlice = append(subSlice, k)
		}
	}
	return pubSlice, subSlice
}

func PersonLogsAt(t time.Time, pubkey string, events []Event) ([]string, []string) {
	pubs := make(map[string]bool)
	subs := make(map[string]bool)
	for _, event := range events {
		if time.Unix(event.Timestamp.Int64(), 0).After(t) {
			break
		}
		if event.To.String() != pubkey {
			continue
		}
		switch event.Action {
		case AddPub:
			pubs[event.Name] = true
		case AddSub:
			subs[event.Name] = true
		case RevokePub:
			delete(pubs, event.Name)
		case RevokeSub:
			delete(subs, event.Name)
		}
	}
	pubSlice := []string{}
	subSlice := []string{}
	for k, v := range pubs {
		if v {
			pubSlice = append(pubSlice, k)
		}
	}
	for k, v := range subs {
		if v {
			subSlice = append(subSlice, k)
		}
	}
	return pubSlice, subSlice
}
