package graph

import (
	"context"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"

	"github.com/project-hermes/hermes-app/server/model"
)

// Resolver is an internal implementation which leverages interfaces for resolving the graph
type Resolver struct {
	diveInt   model.DiveInterface
	sensorInt model.SensorInterface
}

// NewResolver will return a resolver with all the necessary dependencies
func NewResolver(diveInt model.DiveInterface, sensorInt model.SensorInterface) Resolver {
	return Resolver{
		diveInt:   diveInt,
		sensorInt: sensorInt,
	}
}

// Dive returns a resolver for a dive
func (r *Resolver) Dive() DiveResolver {
	return &diveResolver{r}
}

//func (r *Resolver) Sensor() SensorResolver {
//	return &sensorResolver{r}
//}

// Query returns a resolver for all queries
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation returns a resolver for all mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type diveResolver struct{ *Resolver }

//type sensorResolver struct{ *Resolver }
//
//func (r *sensorResolver) Status(ctx context.Context, obj *model.Sensor) (model.SensorStatus, error) {
//	return obj.Status, nil
//}

func (r *diveResolver) StartTime(ctx context.Context, obj *model.Dive) (int, error) {
	ms := obj.StartTime.UnixNano() / int64(time.Millisecond)
	return int(ms), nil
}

func (r *diveResolver) EndTime(ctx context.Context, obj *model.Dive) (int, error) {
	ms := obj.EndTime.UnixNano() / int64(time.Millisecond)
	return int(ms), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Dives(ctx context.Context) ([]model.Dive, error) {
	dives := r.diveInt.List(ctx)

	resolveCtx := graphql.GetResolverContext(ctx)
	log.Printf("Resolve context %+v", resolveCtx)

	cf := graphql.CollectFieldsCtx(ctx, []string{"sensor"})
	log.Printf("GraphQL Collected Fields: %+v", cf)

	fetchSensors := false
	for _, field := range cf {
		if field.Name == "sensor" {
			fetchSensors = true
			break
		}
		log.Printf("field name: %s, alias %s, and definition %+v", field.Name, field.Alias, field.Definition)
	}

	if fetchSensors {
		sensorIDs := map[string]bool{}
		for _, dive := range dives {
			sensorIDs[dive.SensorID] = true
		}

		var ids []string
		for key := range sensorIDs {
			ids = append(ids, key)
		}

		sensors, err := r.sensorInt.GetByIDs(ctx, ids)
		if err != nil {
			return nil, err
		}

		for index, dive := range dives {
			for _, sensor := range sensors {
				if dive.SensorID == sensor.ID {
					dives[index].Sensor = sensor
					break
				}
			}
		}
	}

	return dives, nil
}

func (r *queryResolver) Sensors(ctx context.Context) ([]model.Sensor, error) {
	return r.sensorInt.GetByIDs(ctx, []string{"12345", "23456"})
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SensorCreate(ctx context.Context, sensor model.InputSensor) (model.Sensor, error) {
	return r.sensorInt.Create(ctx, sensor)
}
