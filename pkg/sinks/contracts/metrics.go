package contracts

import (
	"errors"
	"fmt"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
	"time"
)

type Metric struct {
	Key          string
	Tags         []string
	Value        interface{}
	Rate         float64
	MetricType   string
	Destinations []interfaces.Sink
}

var metricTypeMap = map[string][]string{
	"float64":       {consts.HistogramMetricType, consts.GaugeMetricType, consts.DistributionMetricType, consts.TimeInMillisMetricType},
	"int64":         {consts.CountMetricType},
	"string":        {consts.SetMetricType},
	"time.Duration": {consts.TimingMetricType},
	"struct":        {consts.SetMetricType},
}

func NewMetric(key string, value interface{}, tags []string, rate float64, metricType string) (*Metric, error) {
	metric := &Metric{
		Key:        key,
		Tags:       tags,
		Value:      value,
		Rate:       rate,
		MetricType: metricType,
	}

	allowedTypes := metric.getAllowedTypesForValue()
	for _, allowedType := range allowedTypes {
		if allowedType == metricType {
			return metric, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Metric value (%T) is incompatible with metric type %s. Maybe you should use: %v", metric.Value, metric.MetricType, allowedTypes))
}

func (s *Metric) IsFloat64Metric() bool {
	_, ok := s.Value.(float64)
	return ok
}

func (s *Metric) IsInt64Metric() bool {
	_, ok := s.Value.(int64)
	return ok
}

func (s *Metric) IsStringMetric() bool {
	_, ok := s.Value.(string)
	return ok
}

func (s *Metric) IsDurationMetric() bool {
	_, ok := s.Value.(time.Duration)
	return ok
}

func (s *Metric) getAllowedTypesForValue() []string {
	metricType := fmt.Sprintf("%T", s.Value)
	values, ok := metricTypeMap[metricType]
	if ok {
		return values
	}

	return []string{consts.IncrementMetricType, consts.DecrementMetricType}
}
