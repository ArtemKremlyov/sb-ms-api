// Code generated by mockery v2.21.6. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"
	models "gitlab.com/sb-cloud/player-ms-api/internal/models"
)

// PostgresImpl is an autogenerated mock type for the PostgresImpl type
type PostgresImpl struct {
	mock.Mock
}

type PostgresImpl_Expecter struct {
	mock *mock.Mock
}

func (_m *PostgresImpl) EXPECT() *PostgresImpl_Expecter {
	return &PostgresImpl_Expecter{mock: &_m.Mock}
}

// PlaylistAddSong provides a mock function with given fields: playlist, song
func (_m *PostgresImpl) PlaylistAddSong(playlist *models.Playlist, song *models.Song) error {
	ret := _m.Called(playlist, song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist, *models.Song) error); ok {
		r0 = rf(playlist, song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_PlaylistAddSong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistAddSong'
type PostgresImpl_PlaylistAddSong_Call struct {
	*mock.Call
}

// PlaylistAddSong is a helper method to define mock.On call
//   - playlist *models.Playlist
//   - song *models.Song
func (_e *PostgresImpl_Expecter) PlaylistAddSong(playlist interface{}, song interface{}) *PostgresImpl_PlaylistAddSong_Call {
	return &PostgresImpl_PlaylistAddSong_Call{Call: _e.mock.On("PlaylistAddSong", playlist, song)}
}

func (_c *PostgresImpl_PlaylistAddSong_Call) Run(run func(playlist *models.Playlist, song *models.Song)) *PostgresImpl_PlaylistAddSong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist), args[1].(*models.Song))
	})
	return _c
}

func (_c *PostgresImpl_PlaylistAddSong_Call) Return(_a0 error) *PostgresImpl_PlaylistAddSong_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_PlaylistAddSong_Call) RunAndReturn(run func(*models.Playlist, *models.Song) error) *PostgresImpl_PlaylistAddSong_Call {
	_c.Call.Return(run)
	return _c
}

// PlaylistCreate provides a mock function with given fields: playlist
func (_m *PostgresImpl) PlaylistCreate(playlist *models.Playlist) error {
	ret := _m.Called(playlist)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist) error); ok {
		r0 = rf(playlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_PlaylistCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistCreate'
type PostgresImpl_PlaylistCreate_Call struct {
	*mock.Call
}

// PlaylistCreate is a helper method to define mock.On call
//   - playlist *models.Playlist
func (_e *PostgresImpl_Expecter) PlaylistCreate(playlist interface{}) *PostgresImpl_PlaylistCreate_Call {
	return &PostgresImpl_PlaylistCreate_Call{Call: _e.mock.On("PlaylistCreate", playlist)}
}

func (_c *PostgresImpl_PlaylistCreate_Call) Run(run func(playlist *models.Playlist)) *PostgresImpl_PlaylistCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist))
	})
	return _c
}

func (_c *PostgresImpl_PlaylistCreate_Call) Return(_a0 error) *PostgresImpl_PlaylistCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_PlaylistCreate_Call) RunAndReturn(run func(*models.Playlist) error) *PostgresImpl_PlaylistCreate_Call {
	_c.Call.Return(run)
	return _c
}

// PlaylistDelete provides a mock function with given fields: id
func (_m *PostgresImpl) PlaylistDelete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_PlaylistDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistDelete'
type PostgresImpl_PlaylistDelete_Call struct {
	*mock.Call
}

// PlaylistDelete is a helper method to define mock.On call
//   - id uint
func (_e *PostgresImpl_Expecter) PlaylistDelete(id interface{}) *PostgresImpl_PlaylistDelete_Call {
	return &PostgresImpl_PlaylistDelete_Call{Call: _e.mock.On("PlaylistDelete", id)}
}

func (_c *PostgresImpl_PlaylistDelete_Call) Run(run func(id uint)) *PostgresImpl_PlaylistDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *PostgresImpl_PlaylistDelete_Call) Return(_a0 error) *PostgresImpl_PlaylistDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_PlaylistDelete_Call) RunAndReturn(run func(uint) error) *PostgresImpl_PlaylistDelete_Call {
	_c.Call.Return(run)
	return _c
}

