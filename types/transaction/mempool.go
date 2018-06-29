package transaction

import "errors"

// TransactionPool is a collection of pending transactions
type TransactionPool struct {
	Transactions map[string]*Transaction `json:"transaction"`
}

// InitTransactionPool is used to initialize our in-memory transaction pool
func InitTransactionPool() *TransactionPool {
	txs := make(map[string]*Transaction)
	return &TransactionPool{Transactions: txs}
}

// AddTransaction is used to add a transaction to the mempool
func (tp *TransactionPool) AddTransaction(tx *Transaction) error {
	if tx.Mined {
		return errors.New("attempting to insert already mind transaction to mempool")
	}
	// Check to see if the tx is present
	_, present := tp.Transactions[tx.TxHash]
	if present {
		return errors.New("attempting to re-insert transaction into mempool")
	}
	tp.Transactions[tx.TxHash] = tx
	return nil
}
