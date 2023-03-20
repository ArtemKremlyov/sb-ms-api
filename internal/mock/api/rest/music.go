// Code generated by mockery v2.21.6. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"
	models "gitlab.com/sb-cloud/player-ms-api/internal/models"
)

// Music is an autogenerated mock type for the Music type
type Music struct {
	mock.Mock
}

type Music_Expecter struct {
	mock *mock.Mock
}

func (_m *Music) EXPECT() *Music_Expecter {
	return &Music_Expecter{mock: &_m.Mock}
}

// AddSong provides a mock function with given fields: playlist, song
func (_m *Music) AddSong(playlist *models.Playlist, song *models.Song) error {
	ret := _m.Called(playlist, song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist, *models.Song) error); ok {
		r0 = rf(playlist, song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_AddSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSong'
type Music_AddSong_Call struct {
	*mock.Call
}

// AddSong is a helper method to define mock.On call
//   - playlist *models.Playlist
//   - song *models.Song
func (_e *Music_Expecter) AddSong(playlist interface{}, song interface{}) *Music_AddSong_Call {
	return &Music_AddSong_Call{Call: _e.mock.On("AddSong", playlist, song)}
}

func (_c *Music_AddSong_Call) Run(run func(playlist *models.Playlist, song *models.Song)) *Music_AddSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist), args[1].(*models.Song))
	})
	return _c
}

func (_c *Music_AddSong_Call) Return(_a0 error) *Music_AddSong_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_AddSong_Call) RunAndReturn(run func(*models.Playlist, *models.Song) error) *Music_AddSong_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePlaylist provides a mock function with given fields: playlist
func (_m *Music) CreatePlaylist(playlist *models.Playlist) error {
	ret := _m.Called(playlist)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist) error); ok {
		r0 = rf(playlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_CreatePlaylist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePlaylist'
type Music_CreatePlaylist_Call struct {
	*mock.Call
}

// CreatePlaylist is a helper method to define mock.On call
//   - playlist *models.Playlist
func (_e *Music_Expecter) CreatePlaylist(playlist interface{}) *Music_CreatePlaylist_Call {
	return &Music_CreatePlaylist_Call{Call: _e.mock.On("CreatePlaylist", playlist)}
}

func (_c *Music_CreatePlaylist_Call) Run(run func(playlist *models.Playlist)) *Music_CreatePlaylist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist))
	})
	return _c
}

func (_c *Music_CreatePlaylist_Call) Return(_a0 error) *Music_CreatePlaylist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_CreatePlaylist_Call) RunAndReturn(run func(*models.Playlist) error) *Music_CreatePlaylist_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSong provides a mock function with given fields: song
func (_m *Music) CreateSong(song *models.Song) error {
	ret := _m.Called(song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Song) error); ok {
		r0 = rf(song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_CreateSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSong'
type Music_CreateSong_Call struct {
	*mock.Call
}

// CreateSong is a helper method to define mock.On call
//   - song *models.Song
func (_e *Music_Expecter) CreateSong(song interface{}) *Music_CreateSong_Call {
	return &Music_CreateSong_Call{Call: _e.mock.On("CreateSong", song)}
}

func (_c *Music_CreateSong_Call) Run(run func(song *models.Song)) *Music_CreateSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Song))
	})
	return _c
}

