package controller

import (
	"gestbook/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"net/http"
)

func IndexGET(c *gin.Context) {
	posts := model.Showpost()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"post": posts,
	})
}

func ReceivePost(c *gin.Context) {
	c.Request.ParseForm()
	text := c.Request.Form["text"][0]
	model.Insertpost(text)
	c.Redirect(http.StatusMovedPermanently, "/")
}
