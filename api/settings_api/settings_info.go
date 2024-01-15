package settings_api

import (
	"main/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi)SettingsInfoView(c *gin.Context)  {
	res.Ok(map[string]any{}, "settings info", c)
}