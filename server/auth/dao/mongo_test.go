package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()

	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://114.55.95.101:27017/"))

	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")

	if err != nil {
		t.Errorf("faild resolve account id for 123: %v", err)
	} else {
		want := "61d67ef9327cb8b9d64cdfb6"
		if id != want {
			t.Errorf("resolve account id:want: %q, got: %q", want, id)
		}
	}
}
