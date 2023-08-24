package main

import (
	"cleanDesktop/functs"
	"flag"
	"fmt"
	"log"
)

func main() {
	var searchDirectory string
	var rootFileName string

	// Parse command-line arguments
	flag.StringVar(&searchDirectory, "dir", ".", "Search directory")
	flag.StringVar(&rootFileName, "root", "organized_files", "Root file name (name of the folder to store segregated files)")
	flag.Parse()

	log.Printf("Search Directory: %s", searchDirectory)
	log.Printf("Root File Name: %s", rootFileName)

	err := functs.WalkDirectory(searchDirectory, rootFileName)
	if err != nil {
		
		log.Printf("Error: %s", err)
		fmt.Println("An error occurred while organizing files.")
	} else {
		fmt.Println("Files organized successfully.")
	}
}
