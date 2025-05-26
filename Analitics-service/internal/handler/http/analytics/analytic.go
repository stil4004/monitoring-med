package analytic_http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"service/pkg/ctx_log"

	"service/internal/cErrors"
	"service/internal/cconstants"
	analytic_usecase "service/internal/usecase/analytic"
	"service/pkg/model"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AnalyticHandler struct {
	analyticUC analytic_usecase.AnalyticUsecase
}

func NewAnalyticHandler(analyticUC analytic_usecase.AnalyticUsecase) *AnalyticHandler {
	return &AnalyticHandler{
		analyticUC: analyticUC,
	}
}

// Response structures
type (
	StatsResponse struct {
		Status string       `json:"status"`
		Data   []model.Stat `json:"data"`
	}

	ErrorResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	SuccessResponse struct {
		Status string `json:"status"`
	}
)

func (h *AnalyticHandler) GetLastNStats(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			err error
		)

		metricName := c.Query("metric")
		if metricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetLastNStats"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.GetStatsError,
			})
		}

		nStr := c.Query("n", "10")
		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetLastNStats"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "invalid n parameter"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "n must be positive integer",
				Code:    cErrors.GetStatsError,
			})
		}

		stats, err := h.analyticUC.GetLastNStats(ctx, metricName, n)
		if err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetLastNStats"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.GetStatsError,
			})
		}

		return c.Status(http.StatusOK).JSON(StatsResponse{
			Status: "success",
			Data:   stats,
		})
	}
}

func (h *AnalyticHandler) WriteStat(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.WriteStatRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		if req.Name == "" || req.Value.IsZero() {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "name and value are required",
			})
		}

		if err := h.analyticUC.WriteStat(ctx, req.ToStat()); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to write metric",
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": "success",
		})
	}
}

func (h *AnalyticHandler) GetStatsInRange(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			err error
		)

		metricName := c.Query("metric")
		if metricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetStatsInRange"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.GetStatsError,
			})
		}

		startStr := c.Query("start")
		endStr := c.Query("end", time.Now().Format(time.RFC3339))

		start, err := time.Parse(time.RFC3339, startStr)
		if err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetStatsInRange"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "invalid start time format"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "start time must be RFC3339 format",
				Code:    cErrors.GetStatsError,
			})
		}

		end, err := time.Parse(time.RFC3339, endStr)
		if err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetStatsInRange"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "invalid end time format"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "end time must be RFC3339 format",
				Code:    cErrors.GetStatsError,
			})
		}

		if end.Before(start) {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetStatsInRange"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "end time must be after start time"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "end time must be after start time",
				Code:    cErrors.GetStatsError,
			})
		}

		stats, err := h.analyticUC.GetStatsInRange(ctx, metricName, start, end)
		if err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "GetStatsInRange"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.GetStatsError,
			})
		}

		return c.Status(http.StatusOK).JSON(StatsResponse{
			Status: "success",
			Data:   stats,
		})
	}
}
