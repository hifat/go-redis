package service

import (
	"go-redis/internal/domain"
	"log"

	"github.com/google/uuid"
)

type taskService struct {
	taskRepo domain.TaskRepository
}

func NewTaskService(taskRepo domain.TaskRepository) domain.TaskService {
	return &taskService{taskRepo}
}

func (u taskService) Get(res *[]domain.ResponseTask) (err error) {
	if err := u.taskRepo.Get(res); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Show(ID uuid.UUID, res *domain.ResponseTask) (err error) {
	if err := u.taskRepo.Show(ID, res); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Store(req domain.RequestTask) (err error) {
	if err := u.taskRepo.Store(req); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Update(ID uuid.UUID, req domain.RequestTask) (err error) {
	if err := u.taskRepo.Update(ID, req); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) ToggleDone(ID uuid.UUID) (err error) {
	if err := u.taskRepo.ToggleDone(ID); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Delete(ID uuid.UUID) (err error) {
	if err := u.taskRepo.Delete(ID); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
