package middlewares

import (
	"gotools/pkg/constant"
	"gotools/pkg/util"
	"net/http"
)

func RequestId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(constant.RequestId)
		if requestId == "" {
			requestId = util.GenUUID()
		}
		r.Header.Set(constant.RequestId, requestId)
		w.Header().Set(constant.RequestId, requestId)
		next(w, r)
	}
}
