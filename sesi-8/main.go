package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
	"sesi-8/config"
	"time"
)

var path string

func main() {

	err := config.LoadConfig(".env")
	if err != nil {
		log.Println("no file .env")
	}

	appPort := config.GetConfigString("APP_PORT", ":4444")
	path = config.GetConfigString("PUBLIC_PATH", "public/uploads")

	router := fiber.New()

	files := router.Group("/files")
	{
		files.Post("/upload", Upload)
		files.Post("/download", Download)
	}

	log.Println("application is running")
	router.Listen(appPort)
}

func Upload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("error when try to parse FormFile with detail : ", err.Error())
	}

	if file.Size > 1*1024*1024 {
		errBadRequest := fiber.ErrBadRequest
		errBadRequest.Message = "file to big, expected 1 Mb"
		log.Println("error with detail : ", errBadRequest.Error(), "file size :", (file.Size / 1024 / 104), "MB")
		return errBadRequest
	}

	log.Printf("header : %+v\n", file.Header)
	log.Printf("content-type : %+v\n", file.Header.Get("Content-Type"))

	var typeFile string
	typeFile = c.FormValue("type", "")

	log.Println(typeFile)

	if err = os.Mkdir(path+"/"+typeFile, 0666); err != nil {
		if err == os.ErrExist {
			log.Println("already exists")

		} else {
			log.Println("error when try to create directory", typeFile, "with detail", err.Error())
			errInternal := fiber.ErrInternalServerError
			errInternal.Message = err.Error()
			return errInternal
		}
	}

	destination, err := os.Create(path + "/" + typeFile + "/" + file.Filename)
	if err != nil {
		errBadRequest := fiber.ErrBadRequest
		errBadRequest.Message = err.Error()
		//log.Println("error with detail : ", errBadRequest.Error(), "file size :", (file.Size / 1024 / 104), "MB")
		return errBadRequest
	}

	defer destination.Close()

	source, err := file.Open()
	if err != nil {
		errBadRequest := fiber.ErrBadRequest
		errBadRequest.Message = err.Error()
		//log.Println("error with detail : ", errBadRequest.Error(), "file size :", (file.Size / 1024 / 104), "MB")
		return errBadRequest
	}
	defer source.Close()

	if _, err = io.Copy(destination, source); err != nil {
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		//log.Println("error with detail : ", errBadRequest.Error(), "file size :", (file.Size / 1024 / 104), "MB")
		return errInternal
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "file upload successfully",
	})
}

func Download(c *fiber.Ctx) error {
	type request struct {
		Url string `json:"url"`
	}

	var req = request{}

	if err := c.BodyParser(&req); err != nil {
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	now := time.Now().Unix()
	destination, err := os.Create(config.GetConfigString("PUBLIC_PATH_DOWNLOAD	", "public/downloads") + "/" + fmt.Sprintf("%v", now) + ".jpg")

	if err != nil {
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	defer destination.Close()

	resp, err := http.Get(req.Url)
	if err != nil {
		log.Println("error detail : ", err.Error())
		b, _ := json.Marshal(resp.Body)
		log.Println("response : ", string(b))
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	_, err = io.Copy(destination, resp.Body)
	if err != nil {
		errInternal := fiber.ErrInternalServerError
		errInternal.Message = err.Error()
		return errInternal
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "file download successfully",
	})
}
