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

type (
	// TaskID is a unique non-sequential short ID that is
	// is used to identify and process a task.
	TaskID string
	// Task is a model of a todo item. It can be marked as
	// done via the `Done` boolean property.
	Task struct {
		ID       TaskID   `json:"id"`
		Topic    string   `json:"topic"`
		Summary  string   `json:"summary"`
		Priority Priority `json:"priority"`
		Done     bool     `json:"done"`
	}
)

// Do marks this Task as done.
func (t *Task) Do() {
	t.Done = true
}
