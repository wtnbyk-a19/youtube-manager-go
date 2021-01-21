package middlewares

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
)

func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)

	return jwtToken, err
}
