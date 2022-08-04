package consts

const (
	DefaultConfigLocation = "configs/default.yaml"

	DatadogSinkName  = "datadog"
	NewRelicSinkName = "newrelic"

	IncrementMetricType    = "increment"
	DecrementMetricType    = "decrement"
	GaugeMetricType        = "gauge"
	CountMetricType        = "count"
	HistogramMetricType    = "histogram"
	DistributionMetricType = "distribution"
	SetMetricType          = "set"
	TimingMetricType       = "timing"
	TimeInMillisMetricType = "timeInMillis"
	TraceType              = "trace"
)
