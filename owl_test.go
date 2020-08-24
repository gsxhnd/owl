package owl

import (
	"fmt"
	"testing"
)

func init() {
	owl.SetAddr([]string{"local_dev:2379"})
}

func TestGetKeys(t *testing.T) {
	keys, _ := GetKeys("/conf")
	fmt.Print(keys)
}
