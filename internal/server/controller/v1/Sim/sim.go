package Sim

type Sim struct {
	proxyAddress string
}

func NewSim(addr string) *Sim {
	return &Sim{proxyAddress: addr}
}
