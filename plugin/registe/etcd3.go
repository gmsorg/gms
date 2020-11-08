package registe

import (
	"fmt"
	"log"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	etcd "github.com/abronan/valkeyrie/store/etcd/v3"

	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/plugin"
)

func init() {
	etcd.Register()
}

type Etcd3RegistePlugin struct {
	GmsServerName     string
	GmsServiceAddress string
	Etcd3Address      []string
	kv                store.Store
	UpdateInterval    time.Duration
}

func NewEtcd3RegistePlugin(serverName string, etcd3Address []string) plugin.IPlugin {
	return &Etcd3RegistePlugin{
		GmsServerName:  serverName,
		Etcd3Address:   etcd3Address,
		UpdateInterval: time.Minute,
	}
}

func (e *Etcd3RegistePlugin) Start() error {
	kv, err := valkeyrie.NewStore(
		store.ETCDV3,
		e.Etcd3Address,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store etcd3", err)
		return err
	}
	e.kv = kv

	err = e.kv.Put(common.BasePath, []byte(common.BasePath), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[Etcd3RegistePlugin] put BasePath error: %v", err)
		return err
	}

	nodeName := fmt.Sprintf("%v/%v", common.BasePath, e.GmsServerName)
	err = e.kv.Put(nodeName, []byte(e.GmsServerName), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[Etcd3RegistePlugin] put nodeName: %v error: %v", nodeName, err)
	}

	return nil
}

func (e *Etcd3RegistePlugin) Registe(ip string, port int) error {
	// 注册服务
	e.GmsServiceAddress = fmt.Sprintf("%v:%v", ip, port)

	err := e.registeService()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if e.UpdateInterval > 0 {
		ticker := time.NewTicker(e.UpdateInterval)
		go func() {
			for {
				select {
				case <-ticker.C:
					err := e.registeService()
					if err != nil {
						log.Printf("[Registe] error: %v", err)
					}
				}
			}
		}()
	}
	return nil
}

func (e *Etcd3RegistePlugin) registeService() error {
	nodeName := fmt.Sprintf("%v/%v/%v", common.BasePath, e.GmsServerName, e.GmsServiceAddress)
	err := e.kv.Put(nodeName, []byte(e.GmsServiceAddress), &store.WriteOptions{TTL: e.UpdateInterval * 2})
	if err != nil {
		return fmt.Errorf("[RedisRegistePlugin] put nodeName: %v error: %v", nodeName, err)
	}
	return nil
}
