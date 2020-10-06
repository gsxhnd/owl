package owl

import (
	"github.com/stretchr/testify/assert"
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

func resetOwl() {
	owl.filename = ""
	owl.filepath = nil
	owl.value = ""
	owl.key = ""
	owl.config = nil
	owl.client = nil
}

func TestNew(t *testing.T) {
	owlTest := New()
	assert.NotNil(t, owlTest)
	assert.NotSame(t, owlTest, owl)
}

func TestSetRemoteAddr(t *testing.T) {
	tests := []struct {
		name    string
		addr    []string
		wantErr bool
	}{
		{"not_err", []string{"localhost:2379"}, false},
		{"err", []string{"192.168.1.1:2379"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetOwl()
			err := SetRemoteAddr(tt.addr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
	err := SetRemoteAddr([]string{"localhost:2379"})
	assert.NotNil(t, owl.client)
	assert.Nil(t, err)
}
func TestOwl_SetRemoteAddr(t *testing.T) {
	tests := []struct {
		name    string
		addr    []string
		wantErr bool
	}{
		{"not_err", []string{"localhost:2379"}, false},
		{"err", []string{"192.168.1.1:2379"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			err := o.SetRemoteAddr(tt.addr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestPutRemote(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantErr bool
	}{
		{"test_success", "/test", "test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = SetRemoteAddr([]string{"localhost:2379"})
			err := PutRemote(tt.key, tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
func TestOwl_PutRemote(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantErr bool
	}{
		{"test_success", "/test", "test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			_ = o.SetRemoteAddr([]string{"localhost:2379"})
			err := o.PutRemote(tt.key, tt.value)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetRemoteKeys(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		want    []string
		wantErr bool
	}{
		{"test", "/test", []string{"/test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = SetRemoteAddr([]string{"localhost:2379"})
			keys, err := GetRemoteKeys(tt.prefix)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, keys)
			}
		})
	}
}
func TestOwl_GetRemoteKeys(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		want    []string
		wantErr bool
	}{
		{"test", "/test", []string{"/test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			_ = o.SetRemoteAddr([]string{"localhost:2379"})
			keys, err := o.GetRemoteKeys(tt.prefix)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, keys)
			}
		})
	}
}

func TestGetRemote(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{"test_success", "/test", "test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := GetRemote(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, v)
			}
		})
	}
}
func TestOwl_GetRemote(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{"test_success", "/test", "test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			_ = o.SetRemoteAddr([]string{"localhost:2379"})
			v, err := o.GetRemote(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, v)
			}
		})
	}
}

func TestWatcher(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func TestOwl_Watcher(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestAddConfPath(t *testing.T) {
	AddConfPath("1")
	AddConfPath("2")
	assert.Equal(t, 2, len(owl.filepath))
	assert.Equal(t, "1", owl.filepath[0])
	assert.Equal(t, "2", owl.filepath[1])
}
func TestOwl_AddConfPath(t *testing.T) {
	o := New()
	o.AddConfPath("1")
	o.AddConfPath("2")
	assert.Equal(t, 2, len(o.filepath))
	assert.Equal(t, "1", o.filepath[0])
	assert.Equal(t, "2", o.filepath[1])
}

func TestSetConfName(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"a", "test.yaml", "test.yaml"},
		{"b", "test1.yaml", "test1.yaml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetOwl()
			SetConfName(tt.value)
			assert.NotEmpty(t, owl.filename)
			assert.Equal(t, owl.filename, tt.want)
		})
	}
}
func TestOwl_SetConfName(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"a", "test.yaml", "test.yaml"},
		{"b", "test1.yaml", "test1.yaml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			o.SetConfName(tt.value)
			assert.NotEmpty(t, o.filename)
			assert.Equal(t, o.filename, tt.want)
		})
	}
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
func TestOwl_ReadConf(t *testing.T) {
	o := New()
	o.AddConfPath("./mock/")
	o.SetConfName("test.yaml")
	err := o.ReadConf()
	assert.Nil(t, err)
	o.SetConfName("test1.yaml")
	err = o.ReadConf()
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
func TestOwl_ReadInConf(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantErr bool
	}{
		{"", exampleByteYaml, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			err := o.ReadInConf(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
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
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Get(tt.args), tt.want)
		})
	}
}
func TestOwl_Get(t *testing.T) {
	tests := []struct {
		name string
		args string
		want interface{}
	}{
		{name: "name", args: "name", want: "test"},
		{name: "test01", args: "test.test01", want: "test01"},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.Get(tt.args), tt.want)
		})
	}
}

