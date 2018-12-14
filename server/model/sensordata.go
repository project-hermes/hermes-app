package model

import (
	
)

// parseToFloat will take an unknown value and attempt to make it a float64
func parseToFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case int64:
		return float64(v)
	case int:
		return float64(v)
	default:
		return 0.0
	}
}

// MapSensorData will convert map to SensorData object
func MapSensorData(data map[string]interface{}) SensorData {
	return SensorData {
		Depth: parseToFloat(data["depth"]),
		RawPressure: int(data["rawPressure"].(int64)),
		RawTemp: int(data["rawTemp"].(int64)),
		Temp: parseToFloat(data["temp"]),
		Time: int(data["time"].(int64)),
	}
}