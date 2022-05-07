package request

type CreateTask struct {
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status" validate:"oneof=0 1"`
}

type UpdateTask struct {
	Name   string `json:"name" validate:"required"`
	Status int    `json:"status" validate:"oneof=0 1"`
}
