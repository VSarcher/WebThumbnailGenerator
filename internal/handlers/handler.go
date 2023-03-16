package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type ThumbNailRequest struct {
	Urls []string `json:"urls"`
}

type ImgSourceResponse struct {
	Error   int `json:"Error"`
	IsReady int `json:"IsReady"`
}

func SaveThumbnail(c *fiber.Ctx) error {
	todoUrls := new(ThumbNailRequest)
	if err := c.BodyParser(todoUrls); err != nil {
		c.JSON(fiber.Map{"status": "error", "payload": err.Error()})
	}
	fmt.Println(todoUrls)

	for _, todoUrl := range todoUrls.Urls {
		// go func() {
		err := getImage(todoUrl)
		if err != nil {
			c.JSON(fiber.Map{"status": "error", "payload": err.Error()})
		}

	}

	c.JSON(fiber.Map{"payload": "Success"})
	return nil
}

func getImage(imgUrl string) error {
	url1, err := url.Parse("http://free.pagepeeker.com/v2/thumbs.php")
	if err != nil {
		return err
	}
	values := url1.Query()
	values.Add("size", "m")
	values.Add("code", "9b8fa936fd")
	values.Add("url", imgUrl)
	url1.RawQuery = values.Encode()
	req, err := http.NewRequest(
		"GET",
		url1.String(),
		nil,
	)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// save Image as .png form
	imageFileName := fmt.Sprint("image/", imgUrl, ".png")
	ioutil.WriteFile(imageFileName, body, 0666)
	imgResponse := new(ImgSourceResponse)
	err = json.Unmarshal(body, &imgResponse)

	//encode image

	return nil
}
