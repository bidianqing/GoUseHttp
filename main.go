package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Add("Content-Type", "application/json")
	fmt.Fprint(response, "{\"name\":\"tom\"}")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	http.ListenAndServe(":8080", router)
}
