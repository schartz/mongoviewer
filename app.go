package main

import (
	"MongoViewer/connection"
	"context"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// App struct
type App struct {
	ctx         context.Context
	appDBPath   string
	appDB       *bolt.DB
	mongoClient *mongo.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here

	log.SetFlags(log.Lshortfile)
	log.SetPrefix("logger: ")

	log.Println("running setup here")
	a.setup(ctx)
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	defer func(appDB *bolt.DB) {
		err := appDB.Close()
		if err != nil {

		}
	}(a.appDB)
}

func (a *App) setup(ctx context.Context) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configDir := dirname + "/.mongoviewer"
	appDBFile := configDir + "/app.db"

	_, err1 := os.Stat(configDir)
	if os.IsNotExist(err1) {
		log.Println(configDir + " does not exist. Creating.")
		errDir := os.MkdirAll(configDir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

		_, appDBErr := os.Stat(appDBFile)

		if os.IsNotExist(appDBErr) {
			log.Println(appDBFile + " does not exist. Creating.")
			err2 := os.WriteFile(appDBFile, []byte(""), 0600)
			if err2 != nil {
				log.Fatal(err)
			}
		}
	}
	a.appDBPath = appDBFile

	a.appDB, err = bolt.Open(appDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	a.appDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Connections"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		log.Println("Added connections bucket into database")
		return nil
	})

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func (a *App) ConnectionList() []map[string]string {
	connectionList, err := connection.GetConnectionList(a.appDB)
	if err != nil {
		log.Println(err)
	}
	return connectionList
}

func (a *App) TestConnection(connString string) string {
	return connection.TestConnection(connString)
}
func (a *App) AddConnection(connString string) bool {
	return connection.AddConnection(connString, a.appDB)
}

func (a *App) ConnectToDBServer(connString string) []mongo.DatabaseSpecification {

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	a.mongoClient = client

	filter := bson.D{}
	dbList, err := a.mongoClient.ListDatabases(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return dbList.Databases

}
