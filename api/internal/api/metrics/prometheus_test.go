package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kodflow/fizzbuzz/api/internal/api/metrics"
	"github.com/stretchr/testify/assert"
)

func TestStatistics(t *testing.T) {
	app := fiber.New()
	app.Get("/metrics/statistics", metrics.Statistics)

	req := httptest.NewRequest(http.MethodGet, "/metrics/statistics", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request should not return an error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status code should be 200")

}

func TestCounter(t *testing.T) {
	app := fiber.New()
	app.Use(metrics.Counter)
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request should not return an error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status code should be 200")
}
