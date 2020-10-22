package owl

import (
	"github.com/pkg/errors"
	"os"
)

func exists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !stat.IsDir() {
		return true, nil
	} else {
		return false, errors.New("file is not exist")
	}
}
