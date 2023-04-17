package service

import (
	"context"
	"github.com/softlandia/hismap/models"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"github.com/softlandia/hismap/repo/items"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func tConnect(t *testing.T) mongo_db.Connect {
	connect, err := mongo_db.New(context.Background(), "mongodb://192.168.1.60:27017/", "hismap")
	assert.Nil(t, err)
	assert.NotNil(t, connect.DB())
	return connect
}

func tRepo(t *testing.T, connect mongo_db.Connect) items.Repo {
	res := items.New(context.Background(), connect)
	assert.NotNil(t, res)
	return res
}

func TestService_InsertItem(t *testing.T) {
	c := tConnect(t)
	r := tRepo(t, c)
	s := NewService(Config{
		Log: nil,
		Db:  c,
		Ctx: context.Background(),
		Repositories: Repositories{
			Items: r,
		},
	})

	s.Repo.Items.Delete(bson.M{"oid": "service-test"})

	id, err := s.InsertItem(models.Item{
		ID:  primitive.NewObjectID(),
		OID: "service-test",
		Object: models.Object{
			Caption: "-11-",
			StyleId: ">11<",
		},
		Start:  0,
		Stop:   0,
		Xmin:   0,
		Xmax:   0,
		Ymin:   0,
		Ymax:   0,
		Border: models.Border{models.Vertex{X: -100, Y: -200}, models.Vertex{X: 101, Y: 201}},
	})
	assert.Nil(t, err)
	assert.NotEqual(t, primitive.NilObjectID, id)
}

func TestService_GetOneItem(t *testing.T) {
	c := tConnect(t)
	r := tRepo(t, c)
	s := NewService(Config{
		Log: nil,
		Db:  c,
		Ctx: context.Background(),
		Repositories: Repositories{
			Items: r,
		},
	})

	item := models.Item{
		ID:  primitive.NewObjectID(),
		OID: "service-TestService_GetOneItem",
		Object: models.Object{
			Caption: "-11-",
			StyleId: ">11<",
		},
		Start:  -0,
		Stop:   0,
		Xmin:   0,
		Xmax:   0,
		Ymin:   0,
		Ymax:   0,
		Border: models.Border{models.Vertex{X: 0, Y: 0}, models.Vertex{X: 0, Y: 0}},
	}
	_, err := s.InsertItem(item)
	assert.Nil(t, err)

	item.ID = primitive.NewObjectID()
	item.Object.Caption = "-22-"
	_, err = s.InsertItem(item)
	assert.Nil(t, err)

	time.Sleep(time.Millisecond)

	filter := models.ItemsFilter{
		Xmin:  0,
		Xmax:  0,
		Ymin:  0,
		Ymax:  0,
		Start: 0,
		Stop:  0,
	}
	res := s.GetOneItem(filter)
	assert.Equal(t, "-11-", res.Object.Caption)
	assert.Equal(t, 2, len(res.Border))

	s.Repo.Items.Delete(bson.M{"oid": "service-TestService_GetOneItem"})
}
