package owl

import (
	"os"
)

func exists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return !stat.IsDir(), err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
