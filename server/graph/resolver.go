package graph

import (
	"context"

	"github.com/project-hermes/hermes-app/server/model"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context) (model.Hello, error) {
	return model.Hello {
		Name: "World",
	}, nil
}
