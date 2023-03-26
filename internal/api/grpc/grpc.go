package grpc

import (
	"context"
	"gitlab.com/sb-cloud/player-ms-api/internal/music"
	"google.golang.org/protobuf/types/known/durationpb"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "gitlab.com/sb-cloud/player-ms-api/pb/music"

	"gitlab.com/sb-cloud/player-ms-api/internal/models"
)

func New(music *music.Service) Server {
	return Server{
		music: music,
	}
}

type Server struct {
	pb.UnimplementedMusicServiceServer
	music  *music.Service
	server *grpc.Server
}

func (g *Server) Shutdown() {
	g.server.Stop()
}

func (g *Server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	g.server = grpc.NewServer()

	pb.RegisterMusicServiceServer(g.server, g)

	return g.server.Serve(lis)
}

func (g *Server) PlaylistGetAll(ctx context.Context, _ *emptypb.Empty) (*pb.PlaylistAll, error) {
	playlists, err := g.music.GetAllPlaylists()

	if err != nil {
		return nil, err
	}
	data := make([]*pb.Playlist, len(playlists))

	return &pb.PlaylistAll{Data: data}, nil

}

// Получает плейлист по Id
func (g *Server) PlaylistGetById(ctx context.Context, req *pb.GetByIdPlaylistRequest) (*pb.Playlist, error) {

	pl, err := g.music.GetPlaylistByID(uint(req.ID))

	if err != nil {
		return nil, err
	}

	return &pb.Playlist{
		ID:        uint64(pl.ID),
		Name:      pl.Name,
		CreatedAt: pl.CreatedAt.UnixNano(),
		UpdatedAt: pl.UpdatedAt.UnixNano(),
	}, err
}

