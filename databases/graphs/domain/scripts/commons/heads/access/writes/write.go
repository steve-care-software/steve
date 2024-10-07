package writes

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"

type write struct {
	modify permissions.Permissions
	review permissions.Permissions
}

func createWrite(
	modify permissions.Permissions,
) Write {
	return createWriteInternally(modify, nil)
}

func createWriteWithReview(
	modify permissions.Permissions,
	review permissions.Permissions,
) Write {
	return createWriteInternally(modify, review)
}

func createWriteInternally(
	modify permissions.Permissions,
	review permissions.Permissions,
) Write {
	out := write{
		modify: modify,
		review: review,
	}

	return &out
}

// Modify returns the modify
func (obj *write) Modify() permissions.Permissions {
	return obj.modify
}

// HasReview returns true if there is a review, false otherwise
func (obj *write) HasReview() bool {
	return obj.review != nil
}

// Review returns the review, if any
func (obj *write) Review() permissions.Permissions {
	return obj.review
}
