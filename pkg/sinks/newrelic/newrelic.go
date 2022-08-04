package newrelic

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/newrelic/newrelicconfig"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelicSink struct {
	ID          string
	Client      *nr.Application
	Transaction *nr.Transaction
}

func NewNewRelicSink(config interfaces.Config) (interfaces.Sink, error) {
	nrConfigs := config.GetConfigs().(*newrelicconfig.NewRelicConfig)

	app, err := nr.NewApplication(
		nr.ConfigAppName(nrConfigs.ApplicationName),
		nr.ConfigLicense(nrConfigs.LicenseKey),
		nr.ConfigEnabled(true),
		nr.ConfigFromEnvironment())
	if err != nil {
		return nil, err
	}

	return &NewRelicSink{
		ID:     consts.NewRelicSinkName,
		Client: app,
	}, nil
}

func (s *NewRelicSink) SendMetrics(_ ...*contracts.Metric) {
	if s.Transaction != nil {
		s.Transaction.End()
	}
}

func (s *NewRelicSink) GetID() string {
	return s.ID
}

func (s *NewRelicSink) StartTransaction(name string) interface{} {
	tx := s.Client.StartTransaction(name)
	s.Transaction = tx
	return tx
}
