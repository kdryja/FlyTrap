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
	CLIENT         = "0x641c46D43A3c552a76318c12D8Fc2839b913F32F"
	SERVER_PRIVKEY = "privkey.txt"
	CONTRACT       = "0xB8F3CBdd5F161B59b7573e9742791d97CEC414b3"
)

type Blockchain struct {
	instance *Flytrap
	opts     *bind.TransactOpts
	client   *ethclient.Client
}

// Verify function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func Verify(topic, key string, pub bool) (bool, error) {
	b, err := new()
	if err != nil {
		return false, err
	}
	if err := b.setInstance(); err != nil {
		return false, err
	}
	t := [32]byte{}
	copy(t[:], topic)
	if pub {
		return b.instance.Pub(nil, common.HexToAddress(key), t)
	}
	return b.instance.Sub(nil, common.HexToAddress(key), t)
}

func new() (*Blockchain, error) {
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
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return &Blockchain{opts: auth, client: client}, err
}

func (b *Blockchain) setInstance() error {
	address := common.HexToAddress(CONTRACT)
	inst, err := NewFlytrap(address, b.client)
	b.instance = inst
	return err
}

func (b *Blockchain) addPub(key, topic string) error {
	t := [32]byte{}
	copy(t[:], topic)
	_, err := b.instance.AddPub(b.opts, common.HexToAddress(key), t)
	return err
}

func (b *Blockchain) generateContract() error {
	_, _, _, err := DeployFlytrap(b.opts, b.client)
	return err
}
