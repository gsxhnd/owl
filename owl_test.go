package owl

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"reflect"
	"sync"
	"testing"
	"time"
)

func init() {
	owl.SetAddr([]string{"local_dev:2379"})
}

func TestGetKeys(t *testing.T) {
	keys, _ := GetKeys("/conf")
	fmt.Print(keys)
}

func TestAddConfPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAll(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBool(tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetByKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetByKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat64(tt.args.key); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIntSlice(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKeys1(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKeys(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringMap(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringMap(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringMapString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringMapString(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringSlice(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTime(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUint(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUint(tt.args.key); got != tt.want {
				t.Errorf("GetUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeteDuration(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeteDuration(tt.args.key); got != tt.want {
				t.Errorf("GeteDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		conf clientv3.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Owl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_AddConfPath(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_Get(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			got, err := o.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetAll(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetAll(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetBool(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetBool(tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetByKey(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			got, err := o.GetByKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetByKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetFloat64(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetFloat64(tt.args.key); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetInt(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetIntSlice(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetIntSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetInterface(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetKeys(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			got, err := o.GetKeys(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetString(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetStringMap(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetStringMap(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetStringMapString(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetStringMapString(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetStringSlice(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetStringSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetTime(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetTime(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GetUint(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GetUint(tt.args.key); got != tt.want {
				t.Errorf("GetUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_GeteDuration(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if got := o.GeteDuration(tt.args.key); got != tt.want {
				t.Errorf("GeteDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOwl_Put(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
			if err := o.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOwl_ReadConf(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_SetAddr(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		addr []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_SetConfName(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_SetConfig(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		config clientv3.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_SetKey(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestOwl_Watcher(t *testing.T) {
	type fields struct {
		key      string
		value    string
		filename string
		filepath []string
		client   *clientv3.Client
		lock     sync.RWMutex
	}
	type args struct {
		key string
		c   chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Owl{
				key:      tt.fields.key,
				value:    tt.fields.value,
				filename: tt.fields.filename,
				filepath: tt.fields.filepath,
				client:   tt.fields.client,
				lock:     tt.fields.lock,
			}
		})
	}
}

func TestPut(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetAddr(t *testing.T) {
	type args struct {
		addr []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetConfName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetConfig(t *testing.T) {
	type args struct {
		config clientv3.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestWatcher(t *testing.T) {
	type args struct {
		key string
		c   chan string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
