// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: music.proto

package player_ms_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MusicService_GetAllPlaylist_FullMethodName  = "/articles.MusicService/GetAllPlaylist"
	MusicService_GetByIdPlaylist_FullMethodName = "/articles.MusicService/GetByIdPlaylist"
	MusicService_CreatePlaylist_FullMethodName  = "/articles.MusicService/CreatePlaylist"
	MusicService_UpdatePlaylist_FullMethodName  = "/articles.MusicService/UpdatePlaylist"
	MusicService_DeletePlaylist_FullMethodName  = "/articles.MusicService/DeletePlaylist"
	MusicService_GetAllSong_FullMethodName      = "/articles.MusicService/GetAllSong"
	MusicService_GetByIdSong_FullMethodName     = "/articles.MusicService/GetByIdSong"
	MusicService_CreateSong_FullMethodName      = "/articles.MusicService/CreateSong"
	MusicService_UpdateSong_FullMethodName      = "/articles.MusicService/UpdateSong"
	MusicService_DeleteSong_FullMethodName      = "/articles.MusicService/DeleteSong"
)

// MusicServiceClient is the client API for MusicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicServiceClient interface {
	// Возвращает список плейлистов
	GetAllPlaylist(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlaylistAll, error)
	// Получает плейлист по Id
	GetByIdPlaylist(ctx context.Context, in *GetByIdPlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	// Создает плейлист
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Обновляет плейлист
	UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Возвращает список песен
	GetAllSong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SongAll, error)
	// Получает песню по Id
	GetByIdSong(ctx context.Context, in *GetByIdSongRequest, opts ...grpc.CallOption) (*Song, error)
	// Создает песню
	CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Обновляет песню
	UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Song, error)
	// Удаляет
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type musicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicServiceClient(cc grpc.ClientConnInterface) MusicServiceClient {
	return &musicServiceClient{cc}
}

func (c *musicServiceClient) GetAllPlaylist(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlaylistAll, error) {
	out := new(PlaylistAll)
	err := c.cc.Invoke(ctx, MusicService_GetAllPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetByIdPlaylist(ctx context.Context, in *GetByIdPlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_GetByIdPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MusicService_CreatePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_UpdatePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MusicService_DeletePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetAllSong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SongAll, error) {
	out := new(SongAll)
	err := c.cc.Invoke(ctx, MusicService_GetAllSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetByIdSong(ctx context.Context, in *GetByIdSongRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_GetByIdSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MusicService_CreateSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_UpdateSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MusicService_DeleteSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicServiceServer is the server API for MusicService service.
// All implementations must embed UnimplementedMusicServiceServer
// for forward compatibility
type MusicServiceServer interface {
	// Возвращает список плейлистов
	GetAllPlaylist(context.Context, *emptypb.Empty) (*PlaylistAll, error)
	// Получает плейлист по Id
	GetByIdPlaylist(context.Context, *GetByIdPlaylistRequest) (*Playlist, error)
	// Создает плейлист
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*emptypb.Empty, error)
	// Обновляет плейлист
	UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error)
	DeletePlaylist(context.Context, *DeletePlaylistRequest) (*emptypb.Empty, error)
	// Возвращает список песен
	GetAllSong(context.Context, *emptypb.Empty) (*SongAll, error)
	// Получает песню по Id
	GetByIdSong(context.Context, *GetByIdSongRequest) (*Song, error)
	// Создает песню
	CreateSong(context.Context, *CreateSongRequest) (*emptypb.Empty, error)
	// Обновляет песню
	UpdateSong(context.Context, *UpdateSongRequest) (*Song, error)
	// Удаляет
	DeleteSong(context.Context, *DeleteSongRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMusicServiceServer()
}

// UnimplementedMusicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMusicServiceServer struct {
}

func (UnimplementedMusicServiceServer) GetAllPlaylist(context.Context, *emptypb.Empty) (*PlaylistAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) GetByIdPlaylist(context.Context, *GetByIdPlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) DeletePlaylist(context.Context, *DeletePlaylistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) GetAllSong(context.Context, *emptypb.Empty) (*SongAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSong not implemented")
}
func (UnimplementedMusicServiceServer) GetByIdSong(context.Context, *GetByIdSongRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdSong not implemented")
}
func (UnimplementedMusicServiceServer) CreateSong(context.Context, *CreateSongRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSong not implemented")
}
func (UnimplementedMusicServiceServer) UpdateSong(context.Context, *UpdateSongRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedMusicServiceServer) DeleteSong(context.Context, *DeleteSongRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedMusicServiceServer) mustEmbedUnimplementedMusicServiceServer() {}

// UnsafeMusicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicServiceServer will
// result in compilation errors.
type UnsafeMusicServiceServer interface {
	mustEmbedUnimplementedMusicServiceServer()
}

func RegisterMusicServiceServer(s grpc.ServiceRegistrar, srv MusicServiceServer) {
	s.RegisterService(&MusicService_ServiceDesc, srv)
}

func _MusicService_GetAllPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetAllPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetAllPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetAllPlaylist(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetByIdPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetByIdPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetByIdPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetByIdPlaylist(ctx, req.(*GetByIdPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_CreatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, req.(*UpdatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeletePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, req.(*DeletePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetAllSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetAllSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetAllSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetAllSong(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetByIdSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetByIdSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetByIdSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetByIdSong(ctx, req.(*GetByIdSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_CreateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).CreateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_CreateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).CreateSong(ctx, req.(*CreateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdateSong(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicService_ServiceDesc is the grpc.ServiceDesc for MusicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "articles.MusicService",
	HandlerType: (*MusicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPlaylist",
			Handler:    _MusicService_GetAllPlaylist_Handler,
		},
		{
			MethodName: "GetByIdPlaylist",
			Handler:    _MusicService_GetByIdPlaylist_Handler,
		},
		{
			MethodName: "CreatePlaylist",
			Handler:    _MusicService_CreatePlaylist_Handler,
		},
		{
			MethodName: "UpdatePlaylist",
			Handler:    _MusicService_UpdatePlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _MusicService_DeletePlaylist_Handler,
		},
		{
			MethodName: "GetAllSong",
			Handler:    _MusicService_GetAllSong_Handler,
		},
		{
			MethodName: "GetByIdSong",
			Handler:    _MusicService_GetByIdSong_Handler,
		},
		{
			MethodName: "CreateSong",
			Handler:    _MusicService_CreateSong_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _MusicService_UpdateSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _MusicService_DeleteSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "music.proto",
}