func (_c *Music_CreateSong_Call) Return(_a0 error) *Music_CreateSong_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_CreateSong_Call) RunAndReturn(run func(*models.Song) error) *Music_CreateSong_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePlaylist provides a mock function with given fields: id
func (_m *Music) DeletePlaylist(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_DeletePlaylist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePlaylist'
type Music_DeletePlaylist_Call struct {
	*mock.Call
}

// DeletePlaylist is a helper method to define mock.On call
//   - id uint
func (_e *Music_Expecter) DeletePlaylist(id interface{}) *Music_DeletePlaylist_Call {
	return &Music_DeletePlaylist_Call{Call: _e.mock.On("DeletePlaylist", id)}
}

func (_c *Music_DeletePlaylist_Call) Run(run func(id uint)) *Music_DeletePlaylist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *Music_DeletePlaylist_Call) Return(_a0 error) *Music_DeletePlaylist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_DeletePlaylist_Call) RunAndReturn(run func(uint) error) *Music_DeletePlaylist_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSong provides a mock function with given fields: id
func (_m *Music) DeleteSong(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_DeleteSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSong'
type Music_DeleteSong_Call struct {
	*mock.Call
}

// DeleteSong is a helper method to define mock.On call
//   - id uint
func (_e *Music_Expecter) DeleteSong(id interface{}) *Music_DeleteSong_Call {
	return &Music_DeleteSong_Call{Call: _e.mock.On("DeleteSong", id)}
}

func (_c *Music_DeleteSong_Call) Run(run func(id uint)) *Music_DeleteSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *Music_DeleteSong_Call) Return(_a0 error) *Music_DeleteSong_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_DeleteSong_Call) RunAndReturn(run func(uint) error) *Music_DeleteSong_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPlaylists provides a mock function with given fields:
func (_m *Music) GetAllPlaylists() ([]models.Playlist, error) {
	ret := _m.Called()

	var r0 []models.Playlist
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Playlist, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Playlist); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Playlist)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetAllPlaylists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPlaylists'
type Music_GetAllPlaylists_Call struct {
	*mock.Call
}

// GetAllPlaylists is a helper method to define mock.On call
func (_e *Music_Expecter) GetAllPlaylists() *Music_GetAllPlaylists_Call {
	return &Music_GetAllPlaylists_Call{Call: _e.mock.On("GetAllPlaylists")}
}

func (_c *Music_GetAllPlaylists_Call) Run(run func()) *Music_GetAllPlaylists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Music_GetAllPlaylists_Call) Return(_a0 []models.Playlist, _a1 error) *Music_GetAllPlaylists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetAllPlaylists_Call) RunAndReturn(run func() ([]models.Playlist, error)) *Music_GetAllPlaylists_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllSongs provides a mock function with given fields:
func (_m *Music) GetAllSongs() ([]models.Song, error) {
	ret := _m.Called()

	var r0 []models.Song
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Song, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Song); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Song)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetAllSongs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllSongs'
type Music_GetAllSongs_Call struct {
	*mock.Call
}

// GetAllSongs is a helper method to define mock.On call
func (_e *Music_Expecter) GetAllSongs() *Music_GetAllSongs_Call {
	return &Music_GetAllSongs_Call{Call: _e.mock.On("GetAllSongs")}
}

func (_c *Music_GetAllSongs_Call) Run(run func()) *Music_GetAllSongs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Music_GetAllSongs_Call) Return(_a0 []models.Song, _a1 error) *Music_GetAllSongs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetAllSongs_Call) RunAndReturn(run func() ([]models.Song, error)) *Music_GetAllSongs_Call {
	_c.Call.Return(run)
	return _c
}

// GetNextSong provides a mock function with given fields: playlist, songId
func (_m *Music) GetNextSong(playlist *models.Playlist, songId uint) (*models.Song, error) {
	ret := _m.Called(playlist, songId)

	var r0 *models.Song
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Playlist, uint) (*models.Song, error)); ok {
		return rf(playlist, songId)
	}
	if rf, ok := ret.Get(0).(func(*models.Playlist, uint) *models.Song); ok {
		r0 = rf(playlist, songId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Song)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Playlist, uint) error); ok {
		r1 = rf(playlist, songId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetNextSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNextSong'
type Music_GetNextSong_Call struct {
	*mock.Call
}

// GetNextSong is a helper method to define mock.On call
//   - playlist *models.Playlist
//   - songId uint
func (_e *Music_Expecter) GetNextSong(playlist interface{}, songId interface{}) *Music_GetNextSong_Call {
	return &Music_GetNextSong_Call{Call: _e.mock.On("GetNextSong", playlist, songId)}
}

func (_c *Music_GetNextSong_Call) Run(run func(playlist *models.Playlist, songId uint)) *Music_GetNextSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist), args[1].(uint))
	})
	return _c
}

