package http_app

import (
	"context"
	"errors"
	"fmt"
	"service/config"
	meanwindow_http "service/internal/handler/http/alerting/avg"
	analytic_http "service/internal/handler/http/analytics"
	stats_repo "service/internal/repository/stats"
	analytic_usecase "service/internal/usecase/analytic"
	meanwindow_usecase "service/internal/usecase/meanWindow"

	"service/pkg/clients/victoriametrics"
	"service/pkg/ctx_log"

	"time"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewHttpApp(isTest bool) *fx.App {
	return fx.New(
		fx.Provide(
			// Logger
			SetupLogger,
			func(logger *zap.Logger) (context.Context, context.CancelFunc) {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				return ctx_log.ContextWithLogger(ctx, logger), cancel
			},

			// Set shield
			// func() *secure.Shield {
			// 	if isTest {
			// 		return secure.NewShield(os.Getenv("AE_TEST"))
			// 	}
			// 	return secure.NewShield(os.Getenv("AE_KEY"))
			// },

			// Config
			func() (*config.Config, error) {
				var (
					vp  *viper.Viper
					err error
				)

				vp, err = config.LoadConfig(isTest)
				if err != nil {
					return nil, err
				}

				cfg, err := config.ParseConfig(vp)
				if err != nil {
					return nil, err
				}

				// config.DecryptConfig(cfg, key)
				return cfg, nil
			},

			func(cfg *config.Config) (victoriametrics.Client, error) {
				client := victoriametrics.New(cfg.VictoriaMetrics.BaseURL)
				if client == nil {
					return nil, errors.New("can't start client")
				}
				return client, nil
			},

			// Init DB
			// storage.NewDB,

			// clients here

			// Repositories
			stats_repo.New,

			// Usecases
			analytic_usecase.NewAnalyticUsecase,
			meanwindow_usecase.NewMeanWindowUsecase,

			// Handler
			analytic_http.NewAnalyticHandler,
			meanwindow_http.NewMeanWindowHandler,

			func() *fiber.App {
				return fiber.New(fiber.Config{DisableStartupMessage: true})
			},
		),
		fx.Invoke(
			// func(ctx context.Context, qa qa_service.Usecase) {
			// 	go qa.Cache(ctx)
			// },
			// func(ctx context.Context, un un_service.Usecase) {
			// 	go un.Cache(ctx)
			// },
			func(app *fiber.App) {
				app.Use(otelfiber.Middleware())
				app.Use(cors.New(cors.ConfigDefault))
			},

			// Init routes
			analytic_http.MapAnalyticRoutes,
			meanwindow_http.MapMeanWindowRoutes,

			func(
				lc fx.Lifecycle,
				ctx_ context.Context,
				logger *zap.Logger,
				cancel context.CancelFunc,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return nil
					},
					OnStop: func(ctx context.Context) error {
						cancel()
						return nil
					},
				})
			},
			func(
				lc fx.Lifecycle,
				cfg *config.Config,
				log *zap.Logger,
				app *fiber.App,
				cancel context.CancelFunc,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							log.Info("Start server ",
								zap.String("host", cfg.Server.Host),
								zap.String("port", cfg.Server.Port),
							)
							if err := app.Listen(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)); err != nil {
								log.Error("server can't listen", zap.String("error: ", err.Error()))
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						log.Sync()
						cancel()
						return app.Shutdown()
					},
				})
			},
		),
	)
}

func SetupLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "time",
			LevelKey:         "level",
			NameKey:          "logger",
			CallerKey:        "caller",
			MessageKey:       "msg",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			EncodeTime:       zapcore.TimeEncoderOfLayout(time.DateTime),
			EncodeDuration:   zapcore.SecondsDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			SkipLineEnding:   false,
			ConsoleSeparator: " | ",
		},
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
		//InitialFields:     map[string]interface{}{"app": "controller_trx"},
	}

	logger, err := config.Build(zap.WithCaller(false)) // 1 кадр в глубину стека - этого вполне достаточно
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}
