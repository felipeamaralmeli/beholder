package telemetryconfigurator

import (
	"errors"
	"github.com/felipeamaralmeli/beholder/internal/configurationreader"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/datadog"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/datadog/datadogconfig"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/newrelic"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/newrelic/newrelicconfig"
)

type TelemetryConfigurator struct {
	Sinks   []interfaces.Sink
	Configs map[interface{}]interface{}
}

func NewTelemetryConfigurator() *TelemetryConfigurator {
	configs, err := configurationreader.ReadConfigsFromFile(consts.DefaultConfigLocation)
	if err != nil {
		panic(err)
	}
	return &TelemetryConfigurator{
		Configs: configs,
	}
}

func (s *TelemetryConfigurator) ConfigureDefaultSinks() error {
	dDogSink, err := s.configureDatadogSink(s.Configs)
	if err == nil {
		s.appendSink(dDogSink)
	}

	nrSink, err := s.configureNewRelicSink(s.Configs)
	if err == nil {
		s.appendSink(nrSink)
	}

	return nil
}

func (s *TelemetryConfigurator) GetSinkByName(sinkName string) (interfaces.Sink, error) {
	for _, sink := range s.Sinks {
		if sink.GetID() == sinkName {
			return sink, nil
		}
	}

	return nil, errors.New("required sink is not configured or available")
}

func (s *TelemetryConfigurator) ConfigureAndAppendSink(sinkConfigurator func() interfaces.Config, sinkConstructor func(config interfaces.Config) (interfaces.Sink, error)) ([]interfaces.Sink, error) {
	sinkConfig := sinkConfigurator().BuildDefaults(s.Configs)
	sink, err := sinkConstructor(sinkConfig)
	if err == nil {
		s.appendSink(sink)
		return s.Sinks, nil
	}

	return nil, err
}

func (s *TelemetryConfigurator) configureDatadogSink(configs map[interface{}]interface{}) (interfaces.Sink, error) {
	datadogConfig := datadogconfig.NewDatadogConfig().BuildDefaults(configs)
	dDogSink, err := datadog.NewDatadogSink(datadogConfig)
	if err != nil {
		return nil, err
	}

	return dDogSink, nil
}

func (s *TelemetryConfigurator) configureNewRelicSink(configs map[interface{}]interface{}) (interfaces.Sink, error) {
	nrConfig := newrelicconfig.NewNewRelicConfig().BuildDefaults(configs)
	nrSink, err := newrelic.NewNewRelicSink(nrConfig)
	if err != nil {
		return nil, err
	}

	return nrSink, nil
}

func (s *TelemetryConfigurator) appendSink(sink interfaces.Sink) {
	s.Sinks = append(s.Sinks, sink)
}
