package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) model.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) Create(e *model.Event) error {
	_, err := r.db.Exec("INSERT INTO events (user_id, name, description, appointed_date) VALUES ($1, $2, $3, $4)",
		e.UserID, e.Name, e.Description, e.AppointedDate)

	return err
}

func (r *EventRepository) GetAll(userID int64) ([]model.Event, error) {
	rows, err := r.db.Query("SELECT name, description, appointed_date FROM events WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	allEvents := []model.Event{}
	for rows.Next() {
		e := model.Event{}
		err := rows.Scan(&e.Name, &e.Description, &e.AppointedDate)
		if err != nil {
			continue
		}
		allEvents = append(allEvents, e)
	}
	return allEvents, nil
}

func (r *EventRepository) GetByID(eventID int64) (*model.Event, error) {
	e := &model.Event{}
	err := r.db.QueryRow("SELECT name, description, appointed_date FROM events WHERE id = $1", eventID).Scan(
		&e.Name, &e.Description, &e.AppointedDate)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *EventRepository) Update(e *model.Event) error {
	_, err := r.db.Exec("UPDATE events SET name = $1, description = $2, appointed_date = $3 WHERE id = $4",
		e.Name, e.Description, e.AppointedDate, e.ID)
	return err
}

func (r *EventRepository) Delete(eventID int64) error {
	_, err := r.db.Exec("DELETE FROM events WHERE id = $1", eventID)
	return err
}