// Создает плейлист
func (g *Server) PlaylistCreate(ctx context.Context, req *pb.CreatePlaylistRequest) (*emptypb.Empty, error) {
	pl := models.Playlist{
		Name:      req.Name,
		CreatedAt: time.Unix(0, req.CreatedAt),
		UpdatedAt: time.Unix(0, req.UpdatedAt),
	}

	if err := g.music.CreatePlaylist(&pl); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Обновляет плейлист
func (g *Server) PlaylistUpdate(ctx context.Context, req *pb.UpdatePlaylistRequest) (*pb.Playlist, error) {
	pl := &models.Playlist{
		ID:        uint(req.ID),
		Name:      req.Name,
		CreatedAt: time.Unix(0, req.CreatedAt),
		UpdatedAt: time.Unix(0, req.UpdatedAt),
	}
	err := g.music.UpdatePlaylist(pl)
	if err != nil {
		return nil, err
	}

	return &pb.Playlist{
		ID:        uint64(pl.ID),
		Name:      pl.Name,
		CreatedAt: pl.CreatedAt.UnixNano(),
		UpdatedAt: pl.UpdatedAt.UnixNano(),
	}, err
}

// Добавляет песню в плейлист
func (g *Server) PlaylistAddSong(ctx context.Context, req *pb.PlaylistAddSongRequest) (*emptypb.Empty, error) {
	pl, err := g.music.GetPlaylistByID(uint(req.PlaylistID))

	if err != nil {
		return nil, err
	}

	sng, err := g.music.GetSongByID(uint(req.SongID))

	if err != nil {
		return nil, err
	}

	err = g.music.AddSong(pl, sng)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Получает список песен плейлиста
func (g *Server) PlaylistGetSongs(ctx context.Context, req *pb.UpdatePlaylistRequest) (*pb.SongAll, error) {
	pl, err := g.music.GetPlaylistByID(uint(req.ID))

	if err != nil {
		return nil, err
	}

	sng, err := g.music.GetPlaylistSongs(pl)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.Song, len(sng))

	return &pb.SongAll{Data: data}, nil
}

// Удаляет плейлист
func (g *Server) PlaylistDelete(ctx context.Context, req *pb.DeletePlaylistRequest) (*emptypb.Empty, error) {
	err := g.music.DeletePlaylist(uint(req.ID))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

}

// Song

func (g *Server) SongGetAll(ctx context.Context, _ *emptypb.Empty) (*pb.SongAll, error) {
	playlists, err := g.music.GetAllSongs()

	if err != nil {
		return nil, err
	}
	data := make([]*pb.Song, len(playlists))

	return &pb.SongAll{Data: data}, nil

}

// Получает песню по Id
func (g *Server) SongGetById(ctx context.Context, req *pb.GetByIdSongRequest) (*pb.Song, error) {

	pl, err := g.music.GetSongByID(uint(req.ID))

	if err != nil {
		return nil, err
	}

	return &pb.Song{
		ID:          uint64(pl.ID),
		Name:        pl.Name,
		Duration:    durationpb.New(pl.Duration),
		Description: pl.Description,
		CreatedAt:   pl.CreatedAt.UnixNano(),
		UpdatedAt:   pl.UpdatedAt.UnixNano(),
	}, err
}

// Создает песню
func (g *Server) SongCreate(ctx context.Context, req *pb.CreateSongRequest) (*emptypb.Empty, error) {
	pl := models.Song{
		Name:        req.Name,
		Duration:    req.Duration.AsDuration(),
		Description: req.Description,
		CreatedAt:   time.Unix(0, req.CreatedAt),
		UpdatedAt:   time.Unix(0, req.UpdatedAt),
	}

	if err := g.music.CreateSong(&pl); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Обновляет плейлист
func (g *Server) SongUpdate(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Song, error) {
	pl := &models.Song{
		ID:          uint(req.ID),
		Name:        req.Name,
		Duration:    req.Duration.AsDuration(),
		Description: req.Description,
		CreatedAt:   time.Unix(0, req.CreatedAt),
		UpdatedAt:   time.Unix(0, req.UpdatedAt),
	}
	err := g.music.UpdateSong(pl)
	if err != nil {
		return nil, err
	}

	return &pb.Song{
		ID:          uint64(pl.ID),
		Duration:    durationpb.New(pl.Duration),
		Description: pl.Description,
		Name:        pl.Name,
		CreatedAt:   pl.CreatedAt.UnixNano(),
		UpdatedAt:   pl.UpdatedAt.UnixNano(),
	}, err
}

// Получает следующую песню
func (g *Server) SongGetNext(ctx context.Context, req *pb.SongGetNextRequest) (*pb.Song, error) {
	pl, err := g.music.GetPlaylistByID(uint(req.PlaylistID))
	if err != nil {
		return nil, err
	}

	sng, err := g.music.GetNextSong(pl, uint(req.SongID))
	if err != nil {
		return nil, err
	}

	return &pb.Song{
		ID:          uint64(sng.ID),
		Duration:    durationpb.New(sng.Duration),
		Description: sng.Description,
		Name:        sng.Name,
		CreatedAt:   sng.CreatedAt.UnixNano(),
		UpdatedAt:   sng.UpdatedAt.UnixNano(),
	}, err
}

// Получает предыдущую песню
func (g *Server) SongGetPrev(ctx context.Context, req *pb.SongGetPrevRequest) (*pb.Song, error) {
	pl, err := g.music.GetPlaylistByID(uint(req.PlaylistID))
	if err != nil {
		return nil, err
	}

	sng, err := g.music.GetPrevSong(pl, uint(req.SongID))
	if err != nil {
		return nil, err
	}

	return &pb.Song{
		ID:          uint64(sng.ID),
		Duration:    durationpb.New(sng.Duration),
		Name:        sng.Name,
		Description: sng.Description,
		CreatedAt:   sng.CreatedAt.UnixNano(),
		UpdatedAt:   sng.UpdatedAt.UnixNano(),
	}, err
}

// Удаляет песню
func (g *Server) SongDelete(ctx context.Context, req *pb.DeleteSongRequest) (*emptypb.Empty, error) {
	err := g.music.DeleteSong(uint(req.ID))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
