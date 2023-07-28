package main

import (
	"site/internal/middleware"
	"site/scr"

	"github.com/brendonmatos/golive"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	liveServer := golive.NewServer()

	var livePage golive.PageContent
	livePage.Lang = "us"
	livePage.Title = "Example Application for GoLive"

	/*	loggerbsc := golive.NewLoggerBasic()
		loggerbsc.Level = golive.LogDebug
		liveServer.Log = loggerbsc.Log
	*/
	app.Get("/", liveServer.CreateHTMLHandlerWithMiddleware(scr.NewLoginForm(), livePage, middleware.Session))
	app.Get("/ws", websocket.New(liveServer.HandleWSRequest))
	_ = app.Listen(":80")
}
