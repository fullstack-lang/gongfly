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
var __Marker__dummysDeclaration__ models.Marker
var __Marker_time__dummyDeclaration time.Duration

var mutexMarker sync.Mutex

// An MarkerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMarker updateMarker deleteMarker
type MarkerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MarkerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMarker updateMarker
type MarkerInput struct {
	// The Marker to submit or modify
	// in: body
	Marker *orm.MarkerAPI
}

// GetMarkers
//
// swagger:route GET /markers markers getMarkers
//
// # Get all markers
//
// Responses:
// default: genericError
//
//	200: markerDBResponse
func (controller *Controller) GetMarkers(c *gin.Context) {

	// source slice
	var markerDBs []orm.MarkerDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarkers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarker.GetDB()

	query := db.Find(&markerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	markerAPIs := make([]orm.MarkerAPI, 0)

	// for each marker, update fields from the database nullable fields
	for idx := range markerDBs {
		markerDB := &markerDBs[idx]
		_ = markerDB
		var markerAPI orm.MarkerAPI

		// insertion point for updating fields
		markerAPI.ID = markerDB.ID
		markerDB.CopyBasicFieldsToMarker_WOP(&markerAPI.Marker_WOP)
		markerAPI.MarkerPointersEncoding = markerDB.MarkerPointersEncoding
		markerAPIs = append(markerAPIs, markerAPI)
	}

	c.JSON(http.StatusOK, markerAPIs)
}

// PostMarker
//
// swagger:route POST /markers markers postMarker
//
// Creates a marker
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMarker(c *gin.Context) {

	mutexMarker.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMarkers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarker.GetDB()

	// Validate input
	var input orm.MarkerAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create marker
	markerDB := orm.MarkerDB{}
	markerDB.MarkerPointersEncoding = input.MarkerPointersEncoding
	markerDB.CopyBasicFieldsFromMarker_WOP(&input.Marker_WOP)

	query := db.Create(&markerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMarker.CheckoutPhaseOneInstance(&markerDB)
	marker := backRepo.BackRepoMarker.Map_MarkerDBID_MarkerPtr[markerDB.ID]

	if marker != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), marker)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, markerDB)

	mutexMarker.Unlock()
}

// GetMarker
//
// swagger:route GET /markers/{ID} markers getMarker
//
// Gets the details for a marker.
//
// Responses:
// default: genericError
//
//	200: markerDBResponse
func (controller *Controller) GetMarker(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarker", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarker.GetDB()

	// Get markerDB in DB
	var markerDB orm.MarkerDB
	if err := db.First(&markerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var markerAPI orm.MarkerAPI
	markerAPI.ID = markerDB.ID
	markerAPI.MarkerPointersEncoding = markerDB.MarkerPointersEncoding
	markerDB.CopyBasicFieldsToMarker_WOP(&markerAPI.Marker_WOP)

	c.JSON(http.StatusOK, markerAPI)
}

// UpdateMarker
//
// swagger:route PATCH /markers/{ID} markers updateMarker
//
// # Update a marker
//
// Responses:
// default: genericError
//
//	200: markerDBResponse
func (controller *Controller) UpdateMarker(c *gin.Context) {

	mutexMarker.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMarker", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarker.GetDB()

	// Validate input
	var input orm.MarkerAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var markerDB orm.MarkerDB

	// fetch the marker
	query := db.First(&markerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	markerDB.CopyBasicFieldsFromMarker_WOP(&input.Marker_WOP)
	markerDB.MarkerPointersEncoding = input.MarkerPointersEncoding

	query = db.Model(&markerDB).Updates(markerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	markerNew := new(models.Marker)
	markerDB.CopyBasicFieldsToMarker(markerNew)

	// redeem pointers
	markerDB.DecodePointers(backRepo, markerNew)

	// get stage instance from DB instance, and call callback function
	markerOld := backRepo.BackRepoMarker.Map_MarkerDBID_MarkerPtr[markerDB.ID]
	if markerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), markerOld, markerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the markerDB
	c.JSON(http.StatusOK, markerDB)

	mutexMarker.Unlock()
}

// DeleteMarker
//
// swagger:route DELETE /markers/{ID} markers deleteMarker
//
// # Delete a marker
//
// default: genericError
//
//	200: markerDBResponse
func (controller *Controller) DeleteMarker(c *gin.Context) {

	mutexMarker.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMarker", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarker.GetDB()

	// Get model if exist
	var markerDB orm.MarkerDB
	if err := db.First(&markerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&markerDB)

	// get an instance (not staged) from DB instance, and call callback function
	markerDeleted := new(models.Marker)
	markerDB.CopyBasicFieldsToMarker(markerDeleted)

	// get stage instance from DB instance, and call callback function
	markerStaged := backRepo.BackRepoMarker.Map_MarkerDBID_MarkerPtr[markerDB.ID]
	if markerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), markerStaged, markerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexMarker.Unlock()
}
