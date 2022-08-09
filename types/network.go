package types

type NetworkConnections struct {
	ID              string        `json:"id"`
	Status          string        `json:"status"`
	LocalNetworkID  NetworkID     `json:"localNetworkId"`
	RemoteNetworkID NetworkID     `json:"remoteNetworkId"`
	RoutingPolicy   RoutingPolicy `json:"routingPolicy"`
}

type NetworkID struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TenantID string `json:"tenantId"`
}

type RpInfo struct {
	Scheme  string `json:"scheme"`
	DstType string `json:"dstType"`
	DstID   string `json:"dstId"`
}

type RoutingPolicy struct {
	Sen        RpInfo `json:"sen"`
	Crypto     RpInfo `json:"crypto"`
	Signet     RpInfo `json:"signet"`
	SenTest    RpInfo `json:"sen_test"`
	SignetTest RpInfo `json:"signet_test"`
}
