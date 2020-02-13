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

type Action int

const (
	ADDRESS          = "http://localhost:7545"
	SERVER_PRIVKEY   = "privkey.asc"
	FLYTRAP_CONTRACT = "0xAE447730C5296f30a7ba64276F4741104bE228a3"
)

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
	Action    *big.Int
}

// ActionMap shows a map of individual int value to textual description of performed action
var ActionMap = map[int64]string{
	0: "AddTopic",
	1: "AddPub",
	2: "AddSub",
	3: "RevokePub",
	4: "RevokeSub",
}

// VerifyAccess function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func VerifyAccess(topic string, key common.Address, pub bool) (bool, error) {
	b, err := New()
	if err != nil {
		return false, err
	}
	if err := b.SetInstance(common.HexToAddress(FLYTRAP_CONTRACT)); err != nil {

		return false, err
	}
	t := [32]byte{}
	copy(t[:], topic)
	if pub {
		return b.Instance.VerifyPub(nil, key, t)
	}
	return b.Instance.VerifySub(nil, key, t)
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
