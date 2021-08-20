package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"task/level2/calendar/internal/calendar"
	"time"
)

func (s *Server) HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, 400, err)
			return
		}
		err = s.calendar.CreateEvent(body)
		if err != nil {
			s.error(w, r, 503, err)
			return
		}

		s.respond(w, r, 200, map[string]string{"result": "ok"})
	}
}
func (s *Server) HandleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, 400, err)
			return
		}
		err = s.calendar.UpdateEvent(body)
		if err != nil {
			s.error(w, r, 503, err)
		}
		s.respond(w, r, 200, map[string]string{"result": "ok"})
	}
}
func (s *Server) HandleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, 400, err)
			return
		}
		err = s.calendar.DelEvent(body)
		if err != nil {
			s.error(w, r, 503, err)
		}
		s.respond(w, r, 200, map[string]string{"result": "ok"})
	}
}
func (s *Server) HandleGetEvent(period string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const layout = "2006-01-02"
		val := r.URL.Query()
		id, ok := val["id"]
		if !ok {
			s.error(w, r, 400, errors.New("no id argument"))
			return
		}
		evnts := make([]calendar.Event, 0)
		for i := range id {
			t, err := time.Parse(layout, id[i])
			ev := make([]calendar.Event, 0)
			if err != nil {
				s.error(w, r, 400, errors.New("wrong date"))
				return
			}
			switch period {
			case "day":
				ev = s.calendar.GetDay(t)
			case "week":
				ev = s.calendar.GetMonth(t)
			case "month":
				ev = s.calendar.GetWeek(t)
			}
			evnts = append(evnts, ev...)
		}
		if len(evnts) == 0 {
			s.error(w, r, 400, errors.New("empty output"))
			return
		}
		s.respond(w, r, 200, map[string][]calendar.Event{"result": evnts})
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.logger.Error(err)
		}
	}

}
