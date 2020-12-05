package registe

import (
	"testing"

	"github.com/gmsorg/gms/plugin"
)

func TestConsulRegistePlugin_Start(t *testing.T) {
	consulRegistePlugin := NewConsulRegistePlugin("testServerName", []string{"127.0.0.1:8500"})
	consulRegistePlugin.Start()
	registe := consulRegistePlugin.(plugin.IRegistePlugin)
	registe.Registe("127.0.0.1", 1024)

}
