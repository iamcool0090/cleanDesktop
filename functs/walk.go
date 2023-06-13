package functs

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func WalkDirectory(searchDirectory, root string) error {

	// Open the directory
	directory, err := os.Open(searchDirectory)
	if err != nil {
		log.Printf("Error opening directory %s: %s\n", searchDirectory, err)
		return err
	}
	defer directory.Close()

	// Read the directory contents
	fileInfos, err := directory.Readdir(-1)
	if err != nil {
		log.Printf("Error reading directory %s: %s\n", searchDirectory, err)
		return err
	}

	// Process each file in the directory
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			// Get the file extension
			ext := filepath.Ext(fileInfo.Name())
			if ext != ".ink" && fileInfo.Name() != "go-organizer.exe" {

				destinationFolderName := strings.ToUpper(ext[1:]) + "s"
				destinationFilePath := path.Join(root, destinationFolderName, fileInfo.Name())
				sourceFilePath := path.Join(searchDirectory, fileInfo.Name())
				log.Printf("Root: %s, Destination Folder Name: %s", root, destinationFolderName)
				
				err := os.MkdirAll(path.Join(root, destinationFolderName), 0755)
				if err != nil {
					log.Println("Problem creating folder: ", err)
					continue
				}
				
				log.Printf("Source File Path: %s", sourceFilePath)
				log.Printf("Moving to directory: %s", destinationFilePath)
				
				err = MoveFilesToRespectiveDirectories(sourceFilePath, destinationFilePath)
				if err != nil {
					log.Println("Error moving file: ", err)
				} else {
					log.Println("File moved successfully")
				}
			}
		}
	}

	return nil
}
