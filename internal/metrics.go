package internal

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "readsb"
	Aircraft  = "aircraft"
	Stats     = "stats"
)

var (
	AircraftObservedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "observed",
		Namespace: Namespace,
		Subsystem: Aircraft,
		Help:      "Number of aircraft observed",
	})

	AircraftObservedWithPositionGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "observed_with_pos",
		Namespace: Namespace,
		Subsystem: Aircraft,
		Help:      "Number of aircraft observed",
	})

	AircraftObservedWithMLATCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "observed_with_mlat",
		Namespace: Namespace,
		Subsystem: Aircraft,
		Help:      "Number of aircraft observed with MLAT",
	})

	AircraftMaxRange = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "max_range",
		Namespace: Namespace,
		Subsystem: Aircraft,
		Help:      "Max range of observed aircraft",
	})

	MessagesTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "messages_total",
		Namespace: Namespace,
		Subsystem: Aircraft,
		Help:      "Number of Mode-S messages",
	})

	Receiver = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "receiver",
		Namespace: Namespace,
		Help:      "Receiver info",
	}, []string{"version", "lat", "lon"})

	MaxDistanceInMetres = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "max_distance_metres",
		Namespace: Namespace,
		Subsystem: "stats",
		Help:      "Max observed aircraft distance (metres)",
	})

	MaxDistanceInNauticalMiles = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "max_distance_nautical_miles",
		Namespace: Namespace,
		Subsystem: Stats,
		Help:      "Max observed aircraft distance (nautical miles)",
	})

	Signal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "signal",
		Namespace: Namespace,
		Subsystem: Stats,
		Help:      "Signal",
	})

	PeakSignal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "peak_signal",
		Namespace: Namespace,
		Subsystem: Stats,
		Help:      "Peak Signal",
	})

	Noise = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "noise",
		Namespace: Namespace,
		Subsystem: Stats,
		Help:      "Noise",
	})

	StrongSignals = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "strong_signals",
		Namespace: Namespace,
		Subsystem: Stats,
		Help:      "Strong signals",
	})
)

func init() {
	prometheus.MustRegister(AircraftObservedGauge)
	prometheus.MustRegister(AircraftObservedWithPositionGauge)
	prometheus.MustRegister(AircraftObservedWithMLATCount)
	prometheus.MustRegister(MessagesTotal)
	prometheus.MustRegister(AircraftMaxRange)
	prometheus.MustRegister(Receiver)
	prometheus.MustRegister(MaxDistanceInNauticalMiles)
	prometheus.MustRegister(MaxDistanceInMetres)
	prometheus.MustRegister(Signal)
	prometheus.MustRegister(PeakSignal)
	prometheus.MustRegister(Noise)
	prometheus.MustRegister(StrongSignals)
}
