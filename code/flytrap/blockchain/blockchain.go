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
	ADDRESS        = "http://localhost:7545"
	SERVER_PRIVKEY = "privkey.txt"
	CONTRACT       = "0xf3776E4BfEb3856E46a18d336F131BAD20904c04"
)

type Blockchain struct {
	Instance *Flytrap
	Opts     *bind.TransactOpts
	client   *ethclient.Client
}

// Verify function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func Verify(topic, key string, pub bool) (bool, error) {
	b, err := New()
	if err != nil {
		return false, err
	}
	if err := b.SetInstance(common.HexToAddress(CONTRACT)); err != nil {
		return false, err
	}
	t := [32]byte{}
	copy(t[:], topic)
	if pub {
		return b.Instance.VerifyPub(nil, common.HexToAddress(key), t)
	}
	return b.Instance.VerifySub(nil, common.HexToAddress(key), t)
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
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice
	return &Blockchain{Opts: auth, client: client}, err
}

func (b *Blockchain) SetInstance(contract common.Address) error {
	inst, err := NewFlytrap(contract, b.client)
	b.Instance = inst
	return err
}

func (b *Blockchain) AddPub(key, topic string) error {
	t := [32]byte{}
	copy(t[:], topic)
	_, err := b.Instance.AddPub(b.Opts, common.HexToAddress(key), t)
	return err
}

func (b *Blockchain) GenerateContract(cost int) (common.Address, error) {
	addr, _, _, err := DeployFlytrap(b.Opts, b.client, big.NewInt(int64(cost)))
	return addr, err
}
