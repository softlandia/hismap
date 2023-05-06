package models

import (
	"github.com/softlandia/hismap/repo/items"
	"github.com/softlandia/hismap/repo/itemsv2"
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
	X      []int32
	Y      []int32
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

// ToRepoV2 - преобразование типа Item -> itemsv2.Item
func (i Item) ToRepoV2() itemsv2.Item {
	return itemsv2.Item{
		ID:  i.ID,
		OID: i.OID,
		Object: itemsv2.Object{
			Caption: i.Object.Caption,
			StyleId: i.Object.StyleId,
		},
		Start: i.Start,
		Stop:  i.Stop,
		Xmin:  i.Xmin,
		Xmax:  i.Xmax,
		Ymin:  i.Ymin,
		Ymax:  i.Ymax,
		X:     i.X,
		Y:     i.Y,
	}
}

func FooVector(n int) []int32 {
	res := make([]int32, n)
	res[0] = rand.Int31n(1000)

	for i := 1; i < n; i++ {
		d := rand.Int31n(10)
		sign := rand.Int31n(3) - 1
		res[i] = res[i-1] + sign*d
	}
	return res
}

func FooBorder2(n int) Border {
	res := make(Border, n)
	res[0] = Vertex{
		X: rand.Int31n(1000),
		Y: rand.Int31n(1000),
	}
	var dxSign int32 = 1
	var dySign int32 = 1
	for i := 1; i < n; i++ {
		dx := rand.Int31n(10)
		if rand.Int31n(1) == 0 {
			dxSign = -1
		}
		dy := rand.Int31n(10)
		if rand.Int31n(1) == 0 {
			dySign = -1
		}
		res[i] = Vertex{
			X: res[i-1].X + dxSign*dx,
			Y: res[i-1].Y + dySign*dy,
		}
	}
	return res
}
