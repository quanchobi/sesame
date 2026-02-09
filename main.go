package main

import (
	"log"
	"net/http"
)

const port string = "8080"
const root string = "/app"

// Handler function for the readiness endpoint /healthz
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(root+"/", http.StripPrefix(root, http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz/", handlerReadiness)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
