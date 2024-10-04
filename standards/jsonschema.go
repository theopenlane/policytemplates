package standards

import (
	"bytes"
	"embed"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed jsonschema/*.json
var schemas embed.FS

// embeddedLoader is a custom loader that loads the predefined JSON schema from go:embed
type embeddedLoader struct {
	schemaName string
}

// Load loads the predefined JSON schema from go:embed.
// it satisfies the jsonschema.Loader interface, which requires a string argument,
// this is unused in this case because it changes the name to the full "filepath" e.g.
// "file:///Users/sarahfunkhouser/go/src/github.com/theopenlane/policytemplates/standards/jsonschema/standards.json"
// which is not what we want here and it would be convoluted to figure out the correct path
// to the embedded schema
// instead we just use the schemaName from the struct to load the schema from go:embed
func (e embeddedLoader) Load(n string) (any, error) {
	schema, err := schemas.ReadFile(e.schemaName)
	if err != nil {
		return nil, err
	}

	r := bytes.NewBuffer(schema)

	return jsonschema.UnmarshalJSON(r)
}

// loadEmbeddedSchema loads the predefined JSON schema from go:embed
// use this helper function over the `Load` method to load the schema unless you know what you're doing
func loadEmbeddedSchema(n string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()

	// always set the default draft to prevent unknown schema errors
	// with newer JSON schema drafts
	compiler.DefaultDraft(jsonschema.Draft2020)

	// use a custom loader to load the predefined JSON schema from go:embed
	compiler.UseLoader(embeddedLoader{
		schemaName: n,
	})

	// compile the schema
	return compiler.Compile("")
}
