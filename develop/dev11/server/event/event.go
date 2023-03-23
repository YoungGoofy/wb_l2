package event

import (
	"fmt"
	"sync"
	"time"
)

type Eventer interface {
	Create(id int, title string, date time.Time) Note
	Update(userId, noteId int, title string, date time.Time) (Note, error)
	Delete(userId, noteId int) error
	GetForDay(id int, date time.Time) ([]Note, error)
	GetForWeek(id int, date time.Time) ([]Note, error)
	GetForMonth(id int, date time.Time) ([]Note, error)
}

type User struct {
	Id    int `json:"id"`
	Notes []Note
}

type Note struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type taskStore struct {
	sync.Mutex
	Users  map[int][]Note
	NextId int `json:"next_id"`
}

func New() Eventer {
	return &taskStore{
		Users:  make(map[int][]Note),
		NextId: 0}
}

func (ts *taskStore) Create(id int, title string, date time.Time) Note {
	ts.Lock()
	var nextId int
	defer ts.Unlock()
	if len(ts.Users[id]) == 0 {
		nextId = 0
	} else {
		nextId = ts.Users[id][len(ts.Users[id])-1].Id + 1
	}
	note := Note{
		Id:    nextId,
		Title: title,
		Date:  date,
	}
	ts.Users[id] = append(ts.Users[id], note)
	// nextId++
	return note
}

func (ts *taskStore) Update(userId, noteId int, title string, date time.Time) (Note, error) {
	ts.Lock()
	defer ts.Unlock()
	notes, ok := ts.Users[userId]
	if !ok {
		return Note{}, fmt.Errorf("user with id %v does not exist", userId)
	}

	noteIndex := -1
	for i, note := range notes {
		if note.Id == noteId {
			noteIndex = i
			break
		}
	}

	if noteIndex == -1 {
		return Note{}, fmt.Errorf("note with id %v does not exist for user with id %v", noteId, userId)
	}
	notes[noteId].Title = title
	notes[noteId].Date = date

	ts.Users[userId] = notes

	return notes[noteId], nil
}

func (ts *taskStore) Delete(userId, noteId int) error {
	ts.Lock()
	defer ts.Unlock()
	notes, ok := ts.Users[userId]
	if !ok {
		return fmt.Errorf("user with id %v does not exist", userId)
	}

	noteIndex := -1
	for i, note := range notes {
		if note.Id == noteId {
			noteIndex = i
			break
		}
	}

	if noteIndex == -1 {
		return fmt.Errorf("note with id %v does not exist for user with id %v", noteId, userId)
	}

	notes = append(notes[:noteIndex], notes[noteIndex+1:]...)
	ts.Users[userId] = notes
	return nil
}

func (ts *taskStore) GetForDay(id int, date time.Time) ([]Note, error) {
	ts.Lock()
	defer ts.Unlock()

	notes := make([]Note, 0)

	user, ok := ts.Users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %v does not exist", id)
	}

	for _, note := range user {
		if note.Date == date {
			notes = append(notes, note)
		}
	}
	return notes, nil
}

func (ts *taskStore) GetForWeek(id int, date time.Time) ([]Note, error) {
	ts.Lock()
	defer ts.Unlock()

	user, ok := ts.Users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %v does not exist", id)
	}

	startOfWeek := date.Truncate(24*time.Hour).AddDate(0, 0, -int(date.Weekday())+1)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var notes []Note

	for _, note := range user {
		if note.Date.After(startOfWeek) && note.Date.Before(endOfWeek) {
			notes = append(notes, note)
		}
	}
	return notes, nil
}

func (ts *taskStore) GetForMonth(id int, date time.Time) ([]Note, error) {
	ts.Lock()
	defer ts.Unlock()

	user, ok := ts.Users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %v does not exist", id)
	}

	year, month, _ := date.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Создаем слайс для хранения заметок за месяц
	var notes []Note

	for _, note := range user {
		for i := 0; i < 30; i++ {
			if note.Date.After(startOfMonth) && note.Date.Before(endOfMonth) {
				notes = append(notes, note)
			}
		}
	}
	return notes, nil
}
