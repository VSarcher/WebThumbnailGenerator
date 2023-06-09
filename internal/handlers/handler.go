package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/VSarcher/WebThumbnailGenerator/database"
	"github.com/VSarcher/WebThumbnailGenerator/internal/models"
	"github.com/gofiber/fiber/v2"
)

type ThumbNailRequest struct {
	Urls []string `json:"urls"`
}

type ImgSourceResponse struct {
	Error   int `json:"Error"`
	IsReady int `json:"IsReady"`
}

const IMGURLSOURCE = "image/"
const PAGEPEEKER_URL = "http://free.pagepeeker.com/v2/thumbs.php"

func SaveThumbnail(c *fiber.Ctx) error {
	// Get list of urls to retrieve image
	todoUrls := new(ThumbNailRequest)
	if err := c.BodyParser(todoUrls); err != nil || todoUrls.Urls == nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "payload": err.Error()})
	}

	// Channel for send retrieved data of image
	ImageForChannel := make(chan models.ImageInfo)

	// Concurrently get Image from peek site using Goroutine.
	for _, todoUrl := range todoUrls.Urls {
		go func(todoUrl string) {
			dataString, err := getImage(todoUrl)
			if err != nil {
				c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"status": "error", "payload": err.Error()})
			}
			newImage := new(models.ImageInfo)
			newImage.Image = dataString
			newImage.Name = todoUrl
			newImage.Url = todoUrl

			ImageForChannel <- *newImage
		}(todoUrl)
	}

	// save image in DB.
	for range todoUrls.Urls {
		tmp := <-ImageForChannel
		database.DB.Create(&tmp)
	}
	close(ImageForChannel)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"payload": "Success"})
}

// Decode given byte array to string
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Get screenshot from given url and return base64 decoded string
func getImage(imgUrl string) (string, error) {
	// prepare url for get screenshot
	url1, err := url.Parse(PAGEPEEKER_URL)
	if err != nil {
		return "", err
	}
	values := url1.Query()
	values.Add("size", "m")
	values.Add("code", "9b8fa936fd")
	values.Add("url", imgUrl)
	url1.RawQuery = values.Encode()

	// prepare request
	req, err := http.NewRequest(
		"GET",
		url1.String(),
		nil,
	)
	if err != nil {
		return "", err
	}

	// send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// save Image as .png form
	imageFileName := fmt.Sprint(IMGURLSOURCE, imgUrl, ".png")
	ioutil.WriteFile(imageFileName, body, 0666)
	imgResponse := new(ImgSourceResponse)
	err = json.Unmarshal(body, &imgResponse)

	//encode image
	base64Encoding := toBase64(body)
	// fmt.Println("aef", base64Encoding)

	return base64Encoding, nil
}

// func decodeImage(str string) {
// 	imgByte, err := base64.StdEncoding.DecodeString(str)
// 	if err != nil {
// 		return
// 	}
// 	ioutil.WriteFile("sample.png", imgByte, 0666)
// }
