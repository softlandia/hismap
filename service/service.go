package service

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/softlandia/hismap/models"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"github.com/softlandia/hismap/repo/items"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Config - service config
type Config struct {
	Log          *zerolog.Logger
	Db           mongo_db.Connect
	Ctx          context.Context
	Repositories Repositories
}

type Service struct {
	ctx    context.Context
	db     mongo_db.Connect
	Logger *zerolog.Logger
	Repo   Repositories
}

type Repositories struct {
	Items items.Repo
}

// NewService - new Service
func NewService(config Config) *Service {
	return &Service{
		ctx:    config.Ctx,
		db:     config.Db,
		Logger: config.Log,
		Repo:   config.Repositories,
	}
}

// InsertItem - сохранить в базу один элемент карты
func (s *Service) InsertItem(item models.Item) (primitive.ObjectID, error) {
	return s.Repo.Items.InsertOne(item.ToRepo())
}

func (s *Service) GetItems(itemFilter models.ItemsFilter) models.Item {
	s.Repo.Items.Find(itemFilter.ToRepo())
	return models.Item{}
}

func (s *Service) GetOneItem(itemFilter models.ItemsFilter) models.Item {
	res, _ := s.Repo.Items.FindOne(itemFilter.ToRepo())
	return models.RepoToItem(res)
}
