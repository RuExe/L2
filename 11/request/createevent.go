package request

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

type CreateEvent struct {
	UserId int
	Date   time.Time
	Text   string
}

func ParseCreateEvent(req *http.Request) (CreateEvent, error) {
	result := CreateEvent{}
	req.ParseForm()

	userIdString := req.FormValue("user_id")
	if userIdString == "" {
		return result, errors.New("BadRequest")
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return result, errors.New("BadRequest")
	}

	dateString := req.FormValue("date")
	if dateString == "" {
		return result, errors.New("BadRequest")
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return result, errors.New("BadRequest")
	}

	text := req.FormValue("text")
	if text == "" {
		return result, errors.New("BadRequest")
	}
	result.UserId = userId
	result.Date = date
	result.Text = text
	return result, nil
}
