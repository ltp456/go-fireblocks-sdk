package types

type StatusResp struct {
	Success bool `json:"success"`
}

type ReAccountName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PublicKeyInfo struct {
	Status         int           `json:"status"`
	Algorithm      string        `json:"algorithm"`
	DerivationPath []interface{} `json:"derivationPath"`
	PublicKey      string        `json:"publicKey"`
}

type UnspentInputData struct {
	Input         Input  `json:"input"`
	Address       string `json:"address"`
	Amount        string `json:"amount"`
	Confirmations string `json:"confirmations"`
	Status        string `json:"status"`
}

type Input struct {
	TxHash string `json:"txHash"`
	Index  string `json:"index"`
}

type VaultAssetResp struct {
	ID             string `json:"id"`
	Address        string `json:"address"`
	LegacyAddress  string `json:"legacyAddress"`
	Tag            string `json:"tag"`
	EosAccountName string `json:"eosAccountName"`
}

type AssetAddress struct {
	AssetID           string `json:"assetId,omitempty"`
	Address           string `json:"address,omitempty"`
	LegacyAddress     string `json:"legacyAddress,omitempty"`
	Description       string `json:"description,omitempty"`
	Tag               string `json:"tag,omitempty"`
	Type              string `json:"type,omitempty"`
	EnterpriseAddress string `json:"enterpriseAddress,omitempty"`
	Bip44AddressIndex int    `json:"bip44AddressIndex,omitempty"`
}

type CreateVaultAccount struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	HiddenOnUI    bool     `json:"hiddenOnUI,omitempty"`
	CustomerRefID string   `json:"customerRefId,omitempty"`
	AutoFuel      bool     `json:"autoFuel,omitempty"`
	Assets        []Assets `json:"assets,omitempty"`
}

type VaultAccounts struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	HiddenOnUI    bool     `json:"hiddenOnUI,omitempty"`
	CustomerRefID string   `json:"customerRefId,omitempty"`
	AutoFuel      bool     `json:"autoFuel,omitempty"`
	Assets        []Assets `json:"assets,omitempty"`
}

type CreateAddress struct {
	Address           string `json:"address,omitempty"`
	LegacyAddress     string `json:"legacyAddress,omitempty"`
	Tag               string `json:"tag,omitempty"`
	Bip44AddressIndex int    `json:"bip44AddressIndex,omitempty"`
}

type AccountAssert struct {
	ID                   string `json:"id,omitempty"`
	Total                string `json:"total,omitempty"`
	Available            string `json:"available,omitempty"`
	Pending              string `json:"pending,omitempty"`
	LockedAmount         string `json:"lockedAmount,omitempty"`
	Staked               string `json:"staked,omitempty"`
	Frozen               string `json:"frozen,omitempty"`
	TotalStakedCPU       string `json:"totalStakedCPU,omitempty"`
	TotalStakedNetwork   string `json:"totalStakedNetwork,omitempty"`
	SelfStakedCPU        string `json:"selfStakedCPU,omitempty"`
	SelfStakedNetwork    string `json:"selfStakedNetwork,omitempty"`
	PendingRefundCPU     string `json:"pendingRefundCPU,omitempty"`
	PendingRefundNetwork string `json:"pendingRefundNetwork,omitempty"`
	BlockHeight          string `json:"blockHeight,omitempty"`
	BlockHash            string `json:"blockHash,omitempty"`
}

type VaultAccountsPaged struct {
	Accounts    []Accounts `json:"accounts,omitempty"`
	Paging      Paging     `json:"paging,omitempty"`
	PreviousURL string     `json:"previousUrl,omitempty"`
	NextURL     string     `json:"nextUrl,omitempty"`
}

type AccountAssets struct {
	ID                   string `json:"id,omitempty"`
	Total                string `json:"total,omitempty"`
	Pending              string `json:"pending,omitempty"`
	LockedAmount         string `json:"lockedAmount,omitempty"`
	Available            string `json:"available,omitempty"`
	Frozen               string `json:"frozen,omitempty"`
	TotalStakedCPU       string `json:"totalStakedCPU,omitempty"`
	TotalStakedNetwork   string `json:"totalStakedNetwork,omitempty"`
	SelfStakedCPU        string `json:"selfStakedCPU,omitempty"`
	SelfStakedNetwork    string `json:"selfStakedNetwork,omitempty"`
	PendingRefundCPU     string `json:"pendingRefundCPU,omitempty"`
	PendingRefundNetwork string `json:"pendingRefundNetwork,omitempty"`
	BlockHeight          string `json:"blockHeight,omitempty"`
	BlockHash            string `json:"blockHash,omitempty"`
}

type VaultAssets struct {
	ID                   string `json:"id,omitempty"`
	Total                string `json:"total,omitempty"`
	Pending              string `json:"pending,omitempty"`
	LockedAmount         string `json:"lockedAmount,omitempty"`
	Available            string `json:"available,omitempty"`
	Frozen               string `json:"frozen,omitempty"`
	TotalStakedCPU       string `json:"totalStakedCPU,omitempty"`
	TotalStakedNetwork   string `json:"totalStakedNetwork,omitempty"`
	SelfStakedCPU        string `json:"selfStakedCPU,omitempty"`
	SelfStakedNetwork    string `json:"selfStakedNetwork,omitempty"`
	PendingRefundCPU     string `json:"pendingRefundCPU,omitempty"`
	PendingRefundNetwork string `json:"pendingRefundNetwork,omitempty"`
	BlockHeight          string `json:"blockHeight,omitempty"`
	BlockHash            string `json:"blockHash,omitempty"`
}

type Assets struct {
	ID                   string  `json:"id,omitempty"`
	Total                float64 `json:"total,omitempty"`
	Pending              float64 `json:"pending,omitempty"`
	LockedAmount         float64 `json:"lockedAmount,omitempty"`
	Available            float64 `json:"available,omitempty"`
	Frozen               float64 `json:"frozen,omitempty"`
	TotalStakedCPU       float64 `json:"totalStakedCPU,omitempty"`
	TotalStakedNetwork   float64 `json:"totalStakedNetwork,omitempty"`
	SelfStakedCPU        float64 `json:"selfStakedCPU,omitempty"`
	SelfStakedNetwork    float64 `json:"selfStakedNetwork,omitempty"`
	PendingRefundCPU     float64 `json:"pendingRefundCPU,omitempty"`
	PendingRefundNetwork float64 `json:"pendingRefundNetwork,omitempty"`
	BlockHeight          string  `json:"blockHeight,omitempty"`
	BlockHash            string  `json:"blockHash,omitempty"`
}

type Accounts struct {
	ID            string          `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	HiddenOnUI    bool            `json:"hiddenOnUI,omitempty"`
	CustomerRefID string          `json:"customerRefId,omitempty"`
	AutoFuel      bool            `json:"autoFuel,omitempty"`
	Assets        []AccountAssets `json:"assets,omitempty"`
}

type Paging struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}
