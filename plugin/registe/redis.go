package registe

import (
	"log"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/redis"

	"github.com/akkagao/gms/plugin"
)

type RedisRegistePlugin struct {
	GmsServiceAddress string
	RedisAddress      []string
	kv                store.Store
}

func init() {
	redis.Register()
}

func NewRedisRegistePlugin(redisAddress []string) plugin.IPlugin {
	return &RedisRegistePlugin{
		RedisAddress: redisAddress,
	}
}

func (r *RedisRegistePlugin) Start() error {
	kv, err := valkeyrie.NewStore(
		store.REDIS, // or "consul"
		r.RedisAddress,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store redis")
		return err
	}
	r.kv = kv

	// todo
	return nil

}

// todo
// gms#serverName#funcName
func (r *RedisRegistePlugin) Registe(ip string, port int) error {
	// 注册服务
	return nil
}
