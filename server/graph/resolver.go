package graph

import (
	"context"
	"time"

	"github.com/project-hermes/hermes-app/server/model"
)

type Resolver struct {
	diveInt model.DiveInterface
}

// NewResolver will return a resolver with all the necessary dependencies
func NewResolver(diveInt model.DiveInterface) Resolver {
	return Resolver{
		diveInt: diveInt,
	}
}

func (r *Resolver) Dive() DiveResolver {
	return &diveResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type diveResolver struct{ *Resolver }

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
