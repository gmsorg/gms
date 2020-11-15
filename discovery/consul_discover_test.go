package discovery

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/akkagao/gms/plugin"
	"github.com/akkagao/gms/plugin/registe"
)

func TestNewConsulDiscover(t *testing.T) {
	plugins := registe.NewConsulRegistePlugin("gmsDemo", []string{"localhost:8500"})
	plugins.Start()

	registePlugin := plugins.(plugin.IRegistePlugin)
	registePlugin.Registe("127.0.0.1", 1024)

	discover, err := NewConsulDiscover("gmsDemo", []string{"localhost:8500"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(discover.GetServer())

	timeout := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeout.C:
			log.Println(discover.GetServer())
			rand.Seed(time.Now().UnixNano())
			// registePlugin.Registe("127.0.0.1", rand.Intn(6000)+1024)
		}
	}

	select {}
}
