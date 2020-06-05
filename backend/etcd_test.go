package backend

import (
	"github.com/coreos/etcd/clientv3"
	"testing"
	"time"
)

func TestEtcdConn_Get(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			"test",
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := clientv3.Config{
				Endpoints:        []string{"local_dev:2379"},
				AutoSyncInterval: 0,
				DialTimeout:      5 * time.Second,
			}
			e, _ := NewEtcdConn(conf)
			got, err := e.Get("/conf/test.yaml")
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("got: ", got)
		})
	}
}
