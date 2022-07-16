package service

type ProxyOption struct {
	Local  string
	Remote string
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		Remote: ":20021",
		Local:  ":8081",
	}
}
