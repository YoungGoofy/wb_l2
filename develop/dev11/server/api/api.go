package api

import (
	"fmt"
	"net/http"
	"strconv"

	// "strconv"
	"time"

	"github.com/YoungGoofy/wb_l2/develop/dev11/http/middleware"
	"github.com/YoungGoofy/wb_l2/develop/dev11/http/render"
	"github.com/YoungGoofy/wb_l2/develop/dev11/server/event"
)

type CalendarEvent struct {
	event event.Eventer
}

func NewCalendarEvent(event event.Eventer) *CalendarEvent {
	return &CalendarEvent{event: event}
}

func (ce *CalendarEvent) NewRoute() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", middleware.LogMiddleware(ce.Create))
	mux.HandleFunc("/update_event", middleware.LogMiddleware(ce.Update))
	mux.HandleFunc("/delete_event", middleware.LogMiddleware(ce.Delete))
	mux.HandleFunc("/events_for_day", middleware.LogMiddleware(ce.Get))
	mux.HandleFunc("/events_for_week", middleware.LogMiddleware(ce.Get))
	mux.HandleFunc("/events_for_month", middleware.LogMiddleware(ce.Get))

	return mux
}

func (ce *CalendarEvent) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %v", r.Method), "method should be POST")
		return
	}

	uid := r.FormValue("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user id")
		return
	}

	date := r.FormValue("date")
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use DateOnly format")
		return
	}

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "no title provided")
		return
	}

	note := ce.event.Create(user_id, title, d)
	render.JSON(w, r, http.StatusCreated, note)

}

func (ce *CalendarEvent) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be PUT")
		return
	}

	uid := r.FormValue("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user id")
		return
	}

	eid := r.FormValue("event_id")
	event_id, err := strconv.Atoi(eid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse event id")
		return
	}

	date := r.FormValue("date")
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use DateOnly format")
		return
	}

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "no title provided")
		return
	}

	user, err := ce.event.Update(user_id, event_id, title, d)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "bad request")
		return
	}
	render.JSON(w, r, http.StatusAccepted, user)
}

func (ce *CalendarEvent) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be DELETE")
	}

	uid := r.FormValue("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user id")
		return
	}

	eid := r.FormValue("event_id")
	event_id, err := strconv.Atoi(eid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse event id")
		return
	}
	if err := ce.event.Delete(user_id, event_id); err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't delete note")
		return
	}
	render.NoContent(w, r)
}

func (ce *CalendarEvent) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %v", r.Method), "method should be GET")
	}

	uid := r.FormValue("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user id")
		return
	}

	date := r.FormValue("date")
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use DateOnly format")
		return
	}

	events := make([]event.Note, 0)
	switch r.URL.Path {
	case "/events_for_day":
		events, err = ce.event.GetForDay(user_id, d)
	case "/events_for_week":
		events, err = ce.event.GetForWeek(user_id, d)
	case "/events_for_month":
		events, err = ce.event.GetForMonth(user_id, d)
	}

	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't get events")
		return
	}

	if len(events) == 0 {
		render.NoContent(w, r)
		return
	}

	render.JSON(w, r, http.StatusOK, events)

}
