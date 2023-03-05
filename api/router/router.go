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
				playlists.GET("/", crs.PlaylistController.GetAll)
				playlists.POST("/", crs.PlaylistController.Create)
				playlists.POST("/:playlistID/songs/:songId", crs.PlaylistController.AddSong)
				playlists.GET("/:playlistId", crs.PlaylistController.GetById)
				playlists.PUT("/:playlistId", crs.PlaylistController.Update)
				playlists.DELETE("/:playlistId", crs.PlaylistController.Delete)
			}

			songs := v1.Group("/songs")
			{
				songs.GET("/", crs.SongController.GetAll)
				songs.POST("/", crs.SongController.Create)
				songs.GET("/:songId", crs.SongController.GetById)
				songs.PUT("/:songId", crs.SongController.Update)
				songs.DELETE("/:songId", crs.SongController.Delete)
			}
		}
	}

	return router
}
