package stats_repo

import (
	"context"
	"testing"
	"time"

	"service/pkg/clients/victoriametrics"
	"service/pkg/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const vmURL = "http://localhost:8428"

func TestVictoriaMetricsRepository(t *testing.T) {
	ctx := context.Background()
	vmClient := victoriametrics.New(vmURL)
	repo := New(vmClient)

	// Используем уникальные имена метрик для каждого теста
	testMetricName := "test_metric_" + time.Now().Format("20060102150405")
	otherMetricName := "other_metric_" + time.Now().Format("20060102150405")

	// Подготовка тестовых данных
	now := time.Now()
	testMetrics := []model.Stat{
		{Name: testMetricName, Value: decimal.NewFromInt(10), Time: now.Add(-5 * time.Minute)},
		{Name: testMetricName, Value: decimal.NewFromInt(20), Time: now.Add(-3 * time.Minute)},
		{Name: testMetricName, Value: decimal.NewFromInt(30), Time: now.Add(-1 * time.Minute)},
		{Name: otherMetricName, Value: decimal.NewFromInt(100), Time: now.Add(-2 * time.Minute)},
	}

	// Запись тестовых данных
	for _, metric := range testMetrics {
		err := repo.WriteStat(ctx, metric)
		assert.NoError(t, err)
	}

	// Даем время на индексацию
	time.Sleep(10 * time.Second)

	t.Run("WriteStat and CheckMetricExists", func(t *testing.T) {
		exists, err := repo.CheckMetricExists(ctx, testMetricName)
		require.NoError(t, err)
		assert.True(t, exists)

		exists, err = repo.CheckMetricExists(ctx, "non_existent_metric")
		require.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("GetLastNStats", func(t *testing.T) {
		stats, err := repo.GetLastNStats(ctx, testMetricName, 2)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(stats), 2, "should return at least 2 metrics")

		// Проверяем что значения отсортированы по времени (новые сначала)
		assert.True(t, stats[0].Time.After(stats[1].Time))
	})

	t.Run("GetStatsInRange", func(t *testing.T) {
		start := now.Add(-6 * time.Minute)
		end := now.Add(-2 * time.Minute)
		stats, err := repo.GetStatsInRange(ctx, testMetricName, start, end)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(stats), 2, "should return at least 2 metrics in range")

		for _, stat := range stats {
			assert.True(t, stat.Time.After(start) || stat.Time.Equal(start))
			assert.True(t, stat.Time.Before(end) || stat.Time.Equal(end))
		}
	})
}
