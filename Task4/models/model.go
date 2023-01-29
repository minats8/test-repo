package models

type Data struct {
	UserId int    `json : "userId" bson : "userId"`
	Id     int    `json : "id"     bson : "id" `
	Title  string `json : "title"  bson : "title"`
	Body   string `json : "body"   bson " "body"`
}
