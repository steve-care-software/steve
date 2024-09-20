package transfers

import "crypto"

type transfer struct {
	version   uint
	amount    uint64
	publicKey crypto.PublicKey
}

func createTransfer(
	version uint,
	amount uint64,
	publicKey crypto.PublicKey,
) Transfer {
	out := transfer{
		version:   version,
		amount:    amount,
		publicKey: publicKey,
	}

	return &out
}

// Version returns the version
func (obj *transfer) Version() uint {
	return obj.version
}

// Amount returns the amount
func (obj *transfer) Amount() uint64 {
	return obj.amount
}

// PublicKey returns the publicKey
func (obj *transfer) PublicKey() crypto.PublicKey {
	return obj.publicKey
}
