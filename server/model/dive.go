package model

import (
	"time"
)

type Dive struct {
	Name        string      `json:"name"`
	StartTime   time.Time   `json:"startTime"`
	Duration    int         `json:"duration"`
	StartPoint  GeoPoint    `json:"startPoint"`
	DeltaPoints []*GeoPoint `json:"deltaPoints"`
	EndPoint    GeoPoint    `json:"endPoint"`
}