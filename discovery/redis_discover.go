package discovery

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/redis"

	"github.com/akkagao/gms/common"
)

func init() {
	redis.Register()
}

type RedisDiscover struct {
	GmsServerName     string
	GmsServiceAddress []string
	RedisAddress      []string
	kv                store.Store
	address           []string
}

func NewRedisDiscover(serverName string, redisAddress []string) (IDiscover, error) {
	kv, err := valkeyrie.NewStore(
		store.REDIS,
		redisAddress,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[NewRedisDiscover] error: %v", err))
	}

	redisDiscover := &RedisDiscover{
		GmsServerName: serverName,
		RedisAddress:  redisAddress,
		kv:            kv,
		address:       []string{},
	}

	redisDiscover.loadService(kv)

	go redisDiscover.watch()

	return redisDiscover, nil
}

/**
获取服务列表
*/
func (r *RedisDiscover) GetServer() ([]string, error) {
	if r.address == nil || len(r.address) == 0 {
		return nil, errors.New("no server list")
	}
	return r.address, nil
}

func (r *RedisDiscover) watch() {
	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	watchTree, err := r.kv.WatchTree(nodeName, nil, nil)
	if err != nil {
		log.Fatalf("can't watch %v err:%v", nodeName, err)
		return
	}
	
	for {
		select {
		case kvPairs := <-watchTree:
			fmt.Println("watching ...")
			address := []string{}
			for _, pair := range kvPairs {
				serverAddress := strings.TrimPrefix(pair.Key, nodeName+"/")
				address = append(address, serverAddress)
			}
			r.address = address
		}
	}
}

func (r *RedisDiscover) loadService(kv store.Store) {
	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	kvPairs, err := kv.List(nodeName, nil)
	if err != nil {
		log.Fatal("loadService error:", err)
		return
	}

	address := []string{}
	for _, pair := range kvPairs {
		serverAddress := strings.TrimPrefix(pair.Key, nodeName+"/")
		address = append(address, serverAddress)
	}
	r.address = address
}
