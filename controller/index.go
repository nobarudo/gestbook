package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type post struct {
	Id   int    `db:"id"`
	Text string `db:"text"`
}

func DBConnect() (db *sqlx.DB) {
	db, err := sqlx.Open("mysql", "gestbook:password@tcp(localhost:3306)/gestbook")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func IndexGET(c *gin.Context) {

	db := DBConnect()

	posts := []post{}

	db.Select(&posts, "select * from posts")
	log.Println("debug:", posts[0].Text)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"post": posts,
	})
}

func ReceivePost(c *gin.Context) {
	db := DBConnect()
	c.Request.ParseForm()
	text := c.Request.Form["text"][0]
	log.Println("debug:", text)
	db.Queryx("INSERT INTO posts (text) VALUES(?)", text)
	c.Redirect(http.StatusMovedPermanently, "/")
}
