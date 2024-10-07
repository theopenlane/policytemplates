package schema

import (
	"github.com/rs/zerolog/log"
)

// Validate validates the provided standards against a predefined JSON schema.
// Returns nil if the standards are valid, otherwise returns an error
// containing the validation errors.
func Validate[T any](standards Framework[T]) error {
	log.Info().Msg("validating standards against schema")

	schema, err := loadEmbeddedSchema("jsonschema/frameworks.json")
	if err != nil {
		return err
	}

	schemaMap, err := structToMap(standards)
	if err != nil {
		return err
	}

	if err := schema.Validate(schemaMap); err != nil {
		log.Error().Err(err).Msg("standards are invalid")

		return err
	}

	return nil
}
