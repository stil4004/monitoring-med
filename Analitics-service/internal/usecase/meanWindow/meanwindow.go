// meanwindow_usecase/meanwindow.go
package meanwindow_usecase

import (
	"context"
	"fmt"
	"service/internal/cconstants"
	stats_repo "service/internal/repository/stats"
	"service/pkg/ctx_log"
	"sync"
	"time"

	"go.uber.org/zap"
)

type UseCase interface {
	Controller(ctx context.Context)
	AddMetricToMonitoring(ctx context.Context, metricName string) error
	PauseMetricMonitoring(ctx context.Context, metricName string) error
	ResumeMetricMonitoring(ctx context.Context, metricName string) error
	RemoveMetricFromMonitoring(ctx context.Context, metricName string) error
	GetMonitoredMetrics(ctx context.Context) []string
}

type StatsChecker interface {
	CheckMetricExists(ctx context.Context, metricName string) (bool, error)
}

type meanWindowUsecase struct {
	statsChecker        StatsChecker
	monitoredMetrics    map[string]bool
	monitoredMetricsMux *sync.RWMutex
	windowSize          int
}

func NewMeanWindowUsecase(checker stats_repo.Repository) UseCase {
	return &meanWindowUsecase{
		statsChecker:        checker,
		monitoredMetrics:    make(map[string]bool),
		monitoredMetricsMux: &sync.RWMutex{},
		windowSize:          10, // размер окна для анализа
	}
}

func (u *meanWindowUsecase) Controller(ctx context.Context) {
	ticker := time.NewTicker(cconstants.MonitoringCheckInterval * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			u.checkMetrics(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (u *meanWindowUsecase) checkMetrics(ctx context.Context) {
	u.monitoredMetricsMux.RLock()
	defer u.monitoredMetricsMux.RUnlock()

	l := ctx_log.LoggerFromContext(ctx)

	for metric, active := range u.monitoredMetrics {
		if !active {
			continue
		}

		// Здесь должна быть логика проверки метрики
		// Например:
		// values, err := analyticUC.GetLastNStats(ctx, metric, u.windowSize)
		// if err != nil {...}

		// Пример места для обнаружения аномалий:
		l.Info("Checking metric", zap.String("metric", metric))

		// При обнаружении аномалии:
		// l.Warn("Anomaly detected", zap.String("metric", metric))
		// Здесь можно добавить вызов службы уведомлений
	}
}

func (u *meanWindowUsecase) AddMetricToMonitoring(ctx context.Context, metricName string) error {
	exists, err := u.statsChecker.CheckMetricExists(ctx, metricName)
	if err != nil {
		return fmt.Errorf("failed to check metric: %w", err)
	}
	if !exists {
		return fmt.Errorf("metric %s does not exist", metricName)
	}

	u.monitoredMetricsMux.Lock()
	defer u.monitoredMetricsMux.Unlock()

	if u.monitoredMetrics[metricName] {
		return fmt.Errorf("metric %s already monitored", metricName)
	}

	u.monitoredMetrics[metricName] = true
	return nil
}

func (u *meanWindowUsecase) PauseMetricMonitoring(ctx context.Context, metricName string) error {
	u.monitoredMetricsMux.Lock()
	defer u.monitoredMetricsMux.Unlock()

	if !u.monitoredMetrics[metricName] {
		return fmt.Errorf("metric %s not monitored", metricName)
	}

	u.monitoredMetrics[metricName] = false
	return nil
}

func (u *meanWindowUsecase) ResumeMetricMonitoring(ctx context.Context, metricName string) error {
	u.monitoredMetricsMux.Lock()
	defer u.monitoredMetricsMux.Unlock()

	if _, exists := u.monitoredMetrics[metricName]; !exists {
		return fmt.Errorf("metric %s not monitored", metricName)
	}

	u.monitoredMetrics[metricName] = true
	return nil
}

func (u *meanWindowUsecase) RemoveMetricFromMonitoring(ctx context.Context, metricName string) error {
	u.monitoredMetricsMux.Lock()
	defer u.monitoredMetricsMux.Unlock()

	if _, exists := u.monitoredMetrics[metricName]; !exists {
		return fmt.Errorf("metric %s not monitored", metricName)
	}

	delete(u.monitoredMetrics, metricName)
	return nil
}

func (u *meanWindowUsecase) GetMonitoredMetrics(ctx context.Context) []string {
	u.monitoredMetricsMux.RLock()
	defer u.monitoredMetricsMux.RUnlock()

	metrics := make([]string, 0, len(u.monitoredMetrics))
	for metric, active := range u.monitoredMetrics {
		if active {
			metrics = append(metrics, metric)
		}
	}

	return metrics
}
