package model

import (
	"context"
	"log"

	"github.com/project-hermes/hermes-app/server/wrapper"
)

type Sensor struct {
	ID    string `json:"id" firestore:"sensorId,omitempty"`
	Name  string `json:"name" firestore:"name,omitempty"`
	Type  string `json:"type" firestore:"type,omitempty"`
	Model string `json:"model" firestore:"model,omitempty"`
	Status SensorStatus `json:"status" firestore:"status,omitempty"`
}

type SensorInterface interface {
	Create(context.Context, InputSensor) (Sensor, error)
	GetByIDs(context.Context, []string) ([]Sensor, error)
}

type SensorImplementation struct {
	client wrapper.DBClientInterface
}

// NewSensorImplementation will return an object for working with sensors
func NewSensorImplementation(client wrapper.DBClientInterface) SensorInterface {
	return &SensorImplementation{
		client: client,
	}
}

// Create will generate a new sensor with the correct ID
func (si SensorImplementation) Create(ctx context.Context, inputSensor InputSensor) (Sensor, error) {
	docRef := si.client.Collection("sensor").Doc(inputSensor.SensorID)

	sensor := Sensor{
		ID:   inputSensor.SensorID,
		Name: inputSensor.Name,
		Type: inputSensor.Type,
		Model: inputSensor.Model,
		Status: SensorStatusActive,
	}

	if _, err := docRef.Set(ctx, sensor); err != nil {
		log.Printf("An error has occurred: %s", err)
		return Sensor{}, err
	}

	return sensor, nil
}

// GetByIDs will return multiple sensors based on the provided IDs
func (si SensorImplementation) GetByIDs(ctx context.Context, ids []string) ([]Sensor, error) {
	log.Printf("IDs we're fetching %+v", ids)
	collectionRef := si.client.Collection("sensor")
	var docRefs []wrapper.DocRefInterface
	for _, id := range ids {
		docRefs = append(docRefs, collectionRef.Doc(id))
	}

	snapshots, err := si.client.GetAll(ctx, docRefs)
	if err != nil {
		return nil, err
	}



	sensors := make([]Sensor, len(snapshots))
	for _, snapshot := range snapshots {
		var sensor Sensor
		if err := snapshot.DataTo(&sensor); err != nil {
			return nil, err
		} else {
			sensors = append(sensors, sensor)
		}
	}

	return sensors, nil
}
