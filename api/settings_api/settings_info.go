package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/response"
	"net/http"
)

func (SettingsApi) SettingsInfoView(ctx *gin.Context) {
	response.Result(ctx, http.StatusOK, 200, "获取成功", gin.H{"id": 1})
}
