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
var __Satellite__dummysDeclaration__ models.Satellite
var __Satellite_time__dummyDeclaration time.Duration

var mutexSatellite sync.Mutex

// An SatelliteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSatellite updateSatellite deleteSatellite
type SatelliteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SatelliteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSatellite updateSatellite
type SatelliteInput struct {
	// The Satellite to submit or modify
	// in: body
	Satellite *orm.SatelliteAPI
}

// GetSatellites
//
// swagger:route GET /satellites satellites getSatellites
//
// # Get all satellites
//
// Responses:
// default: genericError
//
//	200: satelliteDBResponse
func (controller *Controller) GetSatellites(c *gin.Context) {

	// source slice
	var satelliteDBs []orm.SatelliteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSatellites", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSatellite.GetDB()

	_, err := db.Find(&satelliteDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	satelliteAPIs := make([]orm.SatelliteAPI, 0)

	// for each satellite, update fields from the database nullable fields
	for idx := range satelliteDBs {
		satelliteDB := &satelliteDBs[idx]
		_ = satelliteDB
		var satelliteAPI orm.SatelliteAPI

		// insertion point for updating fields
		satelliteAPI.ID = satelliteDB.ID
		satelliteDB.CopyBasicFieldsToSatellite_WOP(&satelliteAPI.Satellite_WOP)
		satelliteAPI.SatellitePointersEncoding = satelliteDB.SatellitePointersEncoding
		satelliteAPIs = append(satelliteAPIs, satelliteAPI)
	}

	c.JSON(http.StatusOK, satelliteAPIs)
}

// PostSatellite
//
// swagger:route POST /satellites satellites postSatellite
//
// Creates a satellite
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSatellite(c *gin.Context) {

	mutexSatellite.Lock()
	defer mutexSatellite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSatellites", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSatellite.GetDB()

	// Validate input
	var input orm.SatelliteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create satellite
	satelliteDB := orm.SatelliteDB{}
	satelliteDB.SatellitePointersEncoding = input.SatellitePointersEncoding
	satelliteDB.CopyBasicFieldsFromSatellite_WOP(&input.Satellite_WOP)

	_, err = db.Create(&satelliteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSatellite.CheckoutPhaseOneInstance(&satelliteDB)
	satellite := backRepo.BackRepoSatellite.Map_SatelliteDBID_SatellitePtr[satelliteDB.ID]

	if satellite != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), satellite)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, satelliteDB)
}

// GetSatellite
//
// swagger:route GET /satellites/{ID} satellites getSatellite
//
// Gets the details for a satellite.
//
// Responses:
// default: genericError
//
//	200: satelliteDBResponse
func (controller *Controller) GetSatellite(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSatellite", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSatellite.GetDB()

	// Get satelliteDB in DB
	var satelliteDB orm.SatelliteDB
	if _, err := db.First(&satelliteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var satelliteAPI orm.SatelliteAPI
	satelliteAPI.ID = satelliteDB.ID
	satelliteAPI.SatellitePointersEncoding = satelliteDB.SatellitePointersEncoding
	satelliteDB.CopyBasicFieldsToSatellite_WOP(&satelliteAPI.Satellite_WOP)

	c.JSON(http.StatusOK, satelliteAPI)
}

// UpdateSatellite
//
// swagger:route PATCH /satellites/{ID} satellites updateSatellite
//
// # Update a satellite
//
// Responses:
// default: genericError
//
//	200: satelliteDBResponse
func (controller *Controller) UpdateSatellite(c *gin.Context) {

	mutexSatellite.Lock()
	defer mutexSatellite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSatellite", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSatellite.GetDB()

	// Validate input
	var input orm.SatelliteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var satelliteDB orm.SatelliteDB

	// fetch the satellite
	_, err := db.First(&satelliteDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	satelliteDB.CopyBasicFieldsFromSatellite_WOP(&input.Satellite_WOP)
	satelliteDB.SatellitePointersEncoding = input.SatellitePointersEncoding

	db, _ = db.Model(&satelliteDB)
	_, err = db.Updates(&satelliteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	satelliteNew := new(models.Satellite)
	satelliteDB.CopyBasicFieldsToSatellite(satelliteNew)

	// redeem pointers
	satelliteDB.DecodePointers(backRepo, satelliteNew)

	// get stage instance from DB instance, and call callback function
	satelliteOld := backRepo.BackRepoSatellite.Map_SatelliteDBID_SatellitePtr[satelliteDB.ID]
	if satelliteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), satelliteOld, satelliteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the satelliteDB
	c.JSON(http.StatusOK, satelliteDB)
}

// DeleteSatellite
//
// swagger:route DELETE /satellites/{ID} satellites deleteSatellite
//
// # Delete a satellite
//
// default: genericError
//
//	200: satelliteDBResponse
func (controller *Controller) DeleteSatellite(c *gin.Context) {

	mutexSatellite.Lock()
	defer mutexSatellite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSatellite", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSatellite.GetDB()

	// Get model if exist
	var satelliteDB orm.SatelliteDB
	if _, err := db.First(&satelliteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&satelliteDB)

	// get an instance (not staged) from DB instance, and call callback function
	satelliteDeleted := new(models.Satellite)
	satelliteDB.CopyBasicFieldsToSatellite(satelliteDeleted)

	// get stage instance from DB instance, and call callback function
	satelliteStaged := backRepo.BackRepoSatellite.Map_SatelliteDBID_SatellitePtr[satelliteDB.ID]
	if satelliteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), satelliteStaged, satelliteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
