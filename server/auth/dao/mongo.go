package dao

import (
	"context"
	"fmt"

	mgo "coolcar/shared/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, OpenID string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: OpenID,
	},
		mgo.Set(bson.M{
			openIDField: OpenID,
		}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("connot FindOneAndUpdate: %v", err)
	}

	var row mgo.ObjID
	err := res.Decode(&row)

	if err != nil {
		return "", fmt.Errorf("connot decode result: %v", err)
	}

	return row.ID.Hex(), nil
}
