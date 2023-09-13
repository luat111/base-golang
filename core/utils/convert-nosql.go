package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	. "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ToBson(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func GetCurrentBsonTimeStamp() *Timestamp {
	return &Timestamp{T: uint32(time.Now().Unix())}
}

func ConvertToObjectId(str string) (ObjectID, error) {
	objID, err := ObjectIDFromHex(str)
	if err != nil {
		panic(err)
	}

	return objID, err
}

func GetDataFromCusor(cursor *mongo.Cursor) ([]any, error) {
	defer cursor.Close(context.Background())
	var results []any

	for cursor.Next(context.Background()) {
		var result map[string]interface{}

		if err := cursor.Decode(&result); err != nil {
			return []any{}, err
		}

		results = append(results, result)
	}
	return results, nil
}
