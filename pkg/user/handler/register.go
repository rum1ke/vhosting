package handler

import (
	"github.com/gin-gonic/gin"
	sess "vhosting/internal/session"
	"vhosting/pkg/auth"
	"vhosting/pkg/config"
	"vhosting/pkg/logger"
	"vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc user.UserUseCase,
	luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewUserHandler(cfg, uc, luc, auc, suc)

	userRoute := router.Group("/user")
	{
		userRoute.POST("", h.CreateUser)
		userRoute.GET(":id", h.GetUser)
		userRoute.GET("all", h.GetAllUsers)
		userRoute.POST("/change_password", h.UpdateUserPassword)
		userRoute.PATCH(":id", h.PartiallyUpdateUser)
		userRoute.DELETE(":id", h.DeleteUser)
	}
}
