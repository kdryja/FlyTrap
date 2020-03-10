package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/kdryja/thesis/code/flytrap/blockchain"
)

const (
	DEFAULT_CONTRACT = "0x1D199e5D181FC41a7B93e1c2610cFce1409186BF"
)

type Template struct {
	Contract string
}

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/logs", LogsHandler)
	http.HandleFunc("/", MainHandler)
	http.ListenAndServe(":8081", nil)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Template{Contract: DEFAULT_CONTRACT})
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	start, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	end, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
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
