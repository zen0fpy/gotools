package middlewares

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func Tracer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parentSpan opentracing.Span
		tracer := opentracing.GlobalTracer()
		spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(r.RequestURI)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				r.RequestURI,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
			defer parentSpan.Finish()
		}

		ctx := opentracing.ContextWithSpan(r.Context(), parentSpan)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
