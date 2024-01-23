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
var __GongsimCommand__dummysDeclaration__ models.GongsimCommand
var __GongsimCommand_time__dummyDeclaration time.Duration

var mutexGongsimCommand sync.Mutex

// An GongsimCommandID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongsimCommand updateGongsimCommand deleteGongsimCommand
type GongsimCommandID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongsimCommandInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongsimCommand updateGongsimCommand
type GongsimCommandInput struct {
	// The GongsimCommand to submit or modify
	// in: body
	GongsimCommand *orm.GongsimCommandAPI
}

// GetGongsimCommands
//
// swagger:route GET /gongsimcommands gongsimcommands getGongsimCommands
//
// # Get all gongsimcommands
//
// Responses:
// default: genericError
//
//	200: gongsimcommandDBResponse
func (controller *Controller) GetGongsimCommands(c *gin.Context) {

	// source slice
	var gongsimcommandDBs []orm.GongsimCommandDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongsimCommands", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimCommand.GetDB()

	query := db.Find(&gongsimcommandDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongsimcommandAPIs := make([]orm.GongsimCommandAPI, 0)

	// for each gongsimcommand, update fields from the database nullable fields
	for idx := range gongsimcommandDBs {
		gongsimcommandDB := &gongsimcommandDBs[idx]
		_ = gongsimcommandDB
		var gongsimcommandAPI orm.GongsimCommandAPI

		// insertion point for updating fields
		gongsimcommandAPI.ID = gongsimcommandDB.ID
		gongsimcommandDB.CopyBasicFieldsToGongsimCommand_WOP(&gongsimcommandAPI.GongsimCommand_WOP)
		gongsimcommandAPI.GongsimCommandPointersEncoding = gongsimcommandDB.GongsimCommandPointersEncoding
		gongsimcommandAPIs = append(gongsimcommandAPIs, gongsimcommandAPI)
	}

	c.JSON(http.StatusOK, gongsimcommandAPIs)
}

// PostGongsimCommand
//
// swagger:route POST /gongsimcommands gongsimcommands postGongsimCommand
//
// Creates a gongsimcommand
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongsimCommand(c *gin.Context) {

	mutexGongsimCommand.Lock()
	defer mutexGongsimCommand.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongsimCommands", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimCommand.GetDB()

	// Validate input
	var input orm.GongsimCommandAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongsimcommand
	gongsimcommandDB := orm.GongsimCommandDB{}
	gongsimcommandDB.GongsimCommandPointersEncoding = input.GongsimCommandPointersEncoding
	gongsimcommandDB.CopyBasicFieldsFromGongsimCommand_WOP(&input.GongsimCommand_WOP)

	query := db.Create(&gongsimcommandDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongsimCommand.CheckoutPhaseOneInstance(&gongsimcommandDB)
	gongsimcommand := backRepo.BackRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr[gongsimcommandDB.ID]

	if gongsimcommand != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongsimcommand)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongsimcommandDB)
}

// GetGongsimCommand
//
// swagger:route GET /gongsimcommands/{ID} gongsimcommands getGongsimCommand
//
// Gets the details for a gongsimcommand.
//
// Responses:
// default: genericError
//
//	200: gongsimcommandDBResponse
func (controller *Controller) GetGongsimCommand(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongsimCommand", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimCommand.GetDB()

	// Get gongsimcommandDB in DB
	var gongsimcommandDB orm.GongsimCommandDB
	if err := db.First(&gongsimcommandDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongsimcommandAPI orm.GongsimCommandAPI
	gongsimcommandAPI.ID = gongsimcommandDB.ID
	gongsimcommandAPI.GongsimCommandPointersEncoding = gongsimcommandDB.GongsimCommandPointersEncoding
	gongsimcommandDB.CopyBasicFieldsToGongsimCommand_WOP(&gongsimcommandAPI.GongsimCommand_WOP)

	c.JSON(http.StatusOK, gongsimcommandAPI)
}

// UpdateGongsimCommand
//
// swagger:route PATCH /gongsimcommands/{ID} gongsimcommands updateGongsimCommand
//
// # Update a gongsimcommand
//
// Responses:
// default: genericError
//
//	200: gongsimcommandDBResponse
func (controller *Controller) UpdateGongsimCommand(c *gin.Context) {

	mutexGongsimCommand.Lock()
	defer mutexGongsimCommand.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongsimCommand", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimCommand.GetDB()

	// Validate input
	var input orm.GongsimCommandAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongsimcommandDB orm.GongsimCommandDB

	// fetch the gongsimcommand
	query := db.First(&gongsimcommandDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongsimcommandDB.CopyBasicFieldsFromGongsimCommand_WOP(&input.GongsimCommand_WOP)
	gongsimcommandDB.GongsimCommandPointersEncoding = input.GongsimCommandPointersEncoding

	query = db.Model(&gongsimcommandDB).Updates(gongsimcommandDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongsimcommandNew := new(models.GongsimCommand)
	gongsimcommandDB.CopyBasicFieldsToGongsimCommand(gongsimcommandNew)

	// redeem pointers
	gongsimcommandDB.DecodePointers(backRepo, gongsimcommandNew)

	// get stage instance from DB instance, and call callback function
	gongsimcommandOld := backRepo.BackRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr[gongsimcommandDB.ID]
	if gongsimcommandOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongsimcommandOld, gongsimcommandNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongsimcommandDB
	c.JSON(http.StatusOK, gongsimcommandDB)
}

// DeleteGongsimCommand
//
// swagger:route DELETE /gongsimcommands/{ID} gongsimcommands deleteGongsimCommand
//
// # Delete a gongsimcommand
//
// default: genericError
//
//	200: gongsimcommandDBResponse
func (controller *Controller) DeleteGongsimCommand(c *gin.Context) {

	mutexGongsimCommand.Lock()
	defer mutexGongsimCommand.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongsimCommand", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGongsimCommand.GetDB()

	// Get model if exist
	var gongsimcommandDB orm.GongsimCommandDB
	if err := db.First(&gongsimcommandDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongsimcommandDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongsimcommandDeleted := new(models.GongsimCommand)
	gongsimcommandDB.CopyBasicFieldsToGongsimCommand(gongsimcommandDeleted)

	// get stage instance from DB instance, and call callback function
	gongsimcommandStaged := backRepo.BackRepoGongsimCommand.Map_GongsimCommandDBID_GongsimCommandPtr[gongsimcommandDB.ID]
	if gongsimcommandStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongsimcommandStaged, gongsimcommandDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
