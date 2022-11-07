package handler

import (
	"vhosting/internal/info"
	sess "vhosting/internal/session"
	"vhosting/pkg/auth"
	"vhosting/pkg/config"
	sconfig "vhosting/pkg/config_stream"
	"vhosting/pkg/logger"
	"vhosting/pkg/user"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc info.InfoUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewInfoHandler(cfg, scfg, uc, luc, auc, suc, uuc)

	infoRoute := router.Group("/info")
	{
		infoRoute.POST("", h.CreateInfo)
		infoRoute.GET(":id", h.GetInfo)
		infoRoute.GET("all", h.GetAllInfos)
		infoRoute.PATCH(":id", h.PartiallyUpdateInfo)
		infoRoute.DELETE(":id", h.DeleteInfo)
	}
}
