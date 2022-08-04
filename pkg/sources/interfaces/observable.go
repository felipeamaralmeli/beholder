package interfaces

import "github.com/felipeamaralmeli/beholder/pkg/sinks/interfaces"

type Observable interface {
	Register(observer interfaces.Sink)
	Unregister(observer interfaces.Sink)
	Notify()
}
