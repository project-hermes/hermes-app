package wrapper

import (
	"cloud.google.com/go/storage"
	"context"
)

type StorageClientInterface interface {
	Bucket(name string) BucketHandleInterface
	//Close() error
}

type BucketHandleInterface interface {
	Object(name string) ObjectHandleInterface
}

type ObjectHandleInterface interface {

}

// Storage client is a wrapper around the GCS Storage Client
type storageClient struct {
	client *storage.Client
}

// Bucket will fetch the bucket from the Storage System
func (sc storageClient) Bucket(name string) BucketHandleInterface {
	return bucketHandle{
		handle: sc.client.Bucket(name),
	}
}

func NewStorageClient(ctx context.Context) (StorageClientInterface, error) {
	client, err := storage.NewClient(ctx)
	return &storageClient{
		client: client,
	}, err
}

type bucketHandle struct {
	handle *storage.BucketHandle
}

func (bh bucketHandle) Object(name string) ObjectHandleInterface {
	return nil
}

type TestType struct {
	testType TestType
	name string
}