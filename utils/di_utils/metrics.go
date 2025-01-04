package di_utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/teadove/teasutils/utils/logger_utils"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/settings_utils"
)

func runMetricsFromSettingsInBackground(ctx context.Context, container Container) {
	go func() {
		err := runMetrics(ctx, settings_utils.BaseSettings.Metrics.URL, container)
		if err != nil {
			panic(fmt.Sprintf("failed to run metrics http api: %v", err))
		}
	}()
}

func runMetrics(ctx context.Context, url string, container Container) error {
	promHandler := promhttp.Handler()

	server := http.NewServeMux()
	server.HandleFunc("/metrics/", func(writer http.ResponseWriter, request *http.Request) {
		innerCtx := logger_utils.AddLoggerToCtx(request.Context())

		innerCtx, cancel := context.WithTimeout(
			innerCtx,
			settings_utils.BaseSettings.Metrics.RequestTimeout,
		)
		defer cancel()

		promHandler.ServeHTTP(writer, request.WithContext(innerCtx))
	})
	server.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		innerCtx := logger_utils.AddLoggerToCtx(request.Context())

		errFromChecker := checkFromCheckers(innerCtx, container.HealthCheckers())
		if errFromChecker == nil {
			writer.WriteHeader(http.StatusOK)

			_, err := writer.Write([]byte("ok"))
			if err != nil {
				zerolog.Ctx(innerCtx).Error().Stack().Err(err).Msg("failed.to.write.ok")
			}

			return
		}

		writer.WriteHeader(http.StatusServiceUnavailable)

		_, err := writer.Write([]byte(errFromChecker.Error()))
		if err != nil {
			zerolog.Ctx(innerCtx).Error().Stack().Err(err).Msg("failed.to.write.errs")
		}
	})

	zerolog.Ctx(ctx).Debug().Str("url", url).Msg("starting.metrics.server")

	//nolint: gosec // fixed
	err := http.ListenAndServe(url, server)
	if err != nil {
		return errors.Wrap(err, "cannot start metrics server")
	}

	return nil
}
