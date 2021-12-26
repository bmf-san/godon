package godon

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func Serve() {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8081"
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}

	s := http.Server{
		Addr:    ":8080",
		Handler: rp,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
