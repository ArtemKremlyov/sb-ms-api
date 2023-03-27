package rest

type SongChangeTrackRequest struct {
	PlaylistId uint `json:"playlist_id"`
}

type SongChangePlayingRequest struct {
	Playing bool `json:"playing"`
}
