package storage

import (
	"dev11/internal/event"
	"time"
)

type Repository interface {
	CreateEvent(e *event.Event) (int, error)
	UpdateEvent(e *event.Event) error
	DeleteEvent(id int) error
	Events(start, end time.Time) ([]*event.Event, error)
}
