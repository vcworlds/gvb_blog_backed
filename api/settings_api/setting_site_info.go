package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/config"
	"gvb_blog/core"
	"gvb_blog/global"
	"gvb_blog/response"
)

type SettingsUri struct {
	Name string `json:"name" uri:"name" binding:"required"`
}

func (SettingsApi) SettingsSiteInfo(ctx *gin.Context) {
	sn := SettingsUri{}
	if err := ctx.ShouldBindUri(&sn); err != nil {
		response.Fail(ctx, "未找到路由参数信息")
		return
	}
	switch sn.Name {
	case "site":
		response.OkWithData(ctx, global.Config.SiteInfo)
	case "qq":
		response.OkWithData(ctx, global.Config.QQ)
	case "email":
		response.OkWithData(ctx, global.Config.Email)
	case "jwt":
		response.OkWithData(ctx, global.Config.Jwt)
	case "qiniu":
		response.OkWithData(ctx, global.Config.QiNiu)
	default:
		response.Fail(ctx, "暂时没有找到该站点信息")
	}
}

func (SettingsApi) SettingsSiteInfoUpdate(ctx *gin.Context) {
	sn := SettingsUri{}
	if err := ctx.ShouldBindUri(&sn); err != nil {
		response.Fail(ctx, "未找到路由参数信息")
		return
	}
	switch sn.Name {
	case "site":
		var siteInfo config.SiteInfo
		err := ctx.ShouldBindJSON(&siteInfo)
		if err != nil {
			response.OkWithMessage(ctx, "参数错误")
			return
		}
		global.Config.SiteInfo = siteInfo
		err = core.SetYaml()
		if err != nil {
			global.Log.Fatalf("配置文件修改失败", err)
			return
		}
		response.OkWith(ctx)
	case "qq":
		var qq config.QQ
		err := ctx.ShouldBindJSON(&qq)
		if err != nil {
			response.OkWithMessage(ctx, "参数错误")
			return
		}
		global.Config.QQ = qq
		err = core.SetYaml()
		if err != nil {
			global.Log.Fatalf("配置文件修改失败", err)
			return
		}
		response.OkWith(ctx)
	case "email":
		var email config.Email
		if err := ctx.ShouldBindJSON(&email); err != nil {
			response.Fail(ctx, "参数错误")
			return
		}
		global.Config.Email = email
		if err := core.SetYaml(); err != nil {
			response.Fail(ctx, "配置文件修改错误")
			global.Log.Fatal(err)
			return
		}
		response.OkWith(ctx)
	case "jwt":
		var jwt config.Jwt
		err := ctx.ShouldBindJSON(&jwt)
		if err != nil {
			response.OkWithMessage(ctx, "参数错误")
			return
		}
		global.Config.Jwt = jwt
		err = core.SetYaml()
		if err != nil {
			global.Log.Fatalf("配置文件修改失败", err)
			return
		}
		response.OkWith(ctx)
	case "qiniu":
		var QN config.QiNiu
		err := ctx.ShouldBindJSON(&QN)
		if err != nil {
			response.OkWithMessage(ctx, "参数错误")
			return
		}
		global.Config.QiNiu = QN
		err = core.SetYaml()
		if err != nil {
			global.Log.Fatalf("配置文件修改失败", err)
			return
		}
	default:
		response.Fail(ctx, "暂时没有找到该站点信息")
	}

}
