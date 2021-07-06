package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimura/YouTube_Comments/database"
	"github.com/shunyaYoshimura/YouTube_Comments/entity"
	"github.com/shunyaYoshimura/YouTube_Comments/routes"
)

type Application struct{}

func (a Application) CreateApp(r *gin.Engine) {
	configureAppDB()
	configureAPIEndpoint(r)
	configureView(r)
}

func configureAppDB() {
	database.AppConnection()
	conn := database.GetDB()
	entity.MigrateVideo(conn)
}

func configureAPIEndpoint(r *gin.Engine) {
	routes.NewRouter(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "The page you are looking for dosen't exist",
		})
	})
}

func configureView(r *gin.Engine) {
	r.Static("/src", "../dist")
	r.StaticFS("/app", http.Dir("../static"))
}
