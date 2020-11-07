package discovery

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/akkagao/gms/plugin"
	"github.com/akkagao/gms/plugin/registe"
)

func TestNewRedisDiscover(t *testing.T) {
	plugins := registe.NewRedisRegistePlugin("gmsDemo", []string{"127.0.0.1:6379"})
	plugins.Start()

	registePlugin := plugins.(plugin.IRegistePlugin)
	registePlugin.Registe("127.0.0.1", 1024)

	discover, err := NewRedisDiscover("gmsDemo", []string{"127.0.0.1:6379"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(discover.GetServer())

	// time.Sleep(10 * time.Second)
	fmt.Println("registe 1025")
	registePlugin.Registe("127.0.0.1", 1025)
	registePlugin.Registe("127.0.0.1", 1025)

	time.Sleep(3 * time.Second)
	fmt.Println(discover.GetServer())

}
