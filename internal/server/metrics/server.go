package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/go-chi/chi/v5"
	chimid "github.com/go-chi/chi/v5/middleware"

	"github.com/upbound/build-submodule-demo/internal"
	"github.com/upbound/build-submodule-demo/internal/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	opentel "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Server serves the metrics API.
func Server(opts internal.MetricsOptions, logger logging.Logger) (*http.Server, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter), metric.WithResource(resource.NewSchemaless(attribute.String("service.name", "build-submodule-demo"))))

	// Set prometheus exporter as global meter provider to allow access from
	// other packages.
	opentel.SetMeterProvider(provider)

	mr := chi.NewRouter()
	mr.Use(chimid.RequestLogger(&log.Formatter{Log: logger}))
	mr.Handle("/metrics", promhttp.Handler())
	return &http.Server{
		Handler:           mr,
		Addr:              fmt.Sprintf(":%d", opts.MetricsPort),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}, nil
}
