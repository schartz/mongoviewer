package main

import (
	"MongoViewer/app_logging"
	"MongoViewer/connection"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// App struct
type App struct {
	ctx         context.Context
	appDBPath   string
	appDB       *bolt.DB
	mongoClient *mongo.Client
}

var appLogger *logrus.Logger

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	appLogger = app_logging.GetLogger()
	a.setup(ctx)
	appLogger.Println("Setup complete")
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
		appLogger.Fatal(err)
	}

	configDir := dirname + "/.mongoviewer"
	appDBFile := configDir + "/app.db"

	_, err1 := os.Stat(configDir)
	if os.IsNotExist(err1) {
		appLogger.Println(configDir + " does not exist. Creating.")
		errDir := os.MkdirAll(configDir, 0755)
		if errDir != nil {
			appLogger.Fatal(err)
		}

		_, appDBErr := os.Stat(appDBFile)

		if os.IsNotExist(appDBErr) {
			appLogger.Println(appDBFile + " does not exist. Creating.")
			err2 := os.WriteFile(appDBFile, []byte(""), 0600)
			if err2 != nil {
				appLogger.Fatal(err)
			}
		}
	}
	a.appDBPath = appDBFile

	a.appDB, err = bolt.Open(appDBFile, 0600, nil)
	if err != nil {
		appLogger.Fatal(err)
	}

	a.appDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Connections"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		appLogger.Println("Added connections bucket into database")
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
		appLogger.Println(err)
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
	return connection.ConnectToDBServer(connString)
}

func (a *App) GetDBDetails(dbName string) map[string]interface{} {
	return connection.GetDBDetails(dbName)
}
