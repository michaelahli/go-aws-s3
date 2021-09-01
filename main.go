package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/michaelahli/go-aws-s3/utils"
)

func init() {
	godotenv.Load()
}

func main() {
	s3 := utils.NewS3Session()
	s3.UploadObjectbyFilePath("./files/test 1.png")
	log.Println(s3.GetURI())
}
