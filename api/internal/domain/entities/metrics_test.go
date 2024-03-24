package entities_test

import (
	"testing"

	"github.com/kodflow/fizzbuzz/api/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	metrics := entities.NewMetrics()
	assert.NotNil(t, metrics, "Metrics should not be nil")
	assert.Len(t, metrics, 0)
}
