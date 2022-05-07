package dao

import (
	"errors"

	"github.com/yuweiweiouo/coding-exercise/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrTaskNotExist = errors.New("任務不存在")
)

type TaskDao interface {
	GetAll() []model.Task
	Create(task model.Task) (model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id int) error
}

type taskDao struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewTaskDao(db *gorm.DB, l *zap.Logger) TaskDao {
	return &taskDao{
		db:     db,
		logger: l,
	}
}

func (dao taskDao) GetAll() []model.Task {
	result := make([]model.Task, 0, 100)
	if err := dao.db.Find(&result).Error; err != nil {
		dao.logger.Error(err.Error())
	}
	return result
}

func (dao taskDao) Create(task model.Task) (model.Task, error) {
	return task, dao.db.Create(&task).Error
}

func (dao taskDao) Update(task model.Task) (model.Task, error) {
	result := dao.db.Model(&task).
		Omit("id").
		Updates(task)

	if result.Error != nil {
		return task, result.Error
	}

	if result.RowsAffected == 0 {
		return task, ErrTaskNotExist
	}

	return task, nil
}

func (dao taskDao) Delete(id int) error {
	result := dao.db.Delete(&model.Task{}, id)

	if result.RowsAffected == 0 {
		return ErrTaskNotExist
	}

	return result.Error
}
