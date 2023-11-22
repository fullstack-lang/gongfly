// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Radar__dummysDeclaration__ models.Radar
var __Radar_time__dummyDeclaration time.Duration

var mutexRadar sync.Mutex

// An RadarID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRadar updateRadar deleteRadar
type RadarID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RadarInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRadar updateRadar
type RadarInput struct {
	// The Radar to submit or modify
	// in: body
	Radar *orm.RadarAPI
}

// GetRadars
//
// swagger:route GET /radars radars getRadars
//
// # Get all radars
//
// Responses:
// default: genericError
//
//	200: radarDBResponse
func (controller *Controller) GetRadars(c *gin.Context) {

	// source slice
	var radarDBs []orm.RadarDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRadars", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRadar.GetDB()

	query := db.Find(&radarDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	radarAPIs := make([]orm.RadarAPI, 0)

	// for each radar, update fields from the database nullable fields
	for idx := range radarDBs {
		radarDB := &radarDBs[idx]
		_ = radarDB
		var radarAPI orm.RadarAPI

		// insertion point for updating fields
		radarAPI.ID = radarDB.ID
		radarDB.CopyBasicFieldsToRadar_WOP(&radarAPI.Radar_WOP)
		radarAPI.RadarPointersEncoding = radarDB.RadarPointersEncoding
		radarAPIs = append(radarAPIs, radarAPI)
	}

	c.JSON(http.StatusOK, radarAPIs)
}

// PostRadar
//
// swagger:route POST /radars radars postRadar
//
// Creates a radar
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRadar(c *gin.Context) {

	mutexRadar.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRadars", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRadar.GetDB()

	// Validate input
	var input orm.RadarAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create radar
	radarDB := orm.RadarDB{}
	radarDB.RadarPointersEncoding = input.RadarPointersEncoding
	radarDB.CopyBasicFieldsFromRadar_WOP(&input.Radar_WOP)

	query := db.Create(&radarDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRadar.CheckoutPhaseOneInstance(&radarDB)
	radar := backRepo.BackRepoRadar.Map_RadarDBID_RadarPtr[radarDB.ID]

	if radar != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), radar)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, radarDB)

	mutexRadar.Unlock()
}

// GetRadar
//
// swagger:route GET /radars/{ID} radars getRadar
//
// Gets the details for a radar.
//
// Responses:
// default: genericError
//
//	200: radarDBResponse
func (controller *Controller) GetRadar(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRadar", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRadar.GetDB()

	// Get radarDB in DB
	var radarDB orm.RadarDB
	if err := db.First(&radarDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var radarAPI orm.RadarAPI
	radarAPI.ID = radarDB.ID
	radarAPI.RadarPointersEncoding = radarDB.RadarPointersEncoding
	radarDB.CopyBasicFieldsToRadar_WOP(&radarAPI.Radar_WOP)

	c.JSON(http.StatusOK, radarAPI)
}

// UpdateRadar
//
// swagger:route PATCH /radars/{ID} radars updateRadar
//
// # Update a radar
//
// Responses:
// default: genericError
//
//	200: radarDBResponse
func (controller *Controller) UpdateRadar(c *gin.Context) {

	mutexRadar.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRadar", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRadar.GetDB()

	// Validate input
	var input orm.RadarAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var radarDB orm.RadarDB

	// fetch the radar
	query := db.First(&radarDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	radarDB.CopyBasicFieldsFromRadar_WOP(&input.Radar_WOP)
	radarDB.RadarPointersEncoding = input.RadarPointersEncoding

	query = db.Model(&radarDB).Updates(radarDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	radarNew := new(models.Radar)
	radarDB.CopyBasicFieldsToRadar(radarNew)

	// redeem pointers
	radarDB.DecodePointers(backRepo, radarNew)

	// get stage instance from DB instance, and call callback function
	radarOld := backRepo.BackRepoRadar.Map_RadarDBID_RadarPtr[radarDB.ID]
	if radarOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), radarOld, radarNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the radarDB
	c.JSON(http.StatusOK, radarDB)

	mutexRadar.Unlock()
}

// DeleteRadar
//
// swagger:route DELETE /radars/{ID} radars deleteRadar
//
// # Delete a radar
//
// default: genericError
//
//	200: radarDBResponse
func (controller *Controller) DeleteRadar(c *gin.Context) {

	mutexRadar.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRadar", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRadar.GetDB()

	// Get model if exist
	var radarDB orm.RadarDB
	if err := db.First(&radarDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&radarDB)

	// get an instance (not staged) from DB instance, and call callback function
	radarDeleted := new(models.Radar)
	radarDB.CopyBasicFieldsToRadar(radarDeleted)

	// get stage instance from DB instance, and call callback function
	radarStaged := backRepo.BackRepoRadar.Map_RadarDBID_RadarPtr[radarDB.ID]
	if radarStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), radarStaged, radarDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexRadar.Unlock()
}
