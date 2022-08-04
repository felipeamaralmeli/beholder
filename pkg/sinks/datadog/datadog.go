package datadog

import (
	"fmt"
	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/datadog/datadogconfig"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"log"
	"time"
)

type datadog struct {
	ID     string
	Client statsd.ClientInterface
}

func NewDatadogSink(config interfaces.Config) (interfaces.Sink, error) {
	ddogConfig := config.GetConfigs().(*datadogconfig.DatadogConfig)
	connString := fmt.Sprintf("%s:%s", ddogConfig.Addr, ddogConfig.Port)
	client, err := statsd.New(connString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &datadog{
		ID:     consts.DatadogSinkName,
		Client: client,
	}, nil
}

func (s *datadog) GetID() string {
	return s.ID
}

func (s *datadog) SendMetrics(metrics ...*contracts.Metric) {
	for _, metric := range metrics {
		s.getFunctionByMetricType(metric.MetricType)(metric)
	}
}

func (s *datadog) StartTransaction(name string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *datadog) Increment(metric *contracts.Metric) {
	_ = s.Client.Incr(metric.Key, metric.Tags, metric.Rate)
}

func (s *datadog) Decrement(metric *contracts.Metric) {
	_ = s.Client.Decr(metric.Key, metric.Tags, metric.Rate)
}

func (s *datadog) Gauge(metric *contracts.Metric) {
	if metric.IsFloat64Metric() {
		_ = s.Client.Gauge(metric.Key, metric.Value.(float64), metric.Tags, metric.Rate)
	}
}

func (s *datadog) Count(metric *contracts.Metric) {
	if metric.IsInt64Metric() {
		_ = s.Client.Count(metric.Key, metric.Value.(int64), metric.Tags, metric.Rate)
	}
}

func (s *datadog) Histogram(metric *contracts.Metric) {
	if metric.IsFloat64Metric() {
		_ = s.Client.Histogram(metric.Key, metric.Value.(float64), metric.Tags, metric.Rate)
	}
}

func (s *datadog) Distribution(metric *contracts.Metric) {
	if metric.IsFloat64Metric() {
		_ = s.Client.Distribution(metric.Key, metric.Value.(float64), metric.Tags, metric.Rate)
	}
}

func (s *datadog) Set(metric *contracts.Metric) {
	if metric.IsStringMetric() {
		_ = s.Client.Set(metric.Key, metric.Value.(string), metric.Tags, metric.Rate)
	}
}

func (s *datadog) Timing(metric *contracts.Metric) {
	if metric.IsDurationMetric() {
		_ = s.Client.Timing(metric.Key, metric.Value.(time.Duration), metric.Tags, metric.Rate)
	}
}

func (s *datadog) TimeInMilliseconds(metric *contracts.Metric) {
	if metric.IsFloat64Metric() {
		_ = s.Client.TimeInMilliseconds(metric.Key, metric.Value.(float64), metric.Tags, metric.Rate)
	}
}

func (s *datadog) getFunctionByMetricType(metricType string) func(metric *contracts.Metric) {
	switch metricType {
	case consts.IncrementMetricType:
		return s.Increment
	case consts.DecrementMetricType:
		return s.Decrement
	case consts.GaugeMetricType:
		return s.Gauge
	case consts.CountMetricType:
		return s.Count
	case consts.HistogramMetricType:
		return s.Histogram
	case consts.DistributionMetricType:
		return s.Distribution
	case consts.SetMetricType:
		return s.Set
	case consts.TimingMetricType:
		return s.Timing
	case consts.TimeInMillisMetricType:
		return s.TimeInMilliseconds
	default:
		return s.Increment
	}
}
