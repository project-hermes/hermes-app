package model

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"time"
)

type Dive struct {
	Name          string
	StartTime     time.Time
	EndTime       time.Time
	StartLocation appengine.GeoPoint
	EndLocation   appengine.GeoPoint
	Aggregations  []string
}

type DivePublic struct {
	Name             string   `json:"name"`
	StatTime         int64    `json:"startTime"`
	EndTime          int64    `json:"endTime"`
	StartLocationLat float64  `json:"startLocationLat"`
	StartLocationLng float64  `json:"startLocationLng"`
	EndLocationLat   float64  `json:"endLocationLat"`
	EndLocationLng   float64  `json:"endLocationLng"`
	Aggregations     []string `json:"aggregations"`
}

type DiveInterface interface {
	Create(dive Dive, createdAt time.Time) (*datastore.Key, error)
	Get(key *datastore.Key) (Dive, error)
	GetAggregation() ([]byte, error)
	List() ([]Dive, error)
	Public(dive Dive) DivePublic
	Private(dive DivePublic) Dive
	Update(updatedDive Dive, updatedAt time.Time) error
}

type DiveImplementation struct {
	appCtx      context.Context
	firebaseApp *firebase.App
}

func NewDiveImplementation(appEngineCtx context.Context, firebaseApp *firebase.App) DiveInterface {
	return DiveImplementation{appCtx: appEngineCtx, firebaseApp: firebaseApp}
}

func (d DiveImplementation) Create(dive Dive, createdAt time.Time) (*datastore.Key, error) {
	if key, err := datastore.Put(d.appCtx, datastore.NewKey(d.appCtx, "Dive", "", 0, nil), &dive); err != nil {
		log.Errorf(d.appCtx, "Could not create dive")
		return nil, err
	} else {
		return key, nil
	}
}

func (d DiveImplementation) Get(key *datastore.Key) (Dive, error) {
	var tempDive Dive
	log.Infof(d.appCtx, "%v", key.Kind())
	if err := datastore.Get(d.appCtx, key, &tempDive); err != nil {
		log.Errorf(d.appCtx, "Could not get dive %v", err)
		return Dive{}, err
	} else {
		return tempDive, nil
	}
}

func (d DiveImplementation) List() ([]Dive, error) {

	return nil, nil
}

func (d DiveImplementation) Private(dive DivePublic) Dive {
	return Dive{
		Name:          dive.Name,
		StartTime:     time.Unix(dive.StatTime, 0),
		EndTime:       time.Unix(dive.EndTime, 0),
		StartLocation: appengine.GeoPoint{Lat: dive.StartLocationLat, Lng: dive.StartLocationLng},
		EndLocation:   appengine.GeoPoint{Lat: dive.EndLocationLat, Lng: dive.EndLocationLng},
	}
}

func (d DiveImplementation) Public(dive Dive) DivePublic {
	return DivePublic{
		Name:             dive.Name,
		StatTime:         dive.StartTime.Unix(),
		EndTime:          dive.EndTime.Unix(),
		StartLocationLat: dive.StartLocation.Lat,
		StartLocationLng: dive.StartLocation.Lng,
		EndLocationLat:   dive.EndLocation.Lat,
		EndLocationLng:   dive.EndLocation.Lng,
	}
}

func (d DiveImplementation) Update(updatedDive Dive, updatedAt time.Time) error {
	return nil
}

// dive uploads
// dive is saved as raw in gcs
// task is pushed to process dive and create aggregations
// if task fails it will be logged
// task can be restarted with the information in the logs to rerun
func (d DiveImplementation) GetAggregation() ([]byte, error) {
	client, err := d.firebaseApp.Storage(d.appCtx)
	if err != nil {
		log.Errorf(d.appCtx, err.Error())
		return nil, err
	}
	bucket, err := client.Bucket("project-hermes-staging.appspot.com")
	if err != nil {
		log.Errorf(d.appCtx, err.Error())
		return nil, err
	}
	rc, err := bucket.Object("aggregations/config.txt").NewReader(d.appCtx)
	if err != nil {
		log.Errorf(d.appCtx, err.Error())
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Errorf(d.appCtx, err.Error())
		return nil, err
	}
	return data, nil
}
