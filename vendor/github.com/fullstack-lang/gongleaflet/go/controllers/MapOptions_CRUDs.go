// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongleaflet/go/models"
	"github.com/fullstack-lang/gongleaflet/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __MapOptions__dummysDeclaration__ models.MapOptions
var __MapOptions_time__dummyDeclaration time.Duration

var mutexMapOptions sync.Mutex

// An MapOptionsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMapOptions updateMapOptions deleteMapOptions
type MapOptionsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MapOptionsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMapOptions updateMapOptions
type MapOptionsInput struct {
	// The MapOptions to submit or modify
	// in: body
	MapOptions *orm.MapOptionsAPI
}

// GetMapOptionss
//
// swagger:route GET /mapoptionss mapoptionss getMapOptionss
//
// # Get all mapoptionss
//
// Responses:
// default: genericError
//
//	200: mapoptionsDBResponse
func (controller *Controller) GetMapOptionss(c *gin.Context) {

	// source slice
	var mapoptionsDBs []orm.MapOptionsDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapOptionss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapOptions.GetDB()

	query := db.Find(&mapoptionsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	mapoptionsAPIs := make([]orm.MapOptionsAPI, 0)

	// for each mapoptions, update fields from the database nullable fields
	for idx := range mapoptionsDBs {
		mapoptionsDB := &mapoptionsDBs[idx]
		_ = mapoptionsDB
		var mapoptionsAPI orm.MapOptionsAPI

		// insertion point for updating fields
		mapoptionsAPI.ID = mapoptionsDB.ID
		mapoptionsDB.CopyBasicFieldsToMapOptions_WOP(&mapoptionsAPI.MapOptions_WOP)
		mapoptionsAPI.MapOptionsPointersEncoding = mapoptionsDB.MapOptionsPointersEncoding
		mapoptionsAPIs = append(mapoptionsAPIs, mapoptionsAPI)
	}

	c.JSON(http.StatusOK, mapoptionsAPIs)
}

// PostMapOptions
//
// swagger:route POST /mapoptionss mapoptionss postMapOptions
//
// Creates a mapoptions
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMapOptions(c *gin.Context) {

	mutexMapOptions.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMapOptionss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapOptions.GetDB()

	// Validate input
	var input orm.MapOptionsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mapoptions
	mapoptionsDB := orm.MapOptionsDB{}
	mapoptionsDB.MapOptionsPointersEncoding = input.MapOptionsPointersEncoding
	mapoptionsDB.CopyBasicFieldsFromMapOptions_WOP(&input.MapOptions_WOP)

	query := db.Create(&mapoptionsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMapOptions.CheckoutPhaseOneInstance(&mapoptionsDB)
	mapoptions := backRepo.BackRepoMapOptions.Map_MapOptionsDBID_MapOptionsPtr[mapoptionsDB.ID]

	if mapoptions != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mapoptions)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, mapoptionsDB)

	mutexMapOptions.Unlock()
}

// GetMapOptions
//
// swagger:route GET /mapoptionss/{ID} mapoptionss getMapOptions
//
// Gets the details for a mapoptions.
//
// Responses:
// default: genericError
//
//	200: mapoptionsDBResponse
func (controller *Controller) GetMapOptions(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapOptions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapOptions.GetDB()

	// Get mapoptionsDB in DB
	var mapoptionsDB orm.MapOptionsDB
	if err := db.First(&mapoptionsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var mapoptionsAPI orm.MapOptionsAPI
	mapoptionsAPI.ID = mapoptionsDB.ID
	mapoptionsAPI.MapOptionsPointersEncoding = mapoptionsDB.MapOptionsPointersEncoding
	mapoptionsDB.CopyBasicFieldsToMapOptions_WOP(&mapoptionsAPI.MapOptions_WOP)

	c.JSON(http.StatusOK, mapoptionsAPI)
}

// UpdateMapOptions
//
// swagger:route PATCH /mapoptionss/{ID} mapoptionss updateMapOptions
//
// # Update a mapoptions
//
// Responses:
// default: genericError
//
//	200: mapoptionsDBResponse
func (controller *Controller) UpdateMapOptions(c *gin.Context) {

	mutexMapOptions.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMapOptions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapOptions.GetDB()

	// Validate input
	var input orm.MapOptionsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var mapoptionsDB orm.MapOptionsDB

	// fetch the mapoptions
	query := db.First(&mapoptionsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	mapoptionsDB.CopyBasicFieldsFromMapOptions_WOP(&input.MapOptions_WOP)
	mapoptionsDB.MapOptionsPointersEncoding = input.MapOptionsPointersEncoding

	query = db.Model(&mapoptionsDB).Updates(mapoptionsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	mapoptionsNew := new(models.MapOptions)
	mapoptionsDB.CopyBasicFieldsToMapOptions(mapoptionsNew)

	// get stage instance from DB instance, and call callback function
	mapoptionsOld := backRepo.BackRepoMapOptions.Map_MapOptionsDBID_MapOptionsPtr[mapoptionsDB.ID]
	if mapoptionsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), mapoptionsOld, mapoptionsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the mapoptionsDB
	c.JSON(http.StatusOK, mapoptionsDB)

	mutexMapOptions.Unlock()
}

// DeleteMapOptions
//
// swagger:route DELETE /mapoptionss/{ID} mapoptionss deleteMapOptions
//
// # Delete a mapoptions
//
// default: genericError
//
//	200: mapoptionsDBResponse
func (controller *Controller) DeleteMapOptions(c *gin.Context) {

	mutexMapOptions.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMapOptions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapOptions.GetDB()

	// Get model if exist
	var mapoptionsDB orm.MapOptionsDB
	if err := db.First(&mapoptionsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&mapoptionsDB)

	// get an instance (not staged) from DB instance, and call callback function
	mapoptionsDeleted := new(models.MapOptions)
	mapoptionsDB.CopyBasicFieldsToMapOptions(mapoptionsDeleted)

	// get stage instance from DB instance, and call callback function
	mapoptionsStaged := backRepo.BackRepoMapOptions.Map_MapOptionsDBID_MapOptionsPtr[mapoptionsDB.ID]
	if mapoptionsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), mapoptionsStaged, mapoptionsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexMapOptions.Unlock()
}
