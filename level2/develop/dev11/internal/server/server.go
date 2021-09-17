package server

import (
	"net/http"
	"task/level2/calendar/config"
	"task/level2/calendar/internal/calendar"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router   *mux.Router
	logger   *logrus.Logger
	calendar *calendar.Calendar
	serv     *http.Server
}

func NewServer(cfg *config.Config) *Server {
	serv := Server{

		router:   mux.NewRouter(),
		logger:   logrus.New(),
		calendar: calendar.NewCalendar(&cfg.Calendar),
		serv: &http.Server{
			Addr:         cfg.Server.Addr,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	serv.Configue()
	serv.serv.Handler = serv.router
	return &serv
}

func (s *Server) Start(listenerr *chan error) {
	go func() {
		*listenerr <- s.serv.ListenAndServe()
	}()
	s.logger.Info("service start")
}
func (s *Server) Stop() {
	s.serv.Close()
	s.logger.Info("service down")
}
func (s *Server) Configue() {
	s.router.HandleFunc("/create_event", s.HandleCreate()).Methods("POST")
	s.router.HandleFunc("/update_event", s.HandleUpdate()).Methods("POST")
	s.router.HandleFunc("/delete_event", s.HandleDelete()).Methods("POST")
	s.router.HandleFunc("/events_for_day", s.HandleGetEvent("day")).Methods("GET")
	s.router.HandleFunc("/events_for_week", s.HandleGetEvent("week")).Methods("GET")
	s.router.HandleFunc("/events_for_month", s.HandleGetEvent("month")).Methods("GET")
	s.router.Use(s.Middleware)
}
