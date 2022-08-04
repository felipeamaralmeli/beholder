package newrelicconfig

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
)

type NewRelicConfig struct {
	ApplicationName        string
	LicenseKey             string
	CrossApplicationTracer bool
	DistributedTracer      bool
	IgnoredStatusCodes     []int
}

func NewNewRelicConfig() interfaces.Config {
	return &NewRelicConfig{}
}

func (s NewRelicConfig) BuildDefaults(configs map[interface{}]interface{}) interfaces.Config {
	sinkMap := configs["sinks"].(map[string]interface{})
	nrConfigs := sinkMap["newRelic"].(map[string]interface{})
	ignoredStatusCodesStart := nrConfigs["ignoredStatusCodesStart"].(int)
	ignoredStatusCodesEnd := nrConfigs["ignoredStatusCodesEnd"].(int)
	return &NewRelicConfig{
		ApplicationName:        nrConfigs["applicationName"].(string),
		LicenseKey:             nrConfigs["licenseKey"].(string),
		CrossApplicationTracer: nrConfigs["crossApplicationTracer"].(bool),
		DistributedTracer:      nrConfigs["distributedTracer"].(bool),
		IgnoredStatusCodes:     buildIgnoredStatusCodes(ignoredStatusCodesStart, ignoredStatusCodesEnd),
	}
}

func (s *NewRelicConfig) GetConfigs() interface{} {
	return s
}

func buildIgnoredStatusCodes(start, end int) []int {
	ignoredStatusCodes := make([]int, 0)

	for i := start; i <= end; i++ {
		ignoredStatusCodes = append(ignoredStatusCodes, i)
	}

	return ignoredStatusCodes
}
