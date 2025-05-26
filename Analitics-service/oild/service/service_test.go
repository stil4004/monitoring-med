package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMetric(t *testing.T) {
	t.Run("successful addition", func(t *testing.T) {
		svc := NewMonitoringService()
		err := svc.AddMetric("cpu_usage")
		assert.NoError(t, err)
		assert.True(t, svc.metrics["cpu_usage"])
	})

	t.Run("empty name", func(t *testing.T) {
		svc := NewMonitoringService()
		err := svc.AddMetric("")
		assert.Error(t, err)
		assert.Equal(t, "metric name cannot be empty", err.Error())
	})

	t.Run("duplicate metric", func(t *testing.T) {
		svc := NewMonitoringService()
		_ = svc.AddMetric("cpu_usage")
		err := svc.AddMetric("cpu_usage")
		assert.Error(t, err)
		assert.Equal(t, "metric already exists", err.Error())
	})
}

func TestRemoveMetric(t *testing.T) {
	t.Run("successful removal", func(t *testing.T) {
		svc := NewMonitoringService()
		_ = svc.AddMetric("cpu_usage")
		err := svc.RemoveMetric("cpu_usage")
		assert.NoError(t, err)
		assert.False(t, svc.metrics["cpu_usage"])
	})

	t.Run("non-existent metric", func(t *testing.T) {
		svc := NewMonitoringService()
		err := svc.RemoveMetric("unknown")
		assert.Error(t, err)
		assert.Equal(t, "metric not found", err.Error())
	})
}

func TestSetAlertRule(t *testing.T) {
	svc := NewMonitoringService()
	_ = svc.AddMetric("cpu_usage")

	t.Run("valid rule", func(t *testing.T) {
		err := svc.SetAlertRule("cpu_usage", 90.0, "above")
		assert.NoError(t, err)
		assert.Equal(t, 90.0, svc.alertRules["cpu_usage"].Threshold)
	})

	t.Run("non-existent metric", func(t *testing.T) {
		err := svc.SetAlertRule("unknown", 90.0, "above")
		assert.Error(t, err)
	})
}

func TestCheckForAnomalies(t *testing.T) {
	svc := NewMonitoringService()
	_ = svc.AddMetric("cpu_usage")
	_ = svc.SetAlertRule("cpu_usage", 90.0, "above")

	t.Run("no anomaly", func(t *testing.T) {
		data := MetricData{
			Name:   "cpu_usage",
			Values: []float64{85.0, 86.0, 87.0},
		}
		hasAnomaly, err := svc.CheckForAnomalies(data)
		assert.NoError(t, err)
		assert.False(t, hasAnomaly)
	})

	t.Run("anomaly detected", func(t *testing.T) {
		data := MetricData{
			Name:   "cpu_usage",
			Values: []float64{85.0, 91.0, 87.0},
		}
		hasAnomaly, err := svc.CheckForAnomalies(data)
		assert.NoError(t, err)
		assert.True(t, hasAnomaly)
	})

	t.Run("unmonitored metric", func(t *testing.T) {
		data := MetricData{
			Name:   "unknown",
			Values: []float64{85.0, 91.0, 87.0},
		}
		_, err := svc.CheckForAnomalies(data)
		assert.Error(t, err)
	})
}

func TestGetActiveMetrics(t *testing.T) {
	svc := NewMonitoringService()
	_ = svc.AddMetric("cpu_usage")
	_ = svc.AddMetric("memory_usage")

	metrics := svc.GetActiveMetrics()
	assert.Len(t, metrics, 2)
	assert.Contains(t, metrics, "cpu_usage")
	assert.Contains(t, metrics, "memory_usage")
}
