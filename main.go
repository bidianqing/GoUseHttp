package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Add("Content-Type", "application/json")
	response.Write([]byte(`{"name":"tom"}`))
}

func main() {
	// router := httprouter.New()
	// router.GET("/", Index)

	http.Handle("/", http.FileServer(http.Dir("./wwwroot")))
	http.HandleFunc("/http", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("Content-Type", "application/json")
		response.Write([]byte(`{"name":"http"}`))
	})
	http.ListenAndServe(":8080", nil)
}
