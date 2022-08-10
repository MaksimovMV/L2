package memory

import (
	"dev11/internal/event"
	"fmt"
	"sync"
	"time"
)

type Repository struct {
	events map[int]*event.Event
	index  int
	sync.Mutex
}

func NewStorage() Repository {
	return Repository{make(map[int]*event.Event), 0, sync.Mutex{}}
}

func (r *Repository) CreateEvent(e *event.Event) (int, error) {
	r.Lock()
	defer r.Unlock()

	r.index++
	e.ID = r.index
	r.events[r.index] = e
	return r.index, nil
}

func (r *Repository) UpdateEvent(e *event.Event) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.events[e.ID]; !ok {
		return fmt.Errorf("событие отсутствует")
	}
	r.events[e.ID] = e
	return nil
}

func (r *Repository) DeleteEvent(id int) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.events[id]; !ok {
		return fmt.Errorf("событие отсутствует")
	}
	delete(r.events, id)
	return nil
}

func (r *Repository) Events(start, end time.Time) ([]*event.Event, error) {
	var result []*event.Event
	for _, e := range r.events {
		if e.Date.After(start) && e.Date.Before(end) {
			result = append(result, e)
		}
	}
	return result, nil
}
