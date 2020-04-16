package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kdryja/thesis/code/blockchain"
)

const (
	PORT = ":8081"
)

var (
	DEFAULT_CONTRACT = os.Getenv("FLYTRAP_CONTRACT")
)

type Template struct {
	Contract string
}

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/detailed", DetailedHandler)
	http.HandleFunc("/logs", LogsHandler)
	http.HandleFunc("/", MainHandler)

	log.Println("Server is now running under " + PORT)
	http.ListenAndServe(PORT, nil)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Template{Contract: DEFAULT_CONTRACT})
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "invalid payload for start", http.StatusBadRequest)
		return
	}
	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "invalid payload for end", http.StatusBadRequest)
		return
	}

	b, err := blockchain.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs := b.AllLogs(&start, &end, r.FormValue("address"))

	js, err := json.Marshal(logs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DetailedHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "invalid payload for start", http.StatusBadRequest)
		return
	}
	start = start.Add(time.Hour * 24)
	b, err := blockchain.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logs := b.AllLogs(nil, &start, r.FormValue("address"))

	var js []byte
	if topic := r.FormValue("topic"); topic != "" {
		pubs, subs := blockchain.TopicLogsAt(start, topic, logs)
		js, err = json.Marshal(struct {
			Req  string   `json:"req"`
			Pubs []string `json:"pubs"`
			Subs []string `json:"subs"`
		}{
			topic,
			pubs,
			subs,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if person := r.FormValue("person"); person != "" {
		pubs, subs := blockchain.PersonLogsAt(start, person, logs)
		js, err = json.Marshal(struct {
			Req  string   `json:"req"`
			Pubs []string `json:"pubs"`
			Subs []string `json:"subs"`
		}{
			person,
			pubs,
			subs,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
