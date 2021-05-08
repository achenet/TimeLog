// The data_store package store data on the different tasks.
package data_store

type TaskList map[string]Task

type Task struct {
	Name string // Should always be the same as the key for this task in the map.
}

func (tl TaskList) CheckThatAllNamesAreCorrect() bool {
	for taskName, task := range tl {
		if task.Name != taskName {
			return false
		}
	}
	return true
}
