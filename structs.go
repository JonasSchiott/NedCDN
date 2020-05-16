package main

// Header => head
type Header struct {
	Info     string   `json:"info"`
	Statuses []Status `json:"statuses"`
}

//Status => status
type Status struct {
	Name string `json:"name"`
	Info string `json:"info"`
	Code int    `json:"code"`
}

// Shard => shard
type Shard struct {
	ID             int    `json:"id"`
	Guilds         int    `json:"guilds"`
	ExpectedGuilds int    `json:"expected_guilds"`
	StatusText     string `json:"status_text"`
	StatusCode     int    `json:"status_code"`
}
