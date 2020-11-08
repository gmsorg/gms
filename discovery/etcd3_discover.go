package discovery

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	etcd "github.com/abronan/valkeyrie/store/etcd/v3"

	"github.com/akkagao/gms/common"
)

func init() {
	etcd.Register()
}

type Etcd3Discover struct {
	GmsServerName     string
	GmsServiceAddress []string
	Etcd3Address      []string
	kv                store.Store
	address           []string
}

func NewEtcd3Discover(serverName string, etcd3Address []string) (IDiscover, error) {

	kv, err := valkeyrie.NewStore(
		store.ETCDV3,
		etcd3Address,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[NewEtcd3Discover] error: %v", err))
	}

	etcd3Discover := &Etcd3Discover{
		GmsServerName: serverName,
		Etcd3Address:  etcd3Address,
		kv:            kv,
		address:       []string{},
	}

	etcd3Discover.loadService()

	go etcd3Discover.watch()

	return etcd3Discover, nil
}

func (r *Etcd3Discover) DeleteServer(key string) {
	for i := 0; i < len(r.address); i++ {
		if key == r.address[i] {
			r.address = append(r.address[:i], r.address[i+1:]...)
		}
	}
}

/**
获取服务列表
*/
func (r *Etcd3Discover) GetServer() ([]string, error) {
	if r.address == nil || len(r.address) == 0 {
		return nil, errors.New("no server list")
	}
	return r.address, nil
}

func (r *Etcd3Discover) watch() {
	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	watchTree, err := r.kv.WatchTree(nodeName, nil, nil)
	if err != nil {
		log.Fatalf("can't watch %v err:%v", nodeName, err)
		return
	}

	for {
		select {
		case <-watchTree:
			// case kvPairs := <-watchTree:
			// fmt.Println("watching ...", len(kvPairs))
			// address := []string{}
			// for _, pair := range kvPairs {
			// 	serverAddress := strings.TrimPrefix(pair.Key, nodeName+"/")
			// 	address = append(address, serverAddress)
			// }
			// r.address = address
			r.loadService()
		}
	}
}

func (r *Etcd3Discover) loadService() {
	// nodeName := common.BasePath
	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	kvPairs, err := r.kv.List(nodeName, nil)
	if err != nil {
		log.Println("[Etcd3Discover] loadService error: ", err)
		return
	}

	address := []string{}
	for _, pair := range kvPairs {
		serverAddress := strings.TrimPrefix(pair.Key, nodeName+"/")
		address = append(address, serverAddress)
	}
	r.address = address
}
