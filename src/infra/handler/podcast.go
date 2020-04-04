package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/html"
)

// MyPodcastAPI podcast加工に使うためのAPI郡
type MyPodcastAPI struct{}

// GetTemplate 説明文のテンプレを生成します
func (handler *MyPodcastAPI) GetTemplate(c echo.Context) error {
	anondID := c.QueryParam("id")
	numberStr := c.QueryParam("number")
	number, _ := strconv.Atoi(numberStr)

	loc, _ := time.LoadLocation("Asia/Tokyo")
	baseDate := time.Date(2020, 4, 5, 12, 0, 0, 0, loc)
	target := baseDate.AddDate(0, 0, 7*(number-14))
	targetDateString := target.Format("2006/1/2 15:04:05")

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

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	title, _ := traverse(doc)

	res := fmt.Sprintf("<p>%s</p><p>#%d 「%s」について語らう</p><p>はてな匿名ダイアリーをランダムに表示する増田ランダムを利用して，ディスカッションをします．</p><p>今回の記事は「%s」です．</p>", targetDateString, number, title, title)
	return c.HTML(http.StatusOK, res)
}
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}
