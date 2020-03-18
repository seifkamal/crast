package crast

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

func (ll lists) Get(dir string) *List {
	list := ll[dir]
	return &list
}

func (ll *lists) Save(l *List, dir string) {
	(*ll)[dir] = *l
}
