package compensations

import "github.com/steve-care-software/steve/domain/hash"

type compensation struct {
	hash    hash.Hash
	pWrite  *float64
	pReview *float64
}

func createCompensationWithWrite(
	hash hash.Hash,
	pWrite *float64,
) Compensation {
	return createCompensationInternally(hash, pWrite, nil)
}

func createCompensationWithReview(
	hash hash.Hash,
	pReview *float64,
) Compensation {
	return createCompensationInternally(hash, nil, pReview)
}

func createCompensationWithWriteAndReview(
	hash hash.Hash,
	pWrite *float64,
	pReview *float64,
) Compensation {
	return createCompensationInternally(hash, pWrite, pReview)
}

func createCompensationInternally(
	hash hash.Hash,
	pWrite *float64,
	pReview *float64,
) Compensation {
	out := compensation{
		hash:    hash,
		pWrite:  pWrite,
		pReview: pReview,
	}

	return &out
}

// Hash returns the hash
func (obj *compensation) Hash() hash.Hash {
	return obj.hash
}

// HasWrite returns true if there is a write, false otherwise
func (obj *compensation) HasWrite() bool {
	return obj.pWrite != nil
}

// Write returns write, if any
func (obj *compensation) Write() *float64 {
	return obj.pWrite
}

// HasReview returns true if there is a review, false otherwise
func (obj *compensation) HasReview() bool {
	return obj.pReview != nil
}

// Review returns review, if any
func (obj *compensation) Review() *float64 {
	return obj.pReview
}
