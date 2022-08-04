package main

import (
	"github.com/felipeamaralmeli/beholder/internal/configurationreader"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/datadog"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/datadog/datadogconfig"
	"log"
)

func main() {
	defaults, err := configurationreader.ReadConfigsFromFile(consts.DefaultConfigLocation)
	if err != nil {
		panic(err)
	}
	config := datadogconfig.NewDatadogConfig().BuildDefaults(defaults)

	ddogSink, err := datadog.NewDatadogSink(config)
	if err != nil {
		panic(err)
	}

	metric, err := contracts.NewMetric("test", float64(0), nil, 1.0, consts.GaugeMetricType)
	if err != nil {
		log.Fatal(err)
	}

	metrics := []*contracts.Metric{metric}

	ddogSink.SendMetrics(metrics)
}
