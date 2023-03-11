package gpt

type GlobalConfig struct {
	Host                    string
	Proxy                   string
	DefaultConversionConfig ConversionConfig
}

type ConversionConfig struct {
	Model       string
	Temperature float64
	TopP        float64
	MaxTokens   int
	N           int
}
