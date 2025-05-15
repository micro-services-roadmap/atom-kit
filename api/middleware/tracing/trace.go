package tracing

import (
	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
	"github.com/micro-services-roadmap/atom-kit/logx/rich"
	"github.com/micro-services-roadmap/atom-kit/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer oteltrace.Tracer

func init() {
	tracer = sdktrace.NewTracerProvider().Tracer(tracing.TraceName)
}

// TraceHandler trace request
func TraceHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		propagator := otel.GetTextMapPropagator()

		spanName := c.Request.URL.Path
		ctx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		spanCtx, span := tracer.Start(
			ctx,
			spanName,
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest("wpp", spanName, c.Request)...),
		)
		defer span.End()

		c.Request = c.Request.WithContext(spanCtx)
		c.Set("tlog", rich.WithLogger(kg.L).WithContext(spanCtx))
		// convenient for tracking error messages
		propagator.Inject(c, propagation.HeaderCarrier(c.Writer.Header()))

		c.Next()
	}
}
