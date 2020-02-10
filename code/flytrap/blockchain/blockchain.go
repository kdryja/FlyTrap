package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
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
}

// Verify function will inspect smart contract to determine whether the provided public key can publish / subscribe to given topic.
func Verify(topic, key string, pub bool) (bool, error) {
	client, err := ethclient.Dial(ADDRESS)
	if err != nil {
		return false, err
	}
	inst, err := NewFlytrap(common.HexToAddress(CONTRACT), client)
	if err != nil {
		return false, err
	}
	t := [32]byte{}
	copy(t[:], topic)
	if pub {
		return inst.Pub(nil, common.HexToAddress(key), t)
	}
	return inst.Sub(nil, common.HexToAddress(key), t)
}

func main() {
	topic := "MyTopic"
	ok, err := Verify(topic, CLIENT, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Client %s access to %q is %t", CLIENT, topic, ok)
}

func generateContract(client *ethclient.Client) {
	privateKey, err := crypto.LoadECDSA(SERVER_PRIVKEY)
	if err != nil {
		log.Print("ERROR HERE!")
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v, %+v", nonce, gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, _, _, err = DeployFlytrap(auth, client)
	if err != nil {
		log.Fatal(err)
	}
}

func query(client *ethclient.Client) {
	address := common.HexToAddress("0xB8F3CBdd5F161B59b7573e9742791d97CEC414b3")
	instance, err := NewFlytrap(address, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(instance)
}

func adminInstance() (*Blockchain, error) {
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
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(priv)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(CONTRACT)
	instance, err := NewFlytrap(address, client)
	return &Blockchain{instance: instance, opts: auth}, err
}
func (b *Blockchain) addPub(key, topic string) error {
	t := [32]byte{}
	copy(t[:], topic)
	_, err := b.instance.AddPub(b.opts, common.HexToAddress(key), t)
	return err
}
