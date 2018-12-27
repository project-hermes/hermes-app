package wrapper

import (
	"context"

	"cloud.google.com/go/firestore"
)

// DBClientInterface is an interface that mimics the firestore client
type DBClientInterface interface {
	Collection(path string) CollectionRefInterface
}

// DBClient is a wrapper around the firestore client
type DBClient struct {
	client *firestore.Client
}

// Collection will return a collection reference
func (dbc DBClient) Collection(path string) CollectionRefInterface {
	return CollectionRef{
		collection: dbc.client.Collection(path),
	}
}

// NewClient will return a firestore client
func NewClient(ctx context.Context, projectID string) (DBClientInterface, error) {
	client, err := firestore.NewClient(context.Background(), projectID)
	return &DBClient{
		client: client,
	}, err
}

// CollectionRefInterface for firestore
type CollectionRefInterface interface {
	NewDoc() DocRefInterface
	Documents(ctx context.Context) DocumentInteratorInterface
}

// CollectionRef is a wrapper around firestore's CollectionRef
type CollectionRef struct {
	collection *firestore.CollectionRef
}

// NewDoc will generate and return a new document
func (cr CollectionRef) NewDoc() DocRefInterface {
	return cr.collection.NewDoc()
}

// Documents will return an iterator for fetching all of the documents
func (cr CollectionRef) Documents(ctx context.Context) DocumentInteratorInterface {
	return cr.collection.Documents(ctx)
}

// DocRefInterface is a wrapper around DocRef
type DocRefInterface interface {
	Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (*firestore.WriteResult, error)
}

// DocumentInteratorInterface is a wrapper around firestore's document iterator
type DocumentInteratorInterface interface {
	Next() (*firestore.DocumentSnapshot, error)
}
