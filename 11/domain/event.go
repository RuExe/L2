package domain

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Event struct {
	Id     int
	UserId int
	Date   time.Time
	Text   string
}

func NewEvent(userId int, date time.Time, text string) Event {
	return Event{
		Id:     rand.Int(),
		UserId: userId,
		Date:   date,
		Text:   text,
	}
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId int
		Date   string
		Text   string
	}{
		e.UserId,
		e.Date.Format("2006/1/2"),
		e.Text,
	})
}
