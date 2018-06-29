package address

import (
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

/*
Helper utilities for wallet and address related functionality
*/
type Wallet struct {
	Priv ci.PrivKey
	ID   peer.ID
}

func GenerateNewWallet() (*Wallet, error) {
	pk, _, err := ci.GenerateKeyPair(ci.RSA, 2048)
	if err != nil {
		return nil, err
	}
	id, err := peer.IDFromPrivateKey(pk)
	if err != nil {
		return nil, err
	}
	return &Wallet{Priv: pk, ID: id}, nil
}
