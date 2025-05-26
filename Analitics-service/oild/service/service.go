package service

import (
	"errors"
	"sync"
)

type MetricData struct {
	Name   string
	Values []float64
}

type MonitoringService struct {
	metrics    map[string]bool
	mu         sync.RWMutex
	alertRules map[string]AlertRule
}

type AlertRule struct {
	Threshold float64
	Direction string // "above" or "below"
}

func NewMonitoringService() *MonitoringService {
	return &MonitoringService{
		metrics:    make(map[string]bool),
		alertRules: make(map[string]AlertRule),
	}
}

// AddMetric добавляет новую метрику для мониторинга
func (s *MonitoringService) AddMetric(name string) error {
	if name == "" {
		return errors.New("metric name cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.metrics[name] {
		return errors.New("metric already exists")
	}

	s.metrics[name] = true
	return nil
}

// RemoveMetric удаляет метрику из мониторинга
func (s *MonitoringService) RemoveMetric(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.metrics[name] {
		return errors.New("metric not found")
	}

	delete(s.metrics, name)
	delete(s.alertRules, name)
	return nil
}

// SetAlertRule устанавливает правило алертинга для метрики
func (s *MonitoringService) SetAlertRule(metric string, threshold float64, direction string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.metrics[metric] {
		return errors.New("metric not monitored")
	}

	s.alertRules[metric] = AlertRule{
		Threshold: threshold,
		Direction: direction,
	}
	return nil
}

// CheckForAnomalies проверяет метрики на аномалии
func (s *MonitoringService) CheckForAnomalies(metricData MetricData) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.metrics[metricData.Name] {
		return false, errors.New("metric not monitored")
	}

	rule, exists := s.alertRules[metricData.Name]
	if !exists {
		return false, nil
	}

	for _, value := range metricData.Values {
		if rule.Direction == "above" && value > rule.Threshold {
			return true, nil
		}
		if rule.Direction == "below" && value < rule.Threshold {
			return true, nil
		}
	}

	return false, nil
}

// GetActiveMetrics возвращает список активных метрик
func (s *MonitoringService) GetActiveMetrics() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	metrics := make([]string, 0, len(s.metrics))
	for metric := range s.metrics {
		metrics = append(metrics, metric)
	}
	return metrics
}
