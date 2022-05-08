package service

import (
	"reflect"
	"testing"

	"github.com/yuweiweiouo/coding-exercise/internal/model"
	"github.com/yuweiweiouo/coding-exercise/mock"
)

func Test_taskService_GetAll(t *testing.T) {
	mockDao := &mock.TaskDao{}
	serv, _ := CreateTaskService(mockDao)

	tests := []struct {
		name string
		serv TaskService
		want []model.Task
	}{
		{
			name: "取得所有任務",
			serv: serv,
			want: []model.Task{
				{
					Id:     1,
					Name:   "買早餐",
					Status: 1,
				},
				{
					Id:     2,
					Name:   "買午餐",
					Status: 1,
				},
				{
					Id:     3,
					Name:   "買晚餐",
					Status: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDao.On("GetAll").Return(tt.want)
			if got := tt.serv.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_CreateTask(t *testing.T) {
	mockDao := &mock.TaskDao{}
	serv, _ := CreateTaskService(mockDao)

	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		serv    TaskService
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "加任務",
			serv: serv,
			args: args{
				model.Task{
					Name: "買水果",
				},
			},
			want: model.Task{
				Id:     1,
				Name:   "買水果",
				Status: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDao.On("Create", tt.args.task).Return(tt.want, nil)
			got, err := tt.serv.CreateTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskService.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_UpdateTask(t *testing.T) {
	mockDao := &mock.TaskDao{}
	serv, _ := CreateTaskService(mockDao)

	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		serv    TaskService
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "更新任務",
			serv: serv,
			args: args{
				model.Task{
					Id:     3,
					Name:   "買晚餐",
					Status: 1,
				},
			},
			want: model.Task{
				Id:     3,
				Name:   "買晚餐",
				Status: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDao.On("Update", tt.args.task).Return(tt.want, nil)
			got, err := tt.serv.UpdateTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskService.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.UpdateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_DeleteTask(t *testing.T) {
	mockDao := &mock.TaskDao{}
	serv, _ := CreateTaskService(mockDao)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		serv    TaskService
		args    args
		wantErr bool
	}{
		{
			name: "刪除任務",
			serv: serv,
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDao.On("Delete", tt.args.id).Return(nil)
			if err := tt.serv.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("taskService.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
