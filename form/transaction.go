package form

type OneTimeAddress struct {
	Address string `json:"address"`
	tag     string
}

type Source struct {
	Type PeerTypes `json:"type,omitempty"`
	ID   string    `json:"id,omitempty"`
}

type Destination struct {
	Type           PeerTypes      `json:"type,omitempty"`
	ID             string         `json:"id,omitempty"`
	OneTimeAddress OneTimeAddress `json:"oneTimeAddress"`
}
