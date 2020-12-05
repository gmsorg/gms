package discovery

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/consul"

	"github.com/gmsorg/gms/common"
)

func init() {
	consul.Register()
}

type ConsulDiscover struct {
	GmsServerName     string
	GmsServiceAddress []string
	ConsulAddress     []string
	kv                store.Store
	address           []string
}

func NewConsulDiscover(serverName string, consulAddress []string) (IDiscover, error) {

	kv, err := valkeyrie.NewStore(
		store.CONSUL,
		consulAddress,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[NewConsulDiscover] error: %v", err))
	}

	consulDiscover := &ConsulDiscover{
		GmsServerName: serverName,
		ConsulAddress: consulAddress,
		kv:            kv,
		address:       []string{},
	}

	consulDiscover.loadService()

	go consulDiscover.watch()

	return consulDiscover, nil
}

func (r *ConsulDiscover) DeleteServer(key string) {
	for i := 0; i < len(r.address); i++ {
		if key == r.address[i] {
			r.address = append(r.address[:i], r.address[i+1:]...)
		}
	}
}

/**
获取服务列表
*/
func (r *ConsulDiscover) GetServer() ([]string, error) {
	if r.address == nil || len(r.address) == 0 {
		return nil, errors.New("no server list")
	}
	return r.address, nil
}

func (r *ConsulDiscover) watch() {
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
			// log.Println("watching ...", len(kvPairs))
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

func (r *ConsulDiscover) loadService() {
	// nodeName := common.BasePath
	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	kvPairs, err := r.kv.List(nodeName, nil)
	if err != nil {
		log.Println("[ConsulDiscover] loadService error: ", err)
		return
	}

	address := []string{}
	for _, pair := range kvPairs {
		serverAddress := strings.TrimPrefix(pair.Key, nodeName+"/")
		address = append(address, serverAddress)
	}
	r.address = address
}
