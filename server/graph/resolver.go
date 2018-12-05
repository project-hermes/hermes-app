package graph

import (
	"time"
	"context"

	"github.com/project-hermes/hermes-app/server/model"
)

type Resolver struct{}

func (r *Resolver) Dive() DiveResolver {
	return &diveResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type diveResolver struct{ *Resolver }

func (r *diveResolver) StartTime(ctx context.Context, obj *model.Dive) (int, error) {
	ms := time.Now().UnixNano() / int64(time.Millisecond)
	return int(ms), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Dives(ctx context.Context) ([]*model.Dive, error) {
	return []*model.Dive{
		&model.Dive{
			Name: "Test Dive",
			StartTime: time.Now(),
		},
	}, nil
}
