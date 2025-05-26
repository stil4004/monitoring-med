package meanwindow_http

import (
	"context"
	"net/http"

	"service/pkg/ctx_log"
	"service/pkg/utils"

	"service/internal/cErrors"
	"service/internal/cconstants"
	meanwindow_usecase "service/internal/usecase/meanWindow"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MeanWindowHandler struct {
	meanWindowUC meanwindow_usecase.UseCase
}

func NewMeanWindowHandler(meanWindowUC meanwindow_usecase.UseCase) *MeanWindowHandler {
	return &MeanWindowHandler{
		meanWindowUC: meanWindowUC,
	}
}

// Response structures
type (
	MonitoredMetricsResponse struct {
		Status  string   `json:"status"`
		Metrics []string `json:"metrics"`
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

func (h *MeanWindowHandler) AddMetricToMonitoring(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			req struct {
				MetricName string `json:"metric_name"`
			}
		)

		if err := utils.ReadRequest(c, &req); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "AddMetricToMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.BadRequest,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if req.MetricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "AddMetricToMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if err := h.meanWindowUC.AddMetricToMonitoring(ctx, req.MetricName); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "AddMetricToMonitoring"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		return c.Status(http.StatusOK).JSON(SuccessResponse{
			Status: cconstants.OK,
		})
	}
}

func (h *MeanWindowHandler) PauseMetricMonitoring(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			req struct {
				MetricName string `json:"metric_name"`
			}
		)

		if err := utils.ReadRequest(c, &req); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "PauseMetricMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.BadRequest,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if req.MetricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "PauseMetricMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if err := h.meanWindowUC.PauseMetricMonitoring(ctx, req.MetricName); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "PauseMetricMonitoring"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		return c.Status(http.StatusOK).JSON(SuccessResponse{
			Status: cconstants.OK,
		})
	}
}

func (h *MeanWindowHandler) ResumeMetricMonitoring(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			req struct {
				MetricName string `json:"metric_name"`
			}
		)

		if err := utils.ReadRequest(c, &req); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "ResumeMetricMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.BadRequest,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if req.MetricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "ResumeMetricMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if err := h.meanWindowUC.ResumeMetricMonitoring(ctx, req.MetricName); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "ResumeMetricMonitoring"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		return c.Status(http.StatusOK).JSON(SuccessResponse{
			Status: cconstants.OK,
		})
	}
}

func (h *MeanWindowHandler) RemoveMetricFromMonitoring(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			l   = ctx_log.LoggerFromContext(ctx)
			req struct {
				MetricName string `json:"metric_name"`
			}
		)

		if err := utils.ReadRequest(c, &req); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "RemoveMetricFromMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.BadRequest,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if req.MetricName == "" {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "RemoveMetricFromMonitoring"),
				zap.String("STAGE", cconstants.Handler),
				zap.String("ERROR", "metric name is required"))

			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  "error",
				Message: "metric name is required",
				Code:    cErrors.MetricMonitoringError,
			})
		}

		if err := h.meanWindowUC.RemoveMetricFromMonitoring(ctx, req.MetricName); err != nil {
			l.Error("INTERNAL_ERROR",
				zap.String("METHOD", "RemoveMetricFromMonitoring"),
				zap.String("ERROR", err.Error()))

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  "error",
				Message: cconstants.InnerError,
				Code:    cErrors.MetricMonitoringError,
			})
		}

		return c.Status(http.StatusOK).JSON(SuccessResponse{
			Status: cconstants.OK,
		})
	}
}

func (h *MeanWindowHandler) GetMonitoredMetrics(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		metrics := h.meanWindowUC.GetMonitoredMetrics(ctx)

		return c.Status(http.StatusOK).JSON(MonitoredMetricsResponse{
			Status:  "success",
			Metrics: metrics,
		})
	}
}
