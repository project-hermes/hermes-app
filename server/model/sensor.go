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
