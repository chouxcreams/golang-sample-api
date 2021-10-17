package usecase

import (
	"golang-sample-api/domain/model"
	"golang-sample-api/domain/repository"
)

// TaskUsecase task usecaseのinterface
type TaskUsecase interface {
	Create(text *string) error
	FindByID(id int) (*model.Task, error)
	Update(id int, text *string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

// NewTaskUsecase task usecaseのコンストラクタ
func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

// Create taskを保存するときのユースケース
func (tu *taskUsecase) Create(text *string) error {
	task, err := model.NewTask(text)
	if err != nil {
		return err
	}

	return tu.taskRepo.Create(task)
}

// FindByID taskをIDで取得するときのユースケース
func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

// Update taskを更新するときのユースケース
func (tu *taskUsecase) Update(id int, text *string) (*model.Task, error) {
	targetTask, err := tu.taskRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	targetTask.Set(text)

	updatedTask, err := tu.taskRepo.Update(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

// Delete taskを削除するときのユースケース
func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepo.FindById(id)
	if err != nil {
		return err
	}

	return tu.taskRepo.Delete(task)
}
