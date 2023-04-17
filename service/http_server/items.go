package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/softlandia/hismap/service"
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