func (_c *Music_GetNextSong_Call) Return(_a0 *models.Song, _a1 error) *Music_GetNextSong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetNextSong_Call) RunAndReturn(run func(*models.Playlist, uint) (*models.Song, error)) *Music_GetNextSong_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlaylistByID provides a mock function with given fields: id
func (_m *Music) GetPlaylistByID(id uint) (*models.Playlist, error) {
	ret := _m.Called(id)

	var r0 *models.Playlist
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.Playlist, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.Playlist); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Playlist)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetPlaylistByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlaylistByID'
type Music_GetPlaylistByID_Call struct {
	*mock.Call
}

// GetPlaylistByID is a helper method to define mock.On call
//   - id uint
func (_e *Music_Expecter) GetPlaylistByID(id interface{}) *Music_GetPlaylistByID_Call {
	return &Music_GetPlaylistByID_Call{Call: _e.mock.On("GetPlaylistByID", id)}
}

func (_c *Music_GetPlaylistByID_Call) Run(run func(id uint)) *Music_GetPlaylistByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *Music_GetPlaylistByID_Call) Return(_a0 *models.Playlist, _a1 error) *Music_GetPlaylistByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetPlaylistByID_Call) RunAndReturn(run func(uint) (*models.Playlist, error)) *Music_GetPlaylistByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlaylistSongs provides a mock function with given fields: playlist
func (_m *Music) GetPlaylistSongs(playlist *models.Playlist) ([]models.Song, error) {
	ret := _m.Called(playlist)

	var r0 []models.Song
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Playlist) ([]models.Song, error)); ok {
		return rf(playlist)
	}
	if rf, ok := ret.Get(0).(func(*models.Playlist) []models.Song); ok {
		r0 = rf(playlist)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Song)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Playlist) error); ok {
		r1 = rf(playlist)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetPlaylistSongs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlaylistSongs'
type Music_GetPlaylistSongs_Call struct {
	*mock.Call
}

// GetPlaylistSongs is a helper method to define mock.On call
//   - playlist *models.Playlist
func (_e *Music_Expecter) GetPlaylistSongs(playlist interface{}) *Music_GetPlaylistSongs_Call {
	return &Music_GetPlaylistSongs_Call{Call: _e.mock.On("GetPlaylistSongs", playlist)}
}

func (_c *Music_GetPlaylistSongs_Call) Run(run func(playlist *models.Playlist)) *Music_GetPlaylistSongs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist))
	})
	return _c
}

func (_c *Music_GetPlaylistSongs_Call) Return(_a0 []models.Song, _a1 error) *Music_GetPlaylistSongs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetPlaylistSongs_Call) RunAndReturn(run func(*models.Playlist) ([]models.Song, error)) *Music_GetPlaylistSongs_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrevSong provides a mock function with given fields: playlist, songId
func (_m *Music) GetPrevSong(playlist *models.Playlist, songId uint) (*models.Song, error) {
	ret := _m.Called(playlist, songId)

	var r0 *models.Song
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Playlist, uint) (*models.Song, error)); ok {
		return rf(playlist, songId)
	}
	if rf, ok := ret.Get(0).(func(*models.Playlist, uint) *models.Song); ok {
		r0 = rf(playlist, songId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Song)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Playlist, uint) error); ok {
		r1 = rf(playlist, songId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetPrevSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrevSong'
type Music_GetPrevSong_Call struct {
	*mock.Call
}

// GetPrevSong is a helper method to define mock.On call
//   - playlist *models.Playlist
//   - songId uint
func (_e *Music_Expecter) GetPrevSong(playlist interface{}, songId interface{}) *Music_GetPrevSong_Call {
	return &Music_GetPrevSong_Call{Call: _e.mock.On("GetPrevSong", playlist, songId)}
}

func (_c *Music_GetPrevSong_Call) Run(run func(playlist *models.Playlist, songId uint)) *Music_GetPrevSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist), args[1].(uint))
	})
	return _c
}

