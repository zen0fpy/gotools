package middlewares

import (
	"net/http"
	"plough/pkg/constant"
	"plough/pkg/util"
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
