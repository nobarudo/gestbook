package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

type post struct {
	Id   int    `db:"id"`
	Text string `db:"text"`
}

func Showpost() (posts []post) {
	db := DBConnect()

	posts = []post{}

	db.Select(&posts, "select * from posts")
	log.Println("debug:", posts[0].Text)

	return posts
}

func Insertpost(text string) {
	db := DBConnect()
	log.Println("debug:", text)
	db.Queryx("INSERT INTO posts (text) VALUES(?)", text)
}
