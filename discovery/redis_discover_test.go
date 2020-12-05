package discovery

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/gmsorg/gms/plugin"
	"github.com/gmsorg/gms/plugin/registe"
)

func TestNewRedisDiscover(t *testing.T) {
	plugins := registe.NewRedisRegistePlugin("gmsDemo", "127.0.0.1:6379")
	plugins.Start()

	registePlugin := plugins.(plugin.IRegistePlugin)
	registePlugin.Registe("127.0.0.1", 1024)

	discover, err := NewRedisDiscover("gmsDemo", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}

	timeout := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeout.C:
			log.Println("---")
			rand.Seed(time.Now().UnixNano())
			registePlugin.Registe("127.0.0.1", rand.Intn(1000))

			log.Println(discover.GetServer())
		}
	}

	select {}

}
