package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bmf-san/godon"
)

func serveOrigin(name string, port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, name)
	}))
	http.ListenAndServe(port, mux)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	go func() {
		godon.Serve()
	}()

	go func() {
		serveOrigin("web1", ":8081")
		wg.Done()
	}()

	go func() {
		serveOrigin("web2", ":8082")
		wg.Done()
	}()

	go func() {
		serveOrigin("web3", ":8083")
		wg.Done()
	}()

	wg.Wait()
}
