package crast

type Task struct {
	Topic   string `json:"topic"`
	Summary string `json:"summary"`
	Done    bool   `json:"done"`
}
