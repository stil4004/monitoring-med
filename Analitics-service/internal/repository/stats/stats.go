package stats_repo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"service/pkg/clients/victoriametrics"
	"service/pkg/model"

	"github.com/shopspring/decimal"
)

type Repository interface {
	GetLastNStats(ctx context.Context, metricName string, n int) ([]model.Stat, error)
	GetStatsInRange(ctx context.Context, metricName string, start, end time.Time) ([]model.Stat, error)
	CheckMetricExists(ctx context.Context, metricName string) (bool, error)
}

type victoriaMetricsRepo struct {
	client victoriametrics.Client
}

func New(client victoriametrics.Client) Repository {
	return &victoriaMetricsRepo{client: client}
}

func (r *victoriaMetricsRepo) GetLastNStats(
	ctx context.Context,
	metricName string,
	n int,
) ([]model.Stat, error) {
	query := fmt.Sprintf("topk(%d, %s)", n, metricName)
	results, err := r.client.InstantQuery(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to get last %d stats: %w", n, err)
	}

	return convertResults(metricName, results), nil
}

func (r *victoriaMetricsRepo) GetStatsInRange(
	ctx context.Context,
	metricName string,
	start, end time.Time,
) ([]model.Stat, error) {
	step := calculateStep(start, end)
	results, err := r.client.RangeQuery(ctx, metricName, start, end, step)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats in range: %w", err)
	}

	return convertResults(metricName, results), nil
}

func (r *victoriaMetricsRepo) CheckMetricExists(ctx context.Context, metricName string) (bool, error) {
	// Используем простой запрос last_over_time для проверки существования метрики
	query := fmt.Sprintf("last_over_time(%s[1m])", metricName)
	_, err := r.client.InstantQuery(ctx, query, time.Now())

	if err != nil {
		if strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "no such metric") {
			return false, nil
		}
		return false, fmt.Errorf("failed to check metric existence: %w", err)
	}

	return true, nil
}

func calculateStep(start, end time.Time) time.Duration {
	duration := end.Sub(start)
	// Возвращаем шаг 1/100 от диапазона, но не менее 1 секунды
	step := duration / 100
	if step < time.Second {
		return time.Second
	}
	return step
}

func convertResults(metricName string, results []victoriametrics.QueryResult) []model.Stat {
	var stats []model.Stat

	for _, res := range results {
		// Обработка instant query
		if len(res.Value) >= 2 {
			if stat := convertDataPoint(metricName, res.Value); stat != nil {
				stats = append(stats, *stat)
			}
		}

		// Обработка range query
		for _, values := range res.Values {
			if len(values) >= 2 {
				if stat := convertDataPoint(metricName, values); stat != nil {
					stats = append(stats, *stat)
				}
			}
		}
	}

	return stats
}

func convertDataPoint(metricName string, point []interface{}) *model.Stat {
	timestamp, ok := point[0].(float64)
	if !ok {
		return nil
	}

	valueStr, ok := point[1].(string)
	if !ok {
		return nil
	}

	value, err := decimal.NewFromString(valueStr)
	if err != nil {
		return nil
	}

	return &model.Stat{
		Name:  metricName,
		Time:  time.Unix(int64(timestamp), 0),
		Value: value,
	}
}
