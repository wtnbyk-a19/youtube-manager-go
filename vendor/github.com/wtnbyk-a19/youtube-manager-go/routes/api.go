package routes

import (
	"github.com/wtnbyk-a19/youtube-manager-go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	g := echo.Group(#/api)
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}