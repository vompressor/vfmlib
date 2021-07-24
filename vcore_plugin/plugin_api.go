package vcore_plugin

type PluginHelp struct {
	Name              string
	ShortDescription  string
	DetailDescription string
	Type              []string
	TypeDescription   []string
}

type PluginInfo struct {
	PluginName        string
	PluginVersion     string
	PluginDescription string
	Plugin
}

type Plugin interface {
	GetPluginInfo() PluginInfo
	Func(string, ...string) error
	Help() []PluginHelp

	PluginLoaded()
}
