package contracts

import "github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"

type ObservableData struct {
	Metrics []contracts.Metric
	Data    interface{}
}

func (s *ObservableData) ConvertDataToMetrics() {

}
