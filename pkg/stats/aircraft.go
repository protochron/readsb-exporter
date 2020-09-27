package stats

type Aircraft struct {
	Now      float64 `json:"now"`
	Messages int64   `json:"messages"`
	Aircraft []Craft `json:"aircraft"`
}

type Craft struct {
	Hex            string  `json:"hex"`
	Flight         string  `json:"flight"`
	AltBaro        int64   `json:"alt_baro"`
	AltGeom        int64   `json:"alt_geom"`
	Gs             float64 `json:"gs"`
	IAS            int64   `json:"ias"`
	TAS            int64   `json:"tas"`
	Mach           float64 `json:"mach"`
	Track          float64 `json:"track"`
	TrackRate      float64 `json:"track_rate"`
	Roll           float64 `json:"roll"`
	MagHeading     float64 `json:"mag_heading"`
	BaroRate       int64   `json:"baro_rate"`
	GeomRate       int64   `json:"geom_rate"`
	Squawk         string  `json:"squawk"`
	Emergency      string  `json:"emergency"`
	Category       string  `json:"category"`
	NavQNH         float64 `json:"nav_qnh"`
	NavAltitudeMCP int64   `json:"nav_altitude_mcp"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	NIC            int64   `json:"nic"`
	RC             int64   `json:"rc"`
	SeenPos        float64 `json:"seen_pos"`
	Version        int64   `json:"version"`
	NicBaro        int64   `json:"nic_baro"`
	NACP           int64   `json:"nac_p"`
	NACV           int64   `json:"nac_v"`
	SIL            int64   `json:"sil"`
	SILType        string  `json:"sil_type"`
	GVA            int64   `json:"gva"`
	SDA            int64   `json:"sda"`
	Alert          int64   `json:"alert"`
	SPI            int64   `json:"spi"`
	Messages       int64   `json:"messages"`
	Seen           float64 `json:"seen"`
	RSSI           float64 `json:"rssi"`
}
