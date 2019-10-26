package model

type ID string

func NewID(id string) ID {
	return ID(id)
}

func (i ID) String() string {
	return string(i)
}

type Title string

func NewTitle(title string) Title {
	return Title(title)
}

func (t Title) String() string {
	return string(t)
}

type Status int

func (s Status) ToInt() int {
	return int(s)
}

const (
	DOING Status = iota
	DONE
)

type Task struct {
	id     ID
	title  Title
	status Status
}

func NewTask(id string, title string, status Status) *Task {
	i := NewID(id)
	t := NewTitle(title)
	return &Task{id: i, title: t, status: status}
}

func (t *Task) SetStatus(status Status) {
	t.status = status
}

func (t *Task) ID() ID {
	return t.id
}

func (t *Task) Title() Title {
	return t.title
}

func (t *Task) Status() Status {
	return t.status
}
