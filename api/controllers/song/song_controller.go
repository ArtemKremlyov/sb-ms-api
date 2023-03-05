package song

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/song"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (c *Controller) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	songById, err := c.SongService.GetSongByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get songById: server error",
		})
		return
	}

	if songById == nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, songById)
}

func (c *Controller) Create(ctx *gin.Context) {
	var cSong models.Song

	if err := ctx.BindJSON(&cSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	if err := c.SongService.CreateSong(&cSong); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create playlist",
		})
		return
	}

	ctx.JSON(http.StatusCreated, cSong)
}

func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	var cSong models.Song
	if err := ctx.BindJSON(&cSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	cSong.ID = uint(id)
	err = c.SongService.UpdateSong(&cSong)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update playlist",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, cSong)
}

func (c *Controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid song id",
		})
		return
	}

	err = c.SongService.DeleteSong(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete song"})
			return
		}
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
