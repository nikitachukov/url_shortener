package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func HandlersLog(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		newResponseWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(newResponseWriter, r.WithContext(r.Context()))
		duration := time.Since(start)

		log.Printf("uri [%s] method [%s] duration [%v]ms status [%v] size [%d]",
			r.RequestURI,
			r.Method,
			strconv.FormatInt(duration.Milliseconds(), 10),
			newResponseWriter.Status(),
			newResponseWriter.BytesWritten(),
		)
	}

	return http.HandlerFunc(fn)

}
