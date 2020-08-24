package owl

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
	"time"
)

var owl *Owl

func init() {
	owl = new(Owl)
}

// Owl is a lib for get configure value from etcd.
type Owl struct {
	key    string
	value  string
	client *clientv3.Client
	lock   sync.RWMutex
}

// New returns an initialized Owl instance.
func New(conf clientv3.Config) (*Owl, error) {
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
		return nil, err
	}
	return &Owl{
		client: client,
	}, nil
}

// SetConfigName sets configure for the etcd. The
// client include etcd url
func SetConfig(config clientv3.Config) { owl.SetConfig(config) }
func (o *Owl) SetConfig(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		client = nil
	}
	o.client = client
}

// SetAddr sets address for the etcd use default etcd client config.
func SetAddr(addr []string) { owl.SetAddr(addr) }
func (o *Owl) SetAddr(addr []string) {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
	}
	o.client = client
}

// SetKey set config key name in etcd.
func SetKey(key string) { owl.SetKey(key) }
func (o *Owl) SetKey(key string) {
	defer o.lock.Unlock()
	o.lock.Lock()
	o.key = key
	o.value = ""
}

// Get get value from etcd. The config's key was
// stored by SetKey.
func Get() (string, error) { return owl.Get() }
func (o *Owl) Get() (string, error) {
	defer o.lock.Unlock()
	o.lock.Lock()

	key := o.key
	if key == "" {
		return "", errors.New("")
	}
	if o.value != "" {
		return o.value, nil
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := o.client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	for _, v := range resp.Kvs {
		o.value = string(v.Value)
	}
	return o.value, nil
}

// Put value into etcd.
func Put(key, value string) error { return owl.Put(key, value) }
func (o *Owl) Put(key, value string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := o.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetByKey get config content from etcd by key
func GetByKey(key string) (string, error) { return owl.GetByKey(key) }
func (o *Owl) GetByKey(key string) (string, error) {
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

func GetKeys(prefix string) ([]string, error) { return owl.GetKeys(prefix) }
func (o *Owl) GetKeys(prefix string) ([]string, error) {
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

// Watch watch key's value in etcd
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
