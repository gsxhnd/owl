package backend

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type EtcdConn struct {
	EndPoints []string
	Key       string
	Value     string
	client    *clientv3.Client
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

func (e *EtcdConn) Put() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := e.client.Put(ctx, e.Key, e.Value)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
