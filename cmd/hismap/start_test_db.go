package main

import (
	"context"
	"errors"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func testConnect(connect mongo_db.Connect) error {
	id, _ := primitive.ObjectIDFromHex("643c69dfe6bd749f91c11984")
	sr := connect.DB().Collection("test").FindOne(context.Background(), bson.M{"_id": id})
	var res struct {
		Text string `bson:"text"`
	}

	if err := sr.Decode(&res); err != nil {
		return err
	}
	if res.Text != "text" {
		return errors.New("db not hismap")
	}
	return nil
}
