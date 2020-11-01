package plugin

import "log"

type PluginGroup struct {
	plugins        []IPlugin
	registePlugins []IRegistePlugin
}

/**
服务注册插件组执行注册
*/
func (p *PluginGroup) Registe(ip string, port int) {
	for _, plugin := range p.registePlugins {
		plugin.Registe(ip, port)
	}
}

func (p *PluginGroup) AddPlugin(plugin IPlugin) {
	if plugin == nil {
		return
	}

	// 全局插件
	p.plugins = append(p.plugins, plugin)

	// 区分不同类型插件分类保存
	switch tp := plugin.(type) {
	case IRegistePlugin:
		p.registePlugins = append(p.registePlugins, tp)
	}
}

/**
启动所有插件
*/
func (p *PluginGroup) Start() error {
	for _, plugin := range p.plugins {
		err := plugin.Start()
		if err != nil {
			log.Println("[PluginGroup] Start error", err)
			return err
		}
	}
	return nil
}

func NewPluginGroup() IPluginGroup {
	return &PluginGroup{}
}
