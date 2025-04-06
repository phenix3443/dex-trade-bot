package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/phenix3443/dex-trade-bot/pkg/api/router/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1Group := r.Group("/api/v1")
	{
		v1.RegisterWalletRoutes(v1Group)
	}

	return r
}
