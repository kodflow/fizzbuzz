// Package main is the entry point for the application.
package main

//go:generate go run ../internal/docs/generator.go
//go:generate go fmt ../../api/internal/api/api.gen.go

import (
	"github.com/kodflow/fizzbuzz/api/internal/kernel"
	"github.com/kodflow/fizzbuzz/api/internal/kernel/observability/logger"
	"github.com/kodflow/fizzbuzz/api/internal/server"
	"github.com/spf13/cobra"
)

// Helper use Cobra package to create a CLI and give Args gesture
var Helper *cobra.Command = &cobra.Command{
	Use:                   "fizzbuzz",
	Short:                 "Fizzbuzz API Server",
	DisableAutoGenTag:     true,
	DisableFlagsInUseLine: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Info("loading configuration")
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("starting application")
		server.Create().Start()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		logger.Info("waiting for application to shutdown")
		kernel.Wait()
	},
}

func main() {
	Helper.Execute()
}
