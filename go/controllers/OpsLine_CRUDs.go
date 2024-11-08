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
var __OpsLine__dummysDeclaration__ models.OpsLine
var __OpsLine_time__dummyDeclaration time.Duration

var mutexOpsLine sync.Mutex

// An OpsLineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOpsLine updateOpsLine deleteOpsLine
type OpsLineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OpsLineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOpsLine updateOpsLine
type OpsLineInput struct {
	// The OpsLine to submit or modify
	// in: body
	OpsLine *orm.OpsLineAPI
}

// GetOpsLines
//
// swagger:route GET /opslines opslines getOpsLines
//
// # Get all opslines
//
// Responses:
// default: genericError
//
//	200: opslineDBResponse
func (controller *Controller) GetOpsLines(c *gin.Context) {

	// source slice
	var opslineDBs []orm.OpsLineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOpsLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpsLine.GetDB()

	_, err := db.Find(&opslineDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	opslineAPIs := make([]orm.OpsLineAPI, 0)

	// for each opsline, update fields from the database nullable fields
	for idx := range opslineDBs {
		opslineDB := &opslineDBs[idx]
		_ = opslineDB
		var opslineAPI orm.OpsLineAPI

		// insertion point for updating fields
		opslineAPI.ID = opslineDB.ID
		opslineDB.CopyBasicFieldsToOpsLine_WOP(&opslineAPI.OpsLine_WOP)
		opslineAPI.OpsLinePointersEncoding = opslineDB.OpsLinePointersEncoding
		opslineAPIs = append(opslineAPIs, opslineAPI)
	}

	c.JSON(http.StatusOK, opslineAPIs)
}

// PostOpsLine
//
// swagger:route POST /opslines opslines postOpsLine
//
// Creates a opsline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOpsLine(c *gin.Context) {

	mutexOpsLine.Lock()
	defer mutexOpsLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOpsLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpsLine.GetDB()

	// Validate input
	var input orm.OpsLineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create opsline
	opslineDB := orm.OpsLineDB{}
	opslineDB.OpsLinePointersEncoding = input.OpsLinePointersEncoding
	opslineDB.CopyBasicFieldsFromOpsLine_WOP(&input.OpsLine_WOP)

	_, err = db.Create(&opslineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOpsLine.CheckoutPhaseOneInstance(&opslineDB)
	opsline := backRepo.BackRepoOpsLine.Map_OpsLineDBID_OpsLinePtr[opslineDB.ID]

	if opsline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), opsline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, opslineDB)
}

// GetOpsLine
//
// swagger:route GET /opslines/{ID} opslines getOpsLine
//
// Gets the details for a opsline.
//
// Responses:
// default: genericError
//
//	200: opslineDBResponse
func (controller *Controller) GetOpsLine(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOpsLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpsLine.GetDB()

	// Get opslineDB in DB
	var opslineDB orm.OpsLineDB
	if _, err := db.First(&opslineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var opslineAPI orm.OpsLineAPI
	opslineAPI.ID = opslineDB.ID
	opslineAPI.OpsLinePointersEncoding = opslineDB.OpsLinePointersEncoding
	opslineDB.CopyBasicFieldsToOpsLine_WOP(&opslineAPI.OpsLine_WOP)

	c.JSON(http.StatusOK, opslineAPI)
}

// UpdateOpsLine
//
// swagger:route PATCH /opslines/{ID} opslines updateOpsLine
//
// # Update a opsline
//
// Responses:
// default: genericError
//
//	200: opslineDBResponse
func (controller *Controller) UpdateOpsLine(c *gin.Context) {

	mutexOpsLine.Lock()
	defer mutexOpsLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOpsLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpsLine.GetDB()

	// Validate input
	var input orm.OpsLineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var opslineDB orm.OpsLineDB

	// fetch the opsline
	_, err := db.First(&opslineDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	opslineDB.CopyBasicFieldsFromOpsLine_WOP(&input.OpsLine_WOP)
	opslineDB.OpsLinePointersEncoding = input.OpsLinePointersEncoding

	db, _ = db.Model(&opslineDB)
	_, err = db.Updates(&opslineDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	opslineNew := new(models.OpsLine)
	opslineDB.CopyBasicFieldsToOpsLine(opslineNew)

	// redeem pointers
	opslineDB.DecodePointers(backRepo, opslineNew)

	// get stage instance from DB instance, and call callback function
	opslineOld := backRepo.BackRepoOpsLine.Map_OpsLineDBID_OpsLinePtr[opslineDB.ID]
	if opslineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), opslineOld, opslineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the opslineDB
	c.JSON(http.StatusOK, opslineDB)
}

// DeleteOpsLine
//
// swagger:route DELETE /opslines/{ID} opslines deleteOpsLine
//
// # Delete a opsline
//
// default: genericError
//
//	200: opslineDBResponse
func (controller *Controller) DeleteOpsLine(c *gin.Context) {

	mutexOpsLine.Lock()
	defer mutexOpsLine.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOpsLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpsLine.GetDB()

	// Get model if exist
	var opslineDB orm.OpsLineDB
	if _, err := db.First(&opslineDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&opslineDB)

	// get an instance (not staged) from DB instance, and call callback function
	opslineDeleted := new(models.OpsLine)
	opslineDB.CopyBasicFieldsToOpsLine(opslineDeleted)

	// get stage instance from DB instance, and call callback function
	opslineStaged := backRepo.BackRepoOpsLine.Map_OpsLineDBID_OpsLinePtr[opslineDB.ID]
	if opslineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), opslineStaged, opslineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
