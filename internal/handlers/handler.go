package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type ThumbNailRequest struct {
	Urls []string `json:urls`
}

func SaveThumbnail(c *fiber.Ctx) error {
	todoUrls := new(ThumbNailRequest)
	if err := c.BodyParser(todoUrls); err != nil {
		c.JSON(fiber.Map{"status": "error", "payload": err.Error()})
	}
	fmt.Println(todoUrls)

	err := getImage(todoUrls.Urls[0])
	if err != nil {
		c.JSON(fiber.Map{"status": "error", "payload": err.Error()})
	}
	c.JSON(fiber.Map{"payload": "FIBER"})
	return nil
}

func getImage(imgUrl string) error {
	url1, err := url.Parse("http://api.pagepeeker.com/v2/thumbs.php")
	if err != nil {
		return err
	}
	values := url1.Query()
	values.Add("size", "m")
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
	fmt.Println(req)
	return nil
}
