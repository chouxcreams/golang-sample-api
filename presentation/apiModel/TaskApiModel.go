package apiModel

import "golang-sample-api/domain/model"

type TaskCommand struct {
	Text *string `json:"text"`
}

type TaskResponse struct {
	Id        int     `json:"id"`
	Text      *string `json:"text"`
	Completed bool    `json:"completed"`
}

func NewTaskResponse(task *model.Task) TaskResponse {
	return TaskResponse{
		Id:        task.Id,
		Text:      task.Text,
		Completed: task.Completed,
	}
}
