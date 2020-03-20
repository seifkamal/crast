package crast

// Priority level of a task.
type Priority int

// Valid priority levels.
const (
	P1 Priority = iota + 1
	P2
	P3
	P4
)

// IsValid checks whether the attached priority level
// is a valid value (see Priority).
func (p Priority) IsValid() bool {
	switch p {
	case P1, P2, P3, P4:
		return true
	default:
		return false
	}
}

// Task is a model of a todo item. It can be marked as
// done via the `Done` boolean property.
type Task struct {
	Topic    string   `json:"topic"`
	Summary  string   `json:"summary"`
	Priority Priority `json:"priority"`
	Done     bool     `json:"done"`
}

// Do marks this Task as done.
func (t *Task) Do() {
	t.Done = true
}
