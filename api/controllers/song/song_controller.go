package song

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/song"
	"net/http"
)

type Controller struct {
	SongService song.Service
}

func NewSongController(s *song.Service) *Controller {
	return &Controller{
		SongService: *s,
	}
}

func (c *Controller) GetAll(ctx *gin.Context) {
	songs, err := c.SongService.GetAllSongs()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all songs: server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, songs)
}
