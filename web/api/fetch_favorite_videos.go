package api

import (
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/wtnbyk-a19/youtube-manager-go/middlewares"
	"github.com/wtnbyk-a19/youtube-manager-go/models"
	"google.golang.org/api/youtube/v3"
)

func FetchFavoriteVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		user := models.User{}
		dbs.DB.Table("users").Where(models.User{UID: token.UID}).First(&user)

		favorites := []models.Favorite{}
		dbs.DB.Model(&user).Related(favorites)

		videoIds := ""
		for _, f := range favorites {
			if len(videoIds) == 0 {
				videoIds += f.VideoId
			} else {
				videoIds += "," + f.VideoId
			}
		}

		yts := c.Get("yts").(*youtube.Service)

		part := []string{"id", "snippet"}
		call := yts.Videos.List(part).Id(videoIds).MaxResults(10)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