func TestGetString(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{"name", "name", "test"},
		{"test01", "test.test01", "test01"},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetString(tt.key), tt.want)
		})
	}
}
func TestOwl_GetString(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{"name", "name", "test"},
		{"test01", "test.test01", "test01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			o.SetConfName("test.yaml")
			o.AddConfPath("./mock/")
			err := o.ReadConf()
			assert.Nil(t, err)
			assert.Equal(t, o.GetString(tt.key), tt.want)
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want int
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetInt(tt.key), tt.want)
		})
	}
}
func TestOwl_GetInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want int
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.GetInt(tt.key), tt.want)
		})
	}
}

func TestGetUInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want uint
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetUint(tt.key), tt.want)
		})
	}
}
func TestOwl_GetUInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want uint
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.GetUint(tt.key), tt.want)
		})
	}
}

func TestGetFloat64(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want float64
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetFloat64(tt.key), tt.want)
		})
	}
}
func TestOwl_GetFloat64(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want float64
	}{
		{"test_01", "test_int.test01", 1},
		{"test_02", "test_int.test02", 2},
		{"test_03", "test_int.test03", 0},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.GetFloat64(tt.key), tt.want)
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{"test_01", "test_bool.test01", true},
		{"test_02", "test_bool.test02", false},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetBool(tt.key), tt.want)
		})
	}
}
func TestOwl_GetBool(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{"test_01", "test_bool.test01", true},
		{"test_02", "test_bool.test02", false},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.GetBool(tt.key), tt.want)
		})
	}
}

func TestGetStringSlice(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []string
	}{
		{"test_01", "test_string_slice.test01", []string{"test1", "test2"}},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, GetStringSlice(tt.key), tt.want)
		})
	}
}
func TestOwl_GetStringSlice(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []string
	}{
		{"test_01", "test_string_slice.test01", []string{"test1", "test2"}},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.GetStringSlice(tt.key), tt.want)
		})
	}
}

func TestGetIntSlice(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []int
	}{
		{"test_01", "test_int_slice.test01", []int{1, 2}},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetIntSlice(tt.key))
		})
	}
}
func TestOwl_GetIntSlice(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []int
	}{
		{"test_01", "test_int_slice.test01", []int{1, 2}},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, o.GetIntSlice(tt.key))
		})
	}
}

func TestGetStringMap(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want map[string]interface{}
	}{
		{"test_01", "test_string_map", map[string]interface{}{"test01": "test01", "test02": "test02"}},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetStringMap(tt.key))
		})
	}
}
func TestOwl_GetStringMap(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want map[string]interface{}
	}{
		{"test_01", "test_string_map", map[string]interface{}{"test01": "test01", "test02": "test02"}},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, o.GetStringMap(tt.key))
		})
	}
}

func TestGetStringMapString(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want map[string]interface{}
	}{
		{"test_01", "test_string_map", map[string]interface{}{"test01": "test01", "test02": "test02"}},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetStringMapString(tt.key))
		})
	}
}
func TestOwl_GetStringMapString(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want map[string]interface{}
	}{
		{"test_01", "test_string_map", map[string]interface{}{"test01": "test01", "test02": "test02"}},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, o.GetStringMapString(tt.key))
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantNil bool
	}{
		{"1", exampleByteYaml, false},
		{"2", emptyByteYaml, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetOwl()
			err := ReadInConf(tt.content)
			assert.Nil(t, err)
			if tt.wantNil {
				assert.Nil(t, GetAll())
			} else {
				assert.NotNil(t, GetAll())
			}
		})
	}
}
func TestOwl_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantNil bool
	}{
		{"1", exampleByteYaml, false},
		{"2", emptyByteYaml, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			err := o.ReadInConf(tt.content)
			assert.Nil(t, err)
			if tt.wantNil {
				assert.Nil(t, o.GetAll())
			} else {
				assert.NotNil(t, o.GetAll())
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
