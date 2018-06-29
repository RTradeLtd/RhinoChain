package transaction

import "errors"

// TransactionPool is a collection of pending transactions
type TransactionPool struct {
	Transactions map[string]*Transaction `json:"transaction"`
}

// InitTransactionPool is used to initialize our in-memory transaction pool
func InitTransactionPool() *TransactionPool {
	return &TransactionPool{}
}

func (tp *TransactionPool) AddTransaction(tx *Transaction) error {
	if tx.Mined {
		return errors.New("attempting to insert already mind transaction to mempool")
	}
	tp.Transactions[tx.TxHash] = tx
	return nil
}
