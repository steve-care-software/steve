package contexts

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents"
)

type builder struct {
	hashAdapter hash.Adapter
	head        heads.Head
	content     contents.Content
	parent      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		content:     nil,
		parent:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content contents.Content) Builder {
	app.content = content
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent string) Builder {
	app.parent = parent
	return app
}

// Now builds a new Context instance
func (app *builder) Now() (Context, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Context instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Context instance")
	}

	data := [][]byte{
		app.head.Hash().Bytes(),
		app.content.Hash().Bytes(),
	}

	if app.parent != "" {
		data = append(data, []byte(app.parent))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != "" {
		return createContextWithParent(
			*pHash,
			app.head,
			app.content,
			app.parent,
		), nil
	}

	return createContext(*pHash, app.head, app.content), nil
}
