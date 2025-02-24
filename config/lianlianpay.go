package config

type LiangLiangPay struct {
	LLPubKey   string `mapstructure:"ll-pub-key" json:"ll-pub-key" yaml:"ll-pub-key"`
	PrivateKey string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
	PublickKey string `mapstructure:"public-key" json:"public-key" yaml:"public-key"`
	MerchantID string `mapstructure:"merchant-id" json:"merchant-id" yaml:"merchant-id"`
	BaseUrl    string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
}
