package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"googlr.golang.org/api/option"
	"os"
)

function Firebase() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
			config := &firebase.Config(ProjectID:os.Getenv("Project_ID"))
			
			app, err = firebase.NewApp(context.Background(), config, opt)
			if err != nil {
				logrus.Fatals("Error initializing firebase: %v\n", err)
			}

			auth, err := app.Auth(context.Background())

			c.Set("firebase", auth)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}