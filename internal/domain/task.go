package domain

import "github.com/google/uuid"

type TaskRepository interface {
	Get(res *[]ResponseTask) (err error)
	Show(ID uuid.UUID, res *ResponseTask) (err error)
	Store(req RequestTask) (err error)
	Update(ID uuid.UUID, req RequestTask) (err error)
	ToggleDone(ID uuid.UUID) (err error)
	Delete(ID uuid.UUID) (err error)
}

type TaskService interface {
	Get(res *[]ResponseTask) (err error)
	Show(ID uuid.UUID, res *ResponseTask) (err error)
	Store(req RequestTask) (err error)
	Update(ID uuid.UUID, req RequestTask) (err error)
	ToggleDone(ID uuid.UUID) (err error)
	Delete(ID uuid.UUID) (err error)
}

type ResponseTask struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Done bool      `json:"done"`
}

type RequestTask struct {
	Name string `json:"name" binding:"required"`
}
