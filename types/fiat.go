package types

type FiatAccount struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    string       `json:"name"`
	Address string       `json:"address"`
	Assets  []FiatAssets `json:"assets"`
}

type FiatAssets struct {
	ID      string `json:"id"`
	Balance string `json:"balance"`
}
