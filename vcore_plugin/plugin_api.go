package vcore_plugin

type PluginHelp struct {
	Name string
	Description string
	Type []string
}

type PluginFunction struct {
	Name string
	Type []string
	Func func(string, ...interface{}) error
}

type Plugin interface {
	Func(name string, data ...interface{}) error
	
}