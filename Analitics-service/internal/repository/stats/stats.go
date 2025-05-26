package stats_repo

import (
	"context"
	"fmt"
	"sort"
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
	WriteStat(ctx context.Context, stat model.Stat) error
}

type victoriaMetricsRepo struct {
	client victoriametrics.Client
}

func New(client victoriametrics.Client) Repository {
	return &victoriaMetricsRepo{client: client}
}

func (r *victoriaMetricsRepo) WriteStat(ctx context.Context, stat model.Stat) error {
	metric := fmt.Sprintf("%s %s %d", stat.Name, stat.Value.String(), stat.Time.Unix())
	return r.client.WriteMetrics(ctx, metric)
}

func (r *victoriaMetricsRepo) GetLastNStats(
	ctx context.Context,
	metricName string,
	n int,
) ([]model.Stat, error) {
	// Используем range запрос за последний час, чтобы гарантированно получить данные
	query := fmt.Sprintf("sort_desc(%s[1h])", metricName)
	results, err := r.client.RangeQuery(ctx, query, time.Now().Add(-1*time.Hour), time.Now(), time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to get last %d stats: %w", n, err)
	}

	stats := convertResults(metricName, results)
	if len(stats) > n {
		stats = stats[:n]
	}
	return stats, nil
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
	// Используем запрос, который возвращает пустой результат если метрика не существует
	query := fmt.Sprintf("count(%s)", metricName)
	results, err := r.client.InstantQuery(ctx, query, time.Now())
	if err != nil {
		if strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "no such metric") {
			return false, nil
		}
		return false, fmt.Errorf("failed to check metric existence: %w", err)
	}

	// Проверяем, есть ли хоть один результат с ненулевым значением
	for _, res := range results {
		if len(res.Value) >= 2 {
			valueStr, ok := res.Value[1].(string)
			if ok && valueStr != "0" {
				return true, nil
			}
		}
		for _, values := range res.Values {
			if len(values) >= 2 {
				valueStr, ok := values[1].(string)
				if ok && valueStr != "0" {
					return true, nil
				}
			}
		}
	}

	return false, nil
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

	// Сортируем по времени (новые сначала)
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Time.After(stats[j].Time)
	})

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
