package models

import (
	"github.com/softlandia/hismap/repo/items"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Items - набор элементов на карте
type Items []Item

// Item - элемент карты
type Item struct {
	ID     primitive.ObjectID
	OID    string // ссылка на объект исторической карты
	Object Object // описание объекта
	Start  int32  // дата начала
	Stop   int32  // дата окончания
	Xmin   int    // левый край
	Xmax   int    // правый край
	Ymin   int    // нижний край
	Ymax   int    // верхний край
	Border Border // граница
}

// Object - объект исторической карты, логический элемент, может иметь много Item для разного времени и местоположения
type Object struct {
	Caption string // название отображаемое на карте
	StyleId string // uuid стиля
}

// Border - граница Item
type Border []Vertex

// ToRepo - преобразование типа Border -> items.Border
func (b Border) ToRepo() items.Border {
	res := make(items.Border, 0, len(b))
	for _, e := range b {
		res = append(res, items.Vertex{
			X: e.X,
			Y: e.Y,
		})
	}
	return res
}

// RepoToBorder - преобразование типа items.Border -> Border
func RepoToBorder(b items.Border) Border {
	res := make(Border, 0, len(b))
	for _, e := range b {
		res = append(res, Vertex{
			X: e.X,
			Y: e.Y,
		})
	}
	return res
}

// Vertex - элемент границы
type Vertex struct {
	X int32
	Y int32
}

// RepoToItem - преобразование типа items.Item -> Item
func RepoToItem(r items.Item) Item {
	return Item{
		ID:  r.ID,
		OID: r.OID,
		Object: Object{
			Caption: r.Object.Caption,
			StyleId: r.Object.StyleId,
		},
		Start:  r.Start,
		Stop:   r.Stop,
		Xmin:   r.Xmin,
		Xmax:   r.Xmax,
		Ymin:   r.Ymin,
		Ymax:   r.Ymax,
		Border: RepoToBorder(r.Border),
	}
}

// ToRepo - преобразование типа Item -> items.Item
func (i Item) ToRepo() items.Item {
	return items.Item{
		ID:  i.ID,
		OID: i.OID,
		Object: items.Object{
			Caption: i.Object.Caption,
			StyleId: i.Object.StyleId,
		},
		Start:  i.Start,
		Stop:   i.Stop,
		Xmin:   i.Xmin,
		Xmax:   i.Xmax,
		Ymin:   i.Ymin,
		Ymax:   i.Ymax,
		Border: i.Border.ToRepo(),
	}
}
