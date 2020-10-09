package owl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_exists(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    bool
		wantErr bool
	}{
		{"test_exist_file", "./mock/test.yaml", true, false},
		{"test_not_exist_file", "./test1.yaml", false, true},
		{"test_dir", "./mock", false, false},
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
