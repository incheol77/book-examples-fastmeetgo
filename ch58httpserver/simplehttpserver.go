package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello world! I am simple http server"))
	}

	http.ListenAndServe(":80", nil)
}
