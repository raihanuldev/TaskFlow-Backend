package database

// Task Struct

type Task struct {
	ID       int    `json:"id"`
	TaskName string `json:"task_name"`
	Details  string `json:"details"`
	Status   string `json:"status"`
}

var Stores []Task








func init() {
	Stores = append(Stores,
		Task{
			ID:       1,
			TaskName: "Learn Go Basics",
			Details:  "Understand variables, loops, and functions",
			Status:   "todo",
		},
		Task{
			ID:       2,
			TaskName: "Build REST API",
			Details:  "Practice with net/http and JSON handling",
			Status:   "in-progress",
		},
		Task{
			ID:       3,
			TaskName: "Read Go Concurrency",
			Details:  "Learn goroutines and channels",
			Status:   "pending",
		},
		Task{
			ID:       4,
			TaskName: "Implement Task Update Feature",
			Details:  "Add update and delete logic to task API",
			Status:   "done",
		},
	)
}
