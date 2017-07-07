package processors

import (
	"github.com/tmcgaughey/epagent-nozzle/metrics"
	"github.com/cloudfoundry/sonde-go/events"
)

type Processor interface {
	Process(e *events.Envelope) []metrics.WMetric
}
