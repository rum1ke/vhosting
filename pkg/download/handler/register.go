package handler

import (
	"github.com/gin-gonic/gin"
	sess "vhosting/internal/session"
	"vhosting/pkg/auth"
	"vhosting/pkg/config"
	"vhosting/pkg/download"
	"vhosting/pkg/logger"
	"vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc download.DownloadUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewDownloadHandler(cfg, uc, luc, auc, suc, uuc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
