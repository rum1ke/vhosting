package handler

import (
	"github.com/gin-gonic/gin"
	sess "vhosting/internal/session"
	"vhosting/internal/video"
	"vhosting/pkg/auth"
	"vhosting/pkg/config"
	"vhosting/pkg/logger"
	"vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc video.VideoUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewVideoHandler(cfg, uc, luc, auc, suc, uuc)

	videoRoute := router.Group("/video")
	{
		videoRoute.POST("", h.CreateVideo)
		videoRoute.GET(":id", h.GetVideo)
		videoRoute.GET("all", h.GetAllVideos)
		videoRoute.PATCH(":id", h.PartiallyUpdateVideo)
		videoRoute.DELETE(":id", h.DeleteVideo)
	}
}
