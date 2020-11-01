package discovery

import (
	"errors"
	"fmt"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/redis"
)

func init() {
	redis.Register()
}

type RedisDiscover struct {
	GmsServerName     string
	GmsServiceAddress []string
	RedisAddress      []string
	kv                store.Store
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
	}

	go redisDiscover.watch()

	return redisDiscover, nil
}

/**
todo
*/
func (r *RedisDiscover) GetServer() ([]string, error) {
	return []string{"127.0.0.1:1024"}, nil
}

func (r *RedisDiscover) watch() {

}
