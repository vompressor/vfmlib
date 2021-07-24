package vcore_loader

import (
	"errors"
	"fmt"
	"plugin"

	"github.com/vompressor/vfmlib/vcore_plugin"
)

var loadedPlugin = make([]vcore_plugin.PluginInfo, 0)

var commandTable = make(map[string] vcore_plugin.Plugin)

// Load plugin and add plugin info
func Load(path string) (vcore_plugin.Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	sym, err := plug.Lookup("VCore")
	if err != nil {
		return nil, err
	}

	var vcore vcore_plugin.Plugin
	vcore, ok := sym.(vcore_plugin.Plugin)
	if !ok {
		return nil, errors.New("")
	}

	vcore.PluginLoaded()
	info := vcore.GetPluginInfo()
	info.Plugin = vcore
	loadedPlugin = append(loadedPlugin, info)

	for _, n := range vcore.Help() {
		commandTable[info.PluginName + "." + n.Name] = vcore
	}

	return vcore, nil
}

func GetLoadedPluginInfo() []vcore_plugin.PluginInfo {
	return loadedPlugin
}

func PrintLoadedPluginInfo() {
	for _, n := range loadedPlugin {
		fmt.Printf("%s@%s - %s\n", n.PluginName, n.PluginVersion, n.PluginDescription)
	}
}
