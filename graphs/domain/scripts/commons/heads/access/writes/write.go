package writes

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"

type write struct {
	modify permissions.Permission
	review permissions.Permission
}

func createWrite(
	modify permissions.Permission,
) Write {
	return createWriteInternally(modify, nil)
}

func createWriteWithReview(
	modify permissions.Permission,
	review permissions.Permission,
) Write {
	return createWriteInternally(modify, review)
}

func createWriteInternally(
	modify permissions.Permission,
	review permissions.Permission,
) Write {
	out := write{
		modify: modify,
		review: review,
	}

	return &out
}

// Modify returns the modify
func (obj *write) Modify() permissions.Permission {
	return obj.modify
}

// HasReview returns true if there is a review, false otherwise
func (obj *write) HasReview() bool {
	return obj.review != nil
}

// Review returns the review, if any
func (obj *write) Review() permissions.Permission {
	return obj.review
}
