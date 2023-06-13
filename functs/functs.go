package functs

import (
	"log"
	"os"
)

func MoveFilesToRespectiveDirectories(sourcePath, destinationPath string) error {
	log.Printf("Moving file: %s to destination: %s", sourcePath, destinationPath)

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		log.Printf("Failed to open source file: %s", err)
		return err
	}
	sourceFile.Close()

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		log.Printf("Failed to move file: %s", err)
		return err
	}

	log.Println("File moved successfully")
	return nil
}
