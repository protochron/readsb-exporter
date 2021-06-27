# readsb-exporter

A Prometheus exporter for dump1090 metrics.

Based heavily on https://github.com/claws/dump1090-exporter.

## Build
You need to have a recent version of Go installed.

```
make build
```

## Running
Once you have built `readsb-exporter`, you need to configure it with:
- the latitude and longitude of the location of your receiver
- the path to where dump1090 metrics are being written as .json files

```
Usage of ./readsb-exporter:
  -lat float
        Latitude. Values from stats.json will override this.
  -listen-address string
        Listen address (default ":9105")
  -lon float
        Longitude. Values from stats.json will override this.
  -metrics-root string
        Path to .json files (default "/run/readsb")
  -threshold-seconds float
        Filter aircraft to only those seen within the last n seconds. (default 15)
```

As an example:
```
$ readsb-exporter \
  -lat 100 \
  -lon 100 \
  -metrics-root /var/run/readsb
```
