package controllers

import (
	"log"
	"metrics-collector/database"
	"metrics-collector/models"
)

// VideoReproduction API
type VideoReproductionController struct {
	BaseController
}

/*
If you do not use URLMapping Beego will find the function by reflection,
otherwise Beego will find the function with the must faster interface.
However Mapping method doesn't accept parameters on controllers method
*/
func (dc *VideoReproductionController) URLMapping() {
}

// @Title Add a new reproduction entry
// @Description Add a new reproduction entry with their metrics
// @Accept json
// @Success 200 {} models.GenericResponse
// @Param 		body 							body 		models.VideoReproduction		true 	"the JSON body"
// @router /video-reproduction [post]
func (dc *VideoReproductionController) PostReproduction(body models.VideoReproduction) {
	defer dc.handlePanic()
	log.Printf("New reproduction with body:\n%+v", body)
	dc.replyFactory(database.VideoReproductionRepository.StoreMetric(&body))
}
