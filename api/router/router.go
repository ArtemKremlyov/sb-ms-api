package router

import (
	"gitlab.com/sb-cloud/player-ms-api/api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(crs *controllers.Controllers) *gin.Engine {
	router := gin.New()

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			playlists := v1.Group("/playlists")
			{
				playlists.GET("/", crs.PlaylistController)
				playlists.POST("/")
				playlists.GET("/:playlistId")
				playlists.PUT("/:playlistId")
				playlists.DELETE(":/playlistId")
			}
		}
	}

	return router
}
