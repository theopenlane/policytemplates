package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/theopenlane/policytemplates/cmd/prompts"
	"github.com/theopenlane/policytemplates/frameworks/nist80053"
	"github.com/theopenlane/policytemplates/frameworks/nistcsf"
	"github.com/theopenlane/policytemplates/frameworks/soc2"
)

// parseCmd represents the parse command for parsing compliance frameworks from csv files to a standardized JSON format
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse compliance frameworks",
	Long:  `parse compliance frameworks and output the results in a json format.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := parse()
		cobra.CheckErr(err)
	},
}

// init initializes the parse command
func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringP("framework", "f", "soc2", "compliance framework to parse (soc2, nist)")
	parseCmd.Flags().StringP("output", "o", "stdout", "output format (stdout, file)")

	parseCmd.Flags().BoolP("interactive", "i", true, "interactive prompt, set to false to disable")
}

// parse parses the compliance frameworks
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

	log.Info().Str("framework", framework).Str("format", format).Msg("parsing compliance frameworks")

	var (
		controls any
		filename string
	)

	switch framework {
	case "soc2":
		controls, err = soc2.Generate()
		cobra.CheckErr(err)

		filename = "templates/frameworks/soc2-2022.json"
	case "nist-csf":
		controls, err = nistcsf.Generate()
		cobra.CheckErr(err)

		filename = "templates/frameworks/nist-csf-1.1.json"
	case "nist-800-53":
		controls, err = nist80053.Generate()
		cobra.CheckErr(err)

		filename = "templates/frameworks/nist-800-53-5.json"
	default:
		log.Error().Str("framework", framework).Msg("framework not found")

		return fmt.Errorf("framework not supported") //nolint: err113
	}

	if format == "file" {
		err := saveToFile(controls, filename)
		cobra.CheckErr(err)

		log.Info().Str("filename", filename).Msg("frameworks saved to file")

		return nil
	}

	return jsonOutput(controls)
}