// PlaylistGetAll provides a mock function with given fields:
func (_m *PostgresImpl) PlaylistGetAll() ([]models.Playlist, error) {
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

// PostgresImpl_PlaylistGetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistGetAll'
type PostgresImpl_PlaylistGetAll_Call struct {
	*mock.Call
}

// PlaylistGetAll is a helper method to define mock.On call
func (_e *PostgresImpl_Expecter) PlaylistGetAll() *PostgresImpl_PlaylistGetAll_Call {
	return &PostgresImpl_PlaylistGetAll_Call{Call: _e.mock.On("PlaylistGetAll")}
}

func (_c *PostgresImpl_PlaylistGetAll_Call) Run(run func()) *PostgresImpl_PlaylistGetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PostgresImpl_PlaylistGetAll_Call) Return(_a0 []models.Playlist, _a1 error) *PostgresImpl_PlaylistGetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostgresImpl_PlaylistGetAll_Call) RunAndReturn(run func() ([]models.Playlist, error)) *PostgresImpl_PlaylistGetAll_Call {
	_c.Call.Return(run)
	return _c
}

// PlaylistGetByID provides a mock function with given fields: id
func (_m *PostgresImpl) PlaylistGetByID(id uint) (*models.Playlist, error) {
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

// PostgresImpl_PlaylistGetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistGetByID'
type PostgresImpl_PlaylistGetByID_Call struct {
	*mock.Call
}

// PlaylistGetByID is a helper method to define mock.On call
//   - id uint
func (_e *PostgresImpl_Expecter) PlaylistGetByID(id interface{}) *PostgresImpl_PlaylistGetByID_Call {
	return &PostgresImpl_PlaylistGetByID_Call{Call: _e.mock.On("PlaylistGetByID", id)}
}

func (_c *PostgresImpl_PlaylistGetByID_Call) Run(run func(id uint)) *PostgresImpl_PlaylistGetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *PostgresImpl_PlaylistGetByID_Call) Return(_a0 *models.Playlist, _a1 error) *PostgresImpl_PlaylistGetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostgresImpl_PlaylistGetByID_Call) RunAndReturn(run func(uint) (*models.Playlist, error)) *PostgresImpl_PlaylistGetByID_Call {
	_c.Call.Return(run)
	return _c
}

// PlaylistUpdate provides a mock function with given fields: playlist
func (_m *PostgresImpl) PlaylistUpdate(playlist *models.Playlist) error {
	ret := _m.Called(playlist)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Playlist) error); ok {
		r0 = rf(playlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_PlaylistUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaylistUpdate'
type PostgresImpl_PlaylistUpdate_Call struct {
	*mock.Call
}

// PlaylistUpdate is a helper method to define mock.On call
//   - playlist *models.Playlist
func (_e *PostgresImpl_Expecter) PlaylistUpdate(playlist interface{}) *PostgresImpl_PlaylistUpdate_Call {
	return &PostgresImpl_PlaylistUpdate_Call{Call: _e.mock.On("PlaylistUpdate", playlist)}
}

func (_c *PostgresImpl_PlaylistUpdate_Call) Run(run func(playlist *models.Playlist)) *PostgresImpl_PlaylistUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Playlist))
	})
	return _c
}

func (_c *PostgresImpl_PlaylistUpdate_Call) Return(_a0 error) *PostgresImpl_PlaylistUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_PlaylistUpdate_Call) RunAndReturn(run func(*models.Playlist) error) *PostgresImpl_PlaylistUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// SongCreate provides a mock function with given fields: song
func (_m *PostgresImpl) SongCreate(song *models.Song) error {
	ret := _m.Called(song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Song) error); ok {
		r0 = rf(song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_SongCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SongCreate'
type PostgresImpl_SongCreate_Call struct {
	*mock.Call
}

// SongCreate is a helper method to define mock.On call
//   - song *models.Song
func (_e *PostgresImpl_Expecter) SongCreate(song interface{}) *PostgresImpl_SongCreate_Call {
	return &PostgresImpl_SongCreate_Call{Call: _e.mock.On("SongCreate", song)}
}

func (_c *PostgresImpl_SongCreate_Call) Run(run func(song *models.Song)) *PostgresImpl_SongCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Song))
	})
	return _c
}

func (_c *PostgresImpl_SongCreate_Call) Return(_a0 error) *PostgresImpl_SongCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_SongCreate_Call) RunAndReturn(run func(*models.Song) error) *PostgresImpl_SongCreate_Call {
	_c.Call.Return(run)
	return _c
}

