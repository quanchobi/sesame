package main

import (
	"fmt"
	"net/http"
)

const port string = "8080"
const root string = "/"

func main() {
	mux := http.NewServeMux()
	mux.Handle(root, http.FileServer(http.Dir(".")))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
