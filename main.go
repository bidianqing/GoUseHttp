package main

import (
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		response.Write([]byte(`{"name":"http"}`))
	})

	http.ListenAndServe(":8080", server)
}
