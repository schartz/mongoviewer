package connection

import (
	"MongoViewer/app_logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var appLogger = app_logging.GetLogger()

type dbFunctions struct {
	_id   string
	value string
}

func ConnectToDBServer(uriString string) []mongo.DatabaseSpecification {
	ctx := context.TODO()
	InitMongoSession(uriString, logrus.New())
	dbList, err := MONGOSSN.Client.ListDatabases(ctx, bson.D{})
	if err != nil {
		appLogger.Fatal(err)
	}
	return dbList.Databases
}

func GetDBDetails(dbName string) map[string]interface{} {
	collections := listCollections(dbName)
	listUser(dbName)
	dbFunctionList := listFunctions(dbName)

	response := map[string]interface{}{
		"dbFunctions": dbFunctionList,
		"collections": collections,
	}
	return response
}
func listCollections(dbName string) []string {
	// get list of collections
	ctx := context.TODO()
	db := MONGOSSN.Client.Database(dbName)
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		appLogger.Fatal(err)
	}
	println("***")

	var result []string
	for _, item := range collections {
		result = append(result, item)
	}
	return result
}

func listUser(dbName string) {
	command := bson.D{
		{"usersInfo", 1},
		{"filter", bson.D{}},
	}
	result := MONGOSSN.Client.Database(dbName).RunCommand(context.Background(), command)
	if result.Err() != nil {
		appLogger.Error(result.Err())
	}
	println("++++++")
	response, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		appLogger.Error(err)
	}
	fmt.Printf("%s\n", response)
}

func listFunctions(dbName string) []dbFunctions {
	// get list of functions
	functionsCollection := MONGOSSN.Client.Database(dbName).Collection("system.js")
	cursor, err := functionsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		appLogger.Error(err)
	}

	var result []dbFunctions
	for cursor.Next(context.TODO()) {
		var a dbFunctions
		err = cursor.Decode(&a)
		if err != nil {
			appLogger.Error(err)
		}

		result = append(result, a)
	}

	return result
}
