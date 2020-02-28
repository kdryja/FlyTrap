package flytrap

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

const IP_DATABASE = "GeoLite2-Country.mmdb"

func country(address string) [2]byte {
	// For localhost testing, return GB
	if address == "127.0.0.1" {
		return [2]byte{'G', 'B'}
	}
	db, err := geoip2.Open(IP_DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(address)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	out := [2]byte{}
	copy(out[:], record.Country.IsoCode)
	return out
}
