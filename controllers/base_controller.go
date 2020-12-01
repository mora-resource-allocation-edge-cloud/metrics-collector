package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"metrics-collector/models"
	"runtime/debug"
)

type BaseController struct {
	beego.Controller
}

func (bc *BaseController) handlePanic() {
	if r := recover(); r != nil {
		log.Println(r)
		debug.PrintStack()
		status := bc.Ctx.ResponseWriter.Status
		if status == 0 {
			status = 500
			bc.Ctx.ResponseWriter.WriteHeader(status)
		}
		errorMessage := fmt.Sprintf("%v", r)
		bc.Data["json"] = models.CreateResponse(status, errorMessage, "InternalServerError", nil)
		bc.ServeJSON()
	}
}

func (bc *BaseController) replyFactory(array interface{}, err error) {
	defer bc.handlePanic()
	if err != nil {
		panic(err)
	}
	bc.Data["json"] = models.CreateSuccessResponse("", "", array)
	bc.ServeJSON()
}
