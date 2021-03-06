package prismaRepository

import (
	"context"
	"golang-sample-api/domain/model"
	"golang-sample-api/domain/repository"
	"golang-sample-api/prisma/db"
)

type TaskPrismaRepository struct {
	Client *db.PrismaClient
}

func NewTaskRepository(client *db.PrismaClient) repository.TaskRepository {
	return &TaskPrismaRepository{Client: client}
}

// Create taskの保存
func (tr *TaskPrismaRepository) Create(task *model.Task) error {
	_, err := tr.Client.Task.CreateOne(
		db.Task.Text.SetIfPresent(task.Text),
	).Exec(context.Background())
	return err
}

// FindById taskをIDで取得
func (tr *TaskPrismaRepository) FindById(id int) (*model.Task, error) {
	task, err := tr.Client.Task.FindUnique(
		db.Task.ID.Equals(id),
	).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return toDomainModel(task), nil
}

// Update taskの更新
func (tr *TaskPrismaRepository) Update(task *model.Task) (*model.Task, error) {
	_, err := tr.Client.Task.FindUnique(
		db.Task.ID.Equals(task.Id),
	).Update(
		db.Task.Text.SetIfPresent(task.Text),
		db.Task.Completed.Set(task.Completed),
	).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Delete taskの削除
func (tr *TaskPrismaRepository) Delete(task *model.Task) error {
	_, err := tr.Client.Task.FindUnique(
		db.Task.ID.Equals(task.Id),
	).Delete().Exec(context.Background())
	return err
}

// FindAll taskの一覧の取得
func (tr *TaskPrismaRepository) FindAll() ([]model.Task, error) {
	tasks, err := tr.Client.Task.FindMany().Exec(context.Background())
	if err != nil {
		return nil, err
	}
	var modelTasks []model.Task
	for _, task := range tasks {
		modelTasks = append(modelTasks, *toDomainModel(&task))
	}
	return modelTasks, nil
}

func toDomainModel(task *db.TaskModel) *model.Task {
	text, ok := task.Text()
	var modelText *string
	if !ok {
		modelText = nil
	} else {
		modelText = &text
	}
	return &model.Task{
		Id:        task.ID,
		Text:      modelText,
		Completed: task.Completed,
	}
}
