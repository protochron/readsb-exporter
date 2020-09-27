package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/protochron/readsb-exporter/pkg/stats"
	"go.uber.org/zap"
)

var (
	metricRoot       string
	listenAddress    string
	lat              float64
	lon              float64
	thresholdSeconds float64
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&metricRoot, "metrics-root", "/run/readsb", "Path to .json files")
	flag.StringVar(&listenAddress, "listen-address", ":9105", "Listen address")
	flag.Float64Var(&lat, "lat", 0.0, "Latitude. Values from stats.json will override this.")
	flag.Float64Var(&lon, "lon", 0.0, "Longitude. Values from stats.json will override this.")
	flag.Float64Var(&thresholdSeconds, "threshold-seconds", 15.0, "Filter aircraft to only those seen within the last n seconds.")
	flag.Parse()

	c := stats.Collector{
		Logger:     logger,
		MetricRoot: metricRoot,
		Lat:        lat,
		Lon:        lon,
		Threshold:  thresholdSeconds,
	}

	logger.Info("starting exporter", zap.String("root", metricRoot), zap.String("addr", listenAddress), zap.Float64("lat", lat), zap.Float64("lon", lon))

	http.Handle("/metrics", http.HandlerFunc(c.MetricsHandler))
	err = http.ListenAndServe(listenAddress, nil)
	if err != nil {
		logger.Fatal("error while running server", zap.Error(err))
	}
}
