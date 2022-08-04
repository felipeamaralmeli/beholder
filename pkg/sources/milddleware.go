package sources

import (
	"github.com/felipeamaralmeli/beholder/pkg/sinks/consts"
	"github.com/felipeamaralmeli/beholder/pkg/telemetryconfigurator"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

type Adapter func(handler http.Handler) http.Handler

func TelemetryTraceMiddleware(telemetryConfigurator *telemetryconfigurator.TelemetryConfigurator) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sink, err := telemetryConfigurator.GetSinkByName(consts.NewRelicSinkName)
			if err == nil {
				tx := sink.StartTransaction(r.RequestURI).(*newrelic.Transaction)
				defer sink.SendMetrics()
				tx.SetWebRequestHTTP(r)
				writer := tx.SetWebResponse(w)
				h.ServeHTTP(writer, r)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
