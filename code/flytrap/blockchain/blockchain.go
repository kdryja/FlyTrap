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
	FLYTRAP_CONTRACT = "0xD1d784F247a60b119d555fe1eC913886D28A9eaE"
)

const (
	AddTopic = iota
	AddPub
	AddSub
	RevokePub
	RevokeSub
	WrongCountry
	FiveFailures
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
	Name      [32]byte
	From      common.Address
	To        common.Address
	Timestamp *big.Int
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
		"FiveFailures",
	}
	return actions[a]
}

// VerifyAccess function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func VerifyAccess(topic string, key common.Address, country [2]byte, pub bool) (bool, error) {
	b, err := New()
	if err != nil {
		return false, err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {
		return false, err
	}
	t := [32]byte{}
	copy(t[:], topic)
	var ok bool
	var storedCountry [2]byte
	if pub {
		ok, storedCountry, err = b.Instance.VerifyPub(nil, key, t)
	}
	ok, storedCountry, err = b.Instance.VerifyPub(nil, key, t)
	if country != storedCountry {
		copy(t[:], fmt.Sprintf("want: %s, got %s", storedCountry, country))
		if _, err := b.Instance.LogAlert(b.Opts, key, uint8(WrongCountry), t); err != nil {
			return false, err
		}
	}
	return ok && country == storedCountry, err
}

func PersistentLog(reason Action, key common.Address, topic string) error {
	b, err := New()
	if err != nil {
		return err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {
		return err
	}
	t := [32]byte{}
	copy(t[:], topic)
	if _, err := b.Instance.LogAlert(b.Opts, key, uint8(reason), t); err != nil {
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(800000) // in units
	auth.GasPrice = gasPrice
	return &Blockchain{Opts: auth, Client: client}, err
}

func (b *Blockchain) SetInstance(contract common.Address) error {
	inst, err := NewFlytrap(contract, b.Client)
	b.Instance = inst
	return err
}

func (b *Blockchain) AddPub(key, topic string) error {
	t := [32]byte{}
	copy(t[:], topic)
	_, err := b.Instance.AddPub(b.Opts, common.HexToAddress(key), t)
	return err
}

func (b *Blockchain) AddSub(key, topic string) error {
	t := [32]byte{}
	copy(t[:], topic)
	_, err := b.Instance.AddSub(b.Opts, common.HexToAddress(key), t)
	return err
}

func (b *Blockchain) GenerateContract(cost int) (common.Address, error) {
	addr, _, _, err := DeployFlytrap(b.Opts, b.Client, big.NewInt(int64(cost)))
	return addr, err
}
