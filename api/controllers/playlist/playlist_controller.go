package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/playlist"
	"net/http"
	"strconv"
)

type Controller struct {
	playlistService playlist.Service
}

func NewPlaylistController(s *playlist.Service) *Controller {
	return &Controller{
		playlistService: *s,
	}
}

func (c *Controller) GetAll(ctx *gin.Context) {
	playlists, err := c.playlistService.GetAllPlaylists()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all playlist: server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

func (c *Controller) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	playlistByID, err := c.playlistService.GetPlaylistByID(int64(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get playlistByID: server error",
		})
		return
	}

	if playlistByID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, playlistByID)
}

func (c *Controller) Create(ctx *gin.Context) {
	var cPlaylist models.Playlist

	if err := ctx.BindJSON(&cPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	if err := c.playlistService.CreatePlaylist(&cPlaylist); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create playlist",
		})
		return
	}

	ctx.JSON(http.StatusCreated, cPlaylist)
}

func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	var cPlaylist models.Playlist
	if err := ctx.BindJSON(&cPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	cPlaylist.ID = int64(id)
	err = c.playlistService.UpdatePlaylist(&cPlaylist)
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

	ctx.JSON(http.StatusOK, cPlaylist)
}

func (c *Controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	err = c.playlistService.DeletePlaylist(int64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete playlist"})
			return
		}
	}
}
