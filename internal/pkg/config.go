package pkg

type CertKey struct {
	CertFile string `json:"cert-file" mapstructure:"cert-file"`
	KeyFile  string `json:"cert-key" mapstructure:"cert-key"`
	CaFile   string `json:"ca-file" mapstructure:"ca-file"`
}
