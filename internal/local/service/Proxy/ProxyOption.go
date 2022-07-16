package Proxy

type ProxyOption struct {
	Local      string
	Remote     string
	Controller string
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		Local:      ":7887",
		Remote:     ":20021",
		Controller: ":8082",
	}
}
