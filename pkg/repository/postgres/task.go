package postgres

import (
	"toDoApp/pkg/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) model.TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(t *model.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO task (user_id, title, description, is_important, due_date, created_date, is_done) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		t.UserID, t.Title, t.Description, t.IsImportant, t.DueDate, t.CreatedDate, t.IsDone)
	return err
}

func (r *TaskRepository) GetAll(userID int64) ([]model.Task, error) {
	rows, err := r.db.Query("SELECT title, description, is_important, is_done, due_date, created_date FROM task WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	allTasks := []model.Task{}

	for rows.Next() {
		t := model.Task{}
		if err := rows.Scan(
			&t.Title, &t.Description, &t.IsImportant, &t.IsDone, &t.DueDate, &t.CreatedDate); err != nil {
			continue
		}

		allTasks = append(allTasks, t)
	}
	return allTasks, nil
}

func (r *TaskRepository) GetByID(taskID int64) (*model.Task, error) {
	task := &model.Task{}

	err := r.db.QueryRow("SELECT user_id, title, description, is_important, is_done, due_date, created_date FROM task WHERE id = $1", taskID).Scan(
		&task.UserID, &task.Title, &task.Description, &task.IsImportant, &task.IsDone, &task.DueDate, &task.CreatedDate)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) Update(t *model.Task) error {
	_, err := r.db.Exec("UPDATE task SET title = $1, description = $2, is_important = $3, due_date = $4, is_done = $5 WHERE id = $6",
		t.Title, t.Description, t.IsImportant, t.DueDate, t.IsDone, &t.ID)

	return err
}

func (r *TaskRepository) Delete(taskID int64) error {
	_, err := r.db.Exec("DELETE FROM task WHERE id = $1", taskID)
	return err
}
