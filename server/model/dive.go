package model

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	pblatlng "google.golang.org/genproto/googleapis/type/latlng"
)

// Dive is a dive
type Dive struct {
	SensorID   string    `json:"sensorId" firestore:"sensorId,omitempty"`
	StartTime  time.Time `json:"startTime"  firestore:"startTime,omitempty"`
	EndTime    time.Time `json:"duration"  firestore:"endTime,omitempty"`
	StartPoint GeoPoint  `json:"startPoint" firestore:"startPoint,omitempty"`
	EndPoint   GeoPoint  `json:"endPoint" firestore:"endPoint,omitempty"`
}

// MapDive will convert a map into a dive
func MapDive(data map[string]interface{}) Dive {

	startPoint := data["startPoint"].(*pblatlng.LatLng)
	endPoint := data["endPoint"].(*pblatlng.LatLng)

	return Dive {
		SensorID: data["sensorId"].(string),
		StartTime: data["startTime"].(time.Time),
		EndTime: data["endTime"].(time.Time),
		StartPoint: GeoPoint{Lat: startPoint.Latitude, Long: startPoint.Longitude},
		EndPoint: GeoPoint{Lat: endPoint.Latitude, Long: endPoint.Longitude},
	}
}

// DiveInterface is an interface for interacting with dive results
type DiveInterface interface {
	List(context.Context) []*Dive
}

// DiveImplementation is an object for CRUD operations on dives
type DiveImplementation struct {
	client *firestore.Client
}

// NewDiveImplementation will return an object for working with dives
func NewDiveImplementation(client *firestore.Client) DiveInterface {
	return &DiveImplementation{
		client: client,
	}
}

// List will fetch all of the dives
func (di DiveImplementation) List(ctx context.Context) []*Dive {
	iter := di.client.Collection("dive").Documents(context.Background())

	var dives []*Dive
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("failed to iterate: %v", err)
		}

		dive := MapDive(doc.Data())

		dives = append(dives, &dive)
	}

	return dives
}
