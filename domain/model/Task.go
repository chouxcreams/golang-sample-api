package model

type Task struct {
	Id   int
	Text string
}

func NewTask(text string) (*Task, error) {
	return &Task{
		Text: text,
	}, nil
}

func (t *Task) Set(text string) error {
	t.Text = text
	return nil
}
