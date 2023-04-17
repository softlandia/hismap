package mongo_db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - соединение с базой mongodb
type Connect struct {
	db     *mongo.Database
	client *mongo.Client
}

func New(ctx context.Context, connectString string, dbName string) (Connect, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectString))
	if err != nil {
		return Connect{}, err
	}

	if err = client.Connect(ctx); err != nil {
		return Connect{}, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return Connect{}, err
	}

	return Connect{
		db:     client.Database(dbName),
		client: client,
	}, nil
}

// DB -
func (c *Connect) DB() *mongo.Database {
	return c.db
}

// Transaction -
func (c *Connect) Transaction(ctx context.Context, fn func(trCtx context.Context) error, opt ...*options.TransactionOptions) error {
	session, err := c.client.StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(ctx)

	err = session.StartTransaction(opt...)
	if err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = fn(sc)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(sc)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/*// StrToObjectId -
func StrToObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

// ObjectIdToStr -
func ObjectIdToStr(id primitive.ObjectID) string {
	return id.Hex()
}*/
