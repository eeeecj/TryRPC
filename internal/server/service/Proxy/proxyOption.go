package Proxy

type ProxyOption struct {
	Controller string
	Local      string
	Remote     string
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		Remote:     ":20021",
		Local:      ":8081",
		Controller: ":8082",
	}
}
