package flytrap

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kdryja/thesis/code/flytrap/blockchain"
)

type Auth struct {
	tokenMap *sync.Map
}

const (
	TOKEN_LEN = 32
	PROP      = "flytrap"
)

func generateToken() (string, error) {
	b := make([]byte, TOKEN_LEN)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
func (a *Auth) issue(c net.Conn) {
	defer c.Close()
	log.Print("New token request")
	tok, err := func() (string, error) {
		t := make([]byte, 256)
		n, err := c.Read(t)
		if err != nil {
			return "", err
		}
		t = t[:n]

		// Attempting to split the payload into signature and compressed public key
		received := bytes.SplitN(t, []byte{0x00, 0x00}, 2)
		sig := received[0]
		pubBytes, err := crypto.DecompressPubkey(received[1])
		if err != nil {
			return "", err
		}
		sigPub := crypto.PubkeyToAddress(*pubBytes)
		pub, err := crypto.SigToPub(sigPub.Hash().Bytes(), sig)
		if err != nil {
			return "", err
		}
		if crypto.PubkeyToAddress(*pub) != sigPub {
			return "", fmt.Errorf("attempted forged signature submission")
		}
		tok, err := generateToken()
		if err != nil {
			return "", err
		}
		log.Printf("Provided pubkey: %s", sigPub.String())
		return tok, blockchain.RegisterToken(tok, sigPub)
	}()
	if err != nil {
		log.Printf("Registration of new token has failed with: %q", err)
		c.Write([]byte("FAIL"))
		return
	}
	log.Printf("Generated token: %s", tok)
	c.Write([]byte(tok))
	log.Print("Token request closed gracefully")
}

// Verify method checks whether provided token exists in cache.
func (a *Auth) Verify(data []byte, addr string) bool {
	sep := []byte(PROP)
	// Locate flytrap specific property in the payload.
	i := bytes.Index(data, sep)
	if i == -1 {
		log.Print("token not provided!")
		return false
	}
	// Location of the secret key / value property.
	loc := i + len(sep)
	keyLen := binary.BigEndian.Uint16(data[loc : loc+2])
	token := string(data[loc+2 : loc+2+int(keyLen)])
	authTok, ok := a.tokenMap.Load(addr)
	if !ok {
		log.Print(ok)
		return false
	}
	return authTok == token
}

// Server starts a new auth server
func (a *Auth) Server(l net.Listener) {
	log.Printf("Now accepting auth connections under %s", l.Addr())
	a.tokenMap = &sync.Map{}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Print(err)
		}
		go a.issue(c)
	}
}
