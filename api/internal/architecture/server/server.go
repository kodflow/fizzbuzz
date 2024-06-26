// Package server provides a simple HTTP/HTTPS server implementation for serving static and connect API.
package server

import (
	"crypto/tls"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kodflow/fizzbuzz/api/internal/api"
	"github.com/kodflow/fizzbuzz/api/internal/application/observability/logger"
	"github.com/kodflow/fizzbuzz/api/internal/architecture/kernel"
	"github.com/kodflow/fizzbuzz/api/internal/docs"
)

var methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}

// Server is a struct that represents a Fiber server instance with an underlying `fiber.App`, `fiber.Router` for the API endpoints, and a map of `fiber.Router` instances for different API versions.
type Server struct {
	// `app` is an instance of the `fiber.App` structure that represents the underlying Fiber server instance.
	app *fiber.App
	// `certs` is a pointer to a `tls.Config` structure that holds the TLS configuration for the server.
	certs *tls.Config
}

// Start is a method of the `Server` struct that starts the server and listens for incoming HTTP and/or HTTPS requests, depending on the `config` settings.
func (server *Server) Start() {
	// Launch a goroutine to start the HTTP server.
	go server.http()
	// If HTTPS is enabled in the `config`, launch a goroutine to start the HTTPS server as well.
	go server.https()
}

// API is a method of the `Server` struct that registers the provided API with the server.
// It creates a new version of the API router, adds the provided API to the router's namespace, and registers the new router with the server's main `app` instance.
func (server *Server) Register(handlers map[string]fiber.Handler) {
	for url, pathItem := range api.Mapping.Paths {
		for _, method := range methods {
			if operationID, exist := getOperationID(pathItem, method); exist {
				handlersToRegister := getHandlers(operationID, handlers)
				if len(handlersToRegister) > 0 {
					server.app.Add(method, url, handlersToRegister...)
					logger.Infof("Register %v %v with %v", method, url, operationID)
				} else {
					logger.Warnf("Handler not found %v", operationID)
				}
			}
		}
	}
}

func getHandlers(operationID string, handlers map[string]fiber.Handler) []fiber.Handler {
	var handlersToRegister []fiber.Handler
	var operationIDs = strings.Split(operationID, "=>")
	for key, handler := range operationIDs {
		handler = strings.TrimSpace(handler)
		if h, exist := handlers[handler]; exist {
			handlersToRegister = append(handlersToRegister, h)
		} else {
			logger.Warnf("Handler not found %v %v", key, handler)
		}
	}
	return handlersToRegister
}

func getOperationID(pathItem *docs.PathItem, method string) (string, bool) {
	switch method {
	case "GET":
		if pathItem.Get != nil {
			return pathItem.Get.OperationID, true
		}
	case "POST":
		if pathItem.Post != nil {
			return pathItem.Post.OperationID, true
		}
	case "PUT":
		if pathItem.Put != nil {
			return pathItem.Put.OperationID, true
		}
	case "DELETE":
		if pathItem.Delete != nil {
			return pathItem.Delete.OperationID, true
		}
	case "PATCH":
		if pathItem.Patch != nil {
			return pathItem.Patch.OperationID, true
		}
	case "OPTIONS":
		if pathItem.Options != nil {
			return pathItem.Options.OperationID, true
		}
	case "HEAD":
		if pathItem.Head != nil {
			return pathItem.Head.OperationID, true
		}
	}
	return "", false
}

func (server *Server) http() {
	kernel.PANIC <- server.app.Listen(":80")
	logger.Info("Server started on port 80")
}

func (server *Server) https() {
	kernel.PANIC <- server.app.ListenTLSWithCertificate(":443", server.certs.Certificates[0])
	logger.Info("Server started on port 443")
}
