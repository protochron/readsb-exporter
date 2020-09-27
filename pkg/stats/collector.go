package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/protochron/readsb-exporter/internal"
	"github.com/protochron/readsb-exporter/internal/stats"
	"go.uber.org/zap"
)

type Collector struct {
	Logger     *zap.Logger
	MetricRoot string
	Threshold  float64
	Lat        float64
	Lon        float64
	aircraft   *Aircraft
	overall    *Overall
	receiver   *Receiver
}

func (c *Collector) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	aircraft := Aircraft{}
	receiver := Receiver{}
	summary := Overall{}

	aircraftBytes, err := ioutil.ReadFile(path.Join(c.MetricRoot, "aircraft.json"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to read aircraft.json"))
		c.Logger.Error("unable to read aircraft.json", zap.Error(err))
		return
	}

	receiverBytes, err := ioutil.ReadFile(path.Join(c.MetricRoot, "receiver.json"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to read receiver.json"))
		c.Logger.Error("unable to read receiver.json", zap.Error(err))
		return
	}

	statsBytes, err := ioutil.ReadFile(path.Join(c.MetricRoot, "stats.json"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to read stats.json"))
		c.Logger.Error("unable to read stats.json", zap.Error(err))
		return
	}

	err = json.Unmarshal(aircraftBytes, &aircraft)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to unmarshal aircraft.json"))
		c.Logger.Error("error unmarshaling", zap.Error(err))
		return
	}
	err = json.Unmarshal(receiverBytes, &receiver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to unmarshal receiver.json"))
		c.Logger.Error("error unmarshaling", zap.Error(err))
		return
	}
	err = json.Unmarshal(statsBytes, &summary)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("unable to unmarshal stats.json"))
		c.Logger.Error("error unmarshaling", zap.Error(err))
		return
	}

	c.aircraft = &aircraft
	c.receiver = &receiver
	c.overall = &summary

	c.processMetrics()

	promhttp.Handler().ServeHTTP(w, r)
}

func (c *Collector) processMetrics() {
	lat := strconv.FormatFloat(c.receiver.Lat, 'f', -1, 64)
	lon := strconv.FormatFloat(c.receiver.Lon, 'f', -1, 64)
	if c.Lat == 0.0 {
		c.Lat = c.receiver.Lat
	}

	if c.Lon == 0.0 {
		c.Lon = c.receiver.Lon
	}

	internal.Receiver.With(prometheus.Labels{
		"version": c.receiver.Version,
		"lat":     lat,
		"lon":     lon,
	}).Set(1)

	c.processAircraft()
	c.processStats()
}

func (c *Collector) processAircraft() {
	var observed, withPos int
	var maxRange float64

	for _, a := range c.aircraft.Aircraft {
		if a.Seen < c.Threshold {
			observed += 1
		}

		if a.SeenPos < c.Threshold {
			withPos += 1

			distance := stats.Haversine(
				stats.LatLon{Lat: c.Lat, Lon: c.Lon},
				stats.LatLon{Lat: a.Lat, Lon: a.Lon},
			)

			if distance > maxRange {
				maxRange = distance
			}
		}
	}

	// TODO: add MLAT
	internal.AircraftObservedGauge.Set(float64(observed))
	internal.AircraftMaxRange.Set(maxRange)
	internal.AircraftObservedWithPositionGauge.Set(float64(withPos))
	internal.MessagesTotal.Set(float64(c.aircraft.Messages))
}

func (c *Collector) processStats() {
	internal.MaxDistanceInNauticalMiles.Set(c.overall.Last1Min.MaxDistanceInNauticalMiles)
	internal.MaxDistanceInMetres.Set(c.overall.Last1Min.MaxDistanceInMetres)
	internal.Signal.Set(c.overall.Last1Min.Local.Signal)
	internal.PeakSignal.Set(c.overall.Last1Min.Local.PeakSignal)
	internal.Noise.Set(c.overall.Last1Min.Local.Noise)
	internal.StrongSignals.Set(float64(c.overall.Last1Min.Local.StrongSignals))
}
