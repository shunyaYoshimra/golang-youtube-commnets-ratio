package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimura/YouTube_Comments/entity"
	"github.com/shunyaYoshimura/YouTube_Comments/repositories"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	comments  []*youtube.CommentThread
	jaComment int64
	ratio     float64
)

type VideoController struct {
	Repository repositories.VideoRepository
}

func NewVideoController() VideoController {
	return VideoController{
		Repository: repositories.NewVideoRepository(),
	}
}

func (vc *VideoController) Index(c *gin.Context) {
	videos := vc.Repository.RetrieveVideos()
	c.JSON(http.StatusOK, videos)
}

func (vc *VideoController) Create(c *gin.Context) {
	query := c.PostForm("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, "Video IDを入力して下さい")
	} else {
		comments := getComments(query)
		checkJaorNot(comments)
		ratioOfJaComments(comments)
		title := getTitle(query)
		video := entity.Video{
			Ration:  100 - ratio,
			Title:   title,
			VideoID: query,
		}
		fmt.Println(video)
		if err := vc.Repository.Create(&video); err != nil {
			c.JSON(http.StatusBadRequest, "予期せぬエラーが発生しました")
		} else {
			c.JSON(http.StatusOK, video)
		}
	}
}

// youtube api

func newClient() *http.Client {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	}
	return client
}

func newYoutubeService(client *http.Client) *youtube.Service {
	service, err := youtube.New(client)
	if err != nil {
		log.Fatal("Unable to create YouTube service: %v", err)
	}
	return service
}

func getComments(query string) []*youtube.CommentThread {
	service := newYoutubeService(newClient())
	call := service.CommentThreads.List([]string{"id", "snippet"}).VideoId(query).Order("time").MaxResults(100)
	response, err := call.Do()
	if err != nil {
		log.Fatal("%v", err)
	}
	return response.Items
}

func getTitle(query string) string {
	var title string
	service := newYoutubeService(newClient())
	videoListCall := service.Videos.List([]string{"id", "snippet"}).Id(query).MaxResults(25)
	response, err := videoListCall.Do()
	if err != nil {
		panic(err)
	}
	for _, video := range response.Items {
		title = video.Snippet.Title

	}
	return title
}

func checkJaorNot(comments []*youtube.CommentThread) {
	jaComment = 0
	for _, comment := range comments {
		for _, r := range comment.Snippet.TopLevelComment.Snippet.TextOriginal {
			if unicode.In(r, unicode.Hiragana) || unicode.In(r, unicode.Katakana) {
				jaComment++
				break
			}
		}
	}
	fmt.Println(jaComment)
}

func ratioOfJaComments(comments []*youtube.CommentThread) float64 {
	if len(comments) >= 100 {
		ratio = (float64(jaComment) / float64(100)) * 100
	} else {
		ratio = (float64(jaComment) / float64(len(comments))) * 100
	}
	return ratio
}
