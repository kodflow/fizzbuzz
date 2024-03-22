//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	header = `
// Automatically generated by api/generator/api.gen.go, DO NOT EDIT manually
// Package api implements Register method for fiber
package api

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/swag"
)

func init() {
	json.Unmarshal([]byte(doc), Mapping)
}

// API represents a collection of HTTP endpoints grouped by namespace and version.
var (
	Endpoints map[string]fiber.Handler = map[string]func(*fiber.Ctx) error{
`
	footer = `
	}
	Mapping = &docs.Swagger{}
	doc, _  = swag.ReadDoc()
)
`
)

// getGitRoot retourne le chemin du dossier racine du dépôt Git.
// Returns:
// - path: string Chemin du dossier racine Git.
// - err: error Erreur, le cas échéant.
func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func main() {

	wd, err := GetGitRoot()
	if err != nil {
		panic(err)
	}

	var ids []string
	var re *regexp.Regexp = regexp.MustCompile(`// @Id\s+([^\s]+)`)

	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || strings.Contains(path, "vendor") || !strings.HasSuffix(path, ".go") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			matches := re.FindStringSubmatch(scanner.Text())
			if len(matches) > 1 {
				ids = append(ids, matches[1])
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	generatedPath := path.Join(wd, "api", "internal", "api", "api.gen.go")

	os.RemoveAll(generatedPath)
	file, err := os.Create(generatedPath)
	if err != nil {
		panic(err)
	}

	file.WriteString(header)
	for _, id := range ids {
		file.WriteString(fmt.Sprintf("\"%s\": %s,\n", id, id))
	}
	file.WriteString(footer)

	goimportsPath, err := exec.LookPath("goimports")
	if err != nil {
		panic("goimports not found. Please install it using 'go install golang.org/x/tools/cmd/goimports@latest'")
	}

	cmd := exec.Command(goimportsPath, "-w", generatedPath)
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("goimports failed: %s", err))
	}
}
