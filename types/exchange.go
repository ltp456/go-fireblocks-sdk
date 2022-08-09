package types

type ExchangeAccount struct {
	ID                  string            `json:"id"`
	Type                string            `json:"type"`
	Name                string            `json:"name"`
	Status              string            `json:"status"`
	Assets              []Assets          `json:"assets"`
	TradingAccounts     []TradingAccounts `json:"tradingAccounts"`
	IsSubaccount        bool              `json:"isSubaccount"`
	MainAccountID       string            `json:"mainAccountId"`
	FundableAccountType string            `json:"fundableAccountType"`
}

type TradingAccounts struct {
	Type   string   `json:"type"`
	Assets []Assets `json:"assets"`
}
