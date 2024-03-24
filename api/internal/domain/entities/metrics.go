package entities

// Metrics represents the metrics for a specific request.
// It stores the HTTP method, the request path, and the count of requests.
type Metrics struct {
	Method string
	Path   []byte // The path is stored as a byte slice to avoid memory leaks.
	Count  int
}

// NewMetrics creates and returns a new Metrics instance.
// This function is a constructor for the Metrics entity, initializing it with provided values.
//
// Parameters:
// - method: string - The HTTP method of the request (e.g., GET, POST).
// - path: string - The path of the request.
// - count: int - The initial count of requests for the given method and path.
//
// Returns:
// - *Metrics: A pointer to the newly created Metrics instance.
func NewMetrics(method, path string, count int) *Metrics {
	return &Metrics{
		Method: method,
		Path:   []byte(path),
		Count:  count,
	}
}
