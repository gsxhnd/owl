package owl

import (
	"os"
	"testing"
)

func Test_exists(t *testing.T) {
	wd, _ := os.Getwd()
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "mock", args: args{path: "./mock/test.yaml"}, want: true, wantErr: false},
		{name: "mock", args: args{path: wd + "/test.yaml"}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
