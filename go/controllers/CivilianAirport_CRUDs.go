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
var __CivilianAirport__dummysDeclaration__ models.CivilianAirport
var __CivilianAirport_time__dummyDeclaration time.Duration

var mutexCivilianAirport sync.Mutex

// An CivilianAirportID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCivilianAirport updateCivilianAirport deleteCivilianAirport
type CivilianAirportID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CivilianAirportInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCivilianAirport updateCivilianAirport
type CivilianAirportInput struct {
	// The CivilianAirport to submit or modify
	// in: body
	CivilianAirport *orm.CivilianAirportAPI
}

// GetCivilianAirports
//
// swagger:route GET /civilianairports civilianairports getCivilianAirports
//
// # Get all civilianairports
//
// Responses:
// default: genericError
//
//	200: civilianairportDBResponse
func (controller *Controller) GetCivilianAirports(c *gin.Context) {

	// source slice
	var civilianairportDBs []orm.CivilianAirportDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCivilianAirports", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCivilianAirport.GetDB()

	_, err := db.Find(&civilianairportDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	civilianairportAPIs := make([]orm.CivilianAirportAPI, 0)

	// for each civilianairport, update fields from the database nullable fields
	for idx := range civilianairportDBs {
		civilianairportDB := &civilianairportDBs[idx]
		_ = civilianairportDB
		var civilianairportAPI orm.CivilianAirportAPI

		// insertion point for updating fields
		civilianairportAPI.ID = civilianairportDB.ID
		civilianairportDB.CopyBasicFieldsToCivilianAirport_WOP(&civilianairportAPI.CivilianAirport_WOP)
		civilianairportAPI.CivilianAirportPointersEncoding = civilianairportDB.CivilianAirportPointersEncoding
		civilianairportAPIs = append(civilianairportAPIs, civilianairportAPI)
	}

	c.JSON(http.StatusOK, civilianairportAPIs)
}

// PostCivilianAirport
//
// swagger:route POST /civilianairports civilianairports postCivilianAirport
//
// Creates a civilianairport
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCivilianAirport(c *gin.Context) {

	mutexCivilianAirport.Lock()
	defer mutexCivilianAirport.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCivilianAirports", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCivilianAirport.GetDB()

	// Validate input
	var input orm.CivilianAirportAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create civilianairport
	civilianairportDB := orm.CivilianAirportDB{}
	civilianairportDB.CivilianAirportPointersEncoding = input.CivilianAirportPointersEncoding
	civilianairportDB.CopyBasicFieldsFromCivilianAirport_WOP(&input.CivilianAirport_WOP)

	_, err = db.Create(&civilianairportDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCivilianAirport.CheckoutPhaseOneInstance(&civilianairportDB)
	civilianairport := backRepo.BackRepoCivilianAirport.Map_CivilianAirportDBID_CivilianAirportPtr[civilianairportDB.ID]

	if civilianairport != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), civilianairport)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, civilianairportDB)
}

// GetCivilianAirport
//
// swagger:route GET /civilianairports/{ID} civilianairports getCivilianAirport
//
// Gets the details for a civilianairport.
//
// Responses:
// default: genericError
//
//	200: civilianairportDBResponse
func (controller *Controller) GetCivilianAirport(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCivilianAirport", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCivilianAirport.GetDB()

	// Get civilianairportDB in DB
	var civilianairportDB orm.CivilianAirportDB
	if _, err := db.First(&civilianairportDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var civilianairportAPI orm.CivilianAirportAPI
	civilianairportAPI.ID = civilianairportDB.ID
	civilianairportAPI.CivilianAirportPointersEncoding = civilianairportDB.CivilianAirportPointersEncoding
	civilianairportDB.CopyBasicFieldsToCivilianAirport_WOP(&civilianairportAPI.CivilianAirport_WOP)

	c.JSON(http.StatusOK, civilianairportAPI)
}

// UpdateCivilianAirport
//
// swagger:route PATCH /civilianairports/{ID} civilianairports updateCivilianAirport
//
// # Update a civilianairport
//
// Responses:
// default: genericError
//
//	200: civilianairportDBResponse
func (controller *Controller) UpdateCivilianAirport(c *gin.Context) {

	mutexCivilianAirport.Lock()
	defer mutexCivilianAirport.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCivilianAirport", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCivilianAirport.GetDB()

	// Validate input
	var input orm.CivilianAirportAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var civilianairportDB orm.CivilianAirportDB

	// fetch the civilianairport
	_, err := db.First(&civilianairportDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	civilianairportDB.CopyBasicFieldsFromCivilianAirport_WOP(&input.CivilianAirport_WOP)
	civilianairportDB.CivilianAirportPointersEncoding = input.CivilianAirportPointersEncoding

	db, _ = db.Model(&civilianairportDB)
	_, err = db.Updates(civilianairportDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	civilianairportNew := new(models.CivilianAirport)
	civilianairportDB.CopyBasicFieldsToCivilianAirport(civilianairportNew)

	// redeem pointers
	civilianairportDB.DecodePointers(backRepo, civilianairportNew)

	// get stage instance from DB instance, and call callback function
	civilianairportOld := backRepo.BackRepoCivilianAirport.Map_CivilianAirportDBID_CivilianAirportPtr[civilianairportDB.ID]
	if civilianairportOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), civilianairportOld, civilianairportNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the civilianairportDB
	c.JSON(http.StatusOK, civilianairportDB)
}

// DeleteCivilianAirport
//
// swagger:route DELETE /civilianairports/{ID} civilianairports deleteCivilianAirport
//
// # Delete a civilianairport
//
// default: genericError
//
//	200: civilianairportDBResponse
func (controller *Controller) DeleteCivilianAirport(c *gin.Context) {

	mutexCivilianAirport.Lock()
	defer mutexCivilianAirport.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCivilianAirport", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCivilianAirport.GetDB()

	// Get model if exist
	var civilianairportDB orm.CivilianAirportDB
	if _, err := db.First(&civilianairportDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&civilianairportDB)

	// get an instance (not staged) from DB instance, and call callback function
	civilianairportDeleted := new(models.CivilianAirport)
	civilianairportDB.CopyBasicFieldsToCivilianAirport(civilianairportDeleted)

	// get stage instance from DB instance, and call callback function
	civilianairportStaged := backRepo.BackRepoCivilianAirport.Map_CivilianAirportDBID_CivilianAirportPtr[civilianairportDB.ID]
	if civilianairportStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), civilianairportStaged, civilianairportDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
