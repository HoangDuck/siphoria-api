package services

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	response "hotel-booking-api/model/model_func"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type FileUploadService struct {
	Echo *echo.Echo
}

func (fileUpload *FileUploadService) UploadMultipleFilesAPI(c echo.Context) error {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["images"]
	//generate folder name to contain all uploaded images
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	err = os.Mkdir(timestamp, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(filepath.Join(timestamp, file.Filename))
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}
	filesPath, err := ioutil.ReadDir(timestamp)
	if err != nil {
		log.Fatal(err)
	}
	var stu []string
	var resultFile []string

	for _, file := range filesPath {
		stu = append(stu, file.Name())
	}
	var ctx = context.Background()
	var cld, _ = cloudinary.NewFromURL(ConfigInfo.Cloudinary.CloudinaryUrl)

	var wg sync.WaitGroup
	for _, st := range stu {
		wg.Add(1)
		go func(st string) {
			defer wg.Done()
			result, err := cld.Upload.Upload(ctx, timestamp+"/"+st, uploader.UploadParams{PublicID: strings.Split(st, ".")[0]})
			if err != nil {
				log.Fatal(err)
				return
			}
			resultFile = append(resultFile, result.SecureURL)
		}(st)
	}
	wg.Wait()
	defer func() {
		err = os.RemoveAll(timestamp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resultFile)
	}()
	return response.Ok(c, "Tải file lên thành công", resultFile)
}

func UploadMultipleFiles(c echo.Context) []string {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return []string{}
	}
	files := form.File["images"]
	//generate folder name to contain all uploaded images
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	err = os.Mkdir(timestamp, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			continue
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(filepath.Join(timestamp, file.Filename))
		if err != nil {
			continue
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			continue
		}

	}
	filesPath, err := ioutil.ReadDir(timestamp)
	if err != nil {
		log.Fatal(err)
	}
	var stu []string
	var resultFile []string

	for _, file := range filesPath {
		stu = append(stu, file.Name())
	}
	var ctx = context.Background()
	var cld, _ = cloudinary.NewFromURL(ConfigInfo.Cloudinary.CloudinaryUrl)

	var wg sync.WaitGroup
	for _, st := range stu {
		wg.Add(1)
		go func(st string) {
			defer wg.Done()
			result, err := cld.Upload.Upload(ctx, timestamp+"/"+st, uploader.UploadParams{PublicID: strings.Split(st, ".")[0]})
			if err != nil {
				return
			}
			resultFile = append(resultFile, result.SecureURL)
		}(st)
	}
	wg.Wait()
	defer func() {
		err = os.RemoveAll(timestamp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resultFile)
	}()
	return resultFile
}
