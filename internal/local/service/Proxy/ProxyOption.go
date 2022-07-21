package Proxy

type ProxyOption struct {
	Local      string
	Remote     string
	Controller string
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		Local:      ":7887",
		Remote:     "120.24.250.251:20021",
		Controller: "120.24.250.251:8082",
	}
}
