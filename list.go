package crast

import (
	"path/filepath"
	"sort"
)

// List is a slice of Task items. Items can be added and marked as
// done via a List.
type List []Task

// Add appends an item to the list.
func (l *List) Add(t *Task) {
	*l = append(*l, *t)
}

// Get retrieves a Task item that corresponds to the given ID.
func (l List) Get(id TaskID) *Task {
	for _, task := range l {
		if task.ID == id {
			return &task
		}
	}

	return nil
}

// Update replaces an existing task with the given one, provided
// they both have the same ID.
func (l *List) Update(updatedTask *Task) {
	for i, task := range *l {
		if task.ID == updatedTask.ID {
			(*l)[i] = *updatedTask
		}
	}
}

// Do marks any items in the list that correspond to the given IDs
// as done.
func (l *List) Do(ids ...TaskID) {
	for _, id := range ids {
		for i, task := range *l {
			if task.ID == id {
				(*l)[i].Do()
			}
		}
	}
}

// Undo resets the Done flag on any items in the list that
// correspond to the given IDs.
func (l *List) Undo(ids ...TaskID) {
	for _, id := range ids {
		for i, task := range *l {
			if task.ID == id {
				(*l)[i].Undo()
			}
		}
	}
}

// Remove removes an item from the list with an ID matching the
// one given.
func (l *List) Remove(id TaskID) {
	for i, task := range *l {
		if task.ID == id {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
}

// Clear removes all items from the list.
func (l *List) Clear() {
	*l = []Task{}
}

// ByPriority returns a copy of the list sorted by priority.
func (l List) ByPriority() List {
	sort.Slice(l, func(a, b int) bool {
		return l[a].Priority < l[b].Priority
	})

	return l
}

type lists map[string]List

// Add adds an entry to the lists map with the given directory
// as the key. If an entry already exists, it will be overwritten.
func (ll *lists) Add(dir string, list *List) {
	(*ll)[dir] = *list
}

// Get retrieves the list associated with the given directory.
// The search is done recursively over the directory path; If
// no entry is found, the next (parent) directory is used, and
// so on, until one is found. Otherwise, the root directory
// ("/") is returned.
func (ll lists) Get(dir string) (*List, string) {
	for !ll.Has(dir) && dir != "/" {
		dir = filepath.Dir(dir)
	}

	list := ll[dir]
	return &list, dir
}

// Has checks whether the given directory has its own entry
// in the lists map, and returns a boolean value with the
// result.
func (ll lists) Has(dir string) (exists bool) {
	_, exists = ll[dir]
	return
}

// Remove removes a list matching the given directory from
// the list map.
func (ll *lists) Remove(dir string) {
	delete(*ll, dir)
}
