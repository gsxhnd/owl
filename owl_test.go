package owl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	owlTest := New()
	t.Log(&owlTest, &owl)
	if &owlTest == &owl {
		t.Error("use one instance")
	}
}

func TestAddConfPath(t *testing.T) {
	AddConfPath("1")
	AddConfPath("2")
	if len(owl.filepath) != 2 {
		t.Error("length error")
	}
	if owl.filepath[0] != "1" || owl.filepath[1] != "2" {
		t.Error("filepath value error")
	}
}

func TestSetConfName(t *testing.T) {
	SetConfName("test.yaml")
	if owl.filename == "" {
		t.Error("config name is not set")
	}
	if owl.filename != "test.yaml" {
		t.Error("config name is error")
	}
}

func TestReadConf(t *testing.T) {
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	if err != nil {
		t.Error(fmt.Sprintf("%+v", err))
	}

}

func TestGetAll(t *testing.T) {
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	_ = ReadConf()
	conf := GetAll()
	if conf == nil {
		t.Error("conf is nil")
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		args string
		want interface{}
	}{
		{name: "name", args: "name", want: "test"},
		{name: "test01", args: "test.test01", want: "test01"},
	}
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	_ = ReadConf()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
