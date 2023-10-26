package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/open-beagle/awecloud-bmq-server/pkg/conf"
	"github.com/open-beagle/awecloud-bmq-server/pkg/log"
	"github.com/open-beagle/awecloud-bmq-server/pkg/router"
)

func NewAPIServer() *http.Server {
	gin.SetMode(conf.RunMode)
	r := gin.New()
	store := cookie.NewStore([]byte("beagle"))
	r.Use(sessions.Sessions("bmqsession", store))
	loadAPIRouter(r, log.GinLogger(), log.GinRecovery(true))
	conf.Logger.Info("server is starting...", zap.Int("port", conf.API.Port))
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.API.Port),
		Handler: r,
	}
}

// 加载路由
func loadAPIRouter(r *gin.Engine, middleware ...gin.HandlerFunc) {
	r.Use()
	r.Use(gin.Recovery())
	r.Use(middleware...)
	base := r.Group(conf.API.Prefix)
	{
		base.POST("/login", router.Login)
		base.GET("/logout", router.Login)
		base.GET("/login/user", router.LoginUser)
		base.GET("/login/captcha", router.LoginCaptchaGet)
		base.POST("/login/captcha/check", router.LoginCaptchaCheck)
	}
}
