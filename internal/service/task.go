package service

import (
	"context"
	"encoding/json"
	"log"
	"redigo/internal/domain/taskDomain"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type taskService struct {
	taskRepo    taskDomain.TaskRepository
	redisClient *redis.Client
}

func NewTaskService(taskRepo taskDomain.TaskRepository, redisClient *redis.Client) taskDomain.TaskService {
	return &taskService{
		taskRepo,
		redisClient,
	}
}

func (u taskService) MockData(amount int) (err error) {
	if err = u.taskRepo.MockData(amount); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Get(res *[]taskDomain.ResponseTask) (err error) {
	key := "repo:GetTasks"

	// Redis get
	tasksJson, err := u.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(tasksJson), &res)
		if err == nil {
			return err
		}
	}

	if err = u.taskRepo.Get(res); err != nil {
		log.Println(err.Error())
		return err
	}

	// Redis set
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = u.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Show(ID uuid.UUID, res *taskDomain.ResponseTask) (err error) {
	if err := u.taskRepo.Show(ID, res); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Store(req taskDomain.RequestTask) (err error) {
	if err := u.taskRepo.Store(req); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u taskService) Update(ID uuid.UUID, req taskDomain.RequestTask) (err error) {
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
