package vcore_loader

import (
	"errors"
	"fmt"
	"plugin"
	"strings"

	"github.com/vompressor/vfmlib/vcore_plugin"
)

var loadedPlugin = make([]vcore_plugin.PluginInfo, 0)

var commandTable = make(map[string]vcore_plugin.Plugin)

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
		return nil, errors.New("load err")
	}

	vcore.PluginLoaded()
	info := vcore.GetPluginInfo()
	loadedPlugin = append(loadedPlugin, info)

	for _, n := range vcore.Help() {
		commandTable[info.PluginName+"."+n.Name] = vcore
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

func GetCommands() []string {
	x := make([]string, 0)
	for k := range commandTable {
		x = append(x, k)
	}
	return x
}

func PrintCommands() {
	for k := range commandTable {
		println(k)
	}
}

func Call(name string, arg ...string) error {
	if strings.Contains(name, ".") {
		x := commandTable[name]
		if x == nil {
			return errors.New("command not found")
		}

		return x.Func(name, arg...)
	}

	isDup := 0
	var f vcore_plugin.Plugin
	for k, v := range commandTable {
		if strings.Split(k, ".")[1] == name {
			isDup++
			f = v
		}
	}
	if isDup == 0 {
		return errors.New("command not found")
	} else if isDup != 0 {
		return errors.New("command duplicated")
	}
	return f.Func(name, arg...)
}
