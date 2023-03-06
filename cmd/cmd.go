package cmd

import (
	"github.com/go-diydns/utils"
	"github.com/spf13/cobra"

	"github.com/go-diydns/src"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "go-diydns",
	Short: "A simple DIY DNS updater",
	Long:  `A simple DIY DNS updater that updates your DNS records when your public IP changes (or when you want it to)`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintBanner()

		if configPath != "" {
			utils.PrintLog("info", "Using configuration file: "+configPath)
		} else {
			utils.PrintLog("info", "Using default configuration file: "+configPath)
		}

		utils.PrintLog("info", "Starting updater...")

		src.RunUpdate(configPath)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Specify the path to the configuration file",
	Long:  `Specify the path to the configuration file`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of your application",
	Long:  `All software has versions. This is mine.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintLog("info", "Version: "+utils.Version)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config.json", "Path to the configuration file")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
