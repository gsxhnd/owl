package backend

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
	"time"
)

type EtcdConn struct {
	EndPoints []string
	Key       string
	Value     string
	client    *clientv3.Client
	lock      sync.RWMutex
}

func NewEtcdClient(addr []string) (*EtcdConn, error) {
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
	return &EtcdConn{
		EndPoints: addr,
		client:    client,
	}, nil
}

func NewEtcdConn(conf clientv3.Config) (*EtcdConn, error) {
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
		return nil, err
	}
	return &EtcdConn{
		EndPoints: conf.Endpoints,
		client:    client,
	}, nil
}

func (e *EtcdConn) Put() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := e.client.Put(ctx, e.Key, e.Value)
	if err != nil {
		return err
	}
	return nil
}

func (e *EtcdConn) Get() (string, error) {
	defer e.lock.Unlock()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := e.client.Get(ctx, e.Key)
	if err != nil {
		return "", err
	}
	for _, v := range resp.Kvs {
		e.lock.Lock()
		e.Value = string(v.Value)
	}
	return e.Value, nil
}

func (e *EtcdConn) update(v string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.Value = v
}

func (e *EtcdConn) Watcher() string {
	rch := e.client.Watch(context.Background(), e.Key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				e.update(string(ev.Kv.Value))
				return e.Value
			case mvccpb.DELETE:
				e.update("")
				return ""
			default:
				return e.Value
			}
		}
	}
	return ""
}
