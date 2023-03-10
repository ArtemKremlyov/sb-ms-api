package rest

import (
	"github.com/stretchr/testify/assert"
	mock "gitlab.com/sb-cloud/player-ms-api/internal/mock/api/rest"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"testing"
)

func TestMusic_GetAllPlaylists(t *testing.T) {
	expectedPlaylists := []models.Playlist{
		{ID: 1, Name: "Jazz"},
		{ID: 2, Name: "Rock"},
		{ID: 3, Name: "Metal"},
	}

	mock := &mock.Music{}
	mock.On("GetAllPlaylists").Return(expectedPlaylists, nil)

	playlists, err := mock.GetAllPlaylists()
	assert.NoError(t, err)
	assert.Equal(t, expectedPlaylists, playlists)

	mock.AssertExpectations(t)
}

func TestMusic_GetPlaylistById(t *testing.T) {
	id := uint(1)
	expectedPlaylist := &models.Playlist{
		ID:   1,
		Name: "Phonk",
	}

	mock := mock.Music{}
	mock.On("GetPlaylistByID", id).Return(expectedPlaylist, nil)

	playlist, err := mock.GetPlaylistByID(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedPlaylist, playlist)

	mock.AssertExpectations(t)
}

func TestMusic_AddSong(t *testing.T) {
	testPlaylist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	testSong := &models.Song{
		ID:       1,
		Name:     "Song",
		Duration: 1,
	}

	mock := mock.Music{}
	mock.On("AddSong", testPlaylist, testSong).Return(nil)

	err := mock.AddSong(testPlaylist, testSong)

	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func TestMusic_CreatePlaylist(t *testing.T) {
	newPlaylist := &models.Playlist{Name: "Rap"}

	mock := &mock.Music{}
	mock.On("CreatePlaylist", newPlaylist).Return(nil)

	err := mock.CreatePlaylist(newPlaylist)

	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func TestMusic_UpdatePlaylist(t *testing.T) {
	playlist := &models.Playlist{ID: 1, Name: "Playlist"}

	mock := &mock.Music{}
	mock.On("UpdatePlaylist", playlist).Return(nil)

	err := mock.UpdatePlaylist(playlist)

	assert.NoError(t, err)

	mock.AssertExpectations(t)

}

func TestMusic_DeletePlaylist(t *testing.T) {
	mock := &mock.Music{}
	id := uint(1)
	mock.On("DeletePlaylist", id).Return(nil)

	err := mock.DeletePlaylist(id)

	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func TestMusic_GetAllSongs(t *testing.T) {
	expectedSongs := []models.Song{
		{ID: 1, Name: "Track 1", Duration: 30},
		{ID: 2, Name: "Track 2", Duration: 20},
		{ID: 3, Name: "Track 3", Duration: 10},
	}

	mock := &mock.Music{}
	mock.On("GetAllSongs").Return(expectedSongs, nil)

	playlists, err := mock.GetAllSongs()
	assert.NoError(t, err)
	assert.Equal(t, expectedSongs, playlists)

	mock.AssertExpectations(t)
}

func TestMusic_GetSongById(t *testing.T) {
	expectedSong := &models.Song{
		ID:       1,
		Name:     "Track 1",
		Duration: 10,
	}
	id := uint(1)

	mock := mock.Music{}
	mock.On("GetSongByID", id).Return(expectedSong, nil)

	song, err := mock.GetSongByID(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedSong, song)

	mock.AssertExpectations(t)

}

func TestMusic_CreateSong(t *testing.T) {
	newSong := &models.Song{Name: "Rap", Duration: 1}

	mock := &mock.Music{}
	mock.On("CreateSong", newSong).Return(nil)

	err := mock.CreateSong(newSong)

	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func TestMusic_UpdateSong(t *testing.T) {
	song := &models.Song{ID: 1, Name: "Song", Duration: 1}

	mock := &mock.Music{}
	mock.On("UpdateSong", song).Return(nil)

	err := mock.UpdateSong(song)

	assert.NoError(t, err)

	mock.AssertExpectations(t)

}

func TestMusic_DeleteSong(t *testing.T) {
	mock := &mock.Music{}
	id := uint(1)
	mock.On("DeleteSong", id).Return(nil)

	err := mock.DeleteSong(id)

	assert.NoError(t, err)

	mock.AssertExpectations(t)
}
