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
	list := t.GetItems(filter)

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

// fillTestData - заполнить базу элементов тестовыми данными
func (t *itemsTransport) fillTestData(ctx *gin.Context) {
	oid := ctx.Param("OID")
	n := xlib.AtoI(ctx.Query("n"), 0)
	if n == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"n": "invalid"})
		return
	}

	itemList := t.FillTestItems(oid, n)
	if err := t.InsertItems(itemList); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"n":     "invalid",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"n": n})
}

func (t *itemsTransport) deleteOid(ctx *gin.Context) {
	oid := ctx.Param("OID")
	t.Repo.Items.Delete(bson.M{"oid": oid})
	ctx.JSON(http.StatusNoContent, nil)
}
