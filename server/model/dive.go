package model

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Dive is a dive
type Dive struct {
	SensorID   string    `json:"sensorId" firestore:"sensorId,omitempty"`
	StartTime  time.Time `json:"startTime"  firestore:"startTime,omitempty"`
	EndTime    time.Time `json:"duration"  firestore:"endTime,omitempty"`
	StartPoint GeoPoint  `json:"startPoint"`
	EndPoint   GeoPoint  `json:"endPoint"`
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

		var d Dive
		mapError := doc.DataTo(&d)
		if mapError != nil {
			log.Fatalf("unable to map doc to dive: %v", err)
		}

		dives = append(dives, &d)
	}

	return dives
}
