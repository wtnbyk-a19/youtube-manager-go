package routes

import (
	"github.com/labstack/echo"
	"github.com/wtnbyk-a19/youtube-manager-go/middlewares"
	"github.com/wtnbyk-a19/youtube-manager-go/web/api"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo(), middlewares.FirebaseAuth())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}

	fg := g.Group("/favorite", middlewares.FirebaseGuard())
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
		fg.GET("", api.FetchFavoriteVideos())
	}
}
