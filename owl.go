package owl

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	yaml "gopkg.in/yaml.v3"
)

var owl *Owl

func init() {
	owl = New()
}

var (
	FileNotExistError = errors.New("file not exist")
	FileIsDirError    = errors.New("path is dir")
	FilenameError     = errors.New("filename is not set")
)

// Owl is a lib for get configure value from etcd.
type Owl struct {
	key      string
	value    string
	config   map[string]interface{}
	filename string
	filepath []string
	client   *clientv3.Client
	lock     sync.RWMutex
}

// New returns an initialized Owl instance.
func New() *Owl {
	return &Owl{}
}

// SetRemoteAddr set url for the etcd.
func SetRemoteAddr(addr []string) error { return owl.SetRemoteAddr(addr) }
func (o *Owl) SetRemoteAddr(addr []string) error {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = client.Status(ctx, addr[0])
	if err != nil {
		return err
	}
	o.client = client
	return nil
}

// GetRemoteKeys get keys from etcd by prefix
func GetRemoteKeys(prefix string) ([]string, error) { return owl.GetRemoteKeys(prefix) }
func (o *Owl) GetRemoteKeys(prefix string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := o.client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	var keys []string
	for _, v := range resp.Kvs {
		keys = append(keys, string(v.Key))
	}
	return keys, nil
}

// GetRemote get config content from etcd by key
func GetRemote(key string) (string, error) { return owl.GetRemote(key) }
func (o *Owl) GetRemote(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := o.client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	var value string

	for _, v := range resp.Kvs {
		value = string(v.Value)
	}

	return value, nil
}

// PutRemote value into etcd.
func PutRemote(key, value string) error { return owl.PutRemote(key, value) }
func (o *Owl) PutRemote(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := o.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRemote value from etcd.
func DeleteRemote(key string) error { return owl.DeleteRemote(key) }
func (o *Owl) DeleteRemote(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := o.client.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

// Watcher watch key's value in etcd
func Watcher(key string, c chan string) { owl.Watcher(key, c) }
func (o *Owl) Watcher(key string, c chan string) {
	rch := o.client.Watch(context.Background(), key)
	for resp := range rch {
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				c <- string(ev.Kv.Value)
			case mvccpb.DELETE:
				c <- ""
			}
		}
	}
}

func SetConfName(name string) { owl.SetConfName(name) }
func (o *Owl) SetConfName(name string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filename = name
	fmt.Println(o.filename)
}

// AddConfPath adds a path for owl to search for the config file in.
func AddConfPath(path string) { owl.AddConfPath(path) }
func (o *Owl) AddConfPath(path string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filepath = append(o.filepath, path)
}

// GetConfPaths returns a list of configure paths for owl.
func GetConfPaths() []string { return owl.GetConfPath() }
func (o *Owl) GetConfPath() []string {
	var paths []string
	if len(o.filepath) == 0 {
		if o.filename == "" {
			return nil
		} else {
			paths = append(paths, o.filename)
		}
	} else {
		for _, v := range o.filepath {
			paths = append(paths, v+o.filename)
		}
	}
	return paths
}

// ReadConf will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func ReadConf() error { return owl.ReadConf() }
func (o *Owl) ReadConf() error {
	if o.filename == "" {
		return FilenameError
	}

	file, err := o.findConfigFile()
	if err != nil {
		return err
	}

	content, _ := ioutil.ReadFile(file)

	err = yaml.Unmarshal(content, &o.config)
	if err != nil {
		return err
	}
	return nil
}

func (o *Owl) findConfigFile() (string, error) {
	for _, v := range o.filepath {
		stat, err := os.Stat(v + o.filename)
		if err != nil {
			return "", FileNotExistError
		}
		if !stat.IsDir() {
			return v + o.filename, nil
		} else {
			return "", FileIsDirError
		}
	}

	stat, err := os.Stat(o.filename)
	if err != nil {
		return "", FileNotExistError
	}
	if !stat.IsDir() {
		return o.filename, nil
	} else {
		return "", FileIsDirError
	}
}

// ReadInConf will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func ReadInConf(content []byte) error { return owl.ReadInConf(content) }
func (o *Owl) ReadInConf(content []byte) error {
	err := yaml.Unmarshal(content, &o.config)
	if err != nil {
		return err
	}
	return nil
}

// Get returns the value associated with the key as interface.
func Get(key string) interface{} { return owl.Get(key) }
func (o *Owl) Get(key string) interface{} {
	keys := strings.Split(key, ".")
	return o.find(o.config, keys)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string { return owl.GetString(key) }
func (o *Owl) GetString(key string) string {
	return cast.ToString(o.Get(key))
}

// GetInt returns the value associated with the key as int.
func GetInt(key string) int { return owl.GetInt(key) }
func (o *Owl) GetInt(key string) int {
	return cast.ToInt(o.Get(key))
}

// GetInt64 returns the value associated with the key as int64.
func GetInt64(key string) int64 { return owl.GetInt64(key) }
func (o *Owl) GetInt64(key string) int64 {
	return cast.ToInt64(o.Get(key))
}

// GetUint returns the value associated with the key as uint.
func GetUint(key string) uint { return owl.GetUint(key) }
func (o *Owl) GetUint(key string) uint {
	return cast.ToUint(o.Get(key))
}

// GetFloat64 returns the value associated with the key as float64.
func GetFloat64(key string) float64 { return owl.GetFloat64(key) }
func (o *Owl) GetFloat64(key string) float64 {
	return cast.ToFloat64(o.Get(key))
}

// GetBool returns the value associated with the key as bool.
func GetBool(key string) bool { return owl.GetBool(key) }
func (o *Owl) GetBool(key string) bool {
	return cast.ToBool(o.Get(key))
}

// GetStringSlice returns the value associated with the key as string slice.
func GetStringSlice(key string) []string { return owl.GetStringSlice(key) }
func (o *Owl) GetStringSlice(key string) []string {
	return cast.ToStringSlice(o.Get(key))
}

// GetIntSlice returns the value associated with the key as int slice.
func GetIntSlice(key string) []int { return owl.GetIntSlice(key) }
func (o *Owl) GetIntSlice(key string) []int {
	return cast.ToIntSlice(o.Get(key))
}

// GetStringMap returns the value associated with the key as string map.
func GetStringMap(key string) map[string]interface{} { return owl.GetStringMap(key) }
func (o *Owl) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(o.Get(key))
}

func GetStringMapString(key string) map[string]string { return owl.GetStringMapString(key) }
func (o *Owl) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(o.Get(key))
}

// GetAll returns the all value as map.
func GetAll() map[string]interface{}          { return owl.GetAll() }
func (o *Owl) GetAll() map[string]interface{} { return o.config }

func (o *Owl) find(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}
	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}
		switch source[path[0]].(type) {
		case map[interface{}]interface{}:
			return o.find(cast.ToStringMap(source[path[0]]), path[1:])
		case map[string]interface{}:
			return o.find(source[path[0]].(map[string]interface{}), path[1:])
		default:
			return nil
		}
	}
	return nil
}
