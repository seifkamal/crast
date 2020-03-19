package crast

import "path/filepath"

type List []Task

func (l *List) Add(t *Task) {
	*l = append(*l, *t)
}

func (l *List) Do(id int) {
	(*l)[id].Done = true
}

func (l *List) Remove(id int) {
	*l = append((*l)[:id], (*l)[id+1:]...)
}

type lists map[string]List

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

func (ll lists) Has(dir string) (exists bool) {
	_, exists = ll[dir]
	return
}
