package newrelic

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/newrelic/newrelicconfig"
	nr "github.com/newrelic/go-agent"
	"log"
)

type newrelic struct {
	Client nr.Application
}

func NewNewRelic(config interfaces.Config) interfaces.Sink {
	nrConfigs := config.GetConfigs().(newrelicconfig.NewRelicConfig)
	nrConfig := nr.NewConfig(nrConfigs.ApplicationName, nrConfigs.LicenseKey)
	nrConfig.Enabled = true
	nrConfig.CrossApplicationTracer.Enabled = nrConfigs.CrossApplicationTracer
	nrConfig.DistributedTracer.Enabled = nrConfigs.DistributedTracer
	nrConfig.ErrorCollector.IgnoreStatusCodes = append(nrConfig.ErrorCollector.IgnoreStatusCodes, nrConfigs.IgnoredStatusCodes...)

	app, err := nr.NewApplication(nrConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &newrelic{
		Client: app,
	}
}

func (s *newrelic) SendMetrics(metrics []*contracts.Metric) {
	panic("implement me!")
}
