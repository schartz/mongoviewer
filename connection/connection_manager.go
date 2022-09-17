package connection

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoSession struct {
	db     *mongo.Database
	Client *mongo.Client
	logger *logrus.Logger
}

var MONGOSSN *MongoSession

func InitMongoSession(uri string, logger *logrus.Logger) {

	session := connect(uri, logger)
	if session != nil {

		// log statements here as well

		mongoSession := new(MongoSession)
		mongoSession.db = nil
		mongoSession.logger = logger
		mongoSession.Client = session
		MONGOSSN = mongoSession
		return
	}
	logger.Fatalf("Failed to connect to mongo server: %v", uri)
}

func connect(uri string, logger *logrus.Logger) *mongo.Client {
	var connectOnce sync.Once
	var client *mongo.Client
	connectOnce.Do(func() {

		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			logger.Fatal(err)
		}
		err = client.Connect(context.TODO())
		if err != nil {
			logger.Fatal(err)
		}
	})

	return client
}

func (mongoSession *MongoSession) SelectDatabase(dbName string, logger *logrus.Logger) {
	if mongoSession.Client == nil {
		logger.Fatal("Cannot select database in an uninitialized session")
	}

	err := mongoSession.Client.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	mongoSession.db = mongoSession.Client.Database(dbName)
	logger.Info("Successfully connected to database: %v", dbName)

}

func TestConnection(connString string) string {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		appLogger.Println(err)
		return err.Error()
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err2 := client.Connect(ctx)
	if err2 != nil {
		appLogger.Println(err)
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
