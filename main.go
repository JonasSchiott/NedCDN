package main

import "net/http"

func main() {
	http.Handle("/cdn/", http.StripPrefix("/cdn/", http.FileServer(http.Dir("./cdn"))))
	http.ListenAndServe(":8000", nil)
}
