package analytic_http

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func MapAnalyticRoutes(ctx context.Context, router *fiber.App, h *AnalyticHandler) (bool, error) {
	analyticGR := router.Group("/analytics")

	// Stats routes
	analyticGR.Get("/stats/last", h.GetLastNStats(ctx))
	analyticGR.Get("/stats/range", h.GetStatsInRange(ctx))
	router.Post("/stats", h.WriteStat(ctx))

	return true, nil
}
