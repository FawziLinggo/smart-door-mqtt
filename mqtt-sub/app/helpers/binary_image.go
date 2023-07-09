package helpers

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func ImageToBase64(filePath string) (string, error) {
	imageBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(imageBytes)

	return base64String, nil
}

func Base64ToImage(base64String string) error {
	// generate timestamp for image name format 20060102150405 UTC + 7
	outputPath := os.Getenv("PATH_IMAGE")
	timestamp := time.Now().UTC().Add(time.Hour * 7).Format("20060102150405")
	imageName := "image_" + timestamp + ".png"
	outputPath = outputPath + imageName

	imageBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outputPath, imageBytes, 0644)
	if err != nil {
		return err
	}

	err = SavePathToTable(imageName)
	if err != nil {
		return err
	}
	log.Println("Image saved with name: ", imageName)

	return nil
}
