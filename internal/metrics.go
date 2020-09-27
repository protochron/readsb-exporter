package internal

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "readsb"
)

var (
	AircraftObservedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "observed",
		Namespace: Namespace,
		Subsystem: "aircraft",
		Help:      "Number of aircraft observed",
	})

	AircraftObservedWithPositionGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "observed_with_pos",
		Namespace: Namespace,
		Subsystem: "aircraft",
		Help:      "Number of aircraft observed",
	})

	AircraftObservedWithMLATCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "observed_with_mlat",
		Namespace: "readsb",
		Subsystem: Namespace,
		Help:      "Number of aircraft observed with MLAT",
	})

	AircraftMaxRange = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "max_range",
		Namespace: Namespace,
		Subsystem: "aircraft",
		Help:      "Max range of observed aircraft",
	})

	MessagesTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "messages_total",
		Namespace: Namespace,
		Subsystem: "aircraft",
		Help:      "Number of Mode-S messages",
	})

	Receiver = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "receiver",
		Namespace: Namespace,
		Help:      "Receiver info",
	}, []string{"version", "lat", "lon"})
)

func init() {
	prometheus.MustRegister(AircraftObservedGauge)
	prometheus.MustRegister(AircraftObservedWithPositionGauge)
	prometheus.MustRegister(AircraftObservedWithMLATCount)
	prometheus.MustRegister(MessagesTotal)
	prometheus.MustRegister(AircraftMaxRange)
	prometheus.MustRegister(Receiver)
}
