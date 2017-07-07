package processors

import (
	"github.com/ckunst/epagent-nozzle/metrics"
	"github.com/cloudfoundry/sonde-go/events"
)

type Processor interface {
	Process(e *events.Envelope) []metrics.WMetric
}
