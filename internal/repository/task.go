package repository

import (
	"go-redis/internal/domain"
	"go-redis/internal/model/gormModel"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db}
}

func (r taskRepository) Get(res *[]domain.ResponseTask) (err error) {
	return r.db.Model(&gormModel.Task{}).Find(res).Error
}

func (r taskRepository) Show(ID uuid.UUID, res *domain.ResponseTask) (err error) {
	return r.db.Model(&gormModel.Task{}).Where("id = ?", ID).First(&res).Error
}

func (r taskRepository) Store(req domain.RequestTask) (err error) {
	newTask := gormModel.Task{
		Name: req.Name,
	}

	return r.db.Create(&newTask).Error
}

func (r taskRepository) Update(ID uuid.UUID, req domain.RequestTask) (err error) {
	editTask := gormModel.Task{
		Name: req.Name,
	}

	return r.db.Where("id = ?", ID).Updates(editTask).Error
}

func (r taskRepository) ToggleDone(ID uuid.UUID) (err error) {
	return r.db.Model(&gormModel.Task{}).Where("id = ?", ID).Update("done", gorm.Expr("NOT done")).Error
}

func (r taskRepository) Delete(ID uuid.UUID) (err error) {
	return r.db.Model(&gormModel.Task{}).Delete("id = ?", ID).Error
}
