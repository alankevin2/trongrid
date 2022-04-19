package api

const (
	Network_Mainnet = "Mainnet"
	Network_Shasta  = "Shasta"
)

type Network string
type Symbol string

type TronGridV1 struct {
	network Network
	apiKey  string
	baseURL string
	tokens  map[string]string
}

// Request from consumer
type GetTransactionsByAddressRequest struct {
	TRC20        bool
	Limit        string `json:"limit"`         // limit > 0
	OrderBy      string `json:"order_by"`      // "block_timestamp,asc | block_timestamp,desc (default)"
	MinTimestamp string `json:"min_timestamp"` // unix timestamp, unit: second. e.g. 1649755122000
	MaxTimestamp string `json:"max_timestamp"` // unix timestamp, unit: second. e.g. 1649755122000
}

// Response from trongrid
type TrongridTransactionsResp struct {
	Data    []TrongridTransaction `json:"data"`
	Success bool                  `json:"success"`
	Meta    map[string]int        `json:"meta"`
}

type TrongridTransaction struct {
	ID        string                 `json:"transaction_id"`
	TokenInfo map[string]interface{} `json:"token_info"`
	TimeStamp uint64                 `json:"block_timestamp"`
	From      string                 `json:"from"`
	To        string                 `json:"to"`
	Type      string                 `json:"type"`
	Value     string                 `json:"value"`
}
