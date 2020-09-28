package owl

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var exampleYaml = `
name: test
addr: :8080
test:
  test01: test01
  test02: test02
`

var exampleByteYaml = []byte(exampleYaml)

var emptyByteYaml = []byte("")

func TestNew(t *testing.T) {
	owlTest := New()
	assert.NotNil(t, owlTest)
	assert.NotSame(t, owlTest, owl)
}

func TestAddConfPath(t *testing.T) {
	AddConfPath("1")
	AddConfPath("2")
	assert.Equal(t, 2, len(owl.filepath))
	assert.Equal(t, "1", owl.filepath[0])
	assert.Equal(t, "2", owl.filepath[1])
}

func TestSetConfName(t *testing.T) {
	SetConfName("test.yaml")
	assert.NotEmpty(t, owl.filename)
	assert.Equal(t, "test.yaml", owl.filename)
}

func TestReadConf(t *testing.T) {
	AddConfPath("./mock/")
	SetConfName("test.yaml")
	err := ReadConf()
	assert.Nil(t, err)
	SetConfName("test1.yaml")
	err = ReadConf()
	assert.Error(t, err)
}

func TestReadInConf(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "", args: args{content: exampleByteYaml}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadInConf(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("ReadInConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetRemoteAddr(t *testing.T) {
	SetRemoteAddr([]string{"localhost:2379"})
	assert.NotNil(t, owl.client)
}

func TestGetAll(t *testing.T) {
	AddConfPath("./mock/")

	SetConfName("test.yaml")
	_ = ReadConf()
	assert.NotNil(t, GetAll())

	owl.config = nil
	_ = ReadInConf(emptyByteYaml)
	assert.Nil(t, GetAll())
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

func TestOwl_findConfigFile(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		filename string
		want     string
		wantErr  bool
	}{
		{"a", "./mock/", "test.yaml", "./mock/test.yaml", false},
		{"b", "./mock/", "test1.yaml", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			o.filepath = []string{tt.value}
			o.filename = tt.filename
			file, err := o.findConfigFile()
			assert.Equal(t, file, tt.want)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
