package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"urlshortener/types"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromOriginalString(string, string) (types.UrlDb, error)
	GetUrlFromCode(string,string)(types.UrlDb,error)
	UpdateCount(int,string)(error)
	SortDocument() []types.UrlDb
	
}

func ConnectDb() {

	err:=godotenv.Load(".env")

	if err != nil{

		panic(err)
	}

	uri := os.Getenv("DB_HOST")

	fmt.Println("Uri : ",uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))

	fmt.Println("The Value of client is : ",client)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Mgr = manager{connection: client, ctx: ctx, cancel: cancel}
}
