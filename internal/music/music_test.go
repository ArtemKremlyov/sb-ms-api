package music_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	mock "gitlab.com/sb-cloud/player-ms-api/internal/mock/music"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gitlab.com/sb-cloud/player-ms-api/internal/music"
	"testing"
)

type MusicSuite struct {
	suite.Suite
	mockPG  *mock.PostgresImpl
	service *music.Service
}

func (m *MusicSuite) SetupTest() {
	m.mockPG = mock.NewPostgresImpl(m.T())
	m.service = music.NewMusicService(m.mockPG)
}

func (m *MusicSuite) TestMusic_GetAllPlaylists() {
	expectedPlaylists := []models.Playlist{
		{ID: 1, Name: "Jazz"},
		{ID: 2, Name: "Rock"},
		{ID: 3, Name: "Metal"},
	}

	m.mockPG.EXPECT().PlaylistGetAll().Return(expectedPlaylists, nil).Times(1)

	playlist, err := m.service.GetAllPlaylists()
	assert.Equal(m.T(), expectedPlaylists, playlist)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorGetAllPlaylists() {
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().PlaylistGetAll().Return(nil, expectedErrorPostgres).Times(1)

	playlist, err := m.service.GetAllPlaylists()
	assert.Nil(m.T(), playlist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetPlaylistById() {
	id := uint(1)
	expectedPlaylist := &models.Playlist{
		ID:   1,
		Name: "Phonk",
	}

	m.mockPG.EXPECT().PlaylistGetByID(id).Return(expectedPlaylist, nil).Times(1)

	playlist, err := m.service.GetPlaylistByID(id)
	assert.Equal(m.T(), expectedPlaylist, playlist)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorGetPlaylistById() {
	id := uint(1)
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().PlaylistGetByID(id).Return(nil, expectedErrorPostgres).Times(1)

	playlist, err := m.service.GetPlaylistByID(id)
	assert.Nil(m.T(), playlist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_CreatePlaylist() {
	newPlaylist := &models.Playlist{Name: "Rap"}

	m.mockPG.EXPECT().PlaylistCreate(newPlaylist).Return(nil).Times(1)

	err := m.service.CreatePlaylist(newPlaylist)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorCreatePlaylist() {
	newPlaylist := &models.Playlist{Name: "Rap"}
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().PlaylistCreate(newPlaylist).Return(expectedErrorPostgres).Times(1)

	err := m.service.CreatePlaylist(newPlaylist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_UpdatePlaylist() {
	playlist := &models.Playlist{ID: 1, Name: "Playlist"}

	m.mockPG.EXPECT().PlaylistUpdate(playlist).Return(nil).Times(1)

	err := m.service.UpdatePlaylist(playlist)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorUpdatePlaylist() {
	playlist := &models.Playlist{ID: 1, Name: "Playlist"}
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().PlaylistUpdate(playlist).Return(expectedErrorPostgres).Times(1)

	err := m.service.UpdatePlaylist(playlist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetPlaylistSongs() {
	playlist := &models.Playlist{ID: 1, Name: "Playlist"}
	expectedSongs := []models.Song{
		{ID: 1, Name: "Track 1", Duration: 30},
		{ID: 2, Name: "Track 2", Duration: 20},
		{ID: 3, Name: "Track 3", Duration: 10},
	}

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(expectedSongs, nil).Times(1)

	songs, err := m.service.GetPlaylistSongs(playlist)
	assert.NoError(m.T(), err)
	assert.Equal(m.T(), expectedSongs, songs)
}

func (m *MusicSuite) TestMusic_ErrorGetPlaylistSongs() {
	playlist := &models.Playlist{ID: 1, Name: "Playlist"}
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(nil, expectedErrorPostgres).Times(1)

	songs, err := m.service.GetPlaylistSongs(playlist)
	assert.Error(m.T(), err)
	assert.Nil(m.T(), songs)
}

func (m *MusicSuite) TestMusic_DeletePlaylist() {
	id := uint(1)

	m.mockPG.EXPECT().PlaylistDelete(id).Return(nil).Times(1)

	err := m.service.DeletePlaylist(id)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorDeletePlaylist() {
	id := uint(1)
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().PlaylistDelete(id).Return(expectedErrorPostgres).Times(1)

	err := m.service.DeletePlaylist(id)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_AddSong() {
	testPlaylist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	testSong := &models.Song{
		ID:       1,
		Name:     "Song",
		Duration: 1,
	}

	m.mockPG.EXPECT().PlaylistAddSong(testPlaylist, testSong).Return(nil).Times(1)

	err := m.service.AddSong(testPlaylist, testSong)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorAddSong() {
	expectedErrorPostgres := errors.New("something with postgres")
	testPlaylist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	testSong := &models.Song{
		ID:       1,
		Name:     "Song",
		Duration: 1,
	}

	m.mockPG.EXPECT().PlaylistAddSong(testPlaylist, testSong).Return(expectedErrorPostgres).Times(1)

	err := m.service.AddSong(testPlaylist, testSong)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetAllSongs() {
	expectedSongs := []models.Song{
		{ID: 1, Name: "Track 1", Duration: 30},
		{ID: 2, Name: "Track 2", Duration: 20},
		{ID: 3, Name: "Track 3", Duration: 10},
	}

	m.mockPG.EXPECT().SongGetAll().Return(expectedSongs, nil).Times(1)

	song, err := m.service.GetAllSongs()
	assert.Equal(m.T(), expectedSongs, song)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorGetAllSongs() {
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongGetAll().Return(nil, expectedErrorPostgres).Times(1)

	playlist, err := m.service.GetAllSongs()
	assert.Nil(m.T(), playlist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetSongById() {
	expectedSong := &models.Song{
		ID:       1,
		Name:     "Track 1",
		Duration: 10,
	}
	id := uint(1)

	m.mockPG.EXPECT().SongGetByID(id).Return(expectedSong, nil).Times(1)

	song, err := m.service.GetSongByID(id)
	assert.Equal(m.T(), expectedSong, song)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorGetSongById() {
	id := uint(1)
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongGetByID(id).Return(nil, expectedErrorPostgres).Times(1)

	playlist, err := m.service.GetSongByID(id)
	assert.Nil(m.T(), playlist)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetNextSong() {
	playlist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	songId := uint(1)
	songs := []models.Song{
		{ID: 1, Name: "Track 1", Duration: 30},
		{ID: 2, Name: "Track 2", Duration: 20},
		{ID: 3, Name: "Track 3", Duration: 10},
	}

	expectedSong := &models.Song{
		ID: 2, Name: "Track 2", Duration: 20,
	}

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(songs, nil).Times(1)

	song, err := m.service.GetNextSong(playlist, songId)
	assert.NoError(m.T(), err)
	assert.Equal(m.T(), expectedSong, song)
}

func (m *MusicSuite) TestMusic_GetNextSongError() {
	expectedErrorPostgres := errors.New("something with postgres")
	playlist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	songId := uint(2)

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(nil, expectedErrorPostgres).Times(1)

	_, err := m.service.GetNextSong(playlist, songId)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_GetPrevSong() {
	playlist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	songId := uint(2)
	songs := []models.Song{
		{ID: 1, Name: "Track 1", Duration: 30},
		{ID: 2, Name: "Track 2", Duration: 20},
		{ID: 3, Name: "Track 3", Duration: 10},
	}

	expectedSong := &models.Song{
		ID: 1, Name: "Track 1", Duration: 30,
	}

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(songs, nil).Times(1)

	song, err := m.service.GetPrevSong(playlist, songId)
	assert.NoError(m.T(), err)
	assert.Equal(m.T(), expectedSong, song)
}

func (m *MusicSuite) TestMusic_GetPrevSongError() {
	expectedErrorPostgres := errors.New("something with postgres")
	playlist := &models.Playlist{
		ID:   1,
		Name: "Playlist",
	}
	songId := uint(2)

	m.mockPG.EXPECT().SongsGetByPlaylistId(playlist).Return(nil, expectedErrorPostgres).Times(1)

	_, err := m.service.GetPrevSong(playlist, songId)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_CreateSong() {
	newSong := &models.Song{Name: "Rap", Duration: 1}

	m.mockPG.EXPECT().SongCreate(newSong).Return(nil).Times(1)

	err := m.service.CreateSong(newSong)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorCreateSong() {
	newSong := &models.Song{Name: "Rap", Duration: 1}
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongCreate(newSong).Return(expectedErrorPostgres).Times(1)

	err := m.service.CreateSong(newSong)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_UpdateSong() {
	song := &models.Song{ID: 1, Name: "Song", Duration: 1}

	m.mockPG.EXPECT().SongUpdate(song).Return(nil).Times(1)

	err := m.service.UpdateSong(song)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorUpdateSong() {
	song := &models.Song{ID: 1, Name: "Song", Duration: 1}
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongUpdate(song).Return(expectedErrorPostgres).Times(1)

	err := m.service.UpdateSong(song)
	assert.Error(m.T(), err)
}

func (m *MusicSuite) TestMusic_DeleteSong() {
	id := uint(1)

	m.mockPG.EXPECT().SongDelete(id).Return(nil).Times(1)

	err := m.service.DeleteSong(id)
	assert.NoError(m.T(), err)
}

func (m *MusicSuite) TestMusic_ErrorDeleteSong() {
	id := uint(1)
	expectedErrorPostgres := errors.New("something with postgres")

	m.mockPG.EXPECT().SongDelete(id).Return(expectedErrorPostgres).Times(1)

	err := m.service.DeleteSong(id)
	assert.Error(m.T(), err)
}

func TestMusicSuite(t *testing.T) {
	suite.Run(t, &MusicSuite{})
}
