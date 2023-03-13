package rest

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Music interface {
	GetAllPlaylists() ([]models.Playlist, error)
	GetPlaylistByID(id uint) (*models.Playlist, error)
	AddSong(playlist *models.Playlist, song *models.Song) error
	CreatePlaylist(playlist *models.Playlist) error
	UpdatePlaylist(playlist *models.Playlist) error
	DeletePlaylist(id uint) error

	GetAllSongs() ([]models.Song, error)
	GetSongByID(id uint) (*models.Song, error)
	CreateSong(song *models.Song) error
	UpdateSong(song *models.Song) error
	DeleteSong(id uint) error
}

type REST struct {
	music Music
}

func NewREST(music Music) *REST {
	return &REST{music: music}
}

func (r REST) Register(router *gin.Engine) *gin.Engine {
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
				playlists.GET("/", r.PlaylistGetAll)
				playlists.POST("/", r.PlaylistCreate)
				playlists.POST("/:playlistID/songs/:songId", r.PlaylistAddSong)
				playlists.GET("/:playlistId", r.PlaylistGetById)
				playlists.PUT("/:playlistId", r.PlaylistUpdate)
				playlists.DELETE("/:playlistId", r.PlaylistDelete)
			}

			songs := v1.Group("/songs")
			{
				songs.GET("/", r.SongGetAll)
				songs.POST("/", r.SongCreate)
				songs.GET("/:songId", r.SongGetById)
				songs.PUT("/:songId", r.SongUpdate)
				songs.DELETE("/:songId", r.SongDelete)
			}
		}
	}

	return router
}

// Playlists
func (r REST) PlaylistGetAll(ctx *gin.Context) {
	playlists, err := r.music.GetAllPlaylists()

	if err != nil {
		log.Printf("failed to get all playlist: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all playlist: server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

func (r REST) PlaylistGetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("playlistId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	playlistByID, err := r.music.GetPlaylistByID(uint(id))

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			log.Printf("failed to get playlistByID: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get playlistByID: server error",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, playlistByID)
}

func (r REST) PlaylistAddSong(ctx *gin.Context) {
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

	playlistById, err := r.music.GetPlaylistByID(uint(playlistID))
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		} else {
			log.Printf("failed to find playlist: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find playlist"})
		}
		return
	}

	songById, err := r.music.GetSongByID(uint(songID))
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			log.Printf("failed to find song: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find song"})
		}
		return
	}

	err = r.music.AddSong(playlistById, songById)

	if err != nil {
		log.Printf("failed to append song to playlist: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to append song to playlist"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r REST) PlaylistCreate(ctx *gin.Context) {
	var cPlaylist models.Playlist

	if err := ctx.BindJSON(&cPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(cPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := r.music.CreatePlaylist(&cPlaylist); err != nil {
		log.Printf("failed to create playlist: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create playlist",
		})
		return
	}

	ctx.JSON(http.StatusCreated, cPlaylist)
}

func (r REST) PlaylistUpdate(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("playlistId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	validate := validator.New()
	if err := validate.Struct(cPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.music.UpdatePlaylist(&cPlaylist)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		} else {
			log.Printf("failed to update playlist: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update playlist",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, cPlaylist)
}

func (r REST) PlaylistDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("playlistId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid playlist id",
		})
		return
	}

	err = r.music.DeletePlaylist(uint(id))
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "playlist not found"})
		} else {
			log.Printf("failed to delete playlist: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete playlist"})
		}
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

// Songs
func (r REST) SongGetAll(ctx *gin.Context) {
	songs, err := r.music.GetAllSongs()

	if err != nil {
		log.Printf("failed to get all songs: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all songs: server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, songs)
}

func (r REST) SongGetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid song id",
		})
		return
	}

	songById, err := r.music.GetSongByID(uint(id))

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			log.Printf("failed to get songById: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get songById: server error",
			})
		}
		return

	}

	ctx.JSON(http.StatusOK, songById)
}

func (r REST) SongCreate(ctx *gin.Context) {
	var cSong models.Song

	if err := ctx.BindJSON(&cSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(cSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := r.music.CreateSong(&cSong); err != nil {
		log.Printf("failed to create song: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create song",
		})
		return
	}

	ctx.JSON(http.StatusCreated, cSong)
}

func (r REST) SongUpdate(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	validate := validator.New()
	if err := validate.Struct(cSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.music.UpdateSong(&cSong)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			log.Printf("failed to update song: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update song",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, cSong)
}

func (r REST) SongDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("songId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid song id",
		})
		return
	}

	err = r.music.DeleteSong(uint(id))
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		} else {
			log.Printf("failed to delete song: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete song"})
		}
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
