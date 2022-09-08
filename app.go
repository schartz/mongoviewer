package main

import (
	"MongoViewer/connection"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	cwd, err := os.Getwd()
	if err != nil {
		println(err)
	}
	myStr := `[{"conn": "Conn 1"}, {"conn": "conn 2"}]`
	filename := path.Join(cwd, "mylist.json")
	err1 := os.WriteFile(filename, []byte(myStr), 0600)
	if err != nil {
		println(err1)
	}
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func (a *App) ConnectionList() string {
	conns, err := connection.GetConnectionList()
	if err != nil {
		log.Fatal(err)
	}
	return conns
}

func (a *App) TestConnection(connString string) string {
	println(connString)
	println("************************************")
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
