package controllers

import (
	"github.com/kataras/iris"
)

// OkzApp set up new app with routes and statics assets
func OkzApp() *iris.Application {
	app := iris.New()

	app.StaticWeb("/css", "./views/css")
	app.StaticWeb("/js", "./views/js")

	templates := iris.HTML("./views", ".html").Reload(true)
	app.RegisterView(templates)

	app.Get("/", checkIfLoggedIn, home)
	app.Get("/login", blockAuthPages, login)
	app.Post("/login", blockAuthPages, authentification)
	app.Post("/register", blockAuthPages, register)
	app.Get("/ad", detail)
	app.Get("/create", create)
	app.Post("/create", createAdvert)
	app.Get("/user", userPage)
	app.Get("/user/msg", userMessages)
	app.Get("/user/setting", userSetting)
	app.Get("/user/thread", thread)
	app.Post("/verification", verification)
	app.Get("/logout", logout)

	return app
}