package repository

import (
	"fmt"
	"redigo/internal/domain/taskDomain"
	"redigo/internal/model/gormModel"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) taskDomain.TaskRepository {
	return &taskRepository{db}
}

func (r taskRepository) MockData(amount int) (err error) {
	var count int64
	if err = r.db.Model(&gormModel.Task{}).Count(&count).Error; err != nil {
		return err
	}

	tasks := []gormModel.Task{}
	for i := 1; i <= amount; i++ {
		tasks = append(tasks, gormModel.Task{
			Name: fmt.Sprintf("Task %d", i),
		})
	}

	return r.db.Create(&tasks).Error
}

func (r taskRepository) Get(res *[]taskDomain.ResponseTask) (err error) {
	return r.db.Model(&gormModel.Task{}).Order("created_at ASC").Find(res).Error
}

func (r taskRepository) Show(ID uuid.UUID, res *taskDomain.ResponseTask) (err error) {
	return r.db.Model(&gormModel.Task{}).Where("id = ?", ID).First(&res).Error
}

func (r taskRepository) Store(req taskDomain.RequestTask) (err error) {
	newTask := gormModel.Task{
		Name: req.Name,
	}

	return r.db.Create(&newTask).Error
}

func (r taskRepository) Update(ID uuid.UUID, req taskDomain.RequestTask) (err error) {
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
