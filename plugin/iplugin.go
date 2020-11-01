package plugin

type IPluginGroup interface {
	Start() error
	AddPlugin(plugin IPlugin)
	Registe(ip string, port int)
}

type IPlugin interface {
	Start() error
}

type (
	// 服务注册插件接口
	IRegistePlugin interface {
		Registe(ip string, port int) error
	}
)
