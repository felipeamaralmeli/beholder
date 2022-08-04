package interfaces

import "github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"

type Sink interface {
	GetID() string
	SendMetrics(metrics []*contracts.Metric)
}
