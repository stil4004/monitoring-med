package cErrors

import "errors"

type ResponseErrorModel struct {
	InternalCode int64  `json:"internal_error_code,omitempty"`
	StandartCode int64  `json:"error_code,omitempty"`
	Message      string `json:"message,omitempty"`
	Stage        string `json:"stage,omitempty"`
}

// InternalCode
const (
	AddQuestionError        int = 1000
	DeleteQuestionByIDError int = 1001
	GetQuestionByIDError    int = 1002
	GetQuestionsByTypeError int = 1003

	AddUnresolvedQError              int = 2001
	UpdateUnresolvedQError           int = 2002
	DeleteUnresolvedQByIDError       int = 2003
	GetUnresolvedQByIDError          int = 2004
	GetUnresolvedQsByClientNameError int = 2005
	GetUnresolvedQsByTypeError       int = 2006
)

var (
	ErrNoSuchProjectType = errors.New("no such project type")
)

const (
	GetStatsError         = "get stats error"
	WriteStatError        = "get write error"
	MetricMonitoringError = "metric monitoring error"
)
