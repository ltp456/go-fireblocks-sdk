package types

type TxConfirm struct {
	Success      bool     `json:"success"`
	Transactions []string `json:"transactions"`
}

type OperationStatus struct {
	Success        bool   `json:"success"`
	TxStatus       string `json:"txStatus"`
	TxID           string `json:"txId"`
	ReplacedTxHash string `json:"replacedTxHash"`
}

type AddressStatus struct {
	IsValid     bool `json:"isValid"`
	IsActive    bool `json:"isActive"`
	RequiresTag bool `json:"requiresTag"`
}

type EstimateFee struct {
	Low    Fee `json:"low,omitempty"`
	Medium Fee `json:"medium,omitempty"`
	High   Fee `json:"high,omitempty"`
}

type Fee struct {
	FeePerByte  string `json:"feePerByte,omitempty"`
	GasPrice    string `json:"gasPrice,omitempty"`
	GasLimit    string `json:"gasLimit,omitempty"`
	NetworkFee  string `json:"networkFee,omitempty"`
	BaseFee     string `json:"baseFee,omitempty"`
	PriorityFee string `json:"priorityFee,omitempty"`
}

type Status struct {
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type Transaction struct {
	ID                            string             `json:"id,omitempty,omitempty"`
	AssetID                       string             `json:"assetId,omitempty"`
	Source                        Source             `json:"source,omitempty"`
	Destination                   Destination        `json:"destination,omitempty"`
	AmountInfo                    AmountInfo         `json:"amountInfo,omitempty"`
	TreatAsGrossAmount            bool               `json:"treatAsGrossAmount,omitempty"`
	FeeInfo                       FeeInfo            `json:"feeInfo,omitempty"`
	RequestedAmount               float64            `json:"requestedAmount,omitempty"`
	Amount                        float64            `json:"amount,omitempty"`
	NetAmount                     float64            `json:"netAmount,omitempty"`
	AmountUSD                     float64            `json:"amountUSD,omitempty"`
	ServiceFee                    float64            `json:"serviceFee,omitempty"`
	NetworkFee                    float64            `json:"networkFee,omitempty"`
	CreatedAt                     int                `json:"createdAt,omitempty"`
	LastUpdated                   int                `json:"lastUpdated,omitempty"`
	Status                        string             `json:"status,omitempty"`
	TxHash                        string             `json:"txHash,omitempty"`
	Index                         int                `json:"index,omitempty"`
	Tag                           string             `json:"tag,omitempty"`
	SubStatus                     string             `json:"subStatus,omitempty"`
	SourceAddress                 string             `json:"sourceAddress,omitempty"`
	DestinationAddress            string             `json:"destinationAddress,omitempty"`
	DestinationAddressDescription string             `json:"destinationAddressDescription,omitempty"`
	DestinationTag                string             `json:"destinationTag,omitempty"`
	SignedBy                      []string           `json:"signedBy,omitempty"`
	CreatedBy                     string             `json:"createdBy,omitempty"`
	RejectedBy                    string             `json:"rejectedBy,omitempty"`
	AddressType                   string             `json:"addressType,omitempty"`
	Note                          string             `json:"note,omitempty"`
	ExchangeTxID                  string             `json:"exchangeTxId,omitempty"`
	FeeCurrency                   string             `json:"feeCurrency,omitempty"`
	Operation                     string             `json:"operation,omitempty"`
	NetworkRecords                []NetworkRecords   `json:"networkRecords,omitempty"`
	AmlScreeningResult            AmlScreeningResult `json:"amlScreeningResult,omitempty"`
	CustomerRefID                 string             `json:"customerRefId,omitempty"`
	NumOfConfirmations            int                `json:"numOfConfirmations,omitempty"`
	SignedMessages                []SignedMessages   `json:"signedMessages,omitempty"`
	ReplacedTxHash                string             `json:"replacedTxHash,omitempty"`
	ExternalTxID                  string             `json:"externalTxId,omitempty"`
	BlockInfo                     BlockInfo          `json:"blockInfo,omitempty"`
	AuthorizationInfo             AuthorizationInfo  `json:"authorizationInfo,omitempty"`
	ExtraParameters               ExtraParameters    `json:"extraParameters,omitempty"`
}

type Source struct {
	Type    string `json:"type,omitempty"`
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	SubType string `json:"subType,omitempty"`
}

type Destination struct {
	Type    string `json:"type,omitempty"`
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	SubType string `json:"subType,omitempty"`
}

type AmountInfo struct {
	Amount          string `json:"amount,omitempty"`
	RequestedAmount string `json:"requestedAmount,omitempty"`
	NetAmount       string `json:"netAmount,omitempty"`
	AmountUSD       string `json:"amountUSD,omitempty"`
}

type FeeInfo struct {
	NetworkFee string `json:"networkFee,omitempty"`
	ServiceFee string `json:"serviceFee,omitempty"`
}

type NetworkRecords struct {
	Source             Source      `json:"source,omitempty"`
	Destination        Destination `json:"destination,omitempty"`
	TxHash             string      `json:"txHash,omitempty"`
	NetworkFee         string      `json:"networkFee,omitempty"`
	AssetID            string      `json:"assetId,omitempty"`
	NetAmount          string      `json:"netAmount,omitempty"`
	Status             string      `json:"status,omitempty"`
	Type               string      `json:"type,omitempty"`
	DestinationAddress string      `json:"destinationAddress,omitempty"`
	SourceAddress      string      `json:"sourceAddress,omitempty"`
}

type Payload struct {
}

type AmlScreeningResult struct {
	Provider string  `json:"provider,omitempty"`
	Payload  Payload `json:"payload,omitempty"`
}

type Signature struct {
	R string `json:"r,omitempty"`
	S string `json:"s,omitempty"`
	V int    `json:"v,omitempty"`
}

type SignedMessages struct {
	Content        string    `json:"content,omitempty"`
	Algorithm      string    `json:"algorithm,omitempty"`
	DerivationPath []int     `json:"derivationPath,omitempty"`
	Signature      Signature `json:"signature,omitempty"`
	PublicKey      string    `json:"publicKey,omitempty"`
}

type BlockInfo struct {
	BlockHeight string `json:"blockHeight,omitempty"`
	BlockHash   string `json:"blockHash,omitempty"`
}

type Users struct {
	UsedID1 string `json:"usedId1,omitempty"`
	UsedID2 string `json:"usedId2,omitempty"`
	UsedID3 string `json:"usedId3,omitempty"`
}

type Groups struct {
	Th    int   `json:"th,omitempty"`
	Users Users `json:"users,omitempty"`
}

type AuthorizationInfo struct {
	AllowOperatorAsAuthorizer bool     `json:"allowOperatorAsAuthorizer,omitempty"`
	Logic                     string   `json:"logic,omitempty"`
	Groups                    []Groups `json:"groups,omitempty"`
}

type ExtraParameters struct {
}
