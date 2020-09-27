package stats

type Receiver struct {
	Version string  `json:"version"`
	Refresh int64   `json:"refresh"`
	History int64   `json:"history"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Overall struct {
	Latest   Summary `json:"latest"`
	Last1Min Summary `json:"last1min"`
}

type Summary struct {
	Start                      float64 `json:"start"`
	End                        float64 `json:"end"`
	Local                      Local   `json:"local"`
	Remote                     Remote  `json:"remote"`
	CPR                        CPR     `json:"cpr"`
	Tracks                     Tracks  `json:"tracks"`
	AltitudeSuppressed         int64   `json:"altitude_suppressed"`
	Messages                   int64   `json:"messages"`
	MaxDistanceInMetres        float64 `json:"max_distance_in_metres"`
	MaxDistanceInNauticalMiles float64 `json:"max_distance_in_nautical_miles"`
}

type Local struct {
	SamplesProcessed int64   `json:"samples_processed"`
	SamplesDropped   int64   `json:"samples_dropped"`
	ModeAC           int64   `json:"mode_ac"`
	ModeS            int64   `json:"mode_s"`
	Bad              int64   `json:"bad"`
	UnkownICAO       int64   `json:"unkown_icao"`
	Accepted         []int64 `json:"accepted"`
	Signal           float64 `json:"signal"`
	Noise            float64 `json:"noise"`
	PeakSignal       float64 `json:"peak_signal"`
	StrongSignals    int64   `json:"strong_signals"`
}

type Remote struct {
	ModeAC     int64   `json:"mode_ac"`
	ModeS      int64   `json:"mode_s"`
	Bad        int64   `json:"bad"`
	UnkownICAO int64   `json:"unkown_icao"`
	Accepted   []int64 `json:"accepted"`
}

type CPR struct {
	Surface               int64 `json:"surface"`
	Airborne              int64 `json:"airborne"`
	GlobalOK              int64 `json:"global_ok"`
	GlobalBad             int64 `json:"global_bad"`
	GlobalRange           int64 `json:"global_range"`
	GlobalSpeed           int64 `json:"global_speed"`
	GlobalSkipped         int64 `json:"global_skipped"`
	LocalOK               int64 `json:"local_ok"`
	LocalAircraftRelative int64 `json:"local_aircraft_relative"`
	LocalReceiverRelative int64 `json:"local_receiver_relative"`
	LocalSkipped          int64 `json:"local_skipped"`
	LocalRange            int64 `json:"local_range"`
	LocalSpeed            int64 `json:"local_speed"`
	Filtered              int64 `json:"filtered"`
}

type Tracks struct {
	All           int64 `json:"all"`
	SingleMessage int64 `json:"single_message"`
}
