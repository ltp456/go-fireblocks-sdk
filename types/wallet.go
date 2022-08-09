package types

type Wallet struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	CustomerRefID string         `json:"customerRefId"`
	Assets        []WalletAssets `json:"assets"`
}

type WalletAssets struct {
	ID             string `json:"id"`
	Balance        string `json:"balance"`
	LockedAmount   string `json:"lockedAmount"`
	Status         string `json:"status"`
	ActivationTime string `json:"activationTime"`
	Address        string `json:"address"`
	Tag            string `json:"tag"`
}
type AssetInfo struct {
	ID           string `json:"id"`
	Balance      int    `json:"balance"`
	LockedAmount int    `json:"lockedAmount"`
	Status       string `json:"status"`
}
