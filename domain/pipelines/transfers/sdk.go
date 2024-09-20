package transfers

import "crypto"

// Transfer represents a transfer
type Transfer interface {
	Version() uint
	Amount() uint64
	PublicKey() crypto.PublicKey
}
