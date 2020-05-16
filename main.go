package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.Handle("/cdn/", http.StripPrefix("/cdn/", http.FileServer(http.Dir("./cdn"))))
	http.HandleFunc("/api/v1/status", status)
	http.ListenAndServe(":8000", nil)
}

func status(w http.ResponseWriter, r *http.Request) {
	shards := []Shard{
		Shard{ID: 1, Guilds: 1000, ExpectedGuilds: 1002, StatusText: "No Issues", StatusCode: 0},
		Shard{ID: 2, Guilds: 1000, ExpectedGuilds: 1002, StatusText: "No Issues", StatusCode: 0},
		Shard{ID: 3, Guilds: 1000, ExpectedGuilds: 1002, StatusText: "No Issues", StatusCode: 0},
	}

	statuses := []Status{
		Status{
			"API",
			"No Issues",
			0,
		},
	}

	header := Header{
		Info:     "no issues",
		Statuses: statuses,
	}

	json.NewEncoder(w).Encode(struct {
		Header Header  `json:"header"`
		Shards []Shard `json:"shards"`
	}{
		Header: header,
		Shards: shards,
	})
}
