package models

import (
	"time"

	"github.com/torcuata22/rest_api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

// Methods:
func (e Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	events = append(events, e)
	if err != nil {
		return err
	}

	defer stmt.Close()
	id, err := result.LastInsertId()

	e.ID = id
	events = append(events, e)
	return err

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

//.Exec() is used to execute a query that changes things
//.Query() is used to run a query that returns several rows of data
//.QueryRow returns only one row (there is only one row for id)
