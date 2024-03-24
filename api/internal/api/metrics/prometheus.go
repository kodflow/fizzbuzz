package metrics

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kodflow/fizzbuzz/api/internal/application/services"
	"github.com/kodflow/fizzbuzz/api/internal/architecture/persistence"
	"github.com/kodflow/fizzbuzz/api/internal/domain/entities"
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
	allStats, err := service.GetAllRequestStats()
	if err != nil {
		logger.Error(err)
		return sendPrometheusError(c)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	var prometheusFormat strings.Builder

	prometheusFormat.WriteString("# HELP request_count Total number of requests for each request path.\n")
	prometheusFormat.WriteString("# TYPE request_count counter\n")

	var maxHitStats *entities.Metrics
	for _, stats := range allStats {
		prometheusFormat.WriteString(fmt.Sprintf("request_count{method=\"%s\", path=\"%s\"} %d\n",
			stats.Method, stats.Path, stats.Count))
		if maxHitStats == nil || stats.Count > maxHitStats.Count {
			maxHitStats = stats
		}
	}

	// Ajouter la métrique request_max_hit
	if maxHitStats != nil {
		prometheusFormat.WriteString("# HELP request_max_hit Details of the request with the most hits.\n")
		prometheusFormat.WriteString("# TYPE request_max_hit counter\n")
		prometheusFormat.WriteString(fmt.Sprintf("request_max_hit{method=\"%s\", path=\"%s\"} %d\n",
			maxHitStats.Method, maxHitStats.Path, maxHitStats.Count))
	}

	return c.SendString(prometheusFormat.String())
}

func sendPrometheusError(c *fiber.Ctx) error {
	errorMetric := "# HELP request_error Indicates an error occurred while fetching statistics.\n" +
		"# TYPE request_error counter\n" +
		"request_error 1\n"
	return c.SendString(errorMetric)
}

// Internal middleware to increment request counter
func Counter(c *fiber.Ctx) error {
	service.IncrementRequest(c.Method(), c.Path())
	return c.Next()
}
