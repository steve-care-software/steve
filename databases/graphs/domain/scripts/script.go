package scripts

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas"

type script struct {
	schema schemas.Schema
}

func createScriptWithSchema(
	schema schemas.Schema,
) Script {
	return createScriptInternally(schema)
}

func createScriptInternally(
	schema schemas.Schema,
) Script {
	out := script{
		schema: schema,
	}

	return &out
}

// IsSchema returns true if there is a schema, false otherwise
func (obj *script) IsSchema() bool {
	return obj.schema != nil
}

// Schema returns the schema, if any
func (obj *script) Schema() schemas.Schema {
	return obj.schema
}
