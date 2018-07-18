package controllers

import (
	"log"

	"github.com/behouba/OKZ_BETA_0.01/models"
	"github.com/kataras/iris"
)

func home(ctx iris.Context) {
	session := models.Sess.Start(ctx).GetString("email")
	ctx.ViewData("cities", models.Cities)
	ctx.ViewData("categories", models.Categories)
	if session == "" {
		ctx.View("index.html")
		return
	}
	userEmail := models.Sess.Start(ctx).GetString("email")
	user, err := models.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err)
	}
	ctx.ViewData("user", user)
	ctx.View("index.html")
}

func detail(ctx iris.Context) {
	ctx.View("detail.html")
}

func create(ctx iris.Context) {
	session := models.Sess.Start(ctx).GetString("email")
	ctx.ViewData("cities", models.Cities)
	ctx.ViewData("categories", models.Categories)
	if session == "" {
		ctx.View("login.html")
		return
	}
	userEmail := models.Sess.Start(ctx).GetString("email")
	user, err := models.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err)
	}
	ctx.ViewData("user", user)
	ctx.View("create.html")
}

func userPage(ctx iris.Context) {
	ctx.View("profile.html")
}

func userMessages(ctx iris.Context) {
	ctx.View("messages.html")
}

func userSetting(ctx iris.Context) {
	ctx.View("profil-setting.html")
}

func thread(ctx iris.Context) {
	ctx.View("thread.html")
}

func createAdvert(ctx iris.Context) {

}