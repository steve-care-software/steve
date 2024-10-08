package contents

import (
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications"
)

type content struct {
	modification modifications.Modification
	data         []byte
}

func createContent(
	modification modifications.Modification,
	data []byte,
) Content {
	out := content{
		modification: modification,
		data:         data,
	}

	return &out
}

// Modification return the modification
func (obj *content) Modification() modifications.Modification {
	return obj.modification
}

// Data return the data
func (obj *content) Data() []byte {
	return obj.data
}
