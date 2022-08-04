package interfaces

import "github.com/felipeamaralmeli/beholder/pkg/sources/contracts"

type Observable interface {
	GetObservableData() contracts.ObservableData
}
