// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsim/go/models"
	"github.com/fullstack-lang/gongsim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongsimStatus__dummysDeclaration__ models.GongsimStatus
var __GongsimStatus_time__dummyDeclaration time.Duration

var mutexGongsimStatus sync.Mutex

// An GongsimStatusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongsimStatus updateGongsimStatus deleteGongsimStatus
type GongsimStatusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongsimStatusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongsimStatus updateGongsimStatus
type GongsimStatusInput struct {
	// The GongsimStatus to submit or modify
	// in: body
	GongsimStatus *orm.GongsimStatusAPI
}

// GetGongsimStatuss
//
// swagger:route GET /gongsimstatuss gongsimstatuss getGongsimStatuss
//
// # Get all gongsimstatuss
//
// Responses:
// default: genericError
//
//	200: gongsimstatusDBResponse
func (controller *Controller) GetGongsimStatuss(c *gin.Context) {

	// source slice
	var gongsimstatusDBs []orm.GongsimStatusDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongsimStatuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimStatus.GetDB()

	_, err := db.Find(&gongsimstatusDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongsimstatusAPIs := make([]orm.GongsimStatusAPI, 0)

	// for each gongsimstatus, update fields from the database nullable fields
	for idx := range gongsimstatusDBs {
		gongsimstatusDB := &gongsimstatusDBs[idx]
		_ = gongsimstatusDB
		var gongsimstatusAPI orm.GongsimStatusAPI

		// insertion point for updating fields
		gongsimstatusAPI.ID = gongsimstatusDB.ID
		gongsimstatusDB.CopyBasicFieldsToGongsimStatus_WOP(&gongsimstatusAPI.GongsimStatus_WOP)
		gongsimstatusAPI.GongsimStatusPointersEncoding = gongsimstatusDB.GongsimStatusPointersEncoding
		gongsimstatusAPIs = append(gongsimstatusAPIs, gongsimstatusAPI)
	}

	c.JSON(http.StatusOK, gongsimstatusAPIs)
}

// PostGongsimStatus
//
// swagger:route POST /gongsimstatuss gongsimstatuss postGongsimStatus
//
// Creates a gongsimstatus
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongsimStatus(c *gin.Context) {

	mutexGongsimStatus.Lock()
	defer mutexGongsimStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongsimStatuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimStatus.GetDB()

	// Validate input
	var input orm.GongsimStatusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongsimstatus
	gongsimstatusDB := orm.GongsimStatusDB{}
	gongsimstatusDB.GongsimStatusPointersEncoding = input.GongsimStatusPointersEncoding
	gongsimstatusDB.CopyBasicFieldsFromGongsimStatus_WOP(&input.GongsimStatus_WOP)

	_, err = db.Create(&gongsimstatusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongsimStatus.CheckoutPhaseOneInstance(&gongsimstatusDB)
	gongsimstatus := backRepo.BackRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr[gongsimstatusDB.ID]

	if gongsimstatus != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongsimstatus)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongsimstatusDB)
}

// GetGongsimStatus
//
// swagger:route GET /gongsimstatuss/{ID} gongsimstatuss getGongsimStatus
//
// Gets the details for a gongsimstatus.
//
// Responses:
// default: genericError
//
//	200: gongsimstatusDBResponse
func (controller *Controller) GetGongsimStatus(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongsimStatus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimStatus.GetDB()

	// Get gongsimstatusDB in DB
	var gongsimstatusDB orm.GongsimStatusDB
	if _, err := db.First(&gongsimstatusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongsimstatusAPI orm.GongsimStatusAPI
	gongsimstatusAPI.ID = gongsimstatusDB.ID
	gongsimstatusAPI.GongsimStatusPointersEncoding = gongsimstatusDB.GongsimStatusPointersEncoding
	gongsimstatusDB.CopyBasicFieldsToGongsimStatus_WOP(&gongsimstatusAPI.GongsimStatus_WOP)

	c.JSON(http.StatusOK, gongsimstatusAPI)
}

// UpdateGongsimStatus
//
// swagger:route PATCH /gongsimstatuss/{ID} gongsimstatuss updateGongsimStatus
//
// # Update a gongsimstatus
//
// Responses:
// default: genericError
//
//	200: gongsimstatusDBResponse
func (controller *Controller) UpdateGongsimStatus(c *gin.Context) {

	mutexGongsimStatus.Lock()
	defer mutexGongsimStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongsimStatus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimStatus.GetDB()

	// Validate input
	var input orm.GongsimStatusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongsimstatusDB orm.GongsimStatusDB

	// fetch the gongsimstatus
	_, err := db.First(&gongsimstatusDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongsimstatusDB.CopyBasicFieldsFromGongsimStatus_WOP(&input.GongsimStatus_WOP)
	gongsimstatusDB.GongsimStatusPointersEncoding = input.GongsimStatusPointersEncoding

	db, _ = db.Model(&gongsimstatusDB)
	_, err = db.Updates(&gongsimstatusDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongsimstatusNew := new(models.GongsimStatus)
	gongsimstatusDB.CopyBasicFieldsToGongsimStatus(gongsimstatusNew)

	// redeem pointers
	gongsimstatusDB.DecodePointers(backRepo, gongsimstatusNew)

	// get stage instance from DB instance, and call callback function
	gongsimstatusOld := backRepo.BackRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr[gongsimstatusDB.ID]
	if gongsimstatusOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongsimstatusOld, gongsimstatusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongsimstatusDB
	c.JSON(http.StatusOK, gongsimstatusDB)
}

// DeleteGongsimStatus
//
// swagger:route DELETE /gongsimstatuss/{ID} gongsimstatuss deleteGongsimStatus
//
// # Delete a gongsimstatus
//
// default: genericError
//
//	200: gongsimstatusDBResponse
func (controller *Controller) DeleteGongsimStatus(c *gin.Context) {

	mutexGongsimStatus.Lock()
	defer mutexGongsimStatus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongsimStatus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimStatus.GetDB()

	// Get model if exist
	var gongsimstatusDB orm.GongsimStatusDB
	if _, err := db.First(&gongsimstatusDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongsimstatusDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongsimstatusDeleted := new(models.GongsimStatus)
	gongsimstatusDB.CopyBasicFieldsToGongsimStatus(gongsimstatusDeleted)

	// get stage instance from DB instance, and call callback function
	gongsimstatusStaged := backRepo.BackRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr[gongsimstatusDB.ID]
	if gongsimstatusStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongsimstatusStaged, gongsimstatusDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
