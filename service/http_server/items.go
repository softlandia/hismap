package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/softlandia/hismap/models"
	"github.com/softlandia/hismap/service"
	"github.com/softlandia/xlib"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type itemsTransport struct {
	*service.Service
}

func (t *itemsTransport) findOid(ctx *gin.Context) {
	res, err := t.Repo.Items.FindOne(bson.M{"oid": ctx.Param("oid")})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"result": err.Error()})
	}

	ctx.JSON(200, res)
}

// find - выборка объектов на карту
func (t *itemsTransport) find(ctx *gin.Context) {
	Xmin := xlib.AtoI(ctx.Query("xmin"), 0)
	Xmax := xlib.AtoI(ctx.Query("xmax"), 0)
	Ymin := xlib.AtoI(ctx.Query("ymin"), 0)
	Ymax := xlib.AtoI(ctx.Query("ymax"), 0)
	Start := xlib.AtoI(ctx.Query("start"), 0)
	Stop := xlib.AtoI(ctx.Query("stop"), 0)
	filter := models.ItemsFilter{
		Xmin:  Xmin,
		Xmax:  Xmax,
		Ymin:  Ymin,
		Ymax:  Ymax,
		Start: int32(Start),
		Stop:  int32(Stop),
	}
	//list := t.GetItems(filter)
	list := t.Repo.ItemsV2.Find(filter.ToRepo())

	ctx.JSON(200, gin.H{
		"Xmin":  Xmin,
		"Xmax":  Xmax,
		"Ymin":  Ymin,
		"Ymax":  Ymax,
		"Start": Start,
		"Stop":  Stop,
		"count": len(list),
		"data":  list,
	})
}

/*// createFooData - поместить в базу элемент с придуманными координатами
func (t *itemsTransport) createFooData(ctx *gin.Context) {
	oid := ctx.Param("OID")
	if len(oid) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"oid": "invalid"})
		return
	}
	vertex := xlib.AtoI(ctx.Param("VERTEX"), 0) // количество узлов в границе
	if vertex <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"VERTEX": "invalid"})
		return
	}
	var data models.Item
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	data.Border = models.FooBorder2(vertex)
	if _, err := t.InsertOneItem(data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}*/

// createFooData - поместить в базу элемент с придуманными координатами
func (t *itemsTransport) createFooData(ctx *gin.Context) {
	oid := ctx.Param("OID")
	if len(oid) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"oid": "invalid"})
		return
	}
	vertex := xlib.AtoI(ctx.Param("VERTEX"), 0) // количество узлов в границе
	if vertex <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"VERTEX": "invalid"})
		return
	}
	var data models.Item
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	data.X = models.FooVector(vertex)
	data.Y = models.FooVector(vertex)
	if _, err := t.Repo.ItemsV2.InsertOne(data.ToRepoV2()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (t *itemsTransport) deleteOid(ctx *gin.Context) {
	oid := ctx.Param("OID")
	t.Repo.Items.Delete(bson.M{"oid": oid})
	ctx.JSON(http.StatusNoContent, nil)
}
