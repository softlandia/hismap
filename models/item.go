package models

import (
	"github.com/softlandia/hismap/repo/items"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
)

// ItemList - набор элементов на карте
type ItemList []Item

// ToRepo - преобразование типа ItemList -> items.Items
func (il ItemList) ToRepo() items.Items {
	res := make(items.Items, 0, len(il))
	for _, e := range il {
		res = append(res, e.ToRepo())
	}
	return res
}

// RepoToItemList - преобразование типа items.Items -> ItemList
func RepoToItemList(list items.Items) ItemList {
	res := make(ItemList, 0, len(list))
	for _, e := range list {
		res = append(res, RepoToItem(e))
	}
	return res
}

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

func FooItem(oid string) Item {
	start := rand.Int31()
	xMin := int(rand.Int31())
	yMin := int(rand.Int31())
	return Item{
		OID: oid,
		Object: Object{
			Caption: "fill-test",
			StyleId: "fill-test",
		},
		Start:  start,
		Stop:   start + rand.Int31(),
		Xmin:   xMin,
		Xmax:   xMin + rand.Intn(1000),
		Ymin:   yMin,
		Ymax:   yMin + rand.Intn(500),
		Border: FooBorder(rand.Intn(50)),
	}
}

func FooBorder(n int) Border {
	res := make(Border, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, Vertex{
			X: rand.Int31(),
			Y: rand.Int31(),
		})
	}
	return res
}
