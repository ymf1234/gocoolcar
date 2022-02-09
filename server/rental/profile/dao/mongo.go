package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

const (
	accountIDField = "account_id"
	profileField   = "profile"
)

type ProfileRecord struct {
	AccountID string            `bson:"account_id"`
	Profile   *rentalpb.Profile `bson:"profile"`
}

func (m *Mongo) GetProfile(c context.Context, aid id.AccountID) (*rentalpb.Profile, error) {
	res := m.col.FindOne(c, bson.M{
		accountIDField: aid.String(),
	})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var pr ProfileRecord

	err := res.Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode profile record: %v", err)
	}
	return pr.Profile, nil
}
