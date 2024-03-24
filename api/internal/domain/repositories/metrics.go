package repositories

import "github.com/kodflow/fizzbuzz/api/internal/domain/entities"

// MetricsRepository is an interface defining the operations required for managing metrics data.
// It provides an abstraction over the data layer for working with request metrics.
type MetricsRepository interface {
	// IncrementRequestCount increments the count for a specific request method and path.
	// This method is intended to track the number of occurrences of each unique request.
	//
	// Parameters:
	// - method: string - The HTTP method of the request (e.g., GET, POST).
	// - path: string - The path of the request.
	//
	// Returns:
	// - error: An error if the increment operation fails, otherwise nil.
	IncrementRequestCount(method, path string) error

	// GetMostFrequentRequest retrieves the request metrics for the most frequently accessed endpoint.
	// This method is used to identify which request has the highest count.
	//
	// Returns:
	// - *entities.Metrics: The metrics data of the most frequent request.
	// - error: An error if the retrieval fails or if no data is available, otherwise nil.
	GetMostFrequentRequest() (*entities.Metrics, error)

	// GetAllRequestStats retrieves the metrics data for all requests.
	// This method returns a list of metrics for each unique request.
	//
	// Returns:
	// - []*entities.Metrics: A slice of metrics data for all requests.
	// - error: An error if the retrieval fails or if no data is available, otherwise nil.
	GetAllRequestStats() ([]*entities.Metrics, error)
}
