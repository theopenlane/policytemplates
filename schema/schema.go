package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/rs/zerolog/log"
)

const (
	// Version is the current version of the schema
	Version = "draft-0"
	// BaseURL is the base URL for the schema
	BaseURL = "https://theopenlane.io"
)

var (
	// BaseSchemaID is the base schema ID for the schema
	BaseSchemaID = fmt.Sprintf("%s/%s/", BaseURL, Version)
)

const (
	outputFile = "schema/jsonschema/frameworks.json"
)

// namePkg returns the package name of the provided type without the interface{} suffix
func namePkg(r reflect.Type) string {
	return strings.ReplaceAll(r.Name(), "[interface {}]", "")
}

// GenerateAuditFrameworksSchema generates a JSON schema for the audit standards
func GenerateAuditFrameworksSchema() error {
	log.Info().Msg("generating schema")

	r := jsonschema.Reflector{
		Namer:        namePkg,
		BaseSchemaID: jsonschema.ID(BaseSchemaID),
	}

	s := r.Reflect(Framework[any]{})

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	log.Info().Msg("writing schema to file")

	if err := os.WriteFile(outputFile, data, 0600); err != nil { //nolint:mnd
		return err
	}

	log.Info().Str("file location", outputFile).Msg("schema generated successfully")

	return nil
}
