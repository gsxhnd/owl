package owl

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_exists(t *testing.T) {
	wd, _ := os.Getwd()
	tests := []struct {
		name    string
		path    string
		want    bool
		wantErr bool
	}{
		{"mock", "./mock/test.yaml", true, false},
		{"mock", wd + "/test.yaml", false, true},
		{"mock", "./test1.yaml", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := exists(tt.path)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
