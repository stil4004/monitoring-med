package meanwindow_http

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func MapMeanWindowRoutes(ctx context.Context, router *fiber.App, h *MeanWindowHandler) (bool, error) {
	monitoringGR := router.Group("/monitoring")

	// Monitoring routes
	monitoringGR.Post("/add", h.AddMetricToMonitoring(ctx))
	monitoringGR.Post("/pause", h.PauseMetricMonitoring(ctx))
	monitoringGR.Post("/resume", h.ResumeMetricMonitoring(ctx))
	monitoringGR.Delete("/remove", h.RemoveMetricFromMonitoring(ctx))
	monitoringGR.Get("/metrics", h.GetMonitoredMetrics(ctx))

	return true, nil
}
