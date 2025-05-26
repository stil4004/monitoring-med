package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Stat struct {
	Name  string          `json:"name"`
	Time  time.Time       `json:"time"`
	Value decimal.Decimal `json:"value"`
}

type WriteStatRequest struct {
	Name  string          `json:"name"`
	Value decimal.Decimal `json:"value"`
	Time  time.Time       `json:"time,omitempty"`
}

func (r WriteStatRequest) ToStat() Stat {
	if r.Time.IsZero() {
		r.Time = time.Now()
	}
	return Stat{
		Name:  r.Name,
		Time:  r.Time,
		Value: r.Value,
	}
}
