package middlewares

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// return func(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		entry := f.NewLogEntry(r)
// 		ww := NewWrapResponseWriter(w, r.ProtoMajor)

// 		t1 := time.Now()
// 		defer func() {
// 			entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
// 		}()

// 		next.ServeHTTP(ww, WithLogEntry(r, entry))
// 	}
// 	return http.HandlerFunc(fn)
// }
