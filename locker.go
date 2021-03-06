package crast

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Locker is a model of the application lock file used to
// store the data and config.
//
// On unmarshalling, an extra `FilePath` property is added
// to the struct, which contains a path to the associated.
type Locker struct {
	FilePath string `json:"-"`
	Lists    lists  `json:"lists"`
}

// SaveList adds the given list to the list map and writes
// the Locker's content to the associated lock file.
func (l *Locker) SaveList(list *List, dir string) error {
	if len(l.Lists) == 0 {
		l.Lists = make(lists)
	}

	l.Lists.Add(dir, list)

	return l.save()
}

// RemoveList removes a list that matches the given directory
// from the list map.
func (l *Locker) RemoveList(dir string) error {
	l.Lists.Remove(dir)
	return l.save()
}

func (l *Locker) save() error {
	bytes, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(l.FilePath, bytes, 0644)
}

// Dirs returns a slice of directory paths that are associated
// with existing lists.
func (l Locker) Dirs() []string {
	dirs := []string{}
	for dir := range l.Lists {
		dirs = append(dirs, dir)
	}

	return dirs
}

// NewLocker returns a Locker associated with the running
// executable. For example, if the executable lives in
// `/usr/bin/crast`, the lock file will be created next
// to it as `/usr/bin/crast-lock.json`.
func NewLocker() (*Locker, error) {
	exeDir, err := os.Executable()
	if err != nil {
		return nil, err
	}

	filePath := exeDir + "-lock.json"
	locker := &Locker{FilePath: filePath}

	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(file), locker); err != nil {
			return nil, err
		}
	}

	return locker, nil
}
