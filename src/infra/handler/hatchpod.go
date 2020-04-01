package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// MyPodcastAPI podcast加工に使うためのAPI郡
type MyPodcastAPI struct{}

// GetTemplate 説明文のテンプレを生成します
func (handler *MyPodcastAPI) GetTemplate(c echo.Context) error {
	anondID := c.QueryParam("id")
	anondURL := fmt.Sprintf("https://anond.hatelabo.jp/%s", anondID)

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest(
		"GET",
		anondURL,
		nil,
	)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	return c.JSONPretty(http.StatusOK, bodyString, indent)
}
