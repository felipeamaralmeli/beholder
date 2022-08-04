package interfaces

type Config interface {
	BuildDefaults(configs map[interface{}]interface{}) Config
	GetConfigs() interface{}
}
