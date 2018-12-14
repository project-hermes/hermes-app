package model

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	pblatlng "google.golang.org/genproto/googleapis/type/latlng"
)

// FirestoreDive
// type FirestoreDive struct {
// 	SensorID   string    `firestore:"sensorId,omitempty"`
// 	StartTime  time.Time `firestore:"startTime,omitempty"`
// 	EndTime    time.Time `firestore:"endTime,omitempty"`
// 	StartPoint *pblatlng.LatLng  `firestore:"startPoint,omitempty"`
// 	EndPoint   *pblatlng.LatLng  `firestore:"endPoint,omitempty"`
// }

// type FirestoreSensorData struct {
// 	SensorID   string    `firestore:"sensorId,omitempty"`
// 	StartTime  time.Time `firestore:"startTime,omitempty"`
// 	EndTime    time.Time `firestore:"endTime,omitempty"`
// 	StartPoint *pblatlng.LatLng  `firestore:"startPoint,omitempty"`
// 	EndPoint   *pblatlng.LatLng  `firestore:"endPoint,omitempty"`
// }

// Dive is a dive
type Dive struct {
	SensorID   string    `json:"sensorId" firestore:"sensorId,omitempty"`
	StartTime  time.Time `json:"startTime"  firestore:"startTime,omitempty"`
	EndTime    time.Time `json:"duration"  firestore:"endTime,omitempty"`
	StartPoint GeoPoint  `json:"startPoint" firestore:"startPoint,omitempty"`
	EndPoint   GeoPoint  `json:"endPoint" firestore:"endPoint,omitempty"`
	SensorData   []*SensorData  `json:"sensorData" firestore:"sensorData,omitempty"`
}

// MapDive will convert a map into a dive
func MapDive(data map[string]interface{}) Dive {

	startPoint := data["startPoint"].(*pblatlng.LatLng)
	endPoint := data["endPoint"].(*pblatlng.LatLng)

	rawSensorData := data["sensorData"].([]interface{})
	var sensorData []*SensorData
	for _, sd := range rawSensorData {
		mapSd := sd.(map[string]interface{})
		newSd := MapSensorData(mapSd)
		sensorData = append(sensorData, &newSd)
	}

	return Dive {
		SensorID: data["sensorId"].(string),
		SensorData: sensorData,
		StartTime: data["startTime"].(time.Time),
		EndTime: data["endTime"].(time.Time),
		StartPoint: GeoPoint{Lat: startPoint.Latitude, Long: startPoint.Longitude},
		EndPoint: GeoPoint{Lat: endPoint.Latitude, Long: endPoint.Longitude},
	}
}

// DiveInterface is an interface for interacting with dive results
type DiveInterface interface {
	Create(context.Context, Dive) error
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

// Create will create a new dive
func (di DiveImplementation) Create(ctx context.Context, dive Dive) error {
	docRef := di.client.Collection("dive").NewDoc()
	_, err := docRef.Set(ctx, dive)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}

	return nil
}

// List will fetch all of the dives
func (di DiveImplementation) List(ctx context.Context) []*Dive {
	iter := di.client.Collection("dive").Documents(ctx)

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
