package request

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

type UpdateEvent struct {
	Id     int
	UserId int
	Date   time.Time
	Text   string
}

func ParseUpdateEvent(req *http.Request) (UpdateEvent, error) {
	result := UpdateEvent{}
	req.ParseForm()

	idString := req.FormValue("id")
	if idString == "" {
		return result, errors.New("BadRequest")
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return result, errors.New("BadRequest")
	}

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
	result.Id = id
	result.UserId = userId
	result.Date = date
	result.Text = text
	return result, nil
}
