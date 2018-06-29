package transaction

// Transaction is an object representing a transaction on RhinoChain
type Transaction struct {
	TxHash string `json:"tx_hash"`
	To     string `json:"to"`
	From   string `json:"from"`
	Value  int64  `json:"value"`
	Mined  bool   `json:"mined"`
}
