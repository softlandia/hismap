package service

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/softlandia/hismap/models"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"github.com/softlandia/hismap/repo/items"
	"github.com/softlandia/hismap/repo/itemsv2"
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
	Items   items.Repo
	ItemsV2 itemsv2.Repo
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

func (s *Service) InsertItems(items models.ItemList) error {
	return s.Repo.Items.Insert(items.ToRepo())
}

// InsertOneItem - сохранить в базу один элемент карты
func (s *Service) InsertOneItem(item models.Item) (primitive.ObjectID, error) {
	return s.Repo.Items.InsertOne(item.ToRepo())
}

// GetItems - получить выборку элементов на карту
func (s *Service) GetItems(itemFilter models.ItemsFilter) models.ItemList {
	repoItems := s.Repo.Items.Find(itemFilter.ToRepo())
	return models.RepoToItemList(repoItems)
}

func (s *Service) GetOneItem(itemFilter models.ItemsFilter) models.Item {
	res, _ := s.Repo.Items.FindOne(itemFilter.ToRepo())
	return models.RepoToItem(res)
}
