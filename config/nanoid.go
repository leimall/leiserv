package config

type NanoID struct {
	ASCII  string `mapstructure:"ascii" json:"ascii" yaml:"ascii"`    // 是否使用自定义字符集
	Length int    `mapstructure:"length" json:"length" yaml:"length"` // 自定义字符集
}
