package request

type CreateTask struct {
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status" validate:"oneof=0 1"`
}

type UpdateTask struct {
	Id     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status" validate:"oneof=0 1"`
}
