package registe

import (
	"fmt"
	"log"
	"time"

	"github.com/abronan/valkeyrie"
	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/consul"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/plugin"
)

func init() {
	consul.Register()
}

type ConsulRegistePlugin struct {
	GmsServerName     string
	GmsServiceAddress string
	ConsulAddress     []string
	kv                store.Store
	UpdateInterval    time.Duration
}

func (c *ConsulRegistePlugin) Start() error {
	kv, err := valkeyrie.NewStore(
		store.CONSUL,
		c.ConsulAddress,
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store etcd3", err)
		return err
	}
	c.kv = kv

	err = c.kv.Put(common.BasePath, []byte(common.BasePath), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[ConsulRegistePlugin] put BasePath error: %w", err)
		return err
	}

	nodeName := fmt.Sprintf("%v/%v", common.BasePath, c.GmsServerName)
	err = c.kv.Put(nodeName, []byte(c.GmsServerName), &store.WriteOptions{IsDir: true})
	if err != nil {
		fmt.Errorf("[ConsulRegistePlugin] put nodeName: %v error: %w", nodeName, err)
	}

	return nil
}

func (c *ConsulRegistePlugin) Registe(ip string, port int) error {
	// 注册服务
	c.GmsServiceAddress = fmt.Sprintf("%v:%v", ip, port)

	err := c.registeService()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if c.UpdateInterval > 0 {
		ticker := time.NewTicker(c.UpdateInterval)
		go func() {
			for {
				select {
				case <-ticker.C:
					err := c.registeService()
					if err != nil {
						log.Printf("[ConsulRegistePlugin] error: %v", err)
					}
				}
			}
		}()
	}
	return nil
}

func NewConsulRegistePlugin(serverName string, etcd3Address []string) plugin.IPlugin {
	return &ConsulRegistePlugin{
		GmsServerName:  serverName,
		ConsulAddress:  etcd3Address,
		UpdateInterval: time.Minute,
	}
}

func (c *ConsulRegistePlugin) registeService() error {
	nodeName := fmt.Sprintf("%v/%v/%v", common.BasePath, c.GmsServerName, c.GmsServiceAddress)
	err := c.kv.Put(nodeName, []byte(c.GmsServiceAddress), &store.WriteOptions{TTL: c.UpdateInterval * 2})
	if err != nil {
		return fmt.Errorf("[ConsulRegistePlugin] put nodeName: %v error: %w", nodeName, err)
	}
	return nil
}
