package server

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kodflow/fizzbuzz/api/config"
	"github.com/kodflow/fizzbuzz/api/internal/api/status"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	srv := Create(fiber.Config{
		AppName: config.APP_NAME,
		Prefork: false, // Multithreading
	})

	srv.Register(map[string]fiber.Handler{
		"status.HealthCheck": status.HealthCheck,
	})

	go func() {
		srv.Start()
	}()

	assert.NotNil(t, srv)
}
