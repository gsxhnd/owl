package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"123", args{"123", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg)
		})
	}
}
