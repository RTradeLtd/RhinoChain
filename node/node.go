package node

import (
	"github.com/RTradeLtd/RhinoChain/types/address"
	"github.com/RTradeLtd/RhinoChain/types/transaction"
	peer "github.com/libp2p/go-libp2p-peer"
)

type NodeManager struct {
	ID               peer.ID                      `json:"id"`
	TemporalURL      string                       `json:"temporal_url"`      // url we use to connect to temporal
	TemporalUser     string                       `json:"temporal_user"`     // user account we use to connect to temporal
	TemporalPassword string                       `json:"temporal_password"` // password used to authenticate with temporal
	Wallet           *address.Wallet              `json:"wallet"`
	TransactionPool  *transaction.TransactionPool `json:"transaction_pool"`
}

func GenerateNodeManager(url, user, pass string) (*NodeManager, error) {
	wallet, err := address.GenerateNewWallet()
	if err != nil {
		return nil, err
	}
	tp := transaction.InitTransactionPool()
	nm := &NodeManager{
		ID:               wallet.ID,
		TemporalURL:      url,
		TemporalUser:     user,
		TemporalPassword: pass,
		Wallet:           wallet,
		TransactionPool:  tp,
	}

	return nm, nil
}
