package graph

import (
	"context"
	"time"

	"github.com/project-hermes/hermes-app/server/model"
)

// Resolver is an internal implementation which leverages interfaces for resolving the graph
type Resolver struct {
	diveInt model.DiveInterface
	sensorInt model.SensorInterface
}

// NewResolver will return a resolver with all the necessary dependencies
func NewResolver(diveInt model.DiveInterface, sensorInt model.SensorInterface) Resolver {
	return Resolver{
		diveInt: diveInt,
		sensorInt: sensorInt,
	}
}

// Dive returns a resolver for a dive
func (r *Resolver) Dive() DiveResolver {
	return &diveResolver{r}
}

func (r *Resolver) Sensor() SensorResolver {
	return &sensorResolver{r}
}

// Query returns a resolver for all queries
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation returns a resolver for all mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type diveResolver struct{ *Resolver }

type sensorResolver struct{ *Resolver }

func (r *sensorResolver) Status(ctx context.Context, obj *model.Sensor) (model.SensorStatus, error) {
	return obj.Status, nil
}

func (r *diveResolver) StartTime(ctx context.Context, obj *model.Dive) (int, error) {
	ms := obj.StartTime.UnixNano() / int64(time.Millisecond)
	return int(ms), nil
}

func (r *diveResolver) EndTime(ctx context.Context, obj *model.Dive) (int, error) {
	ms := obj.EndTime.UnixNano() / int64(time.Millisecond)
	return int(ms), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Dives(ctx context.Context) ([]*model.Dive, error) {
	return r.diveInt.List(ctx), nil
}

type mutationResolver struct { *Resolver }

func (r *mutationResolver) SensorCreate(ctx context.Context, sensor model.InputSensor) (model.Sensor, error) {
	return r.sensorInt.Create(ctx, sensor)
}

