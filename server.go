package godon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Config struct {
	Proxy    Proxy     `json:"proxy"`
	Backends []Backend `json:"backends"`
}

type Proxy struct {
	Port string `json:"port"`
}

type Backend struct {
	URL string `json:"url"`
}

func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var cfg Config
	json.Unmarshal(data, &cfg)

	var idx int = 0
	maxLen := len(cfg.Backends)
	var mu sync.Mutex
	director := func(req *http.Request) {
		// Round Robin
		mu.Lock()
		backend := cfg.Backends[idx]
		var targetURL *url.URL
		targetURL, err = url.Parse(backend.URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		req.URL = targetURL
		idx++
		if idx == maxLen {
			idx = 0
		}
		mu.Unlock()
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
