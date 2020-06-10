package backend

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
	"time"
)

type etcdConn struct {
	Key    string
	Value  string
	client *clientv3.Client
	lock   sync.RWMutex
}

func NewEtcdClient(addr []string) (*etcdConn, error) {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
		return nil, err
	}
	return &etcdConn{
		client: client,
	}, nil
}

func NewEtcdConn(conf clientv3.Config) (*etcdConn, error) {
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
		return nil, err
	}
	return &etcdConn{
		client: client,
	}, nil
}

func (e *etcdConn) Put(key, value string) error {
	defer e.lock.Unlock()
	e.lock.Lock()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if key == "" {
		key = e.Key
	}
	if value == "" {
		value = e.Value
	} else {
		e.Value = value
	}
	_, err := e.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (e *etcdConn) Get(key string) (string, error) {
	defer e.lock.Unlock()
	e.lock.Lock()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if key == "" {
		key = e.Key
	}
	resp, err := e.client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	for _, v := range resp.Kvs {
		e.Value = string(v.Value)
	}
	return e.Value, nil
}

func (e *etcdConn) update(v string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.Value = v
}

func (e *etcdConn) Watcher(key string, c chan string) {
	fmt.Println("start watch:", key)
	if key == "" {
		key = e.Key
	}
	rch := e.client.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				e.update(string(ev.Kv.Value))
				c <- e.Value
			case mvccpb.DELETE:
				e.update("")
				c <- ""
			default:
			}
		}
	}
}
