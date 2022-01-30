package main

import (
	"L2/11/config"
	"L2/11/request"
	"L2/11/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var (
	eventService = service.NewEventService()
	conf         = config.GetConfig()
)

func main() {
	configureRoutes()
	http.ListenAndServe(conf.ServerConfig.Port, nil)
}

func configureRoutes() {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(res, req)
	}
}

func createEvent(res http.ResponseWriter, req *http.Request) {
	r, err := request.ParseCreateEvent(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = eventService.CreateEvent(r.UserId, r.Date, r.Text)
	if err != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}

func updateEvent(res http.ResponseWriter, req *http.Request) {
	r, err := request.ParseUpdateEvent(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = eventService.UpdateEvent(r.Id, r.UserId, r.Date, r.Text)
	if err != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}

func deleteEvent(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	userIdString := req.FormValue("user_id")
	if userIdString == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	idString := req.FormValue("id")
	if idString == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = eventService.DeleteEvent(userId, id)
	if err != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}

func eventsForDay(res http.ResponseWriter, req *http.Request) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	eventsFor(res, req, start, end)
}

func eventsForWeek(res http.ResponseWriter, req *http.Request) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)
	eventsFor(res, req, start, end)
}

func eventsForMonth(res http.ResponseWriter, req *http.Request) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	eventsFor(res, req, start, end)
}

func eventsFor(res http.ResponseWriter, req *http.Request, start, end time.Time) {
	req.ParseForm()

	userIdString := req.FormValue("user_id")
	if userIdString == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := eventService.UserEventsInRange(userId, start, end)
	if err != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	resp := success(result)
	jsonResp, err := json.Marshal(resp)
	res.Write(jsonResp)
}

func success(result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"result": result,
	}
}

func error(result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error": result,
	}
}
