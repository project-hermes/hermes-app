package types

import (
	"time"
)

// Dive is a dive
type Dive struct {
	SensorID   string        `json:"sensorId" firestore:"sensorId,omitempty"`
	StartTime  time.Time     `json:"startTime"  firestore:"startTime,omitempty"`
	EndTime    time.Time     `json:"duration"  firestore:"endTime,omitempty"`
	StartPoint GeoPoint      `json:"startPoint" firestore:"startPoint,omitempty"`
	EndPoint   GeoPoint      `json:"endPoint" firestore:"endPoint,omitempty"`
	SensorData []*SensorData `json:"sensorData" firestore:"sensorData,omitempty"`
}
