package server

import (
	"net/http"
)

func (s *Server) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Infof("Req: %s %s\n", r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
