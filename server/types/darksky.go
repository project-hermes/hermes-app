package types

// ForecastResponse is the response from the weather service
type ForecastResponse struct {
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
	Currently *DataPoint `json:"currently,omitempty"`
	Minutely  *Forecast  `json:"minutely,omitempty"`
	Hourly    *Forecast  `json:"hourly,omitempty"`
	Daily     *Forecast  `json:"daily,omitempty"`
	Alerts    []*Alert   `json:"alerts,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
}

// Forecast is the weather over a certain period of time
type Forecast struct {
	Data    []DataPoint `json:"data,omitempty"`
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
}

// Alert represent severe weather warnings
type Alert struct {
	Description string   `json:"description,omitempty"`
	Expires     int64    `json:"expires,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Time        int64    `json:"time,omitempty"`
	Title       string   `json:"title,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

// Flags contains extra data about the request
type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}

// DataPoint contains specific data about the weather conditions during a certain period of time
type DataPoint struct {
	ApparentTemperature         float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	NearestStormBearing         float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance        float64 `json:"nearestStormDistance,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	PrecipAccumulation          float64 `json:"precipAccumulation,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 int64   `json:"sunriseTime,omitempty"`
	SunsetTime                  int64   `json:"sunsetTime,omitempty"`
	Temperature                 float64 `json:"temperature,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int64   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int64   `json:"temperatureLowTime,omitempty"`
	Time                        int64   `json:"time,omitempty"`
	UvIndex                     int64   `json:"uvIndex,omitempty"`
	UvIndexTime                 int64   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	WindBearing                 float64 `json:"windBearing,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int64   `json:"windGustTime,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
}
