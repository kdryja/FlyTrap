package flytrap

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kdryja/thesis/code/flytrap/blockchain"
)

func (c *Cache) SaveSummary(t *time.Ticker) {
	for {
		select {
		case <-t.C:
			log.Print("Saving summary for sensitive topics")
			tmpSubs := make(map[string][]string)
			tmpPubs := make(map[string][]string)
			for k, _ := range c.SubSummary {
				tmpSubs[k] = c.subSlice(k)
				delete(c.SubSummary, k)
			}
			for k, _ := range c.PubSummary {
				tmpPubs[k] = c.pubSlice(k)
				delete(c.PubSummary, k)
			}
			out := make(map[string]string)
			jsonSlice, err := json.Marshal(tmpPubs)
			if err != nil {
				log.Fatal(err)
			}
			out["pubs"] = string(jsonSlice)

			jsonSlice, err = json.Marshal(tmpSubs)
			if err != nil {
				log.Fatal(err)
			}
			out["subs"] = string(jsonSlice)

			if len(tmpSubs) == 0 && len(tmpPubs) == 0 {
				log.Print("Nothing to write")
				continue
			}

			marshalLog, err := json.Marshal(out)
			if err != nil {
				log.Fatal(err)
			}
			err = blockchain.SummaryLog(string(marshalLog))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s", marshalLog)
		}
	}
}

func (c *Cache) pubSlice(t string) []string {
	out := []string{}
	for k, v := range c.PubSummary[t] {
		if !v {
			continue
		}
		out = append(out, k)
	}
	return out
}

func (c *Cache) subSlice(t string) []string {
	out := []string{}
	for k, v := range c.SubSummary[t] {
		if !v {
			continue
		}
		out = append(out, k)
	}
	return out
}
