package dao

import "go.mongodb.org/mongo-driver/mongo"

type Mongo struct {
	col *mongo.Collection
}

func newMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("blob"),
	}
}
