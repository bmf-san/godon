package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bmf-san/godon"
)

func serveBackend(name string, port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Backend server name:%v\n", name)
		fmt.Fprintf(w, "Response header:%v\n", r.Header)
	}))
	http.ListenAndServe(port, mux)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	go func() {
		godon.Serve()
		wg.Done()
	}()

	go func() {
		serveBackend("web1", ":8081")
		wg.Done()
	}()

	go func() {
		serveBackend("web2", ":8082")
		wg.Done()
	}()

	go func() {
		serveBackend("web3", ":8083")
		wg.Done()
	}()

	wg.Wait()
}
