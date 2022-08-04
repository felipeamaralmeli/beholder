package contracts

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/contracts"
	"github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"
)

type ObservableData struct {
	Metrics   []*contracts.Metric
	Observers []interfaces.Sink
}

func (s *ObservableData) Register(observer interfaces.Sink) {
	s.Observers = append(s.Observers, observer)
}

func (s *ObservableData) Unregister(observer interfaces.Sink) {
	s.Observers = removeFromSlice(s.Observers, observer)
}

func (s *ObservableData) Notify() {
	for _, observer := range s.Observers {
		observer.SendMetrics(s.Metrics)
	}
}

func removeFromSlice(observerList []interfaces.Sink, observerToRemove interfaces.Sink) []interfaces.Sink {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
