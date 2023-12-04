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

func ImprimirInfoPlaylist(p Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", p.Nome)

	totalDuracao := 0
	for _, musica := range p.Musicas {
		totalDuracao += musica.Duracao
		fmt.Printf("Música: %s\n", musica.Titulo)
		fmt.Printf("Artista: %s\n", musica.Artista)
		fmt.Printf("Duração: %d segundos\n", musica.Duracao)
		fmt.Println("--------------------")
	}

	fmt.Printf("Tempo total da Playlist: %d segundos\n", totalDuracao)
}

func main() {
	playlistExemplo := Playlist{
		Nome: "Minha Playlist",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista A", Duracao: 180},
			{Titulo: "Música 2", Artista: "Artista B", Duracao: 240},
			{Titulo: "Música 3", Artista: "Artista C", Duracao: 200},
		},
	}

	ImprimirInfoPlaylist(playlistExemplo)
}
