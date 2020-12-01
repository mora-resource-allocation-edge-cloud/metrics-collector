package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"log"
	_ "metrics-collector/database"
	_ "metrics-collector/routers"
	"strings"
)

func main() {
	log.Printf("Hello world, by metrics collector!")
	initBeego()
}

func initBeego() {
	log.Printf("Initializing http server")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		scriptName := context.Request.Header.Get("X-Script-Name")
		if strings.HasPrefix(context.Request.RequestURI, scriptName) {
			context.Request.RequestURI = strings.TrimPrefix(context.Request.RequestURI, scriptName)
			context.Request.URL.Path = strings.TrimPrefix(context.Request.URL.Path, scriptName)
		}
	})
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
