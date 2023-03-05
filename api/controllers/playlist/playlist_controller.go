package playlist

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/playlist"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/song"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller struct {
	playlistService playlist.Service
	songService     song.Service
}

func NewPlaylistController(p *playlist.Service, s *song.Service) *Controller {
	return &Controller{
		playlistService: *p,
		songService:     *s,
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
	id, err := strconv.Atoi(ctx.Param("playlistId"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	playlistByID, err := c.playlistService.GetPlaylistByID(uint(id))

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

func (c *Controller) AddSong(ctx *gin.Context) {
	playlistID, err := strconv.Atoi(ctx.Param("playlistId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid playlist ID"})
		return
	}

	// получить идентификатор песни из URL-параметра
	songID, err := strconv.Atoi(ctx.Param("songId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song ID"})
		return
	}

	playlistById, err := c.playlistService.GetPlaylistByID(uint(playlistID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find playlist"})
			return
		}
	}

	songById, err := c.songService.GetSongByID(uint(songID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find song"})
			return
		}
	}

	err = c.playlistService.AddSong(playlistById, songById)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to append song to playlist"})
	}

	ctx.JSON(http.StatusCreated, gin.H{})
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
	id, err := strconv.Atoi(ctx.Param("playlistId"))

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
	cPlaylist.ID = uint(id)
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
	id, err := strconv.Atoi(ctx.Param("playlistId"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	err = c.playlistService.DeletePlaylist(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete playlist"})
			return
		}
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
