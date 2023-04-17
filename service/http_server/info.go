package http_server

import (
	"github.com/softlandia/hismap/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type infoTransport struct {
	*service.Service
}

func (t *infoTransport) about(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct {
		AppType    string `json:"appType"`
		AppVersion string `json:"appVersion"`
		Email      string `json:"email"`
		Time       string `json:"time"`
	}{
		AppType:    "hismap service",
		AppVersion: "0.0.1",
		Email:      "softlandia@gmail.com",
		Time:       time.Now().Format(time.RFC822),
	})

}

func (t *infoTransport) okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
