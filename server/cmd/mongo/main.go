package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://114.55.95.101:27017/?readPreference=primary&appname=mongodb-vscode%200.7.0&ssl=false"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("coolcar").Collection("account")

	findRows(collection, ctx)

}

// 查询一条
func findRow(collection *mongo.Collection, ctx context.Context) {
	res := collection.FindOne(ctx, bson.M{
		"open_id": "123",
	})

	// fmt.Printf("%+v\n", res)
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		OpenID string             `bson:"open_id"`
	}

	err := res.Decode(&row)
	if err != nil {
		panic(err)
	}
	fmt.Printf("zhi %+v\n", row)

}

// 查询多条
func findRows(collection *mongo.Collection, ctx context.Context) {
	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	for res.Next(ctx) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenID string             `bson:"open_id"`
		}

		err = res.Decode(&row)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", row)
	}
}

// 插入多条
func insertMany(collection *mongo.Collection, ctx context.Context) {
	res, err := collection.InsertMany(ctx, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})

	// res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
