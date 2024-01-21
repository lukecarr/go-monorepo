package main

import (
	"fmt"
	"log"
	"net/http"

	"x/osx"

	"github.com/julienschmidt/httprouter"
)

func HelloWorld(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	router := httprouter.New()

	router.GET("/", HelloWorld)

	addr := osx.Getenv("ADDR", ":3000")
	log.Fatal(http.ListenAndServe(addr, router))
}
