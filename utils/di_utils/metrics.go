package di_utils

import (
	"context"
	"fmt"
	"net/http"
	"telemetry-go/presentation_utils/rest_presentation_utils"
	"telemetry-go/utils/settings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

func runMetricsFromSettingsInBackground(ctx context.Context, container Container) {
	go func() {
		err := runMetrics(ctx, settings.Settings.Metrics.URL, container)
		if err != nil {
			panic(fmt.Sprintf("failed to run metrics http api: %v", err))
		}
	}()
}

func runMetrics(ctx context.Context, url string, container Container) error {
	app := rest_presentation_utils.NewGinApp(ctx)
	promHandler := promhttp.Handler()

	app.Any("/metrics/", func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	})

	app.Any("/health", func(c *gin.Context) {
		errs := container.Health(ctx)

		if len(errs) == 0 {
			c.JSON(http.StatusOK, gin.H{"success": true})

			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"success": false, "errors": errs})
	})

	zerolog.Ctx(ctx).Info().Str("url", url).Msg("starting.metrics.server")

	err := app.Run(url)
	if err != nil {
		return errors.Wrap(err, "cannot start metrics server")
	}

	return nil
}
