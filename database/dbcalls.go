package database

import (
	"urlshortener/constant"

	"context"

	"urlshortener/types"

	"urlshortener/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {

	inst := mgr.connection.Database(constant.Database).Collection(collectionName)

	result, err := inst.InsertOne(context.TODO(), data)

	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromOriginalString(originalUrl string, collectionName string) (resp types.UrlDb, err error) {

	inst := mgr.connection.Database(constant.Database).Collection(collectionName)

	inst.FindOne(context.TODO(), bson.M{"long_url": originalUrl}).Decode(&resp)

	return resp, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDb, err error) {

	inst := mgr.connection.Database(constant.Database).Collection(collectionName)

	inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)

	return resp, err
}

func (mgr *manager) UpdateCount(newCount int, originalUrl string) error {

	inst := mgr.connection.Database(constant.Database).Collection(constant.UrlCollection)

	filter := bson.M{"long_url": originalUrl}

	update := bson.M{"$set": bson.M{"count": newCount}}

	_, err := inst.UpdateOne(context.TODO(), filter, update)
	return err

	
}

func (mgr *manager) SortDocument() [] types.UrlDb{

	zapLog := logger.GetLogger()

	collection := mgr.connection.Database(constant.Database).Collection(constant.UrlCollection)

	findOptions := options.Find().SetSort(map[string]int{"count": -1})

	cursor, err := collection.Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {

		zapLog.Error("Error Occured while finding the cursor")
	}

	var urls []types.UrlDb

	for cursor.Next(context.TODO()) {
		
		var url types.UrlDb

		err := cursor.Decode(&url)

		if err != nil {

			zapLog.Error("Error Occured")

		}

		urls = append(urls, url)
	}

	return urls
}
