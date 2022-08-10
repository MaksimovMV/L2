package apiserver

import (
	"dev11/internal/event"
	"dev11/internal/storage"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	Router     *http.ServeMux
	repository storage.Repository
}

type message struct {
	Message string `json:"message"`
}

type errMessage struct {
	ErrMessage string `json:"error"`
}

func NewServer(repository storage.Repository) *server {
	s := &server{
		Router:     http.NewServeMux(),
		repository: repository,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.Router.HandleFunc("/create_event", s.createEvent)
	s.Router.HandleFunc("/update_event", s.updateEvent)
	s.Router.HandleFunc("/delete_event", s.deleteEvent)
	s.Router.HandleFunc("/events_for_day", s.eventsForDay)
	s.Router.HandleFunc("/events_for_week", s.eventsForWeek)
}

func (s *server) createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e, err := s.parseParamsFromJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	id, err := s.repository.CreateEvent(&e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	result := message{"Event id = " + strconv.Itoa(id) + " was created"}
	b, _ := json.Marshal(result)
	_, _ = w.Write(b)
	w.WriteHeader(http.StatusCreated)
}

func (s *server) updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e, err := s.parseParamsFromJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	if err := s.repository.UpdateEvent(&e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	result := message{"Event was updated"}
	b, _ := json.Marshal(result)
	_, _ = w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (s *server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	type params struct {
		ID int `json:"id"`
	}
	p := params{0}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(content, &p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	if err := s.repository.DeleteEvent(p.ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	result := message{"Event was deleted"}
	b, _ := json.Marshal(result)
	_, _ = w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (s *server) eventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e, err := s.parseParamsFromQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}

	start, end := getDay(e.Date)

	events, err := s.repository.Events(start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	b, _ := json.Marshal(events)
	b, _ = json.Marshal(message{string(b)})
	_, _ = w.Write(b)
}

func (s *server) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e, err := s.parseParamsFromQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}

	start, end := getWeek(e.Date)

	events, err := s.repository.Events(start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	b, _ := json.Marshal(events)
	b, _ = json.Marshal(message{string(b)})
	_, _ = w.Write(b)
}

func (s *server) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	e, err := s.parseParamsFromQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}

	start, end := getMonth(e.Date)

	events, err := s.repository.Events(start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(errMessage{err.Error()})
		_, _ = w.Write(b)
		return
	}
	b, _ := json.Marshal(events)
	b, _ = json.Marshal(message{string(b)})
	_, _ = w.Write(b)
}

func (s *server) parseParamsFromJSON(r *http.Request) (event.Event, error) {

	type temp struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Date string `json:"date"`
	}
	t := temp{0, "", ""}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return event.Event{}, err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(content, &t); err != nil {
		return event.Event{}, err
	}

	date, err := time.Parse("2006-01-02", t.Date)
	if err != nil {
		return event.Event{}, err
	}
	e := event.Event{
		ID:   t.ID,
		Name: t.Name,
		Date: date,
	}
	return e, nil
}

func (s *server) parseParamsFromQuery(r *http.Request) (event.Event, error) {
	//http://localhost:8080/create_event?date=2019-09-09&name=shopping%20with%20friends
	var e event.Event
	var err error
	if date := r.URL.Query().Get("date"); date != "" {
		e.Date, err = time.Parse("2006-01-02", date)
		if err != nil {
			return event.Event{}, err
		}
	}

	e.Name = r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	if id != "" {
		e.ID, err = strconv.Atoi(id)
		if err != nil {
			return event.Event{}, err
		}
	}
	return e, nil
}

func getWeek(tm time.Time) (time.Time, time.Time) {
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour), currentZeroDay.Add((7 - weekday) * 24 * time.Hour)
}

func getDay(tm time.Time) (time.Time, time.Time) {
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay, currentZeroDay.Add(24 * time.Hour)
}

func getMonth(tm time.Time) (time.Time, time.Time) {
	year, month, _ := tm.Date()
	firstMonthsDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastMonthsDay := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	return firstMonthsDay, lastMonthsDay
}
