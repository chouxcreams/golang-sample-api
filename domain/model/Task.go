package model

type Task struct {
	Id        int
	Text      *string
	Completed bool
}

func NewTask(text *string) (*Task, error) {
	return &Task{
		Text:      text,
		Completed: false,
	}, nil
}

func (t *Task) Set(text *string) {
	t.Text = text
}

func (t *Task) Complete() {
	t.Completed = true
}
