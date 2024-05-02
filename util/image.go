package util

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

func StoreImageToLocal(image *multipart.FileHeader, owner string) string {
	file, err := image.Open()
	PanicIfError(err)

	tempFile, err := os.CreateTemp("public", "image-"+GenerateImageName(owner)+"*.png")
	PanicIfError(err)
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	PanicIfError(err)

	_, err = tempFile.Write(fileBytes)
	PanicIfError(err)

	fileName := tempFile.Name()
	newFileName := strings.Split(fileName, "\\")

	return newFileName[1]
}

func GenerateImageName(nameUser string) string {
	currentTime := time.Now().UTC().Format("20060102150405.000000000")
	newCurrentTime := strings.ReplaceAll(currentTime, ".", "")
	return strings.ToLower(nameUser + "_" + newCurrentTime)
}

func GetImageUrl(fileName string) string {
	return "http://localhost:3000/image/" + fileName
}
