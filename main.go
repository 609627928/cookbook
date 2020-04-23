package main

import (
	"github.com/609627928/cookbook/db"
	"github.com/609627928/cookbook/handler"
	"github.com/609627928/cookbook/router"
	"github.com/609627928/cookbook/store"
)

func main() {
	r := router.New()
	r_api := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)
	h := handler.NewHandler(us, as)
	h.Register(r_api)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}

