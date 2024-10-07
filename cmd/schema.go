package cmd

import (
	"github.com/spf13/cobra"

	"github.com/theopenlane/policytemplates/schema"
)

// schemaCmd represents the schema command for generating JSON schema for compliance frameworks
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "generate JSON schema for compliance frameworks",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateSchema()
		cobra.CheckErr(err)
	},
}

// init initializes the schema command
func init() {
	rootCmd.AddCommand(schemaCmd)
}

// generateSchema generates the JSON schema for compliance frameworks
func generateSchema() error {
	return schema.GenerateAuditFrameworksSchema()
}
