package main

import (
	_ "net"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		response.Write([]byte(`{"name":"http"}`))
	})

	http.ListenAndServe(":8081", handler)

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: handler,
	// }

	// l, _ := net.Listen("tcp", ":8080")
	// server.Serve(l)
}
