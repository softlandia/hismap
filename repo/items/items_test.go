package items

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
)

func tRepo(t *testing.T) Repo {
	connect, err := mongo_db.New(context.Background(), "mongodb://192.168.1.60:27017/", "hismap")
	assert.Nil(t, err)
	assert.NotNil(t, connect.DB())
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	res := New(context.Background(), &logger, connect)
	assert.NotNil(t, res)
	return res
}

func TestRepo_InsertOne(t *testing.T) {
	r := tRepo(t)
	item := Item{
		ID:  primitive.NewObjectID(),
		OID: "TMP",
		Object: Object{
			Caption: "TMP_TEST",
			StyleId: "-",
		},
		Start:  0,
		Stop:   0,
		Xmin:   0,
		Xmax:   0,
		Ymax:   0,
		Ymin:   0,
		Border: Border{Vertex{0, 0}, Vertex{1, 1}},
	}
	id, err := r.InsertOne(item)
	assert.Nil(t, err)
	assert.NotEqual(t, primitive.NilObjectID, id)
	r.DEL(id)
}

func TestRepo_Insert(t *testing.T) {
	r := tRepo(t)
	r.Delete(bson.M{"oid": "TestRepo_Insert"})

	items := Items{
		Item{
			OID: "TestRepo_Insert",
		},
		Item{
			OID: "TestRepo_Insert",
		},
	}
	err := r.Insert(items)
	assert.Nil(t, err)
}
