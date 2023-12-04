package main

import "fmt"

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  int
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func PlaylistsComMusica(titulo string, playlists []Playlist) []Playlist {
	var playlistsComMusica []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				playlistsComMusica = append(playlistsComMusica, playlist)
				break
			}
		}
	}

	return playlistsComMusica
}

func main() {
	playlist1 := Playlist{
		Nome: "Playlist A",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista A", Duracao: 180},
			{Titulo: "Música 2", Artista: "Artista B", Duracao: 240},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist B",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista A", Duracao: 180},
			{Titulo: "Música 3", Artista: "Artista C", Duracao: 200},
		},
	}

	playlists := []Playlist{playlist1, playlist2}

	tituloBuscado := "Música 1"
	playlistsEncontradas := PlaylistsComMusica(tituloBuscado, playlists)

	fmt.Printf("Playlists com a música '%s':\n", tituloBuscado)
	for _, playlist := range playlistsEncontradas {
		fmt.Printf("%s\n", playlist.Nome)
	}
}
