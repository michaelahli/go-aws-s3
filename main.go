package main

import (
	"awss3/utils"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	s3 := utils.NewS3Session()
	s3.UploadObjectbyFilePath("./files/test 1.png")
	log.Println(s3.GetURI())
}
