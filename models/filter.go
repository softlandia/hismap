package models

import "go.mongodb.org/mongo-driver/bson"

// ItemsFilter - фильтр выбора объектов для отображения
type ItemsFilter struct {
	Xmin  int
	Xmax  int
	Ymin  int
	Ymax  int
	Start int32
	Stop  int32
}

func (i ItemsFilter) ToRepo() bson.M {
	return bson.M{
		"xmin":  bson.M{"$gte": i.Xmin},
		"xmax":  bson.M{"$lte": i.Xmax},
		"ymin":  bson.M{"$gte": i.Ymin},
		"ymax":  bson.M{"$lte": i.Ymax},
		"start": bson.M{"$gte": i.Start},
		"stop":  bson.M{"$lte": i.Stop},
	}
}
