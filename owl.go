package owl

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

var owl *Owl

func init() {
	owl = New()
}

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

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := o.client.Put(ctx, key, value)
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
			default:
			}
		}
	}
}

func SetConfName(name string) { owl.SetConfName(name) }
func (o *Owl) SetConfName(name string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filename = name
}

// AddConfPath adds a path for owl to search for the config file in.
func AddConfPath(path string) { owl.AddConfPath(path) }
func (o *Owl) AddConfPath(path string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filepath = append(o.filepath, path)
}

// ReadConf will read a configuration file, setting existing keys to nil if the
// key does not exist in the file.
func ReadConf() error { return owl.ReadConf() }
func (o *Owl) ReadConf() error {
	if o.filename == "" {
		return errors.WithStack(errors.New("config name not set"))
	}

	file, err := o.findConfigFile()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &o.config)
	if err != nil {
		return err
	}
	return nil
}

func (o *Owl) findConfigFile() (string, error) {
	if o.filepath != nil {
		for _, v := range o.filepath {
			exist, err := exists(v + o.filename)
			if !exist && err != nil {
				return "", err
			} else {
				return v + o.filename, nil
			}
		}
	} else {
		exist, err := exists(o.filename)
		if !exist && err != nil {
			return "", err
		} else {
			return o.filename, nil
		}
	}
	return "", errors.New("file not exist")
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

func GetInt(key string) int { return owl.GetInt(key) }
func (o *Owl) GetInt(key string) int {
	return cast.ToInt(o.Get(key))
}

func GetInt64(key string) int64 { return owl.GetInt64(key) }
func (o *Owl) GetInt64(key string) int64 {
	return cast.ToInt64(o.Get(key))
}

func GetUint(key string) uint { return owl.GetUint(key) }
func (o *Owl) GetUint(key string) uint {
	return cast.ToUint(o.Get(key))
}

func GetFloat64(key string) float64 { return owl.GetFloat64(key) }
func (o *Owl) GetFloat64(key string) float64 {
	return cast.ToFloat64(o.Get(key))
}

func GetBool(key string) bool { return owl.GetBool(key) }
func (o *Owl) GetBool(key string) bool {
	return cast.ToBool(o.Get(key))
}

func GetStringSlice(key string) []string { return owl.GetStringSlice(key) }
func (o *Owl) GetStringSlice(key string) []string {
	return cast.ToStringSlice(o.Get(key))
}

func GetIntSlice(key string) []int { return owl.GetIntSlice(key) }
func (o *Owl) GetIntSlice(key string) []int {
	return cast.ToIntSlice(o.Get(key))
}

func GetStringMap(key string) map[string]interface{} { return owl.GetStringMap(key) }
func (o *Owl) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(o.Get(key))
}

func GetStringMapString(key string) map[string]string { return owl.GetStringMapString(key) }
func (o *Owl) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(o.Get(key))
}

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
