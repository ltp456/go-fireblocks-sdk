package types

type SupportedAssets struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	ContractAddress string `json:"contractAddress"`
}
