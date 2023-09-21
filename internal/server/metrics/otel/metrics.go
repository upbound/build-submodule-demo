package otel

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.opencensus.io/metric/metricdata"
	opentel "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/upbound/build-submodule-demo/internal/generics"
)

const (
	Success   = "success"
	Malformed = "malformed"
	Rejected  = "rejected"
)

var (
	meter = opentel.GetMeterProvider().Meter("build-submodule-demo")

	reqStarted = generics.Must(meter.Int64Counter("http.request.started.total",
		metric.WithDescription("Total number of http requests started."),
		metric.WithUnit(string(metricdata.UnitDimensionless))))

	reqCompleted = generics.Must(meter.Int64Counter("http.request.completed.total",
		metric.WithDescription("Total number of http requests completed."),
		metric.WithUnit(string(metricdata.UnitDimensionless))))

	reqDuration = generics.Must(meter.Float64Histogram("http.request.duration.ms",
		metric.WithDescription("Time between receiving and responding to an http request."),
		metric.WithUnit(string(metricdata.UnitMilliseconds))))

	productMetricSubmitted = generics.Must(meter.Int64Counter("prodmetric.submitted",
		metric.WithDescription("Total number of product metrics submitted."),
		metric.WithUnit(string(metricdata.UnitDimensionless))))
)

// Middleware records metrics for HTTP handlers.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqStarted.Add(r.Context(), 1, metric.WithAttributes(HTTPServerMetricAttributesFromHTTPRequest(r)...))
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		t1 := time.Now()
		defer func() {
			reqCompleted.Add(context.Background(), 1, metric.WithAttributes(HTTPServerMetricAttributesFromHTTPResponse(r, ww)...))
			reqDuration.Record(context.Background(), float64(time.Since(t1).Milliseconds()), metric.WithAttributes(HTTPServerMetricAttributesFromHTTPResponse(r, ww)...))
		}()
		next.ServeHTTP(ww, r)
	})
}

// HTTPServerMetricAttributesFromHTTPRequest constructs default attributes for
// an HTTP request.
func HTTPServerMetricAttributesFromHTTPRequest(r *http.Request) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("http.method", r.Method),
		attribute.String("http.host", r.Host),
		attribute.Bool("http.tls", r.TLS != nil),
	}
}

// HTTPServerMetricAttributesFromHTTPResponse constructs default attributes for
// an HTTP response.
func HTTPServerMetricAttributesFromHTTPResponse(r *http.Request, w middleware.WrapResponseWriter) []attribute.KeyValue {
	return append(HTTPServerMetricAttributesFromHTTPRequest(r), attribute.Int("http.status_code", w.Status()))
}

// ProductMetricSubmit records an product metric submission.
func ProductMetricSubmit(ctx context.Context, account, repository string, success bool) {
	productMetricSubmitted.Add(ctx, 1, metric.WithAttributes([]attribute.KeyValue{
		attribute.String("prodmetric.account", account),
		attribute.String("prodmetric.repository", repository),
		attribute.Bool("prodmetric.success", success),
	}...))
}
