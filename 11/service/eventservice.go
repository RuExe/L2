package service

import (
	"L2/11/domain"
	"errors"
	"time"
)

type EventService struct {
	events map[int]map[int]domain.Event
}

func NewEventService() EventService {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 32, 0, 0, time.UTC)
	return EventService{
		events: map[int]map[int]domain.Event{
			1: {
				1: domain.NewEvent(1, time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-4, 32, 0, 0, time.UTC), "Оч важное мероприятие 1"),
				2: domain.NewEvent(1, date, "Оч важное мероприятие 2"),
				3: domain.NewEvent(1, date.AddDate(0, 0, 6), "Оч важное мероприятие 3"),
				4: domain.NewEvent(1, date.AddDate(0, 0, 16), "Оч важное мероприятие 4"),
				5: domain.NewEvent(1, date.AddDate(0, 1, 0), "Оч важное мероприятие 5"),
			},
		},
	}
}

func (s *EventService) CreateEvent(userId int, date time.Time, text string) (int, error) {
	event := domain.NewEvent(userId, date, text)
	s.events[userId][event.Id] = event
	return event.Id, nil
}

func (s *EventService) UpdateEvent(id, userId int, date time.Time, text string) error {
	userEvents, ok := s.events[userId]
	if !ok {
		return errors.New("current user doesn't exist")
	}

	event, ok := userEvents[id]
	if !ok {
		return errors.New("current event doesn't exist")
	}

	s.events[userId][id] = domain.Event{
		Id:     event.Id,
		UserId: event.UserId,
		Date:   date,
		Text:   text,
	}
	return nil
}

func (s *EventService) DeleteEvent(userId, id int) error {
	userEvents, ok := s.events[userId]
	if !ok {
		return errors.New("current user doesn't exist")
	}

	if _, ok = userEvents[id]; !ok {
		return errors.New("current event doesn't exist")
	}

	delete(userEvents, id)
	return nil
}

func (s *EventService) UserEventsInRange(userId int, start, end time.Time) ([]domain.Event, error) {
	res := make([]domain.Event, 0)
	userEvents, ok := s.events[userId]
	if !ok {
		return res, errors.New("current user doesn't exist")
	}

	for _, v := range userEvents {
		if v.Date.After(start) && v.Date.Before(end) {
			res = append(res, v)
		}
	}
	return res, nil
}
