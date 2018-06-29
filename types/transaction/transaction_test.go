package transaction

import "testing"

func TestTransaction(t *testing.T) {
	tp := InitTransactionPool()
	tx := GenerateTransaction("to", "from", 1)
	err := tp.AddTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	err = tp.AddTransaction(tx)
	if err == nil {
		t.Fatal("no error encountered when one shouldve been")
	}
}
