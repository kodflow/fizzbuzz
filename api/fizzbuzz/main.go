// Package main is the entry point for the application.
package main

//go:generate go run ../components/api/gen.go
//go:generate go fmt ../../api/internal/api/api.gen.go

import (
	"github.com/kodmain/fizzbuzz/api/config"
)

func main() {
	config.Helper.Execute()
}