// SongDelete provides a mock function with given fields: id
func (_m *PostgresImpl) SongDelete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_SongDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SongDelete'
type PostgresImpl_SongDelete_Call struct {
	*mock.Call
}

// SongDelete is a helper method to define mock.On call
//   - id uint
func (_e *PostgresImpl_Expecter) SongDelete(id interface{}) *PostgresImpl_SongDelete_Call {
	return &PostgresImpl_SongDelete_Call{Call: _e.mock.On("SongDelete", id)}
}

func (_c *PostgresImpl_SongDelete_Call) Run(run func(id uint)) *PostgresImpl_SongDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *PostgresImpl_SongDelete_Call) Return(_a0 error) *PostgresImpl_SongDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_SongDelete_Call) RunAndReturn(run func(uint) error) *PostgresImpl_SongDelete_Call {
	_c.Call.Return(run)
	return _c
}

// SongGetAll provides a mock function with given fields:
func (_m *PostgresImpl) SongGetAll() ([]models.Song, error) {
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

// PostgresImpl_SongGetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SongGetAll'
type PostgresImpl_SongGetAll_Call struct {
	*mock.Call
}

// SongGetAll is a helper method to define mock.On call
func (_e *PostgresImpl_Expecter) SongGetAll() *PostgresImpl_SongGetAll_Call {
	return &PostgresImpl_SongGetAll_Call{Call: _e.mock.On("SongGetAll")}
}

func (_c *PostgresImpl_SongGetAll_Call) Run(run func()) *PostgresImpl_SongGetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PostgresImpl_SongGetAll_Call) Return(_a0 []models.Song, _a1 error) *PostgresImpl_SongGetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostgresImpl_SongGetAll_Call) RunAndReturn(run func() ([]models.Song, error)) *PostgresImpl_SongGetAll_Call {
	_c.Call.Return(run)
	return _c
}

// SongGetByID provides a mock function with given fields: id
func (_m *PostgresImpl) SongGetByID(id uint) (*models.Song, error) {
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

// PostgresImpl_SongGetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SongGetByID'
type PostgresImpl_SongGetByID_Call struct {
	*mock.Call
}

// SongGetByID is a helper method to define mock.On call
//   - id uint
func (_e *PostgresImpl_Expecter) SongGetByID(id interface{}) *PostgresImpl_SongGetByID_Call {
	return &PostgresImpl_SongGetByID_Call{Call: _e.mock.On("SongGetByID", id)}
}

func (_c *PostgresImpl_SongGetByID_Call) Run(run func(id uint)) *PostgresImpl_SongGetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *PostgresImpl_SongGetByID_Call) Return(_a0 *models.Song, _a1 error) *PostgresImpl_SongGetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostgresImpl_SongGetByID_Call) RunAndReturn(run func(uint) (*models.Song, error)) *PostgresImpl_SongGetByID_Call {
	_c.Call.Return(run)
	return _c
}

// SongUpdate provides a mock function with given fields: song
func (_m *PostgresImpl) SongUpdate(song *models.Song) error {
	ret := _m.Called(song)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Song) error); ok {
		r0 = rf(song)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostgresImpl_SongUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SongUpdate'
type PostgresImpl_SongUpdate_Call struct {
	*mock.Call
}

// SongUpdate is a helper method to define mock.On call
//   - song *models.Song
func (_e *PostgresImpl_Expecter) SongUpdate(song interface{}) *PostgresImpl_SongUpdate_Call {
	return &PostgresImpl_SongUpdate_Call{Call: _e.mock.On("SongUpdate", song)}
}

func (_c *PostgresImpl_SongUpdate_Call) Run(run func(song *models.Song)) *PostgresImpl_SongUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Song))
	})
	return _c
}

func (_c *PostgresImpl_SongUpdate_Call) Return(_a0 error) *PostgresImpl_SongUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostgresImpl_SongUpdate_Call) RunAndReturn(run func(*models.Song) error) *PostgresImpl_SongUpdate_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPostgresImpl interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostgresImpl creates a new instance of PostgresImpl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostgresImpl(t mockConstructorTestingTNewPostgresImpl) *PostgresImpl {
	mock := &PostgresImpl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
