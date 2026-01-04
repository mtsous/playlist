package http

import (
	"log"
	"net/http"
)

func (s *Server) registerPlaylistRoutes(r *http.ServeMux) *http.ServeMux {
	r.HandleFunc("/", s.SearchPlaylist)

	return r
}

func (s *Server) SearchPlaylist(w http.ResponseWriter, r *http.Request) {
	response := []byte("search handler")

	if _, err := w.Write(response); err != nil {
		log.Print("err")
	}
}
