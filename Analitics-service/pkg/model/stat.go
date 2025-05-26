package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Stat struct {
	Name  string
	Time  time.Time
	Value decimal.Decimal
}
