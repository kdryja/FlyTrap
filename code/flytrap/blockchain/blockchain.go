package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ADDRESS          = "http://localhost:7545"
	SERVER_PRIVKEY   = "privkey.asc"
	FLYTRAP_CONTRACT = "0x7A84fE8FC721Dd0550F3f625d190399f7F325722"
)

const (
	AddTopic = iota
	AddPub
	AddSub
	RevokePub
	RevokeSub
	WrongCountry
	Banned
	Summary
)

type Action int

// Blockchain is a struct containing runtime parameters used for communication with blockchain
type Blockchain struct {
	Instance *Flytrap
	Opts     *bind.TransactOpts
	Client   *ethclient.Client
}

// Event is a struct containing data about event emitted on blockchain
type Event struct {
	Name, Reason string
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Action
}

func (a Action) String() string {
	actions := [...]string{
		"AddTopic",
		"AddPub",
		"AddSub",
		"RevokePub",
		"RevokeSub",
		"WrongCountry",
		"Banned",
		"Summary",
	}
	return actions[a]
}

// VerifyAccess function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func VerifyAccess(topic, ip string, key common.Address, country [2]byte, pub bool) (bool, bool, error) {
	b, err := New()
	if err != nil {
		return false, false, err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {
		return false, false, err
	}
	var ok, sensitive bool
	var storedCountry [2]byte
	if pub {
		ok, storedCountry, sensitive, err = b.Instance.VerifyPub(nil, key, topic)
	} else {
		ok, storedCountry, sensitive, err = b.Instance.VerifySub(nil, key, topic)
	}
	if country != storedCountry {
		if _, err := b.Instance.LogAlert(b.Opts, key, uint8(WrongCountry), topic, fmt.Sprintf("want: %s, got %s. ip: %q", storedCountry, country, ip)); err != nil {
			return false, false, err
		}
	}
	return ok && country == storedCountry, sensitive, err
}

func SummaryLog(summary string) error {
	b, err := New()
	if err != nil {
		return err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {
		return err
	}
	if _, err := b.Instance.LogAlert(b.Opts, common.HexToAddress(FLYTRAP_CONTRACT), uint8(Summary), "na", summary); err != nil {
		return err
	}
	return nil
}

func PersistentLog(reason Action, key common.Address, topic, ip string) error {
	b, err := New()
	if err != nil {
		return err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {
		return err
	}
	if _, err := b.Instance.LogAlert(b.Opts, key, uint8(reason), topic, fmt.Sprintf("too many failed attempts from IP %q", ip)); err != nil {
		return err
	}
	return nil
}

func New() (*Blockchain, error) {
	client, err := ethclient.Dial(ADDRESS)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.LoadECDSA(SERVER_PRIVKEY)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)

	auth := bind.NewKeyedTransactor(priv)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(2500000) // in units
	auth.GasPrice = gasPrice
	return &Blockchain{Opts: auth, Client: client}, err
}

func (b *Blockchain) SetInstance(contract common.Address) error {
	inst, err := NewFlytrap(contract, b.Client)
	b.Instance = inst
	return err
}

func (b *Blockchain) AddPub(key, topic, reason string) error {
	_, err := b.Instance.AddPub(b.Opts, common.HexToAddress(key), topic, reason)
	return err
}

func (b *Blockchain) AddSub(key, topic, reason string) error {
	_, err := b.Instance.AddSub(b.Opts, common.HexToAddress(key), topic, reason)
	return err
}

func (b *Blockchain) GenerateContract(cost int) (common.Address, error) {
	addr, _, _, err := DeployFlytrap(b.Opts, b.Client, big.NewInt(int64(cost)))
	return addr, err
}
