package service

import (
	"github.com/yuweiweiouo/coding-exercise/internal/dao"
	"github.com/yuweiweiouo/coding-exercise/internal/model"
)

type TaskService interface {
	GetAll() []model.Task
	CreateTask(task model.Task) (model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(id int) error
}

func NewTaskService(taskDao dao.TaskDao) TaskService {
	return &taskService{
		taskDao: taskDao,
	}
}

type taskService struct {
	taskDao dao.TaskDao
}

func (serv *taskService) GetAll() []model.Task {
	return serv.taskDao.GetAll()
}

func (serv *taskService) CreateTask(task model.Task) (model.Task, error) {
	return serv.taskDao.Create(task)
}

func (serv *taskService) UpdateTask(task model.Task) (model.Task, error) {
	return serv.taskDao.Update(task)
}

func (serv *taskService) DeleteTask(id int) error {
	return serv.taskDao.Delete(id)
}
