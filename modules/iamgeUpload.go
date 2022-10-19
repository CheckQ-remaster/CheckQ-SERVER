package modules

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

func ImageUpload(file *multipart.File, header *multipart.FileHeader) (string, error) {

	fileName := header.Filename

	osFile, err := os.Create("./image/" + fileName)
	if err != nil {
		log.Println("image os create error")
		return "", err
	}

	write, err := io.Copy(osFile, *file)
	if write == 0 || err != nil {
		log.Println("image os copy error")
		return "", err
	}

	fileName = "/Users/sunghan/Documents/GitHub/CheckQ-SERVER/image/" + fileName

	return fileName, nil
}
