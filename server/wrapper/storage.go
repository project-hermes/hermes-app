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
type StorageClient struct {
	client *storage.Client
}

// Bucket will fetch the bucket from the Storage System
func (sc StorageClient) Bucket(name string) BucketHandleInterface {
	return BucketHandle{
		handle: sc.client.Bucket(name),
	}
}

func NewStorageClient(ctx context.Context) (StorageClientInterface, error) {
	client, err := storage.NewClient(ctx)
	return &StorageClient{
		client: client,
	}, err
}

type BucketHandle struct {
	handle *storage.BucketHandle
}

func (bh BucketHandle) Object(name string) ObjectHandleInterface {
	return nil
}