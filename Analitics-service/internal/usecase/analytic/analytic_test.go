package analytic_usecase

import (
	"context"
	"testing"
	"time"

	"service/pkg/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mock_repo "service/internal/usecase/analytic/mocks"
)

func TestWriteStat(t *testing.T) {
	mockRepo := mock_repo.NewStatsRepository(t)
	uc := NewAnalyticUsecase(mockRepo)

	ctx := context.Background()
	testMetric := "test_metric"
	now := time.Now()

	values := []decimal.Decimal{
		decimal.NewFromInt(10),
		decimal.NewFromInt(20),
		decimal.NewFromInt(30),
	}

	testCases := []struct {
		name        string
		stat        model.Stat
		expectedErr error
	}{
		{
			name: "successful write",
			stat: model.Stat{
				Name:  testMetric,
				Value: values[0],
				Time:  now,
			},
			expectedErr: nil,
		},
		{
			name: "repository error",
			stat: model.Stat{
				Name:  testMetric,
				Value: values[1],
				Time:  now,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "empty metric name",
			stat: model.Stat{
				Name:  "",
				Value: values[2],
				Time:  now,
			},
			expectedErr: model.ErrEmptyMetricName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.stat.Name != "" {
				mockRepo.EXPECT().
					WriteStat(ctx, tc.stat).
					Return(tc.expectedErr)
			}
			err := uc.WriteStat(ctx, tc.stat)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetLastNStats(t *testing.T) {
	mockRepo := mock_repo.NewStatsRepository(t)
	uc := NewAnalyticUsecase(mockRepo)

	ctx := context.Background()
	testMetric := "test_metric"
	now := time.Now()

	testCases := []struct {
		name        string
		n           int
		expected    []model.Stat
		expectedErr error
	}{
		{
			name: "successful get last 2 stats",
			n:    2,
			expected: []model.Stat{
				{Name: testMetric, Value: decimal.NewFromInt(30), Time: now.Add(-1 * time.Minute)},
				{Name: testMetric, Value: decimal.NewFromInt(20), Time: now.Add(-3 * time.Minute)},
			},
			expectedErr: nil,
		},
		{
			name:        "repository error",
			n:           3,
			expected:    nil,
			expectedErr: assert.AnError,
		},
		{
			name:        "invalid n value",
			n:           0,
			expected:    nil,
			expectedErr: model.ErrNMustBePositive,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.n > 0 {
				mockRepo.EXPECT().
					GetLastNStats(ctx, testMetric, tc.n).
					Return(tc.expected, tc.expectedErr)
			}

			got, err := uc.GetLastNStats(ctx, testMetric, tc.n)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}

func TestGetStatsInRange(t *testing.T) {
	mockRepo := mock_repo.NewStatsRepository(t)
	uc := NewAnalyticUsecase(mockRepo)

	ctx := context.Background()
	testMetric := "test_metric"
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	testCases := []struct {
		name        string
		start       time.Time
		end         time.Time
		expected    []model.Stat
		expectedErr error
	}{
		{
			name:  "successful get in range",
			start: start,
			end:   end,
			expected: []model.Stat{
				{Name: testMetric, Value: decimal.NewFromInt(10), Time: start.Add(30 * time.Minute)},
				{Name: testMetric, Value: decimal.NewFromInt(20), Time: start.Add(45 * time.Minute)},
			},
			expectedErr: nil,
		},
		{
			name:        "repository error",
			start:       start,
			end:         end,
			expected:    nil,
			expectedErr: assert.AnError,
		},
		{
			name:        "invalid time range",
			start:       end,
			end:         start,
			expected:    nil,
			expectedErr: model.ErrWrongTimeInterval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.start.Before(tc.end) {
				mockRepo.EXPECT().
					GetStatsInRange(ctx, testMetric, tc.start, tc.end).
					Return(tc.expected, tc.expectedErr)
			}

			got, err := uc.GetStatsInRange(ctx, testMetric, tc.start, tc.end)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}
