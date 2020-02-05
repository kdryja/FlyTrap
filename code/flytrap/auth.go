package flytrap

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"log"
	"net"
	"strings"
	"sync"
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
	t := make([]byte, 256)
	n, err := c.Read(t)
	if err != nil {
		log.Print(err)
		return
	}
	t = t[:n]
	// TODO(kdryja) Verify provided password here
	if string(t) != "secret_password" {
		log.Print("Failed auth!")
		return
	}
	addr := c.RemoteAddr().String()
	addr = addr[:strings.LastIndex(addr, ":")]
	tok, err := generateToken()
	if err != nil {
		log.Print(err)
		return
	}
	a.tokenMap.Store(addr, tok)
	c.Write([]byte(tok))
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
	log.Printf("Token has been found: %s", token)
	authTok, ok := a.tokenMap.Load(addr)
	if !ok {
		log.Print(ok)
		return false
	}
	log.Print(authTok)
	log.Print(token)
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
