// Package config provides functions/var/const for loading and accessing configuration settings for the application.
package config

import (
	"github.com/kodmain/fizzbuzz/api/internal/kernel"
	"github.com/kodmain/fizzbuzz/api/internal/kernel/observability/logger"
	"github.com/spf13/cobra"
)

// Helper use Cobra package to create a CLI and give Args gesture
var Helper *cobra.Command = &cobra.Command{
	Use:                   "fizzbuzz",
	Short:                 "Fizzbuzz API Server",
	DisableAutoGenTag:     true,
	DisableFlagsInUseLine: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("loading configuration")
		// TODO loading configuration settings
		// Exit if no configuration settings are found or fail to load them
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("starting application")
		// TODO start the application
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		kernel.Wait()
	},
}
