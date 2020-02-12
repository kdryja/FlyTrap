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
	ADDRESS             = "http://localhost:7545"
	SERVER_PRIVKEY      = "privkey.asc"
	FLYTRAP_CONTRACT    = "0xa53cEaD0ABC76B24F8966E873bF3bA5dba201735"
	AUTHORIZER_CONTRACT = "0xe260dE168679c271166E2D6081645cBe66Ba1511"
)

type Blockchain struct {
	Instance *Flytrap
	Opts     *bind.TransactOpts
	Client   *ethclient.Client
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

func RegisterToken(token string, address common.Address) error {
	b, err := New()
	if err != nil {
		return err
	}
	t := [32]byte{}
	copy(t[:], token)
	inst, err := NewAuthorizer(common.HexToAddress(AUTHORIZER_CONTRACT), b.Client)
	if err != nil {
		return err
	}
	_, err = inst.RegisterToken(b.Opts, t, address)
	return err
}

func RetrievePubkey(token string) (common.Address, error) {
	b, err := New()
	if err != nil {
		return common.Address{}, err
	}
	t := [32]byte{}
	copy(t[:], token)
	inst, err := NewAuthorizer(common.HexToAddress(AUTHORIZER_CONTRACT), b.Client)
	if err != nil {
		return common.Address{}, err
	}
	return inst.RetrieveKey(nil, t)
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
