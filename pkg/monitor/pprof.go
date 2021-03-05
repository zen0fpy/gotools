package monitor

import (
	"github.com/tal-tech/go-zero/rest"
	"net/http"
	"net/http/pprof"
)

func pprofHandler(h http.HandlerFunc) http.HandlerFunc {

	handler := http.HandlerFunc(h)
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func PprofRoutes() []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/",
			Handler: pprof.Index,
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/symbol",
			Handler: pprofHandler(pprof.Symbol),
		},
		{
			Method:  http.MethodPost,
			Path:    "/debug/pprof/symbol",
			Handler: pprofHandler(pprof.Symbol),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/cmdline",
			Handler: pprof.Cmdline,
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/profile",
			Handler: pprof.Profile,
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/trace",
			Handler: pprof.Trace,
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/allocs",
			Handler: pprofHandler(pprof.Handler("allocs").ServeHTTP),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/goroutine",
			Handler: pprofHandler(pprof.Handler("goroutine").ServeHTTP),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/heap",
			Handler: pprofHandler(pprof.Handler("heap").ServeHTTP),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/mutex",
			Handler: pprofHandler(pprof.Handler("mutex").ServeHTTP),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/block",
			Handler: pprofHandler(pprof.Handler("block").ServeHTTP),
		},
		{
			Method:  http.MethodGet,
			Path:    "/debug/pprof/threadcreate",
			Handler: pprofHandler(pprof.Handler("threadcreate").ServeHTTP),
		},
	}
}
