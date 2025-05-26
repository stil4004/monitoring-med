//go:build integration

package stats_repo

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"service/pkg/clients/victoriametrics"
	"service/pkg/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

const vmURL = "http://localhost:8428"

func TestVictoriaMetricsRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()
	client := victoriametrics.New(vmURL)
	repo := New(client)

	now := time.Now().UTC()
	testMetrics := []model.Stat{
		{Name: "test_metric", Value: decimal.NewFromInt(10), Time: now.Add(-5 * time.Minute)},
		{Name: "test_metric", Value: decimal.NewFromInt(20), Time: now.Add(-3 * time.Minute)},
		{Name: "test_metric", Value: decimal.NewFromInt(30), Time: now.Add(-1 * time.Minute)},
	}

	// Запись тестовых данных
	require.NoError(t, writeTestData(ctx, client, testMetrics))
	time.Sleep(1 * time.Second) // Ждем индексации

	t.Run("GetLastNStats", func(t *testing.T) {
		stats, err := repo.GetLastNStats(ctx, "test_metric", 2)
		require.NoError(t, err)
		require.Len(t, stats, 2)

		// Проверяем что получили последние 2 значения
		expected := []int64{30, 20}
		for i, val := range expected {
			require.True(t, stats[i].Value.Equal(decimal.NewFromInt(val)),
				"Expected %d, got %s", val, stats[i].Value)
		}
	})

	t.Run("GetStatsInRange", func(t *testing.T) {
		start := now.Add(-6 * time.Minute)
		end := now.Add(-2 * time.Minute)

		stats, err := repo.GetStatsInRange(ctx, "test_metric", start, end)
		require.NoError(t, err)
		require.Len(t, stats, 2)

		// Проверяем что попали только метрики в диапазоне
		expected := []int64{20, 10}
		for i, val := range expected {
			require.True(t, stats[i].Value.Equal(decimal.NewFromInt(val)))
			require.True(t, stats[i].Time.After(start))
			require.True(t, stats[i].Time.Before(end))
		}
	})
}

func writeTestData(ctx context.Context, client victoriametrics.Client, metrics []model.Stat) error {
	var builder strings.Builder
	for _, m := range metrics {
		builder.WriteString(fmt.Sprintf("%s %s %d\n",
			m.Name,
			m.Value.String(),
			m.Time.Unix(),
		))
	}
	return client.WriteMetrics(ctx, builder.String())
}
