package owl

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
	"time"
)

var owl *Owl

func init() {
	owl = new(Owl)
}

type Owl struct {
	key    string
	value  string
	config clientv3.Config
	client *clientv3.Client
	lock   sync.RWMutex
	vc     *chan string
}

func New(key string, conf clientv3.Config) *Owl {
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
		return nil
	}
	c := make(chan string)
	return &Owl{
		key:    key,
		client: client,
		vc:     &c,
	}
}

func Default(addr []string) *Owl {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
	}
	return &Owl{
		client: client,
	}

}

func SetConfig(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		client = nil
	}
	owl.client = client
}
func SetAddr(addr []string) {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
	}
	owl.client = client
}
func SetKey(key string) { owl.SetKey(key) }
func (o *Owl) SetKey(key string) {
	defer o.lock.Unlock()
	o.lock.Lock()
	o.key = key
}

func Get(key string) (string, error)          { return owl.get(key) }
func (o *Owl) Get(key string) (string, error) { return o.get(key) }
func (o *Owl) get(key string) (string, error) {
	defer o.lock.Unlock()
	o.lock.Lock()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if key == "" {
		key = o.key
	}
	resp, err := o.client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	for _, v := range resp.Kvs {
		o.value = string(v.Value)
	}
	return o.value, nil
}

func Put(key, value string) error { return owl.Put(key, value) }
func (o *Owl) Put(key, value string) error {
	defer o.lock.Unlock()
	o.lock.Lock()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if key == "" {
		key = o.key
	}
	if value == "" {
		value = o.value
	} else {
		o.value = value
	}
	_, err := o.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}
func (o *Owl) GetValue() string {
	return o.value
}
func (o *Owl) GetKey() string {
	return o.key
}

func (o *Owl) update(v string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.value = v
}
func (o *Owl) Watcher(key string, c chan string) {
	fmt.Println("start watch:", key)
	if key == "" {
		key = o.key
	}
	rch := o.client.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				o.update(string(ev.Kv.Value))
				c <- o.value
			case mvccpb.DELETE:
				o.update("")
				c <- ""
			default:
			}
		}
	}
}
