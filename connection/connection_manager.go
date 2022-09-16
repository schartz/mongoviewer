package connection

import (
	"context"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetConnectionList(db *bolt.DB) ([]map[string]string, error) {
	var connectionList []map[string]string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Connections"))
		cursor := b.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {

			connectionList = append(connectionList, map[string]string{"name": string(key), "uri": string(value)})

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return connectionList, nil

}

func TestConnection(connString string) string {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err2 := client.Connect(ctx)
	if err2 != nil {
		log.Println(err)
		return err.Error()
	}

	msg := "yes"

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, ctx)
	return msg
}

func AddConnection(connString string, db *bolt.DB) bool {
	var connection map[string]string
	err := json.Unmarshal([]byte(connString), &connection)
	if err != nil {
		return false
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Connections"))
		if err != nil {
			return err
		}

		if err := b.Put([]byte(connection["name"]), []byte(connection["uri"])); err != nil {
			log.Fatal(err)
		}

		return nil
	})
	if err != nil {
		return false
	}
	return true
}

func ConnectToDBServer(uriString string, appMongoClientHandle *mongo.Client) []mongo.DatabaseSpecification {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uriString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	appMongoClientHandle = client

	filter := bson.D{}
	dbList, err := appMongoClientHandle.ListDatabases(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return dbList.Databases
}
