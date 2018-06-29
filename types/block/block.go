package block

import (
	"time"

	"github.com/RTradeLtd/RhinoChain/types/transaction"
)

// Block is a struct representing a block
type Block struct {
	ParentBlockHash string                     `json:"parent_block_hash"`
	BlockHash       string                     `json:"block_hash"`
	Transactions    []*transaction.Transaction `json:"transactions"`
	Miner           string                     `json:"miner"`
	Timestamp       *time.Time                 `json:"timestamp"`
}
