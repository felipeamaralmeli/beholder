package interfaces

import "github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"

type Sink interface {
	SendMetrics(metrics []*contracts.Metric)
}
