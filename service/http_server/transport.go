package http_server

import (
	"github.com/softlandia/hismap/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

// MsgError - упаковка ответа ошибки
type MsgError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

// MainHandler - ПУБЛИЧНЫЕ обработчики
func MainHandler(s *service.Service) *gin.Engine {
	info := infoTransport{s}

	g := gin.Default()
	//g.Use(MetricsMiddleware())

	api := g.Group("/api/v1")

	api.GET("/about", info.about)

	return g
}

// AdminHandler - АДМИНСКИЕ обработчики
func AdminHandler(s *service.Service) *gin.Engine {
	items := itemsTransport{s}

	g := gin.Default()
	//g.Use(MetricsMiddleware())

	api := g.Group("/api/admin")

	api.GET("/item/find/oid/:oid", items.findOid)

	return g
}

/*func MetricsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startedAt := time.Now()
		ctx.Next()
		matchedPath := ctx.GetString("matchedPath")
		statusCode := ctx.Writer.Status()
		metrics.HttpCounter.WithLabelValues(fmt.Sprint(statusCode), ctx.Request.Method, matchedPath).Inc()
		dur := time.Since(startedAt)
		metrics.HttpDuration.WithLabelValues(fmt.Sprint(statusCode), ctx.Request.Method, matchedPath).Observe(dur.Seconds())
	}
}*/

// InternalHandler - служебные обработчики
func InternalHandler(_ *zerolog.Logger) *http.ServeMux {
	info := infoTransport{}

	r := http.NewServeMux()
	r.HandleFunc("/readiness", info.okHandler)
	r.HandleFunc("/liveness", info.okHandler)
	r.Handle("/metrics", promhttp.Handler())

	return r
}
