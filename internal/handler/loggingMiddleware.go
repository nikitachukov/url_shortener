package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/nikitachukov/url_shortener.git/internal/logger"
)

func LoggingHandlersMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		newResponseWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(newResponseWriter, r.WithContext(r.Context()))

		logger.Log.Sugar().Infof("uri [%s] method [%s] duration [%v]ms status [%v] size [%d]",
			r.RequestURI,
			r.Method,
			strconv.FormatInt(time.Since(start).Milliseconds(), 10),
			newResponseWriter.Status(),
			newResponseWriter.BytesWritten(),
		)
	}

	return http.HandlerFunc(fn)

}
