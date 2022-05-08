package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yuweiweiouo/coding-exercise/internal/model"
	"github.com/yuweiweiouo/coding-exercise/internal/model/request"
	"github.com/yuweiweiouo/coding-exercise/internal/service"
	"go.uber.org/zap"
)

type TaskController interface {
	All(*gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type taskController struct {
	baseController
	service service.TaskService
}

func NewTaskController(l *zap.Logger, s service.TaskService) TaskController {
	ctl := &taskController{
		service: s,
	}
	ctl.logger = l
	return ctl
}

func (ctl taskController) All(ctx *gin.Context) {
	ctl.success(ctx, http.StatusOK, ctl.service.GetAll())
}

func (ctl taskController) Create(ctx *gin.Context) {
	var payload request.CreateTask
	if err := ctx.BindJSON(&payload); err != nil {
		ctl.error(ctx, http.StatusBadRequest, ErrInvaildData)
		return
	}
	if err := validator.New().Struct(&payload); err != nil {
		ctl.error(ctx, http.StatusBadRequest, ErrInvaildData)
		return
	}

	task := model.Task{
		Name:   payload.Name,
		Status: payload.Status,
	}

	task, err := ctl.service.CreateTask(task)
	if err != nil {
		ctl.error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctl.success(ctx, http.StatusCreated, task)
}

func (ctl taskController) Update(ctx *gin.Context) {
	var payload request.UpdateTask
	if err := ctx.BindJSON(&payload); err != nil {
		ctl.error(ctx, http.StatusBadRequest, ErrInvaildData)
		return
	}
	if err := validator.New().Struct(&payload); err != nil {
		ctl.error(ctx, http.StatusBadRequest, ErrInvaildData)
		return
	}

	task := model.Task{
		Id:     payload.Id,
		Name:   payload.Name,
		Status: payload.Status,
	}

	task, err := ctl.service.UpdateTask(task)
	if err != nil {
		ctl.error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctl.success(ctx, http.StatusOK, task)
}

func (ctl taskController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctl.error(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctl.service.DeleteTask(id); err != nil {
		ctl.error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctl.success(ctx, http.StatusOK, "success")
}
