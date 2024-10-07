package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const appName = "policytemplates"

var (
	cfgFile string
	config  *koanf.Koanf
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "policytemplates",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfiguration(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})

	config = koanf.New(".")

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// load the flags to ensure we know the correct config file path
	initConfiguration(rootCmd)

	// load the config file and env vars
	loadConfigFile()

	// reload because flags and env vars take precedence over file
	initConfiguration(rootCmd)
}

// initConfiguration loads the configuration from the command flags of the given cobra command
func initConfiguration(cmd *cobra.Command) {
	loadEnvVars()

	loadFlags(cmd)
}

// loadConfigFile loads the configuration from the config file
func loadConfigFile() {
	if cfgFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		cfgFile = filepath.Join(home, "."+appName+".yaml")
	}

	// If the config file does not exist, do nothing
	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	err := config.Load(file.Provider(cfgFile), yaml.Parser())

	cobra.CheckErr(err)
}

// loadEnvVars loads the configuration from the environment variables
func loadEnvVars() {
	err := config.Load(env.ProviderWithValue("POLICY_TEMPLATES_", ".", func(s string, v string) (string, interface{}) {
		key := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "POLICY_TEMPLATES_")), "_", ".")

		if strings.Contains(v, ",") {
			return key, strings.Split(v, ",")
		}

		return key, v
	}), nil)

	cobra.CheckErr(err)
}

// loadFlags loads the configuration from the command flags of the given cobra command
func loadFlags(cmd *cobra.Command) {
	err := config.Load(posflag.Provider(cmd.Flags(), config.Delim(), config), nil)

	cobra.CheckErr(err)
}
