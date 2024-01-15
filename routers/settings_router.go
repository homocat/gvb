package routers

import (
	"main/api"
)

func (r RouterGroup)SettingsRouter()  {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings", settingsApi.SettingsInfoView)
}