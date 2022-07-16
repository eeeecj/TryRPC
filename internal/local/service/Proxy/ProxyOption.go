package Proxy

type ProxyOption struct {
	Local  string
	Remote string
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		Local:  ":7887",
		Remote: ":20021",
	}
}
