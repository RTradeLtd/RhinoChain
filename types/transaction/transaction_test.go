package transaction

import "testing"

func TestTransaction(t *testing.T) {
	tp := InitTransactionPool()
	tx := &Transaction{
		TxHash: "1",
		To:     "to",
		From:   "from",
		Value:  0,
		Mined:  false,
	}
	err := tp.AddTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	err = tp.AddTransaction(tx)
	if err == nil {
		t.Fatal("no error encountered when one shouldve been")
	}
}
