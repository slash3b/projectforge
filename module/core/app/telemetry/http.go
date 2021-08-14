package telemetry

import (
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var _ propagation.TextMapCarrier = (*headerCarrier)(nil)

type headerCarrier struct {
	h *fasthttp.RequestHeader
}

func (hc headerCarrier) Get(key string) string {
	return string(hc.h.Peek(key))
}

func (hc headerCarrier) Set(key string, value string) {
	hc.h.Set(key, value)
}

func (hc headerCarrier) Keys() []string {
	var keys []string
	hc.h.VisitAll(func(key []byte, _ []byte) {
		keys = append(keys, string(key))
	})
	return keys
}

func ExtractHeaders(ctx *fasthttp.RequestCtx, logger *zap.SugaredLogger) *fasthttp.RequestCtx {
	nc := otel.GetTextMapPropagator().Extract(ctx, headerCarrier{h: &ctx.Request.Header})
	http, ok := nc.(*fasthttp.RequestCtx)
	if !ok {
		logger.Warnf("unable to extract http tracing headers")
		return ctx
	}
	return http
}

func InjectHTTP(ctx *fasthttp.RequestCtx, span trace.Span) {
	span.SetAttributes(
		semconv.HTTPHostKey.String(string(ctx.Request.Host())),
		semconv.HTTPMethodKey.String(string(ctx.Method())),
		semconv.HTTPURLKey.String(string(ctx.Request.RequestURI())),
		semconv.HTTPSchemeKey.String(string(ctx.Request.URI().Scheme())),
	)
	if b := ctx.Request.Header.Peek("User-Agent"); len(b) > 0 {
		span.SetAttributes(semconv.HTTPUserAgentKey.String(string(b)))
	}
	if b := ctx.Request.Header.Peek("Content-Length"); len(b) > 0 {
		span.SetAttributes(semconv.HTTPRequestContentLengthKey.String(string(b)))
	}
	span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(ctx.Response.StatusCode())...)
	span.SetStatus(semconv.SpanStatusFromHTTPStatusCode(ctx.Response.StatusCode()))
}
