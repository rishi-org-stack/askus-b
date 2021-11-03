package main

import (
	"askUs/v1/package/api"
	"askUs/v1/util/auth"
	"askUs/v1/util/config"
	"askUs/v1/util/db"
	mid "askUs/v1/util/middleware"
	"askUs/v1/util/server"
	"context"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	jwtService, err := auth.Init(env)
	handleError(err)
	ap := api.Init(client, jwtService, env, mid.JwtAuth(jwtService))
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}

func handleError(e error) {

	if e != nil {
		fmt.Println(e.Error())
	}

}

// package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"cloud.google.com/go/storage"
// 	firebase "firebase.google.com/go"
// 	"google.golang.org/api/cloudkms/v1"
// 	"google.golang.org/api/iterator"
// 	"google.golang.org/api/option"
// )

// func implicit() {
// 	ctx := context.Background()

// 	// For API packages whose import path is starting with "cloud.google.com/go",
// 	// such as cloud.google.com/go/storage in this case, if there are no credentials
// 	// provided, the client library will look for credentials in the environment.
// 	storageClient, err := storage.NewClient(ctx,
// 		option.WithCredentialsFile("./creds.json"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer storageClient.Close(``)

// 	it := storageClient.Buckets(ctx, "askusstore-c6e12")
// 	for {
// 		bucketAttrs, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(bucketAttrs.Name)
// 	}

// 	// For packages whose import path is starting with "google.golang.org/api",
// 	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
// 	kmsService, err := cloudkms.NewService(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_ = kmsService
// }

// // func createObject(bucket *storage.BucketHandle, nameOfObject string) error {
// // 	err := bucket.Create(context.Background(), "askusstore-c6e12", &storage.BucketAttrs{
// // 		Name: nameOfObject,
// // 	})
// // 	return err
// // }
// func post(object *storage.ObjectHandle) {
// 	ctx := context.Background()
// 	writer := object.NewWriter(ctx)
// 	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": "id1091234567890"}
// 	defer writer.Close()

// 	file, err := os.ReadFile("./ok.txt")
// 	if err != nil {
// 		fmt.Printf("error initializing app: %v", err)
// 		return
// 	}
// 	if _, err := io.Copy(writer, bytes.NewReader(file)); err != nil {
// 		fmt.Println("57:" + err.Error())
// 		return
// 	}
// 	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
// 		fmt.Println("61:" + err.Error())
// 		return
// 	}
// }
// func upload(object *storage.ObjectHandle) {
// 	writer := object.NewWriter(context.Background())

// 	// Set the attribute
// 	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": "iiidiwjSOHruignxnj"}
// 	defer writer.Close()
// 	file, err := os.ReadFile("./ok.txt")
// 	if err != nil {
// 		fmt.Printf("error initializing app: %v", err)
// 		return
// 	}
// 	if _, err := io.Copy(writer, bytes.NewReader((file))); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// func download(object *storage.ObjectHandle) {
// 	rc, err := object.NewReader(context.Background())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer rc.Close()

// 	data, err := ioutil.ReadAll(rc)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("result")
// 	fmt.Println(string(data))
// }
// func main() {

// 	ctx := context.Background()
// 	opt := option.WithCredentialsFile("./creds.json")
// 	app, err := firebase.NewApp(ctx, nil, opt)
// 	if err != nil {
// 		fmt.Printf("error initializing app: %v", err)
// 		return
// 	}
// 	fmt.Println(app)
// 	stor, err := app.Storage(ctx)
// 	if err != nil {
// 		fmt.Printf("error initializing app: %v", err)
// 		return
// 	}
// 	buc, err := stor.Bucket("askusstore-c6e12.appspot.com")
// 	if err != nil {
// 		fmt.Printf("error initializing app: %v", err)
// 		return
// 	}
// 	object := buc.Object("profile/prof")

// 	post(object)
// }
