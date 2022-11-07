package handler

import (
	"github.com/gin-gonic/gin"
	sess "vhosting/internal/session"
	"vhosting/pkg/auth"
	"vhosting/pkg/config"
	"vhosting/pkg/logger"
	"vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc auth.AuthUseCase, uuc user.UserUseCase,
	suc sess.SessUseCase, luc logger.LogUseCase) {
	h := NewAuthHandler(cfg, uc, uuc, suc, luc)

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/signin", h.SignIn)
		authRoute.POST("/change_password", h.ChangePassword)
		authRoute.GET("/signout", h.SignOut)
	}
}
