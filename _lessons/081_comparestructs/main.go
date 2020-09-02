package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string

	// songTitles  []string
	// songArtists []string
	songs []song
}

func main() {
	// song1 := song{title: "wonderwall", artist: "oasis"}
	// song2 := song{title: "super sonic", artist: "oasis"}

	rock := playlist{
		genre: "indie rock",
		songs: []song{
			song{title: "wonderwall", artist: "oasis"},
			song{title: "super sonic", artist: "oasis"},
		},
	}

	// clone:=rock
	// if rock == clone {} // it contains a field of incomparable type

	song := rock.songs[0] // it's a clone
	song.title = "live forever"

	rock.songs[0].title = "live forever"

	fmt.Printf("%+v\n%+v\n\n", song, rock.songs[0])

	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs { // s is a clone
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	// song1 = song2

	// fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// // if song1.title == song2.title && song1.artist == song2.artist {
	// if song1 == song2 {
	// 	fmt.Println("songs are equal")
	// } else {
	// 	fmt.Println("songs are not equal")
	// }
}
