package registe

import (
	"fmt"
	"log"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/redis"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/plugin"
)

type RedisRegistePlugin struct {
	GmsServerName     string
	GmsServiceAddress string
	RedisAddress      []string
	kv                store.Store
	UpdateInterval    time.Duration
}

func init() {
	redis.Register()
}

func NewRedisRegistePlugin(serverName string, redisAddress string) plugin.IPlugin {
	return &RedisRegistePlugin{
		GmsServerName:  serverName,
		RedisAddress:   []string{redisAddress},
		UpdateInterval: time.Minute,
	}
}

func (r *RedisRegistePlugin) Start() error {
	kv, err := valkeyrie.NewStore(
		store.REDIS,
		r.RedisAddress,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store redis", err)
		return err
	}
	r.kv = kv

	err = r.kv.Put(common.BasePath, []byte(common.BasePath), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[RedisRegistePlugin] put BasePath error: %w", err)
		return err
	}

	nodeName := fmt.Sprintf("%v/%v", common.BasePath, r.GmsServerName)
	err = r.kv.Put(nodeName, []byte(r.GmsServerName), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[RedisRegistePlugin] put nodeName: %v error: %w", nodeName, err)
	}

	return nil

}

// gms/serverName/serviceAddress
func (r *RedisRegistePlugin) Registe(ip string, port int) error {
	// 注册服务
	r.GmsServiceAddress = fmt.Sprintf("%v:%v", ip, port)

	err := r.registeService()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if r.UpdateInterval > 0 {
		ticker := time.NewTicker(r.UpdateInterval)
		go func() {
			for {
				select {
				case <-ticker.C:
					err := r.registeService()
					if err != nil {
						log.Printf("[Registe] error: %v", err)
					}
				}
			}
		}()
	}
	return nil
}

func (r *RedisRegistePlugin) registeService() error {
	nodeName := fmt.Sprintf("%v/%v/%v", common.BasePath, r.GmsServerName, r.GmsServiceAddress)
	err := r.kv.Put(nodeName, []byte(r.GmsServiceAddress), &store.WriteOptions{TTL: r.UpdateInterval * 2})
	if err != nil {
		return fmt.Errorf("[RedisRegistePlugin] put nodeName: %v error: %w", nodeName, err)
	}
	return nil
}
