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
	WriteStat(ctx context.Context, stat model.Stat) error
}

//go:generate go run github.com/vektra/mockery/v2 --name StatsRepository --with-expecter
type StatsRepository interface {
	GetLastNStats(ctx context.Context, metricName string, n int) ([]model.Stat, error)
	GetStatsInRange(ctx context.Context, metricName string, start, end time.Time) ([]model.Stat, error)
	CheckMetricExists(ctx context.Context, metricName string) (bool, error)
	WriteStat(ctx context.Context, stat model.Stat) error
}

type analyticUsecase struct {
	repo stats_repo.Repository
}

func NewAnalyticUsecase(repo stats_repo.Repository) AnalyticUsecase {
	return &analyticUsecase{repo: repo}
}

func (u *analyticUsecase) WriteStat(ctx context.Context, stat model.Stat) error {
	if stat.Name == "" {
		return model.ErrEmptyMetricName
	}
	return u.repo.WriteStat(ctx, stat)
}

func (a *analyticUsecase) GetLastNStats(ctx context.Context, metricName string, n int) ([]model.Stat, error) {
	if n <= 0 {
		return nil, model.ErrNMustBePositive
	}
	return a.repo.GetLastNStats(ctx, metricName, n)
}

func (a *analyticUsecase) GetStatsInRange(ctx context.Context, metricName string, start, end time.Time) ([]model.Stat, error) {
	if start.After(end) {
		return nil, model.ErrWrongTimeInterval
	}
	return a.repo.GetStatsInRange(ctx, metricName, start, end)
}
