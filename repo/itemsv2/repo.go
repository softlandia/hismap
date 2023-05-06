package itemsv2

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	Test() error
	GET(id primitive.ObjectID) Item
	DEL(id primitive.ObjectID)
	Insert(items Items) error
	Delete(filter bson.M)
	Find(filter bson.M) Items
	FindOne(filter bson.M) (res Item, err error)
	InsertOne(item Item) (primitive.ObjectID, error)
}

const TableName = "itemsv2"

type repo struct {
	mongo_db.Connect
	ctx    context.Context
	logger *zerolog.Logger
	Table  *mongo.Collection
}

func New(ctx context.Context, logger *zerolog.Logger, connect mongo_db.Connect) Repo {
	return &repo{
		Connect: connect,
		ctx:     ctx,
		Table:   connect.DB().Collection(TableName),
		logger:  logger,
	}
}

func (r repo) Test() error {
	res, err := r.FindOne(bson.M{"oid": "test"})
	if err != nil {
		return err
	}
	if res.Xmax > -99925 {
		return errors.New("invalid items table")
	}
	return nil
}

func (r repo) DEL(id primitive.ObjectID) {
	_, _ = r.Table.DeleteOne(r.ctx, bson.M{"_id": id})
}

func (r repo) Delete(filter bson.M) {
	_, _ = r.Table.DeleteMany(r.ctx, filter)
}

func (r repo) GET(id primitive.ObjectID) Item {
	sr := r.Table.FindOne(r.ctx, bson.M{"_id": id})
	var res Item
	_ = sr.Decode(&res)
	return res
}

func (r repo) InsertOne(item Item) (primitive.ObjectID, error) {
	res, err := r.Table.InsertOne(r.ctx, &item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r repo) Insert(items Items) error {
	array := make([]interface{}, 0, len(items))
	for _, e := range items {
		array = append(array, e)
	}
	_, err := r.Table.InsertMany(r.ctx, array)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) FindOne(filter bson.M) (res Item, err error) {
	sr := r.Table.FindOne(r.ctx, filter)
	err = sr.Decode(&res)
	return res, err
}

func (r repo) Find(filter bson.M) Items {
	cur, err := r.Table.Find(r.ctx, filter)
	if err != nil {
		return Items{}
	}
	res := make(Items, 0)
	if err := cur.All(r.ctx, &res); err != nil {
		r.logger.Error().Err(err).Msg("repo_items")
		return Items{}
	}
	return res
}
