package wrapper

import (
	"context"

	"cloud.google.com/go/firestore"
)

// DBClientInterface is an interface that mimics the firestore client
type DBClientInterface interface {
	Collection(path string) CollectionRefInterface
	GetAll(context.Context, []DocRefInterface) ([]*firestore.DocumentSnapshot, error)
}

// dbClient is a wrapper around the firestore client
type dbClient struct {
	client *firestore.Client
}

// Collection will return a collection reference
func (dbc dbClient) Collection(path string) CollectionRefInterface {
	return collectionRef{
		collection: dbc.client.Collection(path),
	}
}

// GetAll will return snapshots of all the document references
func (dbc dbClient) GetAll(ctx context.Context, docRefs []DocRefInterface) ([]*firestore.DocumentSnapshot, error) {
	var refs []*firestore.DocumentRef
	for _, ref := range docRefs {
		refs = append(refs, ref.(*documentRef).docRef)
	}

	return dbc.client.GetAll(ctx, refs)
}

// NewClient will return a firestore client
func NewClient(ctx context.Context, projectID string) (DBClientInterface, error) {
	client, err := firestore.NewClient(ctx, projectID)
	return &dbClient{
		client: client,
	}, err
}

// CollectionRefInterface for firestore
type CollectionRefInterface interface {
	Doc(string) DocRefInterface
	NewDoc() DocRefInterface
	Documents(ctx context.Context) DocumentInteratorInterface
}

// collectionRef is a wrapper around firestore's collectionRef
type collectionRef struct {
	collection *firestore.CollectionRef
}

// Doc will generate a new document with the specified ID
func (cr collectionRef) Doc(id string) DocRefInterface {
	return &documentRef{
		docRef: cr.collection.Doc(id),
	}
}

// NewDoc will generate and return a new document
func (cr collectionRef) NewDoc() DocRefInterface {
	return &documentRef{
		docRef: cr.collection.NewDoc(),
	}
}

// Documents will return an iterator for fetching all of the documents
func (cr collectionRef) Documents(ctx context.Context) DocumentInteratorInterface {
	return cr.collection.Documents(ctx)
}

// DocRefInterface is a wrapper around DocRef
type DocRefInterface interface {
	ID() string
	Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (*firestore.WriteResult, error)
}

// documentRef is a wrapper around firestore.documentRef
type documentRef struct {
	docRef *firestore.DocumentRef
}

// Set will set the new object in firestore
func (dr documentRef) Set(ctx context.Context, data interface{}, opts ...firestore.SetOption) (*firestore.WriteResult, error) {
	return dr.docRef.Set(ctx, data, opts...)
}

// ID will return the ID of the document reference
func (dr documentRef) ID() string {
	return dr.docRef.ID
}

// DocumentInteratorInterface is a wrapper around firestore's document iterator
type DocumentInteratorInterface interface {
	Next() (*firestore.DocumentSnapshot, error)
}
