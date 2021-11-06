package store

import (
	"askUs/v1/package/asset"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"google.golang.org/api/option"
)

type Storage struct {
}

// var bucket *storage.BucketHandle

func (s *Storage) New() *storage.BucketHandle {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./creds.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return nil
	}
	stor, err := app.Storage(ctx)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return nil
	}
	buc, err := stor.Bucket("askusstore-c6e12.appspot.com")
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return nil
	}
	return buc
}

func Init() asset.Store {
	return &Storage{}
}

func (store *Storage) Post(ctx context.Context, bucket *storage.BucketHandle, fileUrl string, reader io.Reader) error {
	id, err := gonanoid.New()
	if err != nil {
		fmt.Println("50")
		return err
	}
	object := bucket.Object(fileUrl)
	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id}
	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		fmt.Println("57")
		return err
	}
	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != storage.ErrObjectNotExist && err != nil {
		fmt.Println("61")
		return err
	}
	return nil
}

func (store *Storage) Get(ctx context.Context, bucket *storage.BucketHandle, fileUrl string) ([]byte, error) {
	object := bucket.Object(fileUrl)
	rc, err := object.NewReader(context.Background())
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return []byte{}, err

	}
	return data, nil
}

func (store *Storage) Delete(ctx context.Context, bucket *storage.BucketHandle, fileUrl string) error {
	object := bucket.Object(fileUrl)
	err := object.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
