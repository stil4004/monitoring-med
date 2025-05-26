package analytic_usecase

import (
	"context"
	stats_repo "service/internal/repository/stats"
	"service/pkg/model"
	"time"
)

type AnalyticUsecase interface {
	GetLastNStats(ctx context.Context, metricName string, n int) ([]model.Stat, error)
	GetStatsInRange(ctx context.Context, metricName string, start, end time.Time) ([]model.Stat, error)
}

type analyticUsecase struct {
	repo stats_repo.Repository
}

func NewAnalyticUsecase(repo stats_repo.Repository) AnalyticUsecase {
	return &analyticUsecase{repo: repo}
}

func (a *analyticUsecase) GetLastNStats(ctx context.Context, metricName string, n int) ([]model.Stat, error) {
	return a.repo.GetLastNStats(ctx, metricName, n)
}

func (a *analyticUsecase) GetStatsInRange(ctx context.Context, metricName string, start, end time.Time) ([]model.Stat, error) {
	return a.repo.GetStatsInRange(ctx, metricName, start, end)
}
