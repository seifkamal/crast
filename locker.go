package crast

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Locker struct {
	FilePath string `json:"-"`
	Lists    lists  `json:"lists"`
}

func (l *Locker) SaveList(list *List, dir string) error {
	if len(l.Lists) == 0 {
		l.Lists = make(lists)
	}

	l.Lists.Add(dir, list)

	bytes, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(l.FilePath, bytes, 0644)
}

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
