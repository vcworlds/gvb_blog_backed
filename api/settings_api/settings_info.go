package settings_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "你好gin"})
}
