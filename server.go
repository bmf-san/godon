package godon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var cfg Config
	json.Unmarshal(data, &cfg)

	director := func(req *http.Request) {
		// Randomly select backend to load balance
		backends := cfg.Backends
		var targetURL *url.URL
		for i, b := range backends {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(3) == i {
				targetURL, err = url.Parse(b.URL)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		}
		req.URL = targetURL
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
