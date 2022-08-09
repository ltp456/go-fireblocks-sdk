package types

type GasStation struct {
	Balance       Balance       `json:"balance"`
	Configuration Configuration `json:"configuration"`
}

type Balance struct {
}

type Configuration struct {
	GasThreshold string      `json:"gasThreshold"`
	GasCap       string      `json:"gasCap"`
	MaxGasPrice  interface{} `json:"maxGasPrice"`
}
