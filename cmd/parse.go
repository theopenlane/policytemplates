package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/theopenlane/policytemplates/cmd/prompts"
	"github.com/theopenlane/policytemplates/standards"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse compliance standards",
	Long:  `parse compliance standards and output the results in a json format.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := parse()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringP("framework", "f", "soc2", "compliance framework to parse (soc2, nist)")
	parseCmd.Flags().StringP("output", "o", "stdout", "output format (stdout, file)")

	parseCmd.Flags().BoolP("interactive", "i", true, "interactive prompt, set to false to disable")
}

func parse() error {
	framework := config.String("framework")
	format := config.String("output")

	var err error
	if config.Bool("interactive") {
		framework, err = prompts.Frameworks()
		cobra.CheckErr(err)

		format, err = prompts.Formats()
		cobra.CheckErr(err)
	}

	log.Info().Str("framework", framework).Str("format", format).Msg("parsing compliance standards")

	var (
		controls any
		filename string
	)

	switch framework {
	case "soc2":
		controls, err = standards.GenerateSOC2Standards()
		cobra.CheckErr(err)

		filename = "templates/standards/soc2-2022.json"
	case "nist-csf":
		controls, err = standards.GenerateNISTCSFStandards()
		cobra.CheckErr(err)

		filename = "templates/standards/nist-csf-1.1.json"
	case "nist-800-53":
		controls, err = standards.GenerateNist80053Standards()
		cobra.CheckErr(err)

		filename = "templates/standards/nist-800-53-5.json"
	default:
		log.Error().Str("framework", framework).Msg("framework not found")

		return fmt.Errorf("framework not supported") //nolint: err113
	}

	if format == "file" {
		err := saveToFile(controls, filename)
		cobra.CheckErr(err)

		log.Info().Str("filename", filename).Msg("standards saved to file")

		return nil
	}

	return jsonOutput(controls)
}
