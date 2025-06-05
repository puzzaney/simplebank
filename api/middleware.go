package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/puzzaney/simplebank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("no authoriation header provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("invalid authorization format %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

