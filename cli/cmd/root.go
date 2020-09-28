package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Configuration file provided
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "affinityctl",
	Short: "Affinity",
	Long: strings.TrimSpace(`
Affinity CLI Tool.

General tools to facilitate integration with the
Affinity digital identity services.`),
}

// Execute provides the main entry point for the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
	// Used for ENV variables prefix and home directories
	var appIdentifier string = "affinityctl"

	// ENV
	viper.SetEnvPrefix(appIdentifier)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Configuration file
	viper.AddConfigPath(fmt.Sprintf("/etc/%s", appIdentifier))
	viper.AddConfigPath(fmt.Sprintf("$HOME/%s", appIdentifier))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", appIdentifier))
	viper.AddConfigPath(".")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	// Read configuration file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("failed to read configuration file: %s\n", err.Error())
		}
	}
}
