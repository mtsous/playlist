package main

import "github.com/mtsous/playlist/http"

func main() {
	server := http.NewServer()

	server.Open()
}
