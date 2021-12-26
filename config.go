package godon

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
