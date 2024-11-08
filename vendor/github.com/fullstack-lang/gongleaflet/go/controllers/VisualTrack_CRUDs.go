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
var __VisualTrack__dummysDeclaration__ models.VisualTrack
var __VisualTrack_time__dummyDeclaration time.Duration

var mutexVisualTrack sync.Mutex

// An VisualTrackID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVisualTrack updateVisualTrack deleteVisualTrack
type VisualTrackID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VisualTrackInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVisualTrack updateVisualTrack
type VisualTrackInput struct {
	// The VisualTrack to submit or modify
	// in: body
	VisualTrack *orm.VisualTrackAPI
}

// GetVisualTracks
//
// swagger:route GET /visualtracks visualtracks getVisualTracks
//
// # Get all visualtracks
//
// Responses:
// default: genericError
//
//	200: visualtrackDBResponse
func (controller *Controller) GetVisualTracks(c *gin.Context) {

	// source slice
	var visualtrackDBs []orm.VisualTrackDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVisualTracks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVisualTrack.GetDB()

	_, err := db.Find(&visualtrackDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	visualtrackAPIs := make([]orm.VisualTrackAPI, 0)

	// for each visualtrack, update fields from the database nullable fields
	for idx := range visualtrackDBs {
		visualtrackDB := &visualtrackDBs[idx]
		_ = visualtrackDB
		var visualtrackAPI orm.VisualTrackAPI

		// insertion point for updating fields
		visualtrackAPI.ID = visualtrackDB.ID
		visualtrackDB.CopyBasicFieldsToVisualTrack_WOP(&visualtrackAPI.VisualTrack_WOP)
		visualtrackAPI.VisualTrackPointersEncoding = visualtrackDB.VisualTrackPointersEncoding
		visualtrackAPIs = append(visualtrackAPIs, visualtrackAPI)
	}

	c.JSON(http.StatusOK, visualtrackAPIs)
}

// PostVisualTrack
//
// swagger:route POST /visualtracks visualtracks postVisualTrack
//
// Creates a visualtrack
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVisualTrack(c *gin.Context) {

	mutexVisualTrack.Lock()
	defer mutexVisualTrack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVisualTracks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVisualTrack.GetDB()

	// Validate input
	var input orm.VisualTrackAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create visualtrack
	visualtrackDB := orm.VisualTrackDB{}
	visualtrackDB.VisualTrackPointersEncoding = input.VisualTrackPointersEncoding
	visualtrackDB.CopyBasicFieldsFromVisualTrack_WOP(&input.VisualTrack_WOP)

	_, err = db.Create(&visualtrackDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVisualTrack.CheckoutPhaseOneInstance(&visualtrackDB)
	visualtrack := backRepo.BackRepoVisualTrack.Map_VisualTrackDBID_VisualTrackPtr[visualtrackDB.ID]

	if visualtrack != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), visualtrack)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, visualtrackDB)
}

// GetVisualTrack
//
// swagger:route GET /visualtracks/{ID} visualtracks getVisualTrack
//
// Gets the details for a visualtrack.
//
// Responses:
// default: genericError
//
//	200: visualtrackDBResponse
func (controller *Controller) GetVisualTrack(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVisualTrack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVisualTrack.GetDB()

	// Get visualtrackDB in DB
	var visualtrackDB orm.VisualTrackDB
	if _, err := db.First(&visualtrackDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var visualtrackAPI orm.VisualTrackAPI
	visualtrackAPI.ID = visualtrackDB.ID
	visualtrackAPI.VisualTrackPointersEncoding = visualtrackDB.VisualTrackPointersEncoding
	visualtrackDB.CopyBasicFieldsToVisualTrack_WOP(&visualtrackAPI.VisualTrack_WOP)

	c.JSON(http.StatusOK, visualtrackAPI)
}

// UpdateVisualTrack
//
// swagger:route PATCH /visualtracks/{ID} visualtracks updateVisualTrack
//
// # Update a visualtrack
//
// Responses:
// default: genericError
//
//	200: visualtrackDBResponse
func (controller *Controller) UpdateVisualTrack(c *gin.Context) {

	mutexVisualTrack.Lock()
	defer mutexVisualTrack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVisualTrack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVisualTrack.GetDB()

	// Validate input
	var input orm.VisualTrackAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var visualtrackDB orm.VisualTrackDB

	// fetch the visualtrack
	_, err := db.First(&visualtrackDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	visualtrackDB.CopyBasicFieldsFromVisualTrack_WOP(&input.VisualTrack_WOP)
	visualtrackDB.VisualTrackPointersEncoding = input.VisualTrackPointersEncoding

	db, _ = db.Model(&visualtrackDB)
	_, err = db.Updates(&visualtrackDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	visualtrackNew := new(models.VisualTrack)
	visualtrackDB.CopyBasicFieldsToVisualTrack(visualtrackNew)

	// redeem pointers
	visualtrackDB.DecodePointers(backRepo, visualtrackNew)

	// get stage instance from DB instance, and call callback function
	visualtrackOld := backRepo.BackRepoVisualTrack.Map_VisualTrackDBID_VisualTrackPtr[visualtrackDB.ID]
	if visualtrackOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), visualtrackOld, visualtrackNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the visualtrackDB
	c.JSON(http.StatusOK, visualtrackDB)
}

// DeleteVisualTrack
//
// swagger:route DELETE /visualtracks/{ID} visualtracks deleteVisualTrack
//
// # Delete a visualtrack
//
// default: genericError
//
//	200: visualtrackDBResponse
func (controller *Controller) DeleteVisualTrack(c *gin.Context) {

	mutexVisualTrack.Lock()
	defer mutexVisualTrack.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVisualTrack", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVisualTrack.GetDB()

	// Get model if exist
	var visualtrackDB orm.VisualTrackDB
	if _, err := db.First(&visualtrackDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&visualtrackDB)

	// get an instance (not staged) from DB instance, and call callback function
	visualtrackDeleted := new(models.VisualTrack)
	visualtrackDB.CopyBasicFieldsToVisualTrack(visualtrackDeleted)

	// get stage instance from DB instance, and call callback function
	visualtrackStaged := backRepo.BackRepoVisualTrack.Map_VisualTrackDBID_VisualTrackPtr[visualtrackDB.ID]
	if visualtrackStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), visualtrackStaged, visualtrackDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
