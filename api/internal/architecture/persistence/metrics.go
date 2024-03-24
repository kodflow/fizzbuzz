package persistence

import (
	"errors"
	"sync"

	"github.com/kodflow/fizzbuzz/api/internal/domain/entities"
)

// MetricsRepository is a repository responsible for managing metrics data.
// It provides functionalities to increment request counts and retrieve the most frequent request.
type MetricsRepository struct {
	// mu synchronizes access to the requestCounter.
	mu sync.Mutex
	// requestCounter stores the count of requests for each method-path combination.
	requestCounter map[string]*entities.Metrics
}

// NewMetricsRepository creates and returns a new instance of MetricsRepository.
//
// Returns:
// - *MetricsRepository: A pointer to the newly created MetricsRepository instance.
func NewMetricsRepository() *MetricsRepository {
	return &MetricsRepository{
		requestCounter: make(map[string]*entities.Metrics),
	}
}

// IncrementRequestCount increments the count for a given request method and path.
// It safely updates the requestCounter map ensuring concurrent access is properly handled.
//
// Parameters:
// - method: string - The HTTP method of the request (e.g., GET, POST).
// - path: string - The path of the request.
//
// Returns:
// - error: An error if the increment operation fails, otherwise nil.
func (r *MetricsRepository) IncrementRequestCount(method, path string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := method + ":" + path
	stats, exists := r.requestCounter[key]

	if exists {
		stats.Count++
	} else {
		r.requestCounter[key] = entities.NewMetrics(method, path, 1)
	}

	return nil
}

// GetMostFrequentRequest retrieves the metrics for the most frequently made request.
// It iterates through the requestCounter map to find the request with the highest count.
//
// Returns:
// - *entities.Metrics: The metrics of the most frequent request.
// - error: An error if no requests have been made yet, otherwise nil.
func (r *MetricsRepository) GetMostFrequentRequest() (*entities.Metrics, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var mostFrequent *entities.Metrics

	for _, stats := range r.requestCounter {
		if mostFrequent == nil || stats.Count > mostFrequent.Count {
			mostFrequent = stats
		}
	}

	if mostFrequent == nil {
		return nil, errors.New("no request has been made yet")
	}

	return mostFrequent, nil
}

// GetAllRequestStats retrieves the metrics data for all requests.
// It returns a slice containing the metrics for each unique request.
//
// Returns:
// - []*entities.Metrics: A slice of metrics data for all requests.
// - error: An error if no requests have been made yet, otherwise nil.
func (r *MetricsRepository) GetAllRequestStats() ([]*entities.Metrics, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.requestCounter) == 0 {
		return nil, errors.New("no request has been made yet")
	}

	var allStats []*entities.Metrics

	for _, stats := range r.requestCounter {
		allStats = append(allStats, stats)
	}

	return allStats, nil
}
