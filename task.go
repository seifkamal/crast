package crast

// Task is a model of a todo item. It can be marked as
// done via the `Done` boolean property.
type Task struct {
	Topic   string `json:"topic"`
	Summary string `json:"summary"`
	Done    bool   `json:"done"`
}
