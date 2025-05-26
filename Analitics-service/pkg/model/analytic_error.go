package model

import "errors"

var (
	ErrWrongTimeInterval = errors.New("start time must be before end time")
	ErrEmptyMetricName   = errors.New("metric name cannot be empty")
	ErrNMustBePositive   = errors.New("n must be positive")
)
