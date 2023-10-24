package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/open-beagle/awecloud-bmq-server/pkg/conf"
	"github.com/open-beagle/awecloud-bmq-server/pkg/log"
	"github.com/open-beagle/awecloud-bmq-server/pkg/middleware/header"
)

func NewAPIServer() *http.Server {
	// set run mod
	gin.SetMode(conf.RunMode)
	// load gin router
	r := gin.New()
	loadAPIRouter(r, log.GinLogger(), log.GinRecovery(true))
	conf.Logger.Info("server is starting...", zap.Int("port", conf.API.Port))
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.API.Port),
		Handler: r,
		// ReadTimeout:    30 * time.Second,
		// WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

// 加载路由
func loadAPIRouter(r *gin.Engine, middleware ...gin.HandlerFunc) {
	r.Use()
	r.Use(gin.Recovery())
	r.Use(header.NoCache)
	r.Use(header.Options)
	r.Use(header.Secure)
	r.Use(middleware...)
	// base := r.Group(conf.APIPrefix)
	// {
	// 	base.GET("/health", controller.Health) // 健康检查
	// 	List(base)                             //下拉列表
	// 	Secure(base)                           //安全组
	// 	Excel(base)                            //表格
	// }

}
