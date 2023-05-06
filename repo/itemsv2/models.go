package itemsv2

import "go.mongodb.org/mongo-driver/bson/primitive"

// Items - набор элементов карты
type Items []Item

// Item - запись в базе одного элемента
type Item struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	OID    string             `bson:"oid"`
	Object Object             `bson:"object"`
	Start  int32              `bson:"start"`
	Stop   int32              `bson:"stop"`
	Xmin   int                `bson:"xmin"`
	Xmax   int                `bson:"xmax"`
	Ymax   int                `bson:"ymax"`
	Ymin   int                `bson:"ymin"`
	X      []int32            `bson:"x"`
	Y      []int32            `bson:"y"`
}

// Object - объект исторической карты, логический элемент, может иметь много Item для разного времени и местоположения
type Object struct {
	Caption string `bson:"caption"`  // название отображаемое на карте
	StyleId string `bson:"style_id"` // uuid стиля
}
