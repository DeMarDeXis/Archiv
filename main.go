package main

import "Archivator_Go/cmd"

func main() {
	// archiver pack vlc <path to file> -out /path/to/packed-file
	cmd.Execute()
}

// bugs:
// Перезаписывает в себя vlc:
// ./archivator_go unpack vlc ejemplo.vlc
// ./archivator_go pack vlc ejemplo.txt
//
