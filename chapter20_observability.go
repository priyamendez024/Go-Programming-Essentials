// Chapter 20: Logging, Metrics & Distributed Tracing
package main

import (
    "context"
    "fmt"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.uber.org/zap"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func initTracer() func() {
    exp, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
    tp := trace.NewTracerProvider(trace.WithBatcher(exp))
    otel.SetTracerProvider(tp)
    return func() { _ = tp.Shutdown(context.Background()) }
}

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    logger.Info("Service starting")

    httpRequests := prometheus.NewCounter(prometheus.CounterOpts{
        Name: "http_requests_total", Help: "Total HTTP requests",
    })
    prometheus.MustRegister(httpRequests)
    http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":2112", nil)

    tracer := otel.Tracer("my-service")
    ctx, span := tracer.Start(context.Background(), "main")
    fmt.Println("Tracing span:", span.Name())
    span.End()
}
