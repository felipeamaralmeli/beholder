package datadogconfig

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"strconv"
)

type DatadogConfig struct {
	Addr string
	Port string
	Tags []string
}

func NewDatadogConfig() interfaces.Config {
	return &DatadogConfig{}
}

func (s *DatadogConfig) BuildDefaults(configs map[interface{}]interface{}) interfaces.Config {
	sinkMap := configs["sinks"].(map[string]interface{})
	datadogConfigs := sinkMap["datadog"].(map[string]interface{})
	return &DatadogConfig{
		Addr: datadogConfigs["addr"].(string),
		Port: strconv.Itoa(datadogConfigs["port"].(int)),
		Tags: s.buildTags(datadogConfigs["tags"].([]interface{})),
	}
}

func (s *DatadogConfig) GetConfigs() interface{} {
	return s
}

func (s *DatadogConfig) buildTags(tags []interface{}) []string {
	stringTags := make([]string, 0)
	for _, tag := range tags {
		stringTags = append(stringTags, tag.(string))
	}

	return stringTags
}
