package godon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var cfg Config
	json.Unmarshal(data, &cfg)

	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8081"
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	s := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: rp,
	}
	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