func (_c *Music_GetPrevSong_Call) Return(_a0 *models.Song, _a1 error) *Music_GetPrevSong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetPrevSong_Call) RunAndReturn(run func(*models.Playlist, uint) (*models.Song, error)) *Music_GetPrevSong_Call {
	_c.Call.Return(run)
	return _c
}

// GetSongByID provides a mock function with given fields: id
func (_m *Music) GetSongByID(id uint) (*models.Song, error) {
	ret := _m.Called(id)

	var r0 *models.Song
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.Song, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.Song); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Song)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Music_GetSongByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSongByID'
type Music_GetSongByID_Call struct {
	*mock.Call
}

// GetSongByID is a helper method to define mock.On call
//   - id uint
func (_e *Music_Expecter) GetSongByID(id interface{}) *Music_GetSongByID_Call {
	return &Music_GetSongByID_Call{Call: _e.mock.On("GetSongByID", id)}
}

func (_c *Music_GetSongByID_Call) Run(run func(id uint)) *Music_GetSongByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *Music_GetSongByID_Call) Return(_a0 *models.Song, _a1 error) *Music_GetSongByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Music_GetSongByID_Call) RunAndReturn(run func(uint) (*models.Song, error)) *Music_GetSongByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePlaylist provides a mock function with given fields: playlist
func (_m *Music) UpdatePlaylist(playlist *models.Playlist) error {
	ret := _m.Called(playlist)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist) error); ok {
		r0 = rf(playlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_UpdatePlaylist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePlaylist'
type Music_UpdatePlaylist_Call struct {
	*mock.Call
}

// UpdatePlaylist is a helper method to define mock.On call
//   - playlist *models.Playlist
func (_e *Music_Expecter) UpdatePlaylist(playlist interface{}) *Music_UpdatePlaylist_Call {
	return &Music_UpdatePlaylist_Call{Call: _e.mock.On("UpdatePlaylist", playlist)}
}

func (_c *Music_UpdatePlaylist_Call) Run(run func(playlist *models.Playlist)) *Music_UpdatePlaylist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist))
	})
	return _c
}

func (_c *Music_UpdatePlaylist_Call) Return(_a0 error) *Music_UpdatePlaylist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_UpdatePlaylist_Call) RunAndReturn(run func(*models.Playlist) error) *Music_UpdatePlaylist_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSong provides a mock function with given fields: song
func (_m *Music) UpdateSong(song *models.Song) error {
	ret := _m.Called(song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Song) error); ok {
		r0 = rf(song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Music_UpdateSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSong'
type Music_UpdateSong_Call struct {
	*mock.Call
}

// UpdateSong is a helper method to define mock.On call
//   - song *models.Song
func (_e *Music_Expecter) UpdateSong(song interface{}) *Music_UpdateSong_Call {
	return &Music_UpdateSong_Call{Call: _e.mock.On("UpdateSong", song)}
}

func (_c *Music_UpdateSong_Call) Run(run func(song *models.Song)) *Music_UpdateSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Song))
	})
	return _c
}

func (_c *Music_UpdateSong_Call) Return(_a0 error) *Music_UpdateSong_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Music_UpdateSong_Call) RunAndReturn(run func(*models.Song) error) *Music_UpdateSong_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMusic interface {
	mock.TestingT
	Cleanup(func())
}

// NewMusic creates a new instance of Music. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMusic(t mockConstructorTestingTNewMusic) *Music {
	mock := &Music{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
