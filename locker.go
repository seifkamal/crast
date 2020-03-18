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

func (l *Locker) Save(list List, dir string) error {
	l.Lists[dir] = list

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
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		ioutil.WriteFile(filePath, []byte("{\"lists\": {}}"), 0644)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	locker := &Locker{FilePath: filePath}
	if err := json.Unmarshal([]byte(file), locker); err != nil {
		return nil, err
	}

	return locker, nil
}
