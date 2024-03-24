package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodflow/fizzbuzz/api/internal/application/services"
	"github.com/kodflow/fizzbuzz/api/internal/architecture/persistence"
	"github.com/kodflow/fizzbuzz/api/internal/architecture/serializers/prom"
	"github.com/kodflow/fizzbuzz/api/internal/kernel/observability/logger"
)

var repository = persistence.NewMetricsRepository()
var service = services.NewMetricsService(repository)

// @Summary      Show request metrics.
// @Description  Retrieves data for the most frequent request.
// @Tags         Metrics
// @Accept       */*
// @Produce      application/json
// @Success      200  {object}  entities.Metrics  "Statistics of the most frequent request"
// @Failure      404  {string}  nil           	  "No data available"
// @Router       /metrics/statistics [get]
// @Id           metrics.Statistics
func Statistics(c *fiber.Ctx) error {
	allHits, err := service.GetAllRequestStats()
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	maxHits, err := service.GetMostFrequentRequest()
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	hits, err := prom.NewMetricMeta("request_count", "Total number of requests for each request path.", "counter", allHits)
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	max, err := prom.NewMetricMeta("request_max_hit", "Details of the request with the most hits.", "counter", maxHits)
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	promBytes, err := prom.Marshal(hits, max)
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	return c.SendString(string(promBytes))
}

func sendPrometheusError(c *fiber.Ctx) error {
	data := prom.NewSimpleMeta("request_error", "Indicates an error occurred while fetching statistics.", "counter", 1)
	promBytes, _ := prom.Marshal(data)
	return c.SendString(string(promBytes))
}

func Counter(c *fiber.Ctx) error {
	service.IncrementRequest(c.Method(), c.Path())
	return c.Next()
}
