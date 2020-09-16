package questioncontroller

import (
	"net/http"
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	"github.com/labstack/echo/v4"
)

type requestPost struct {
	UserID     int    `json:"user_id"`
	SentenceID int    `json:"sentence_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PageNum    int    `json:"page_num"`
	RowNum     int    `json:"row_num"`
}

//Post ...質問投稿
func Post(c echo.Context) error {
	request := new(requestPost)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"HasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	var que question.Question
	//構造体に入れる
	que.UserID = request.UserID
	que.SentenceID = request.SentenceID
	que.Title = request.Title
	que.Content = request.Content
	que.PageNum = request.PageNum
	que.RowNum = request.RowNum
	jst, _ := time.LoadLocation("Asia/Tokyo")
	que.CreatedAt = time.Now().In(jst)

	//DB処理
	hasSuccess := question.Create(que)
	returnValue := map[string]bool{"HasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}
