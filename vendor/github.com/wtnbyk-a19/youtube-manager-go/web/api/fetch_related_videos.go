package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchRelatedVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)

		videoId := c.Param("id")

		part := []string{"id", "snippet"}
		call := yts.Search.List(part).RelatedToVideoId(videoId).Type("video").MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API : %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
