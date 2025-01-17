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
var __LayerGroupUse__dummysDeclaration__ models.LayerGroupUse
var __LayerGroupUse_time__dummyDeclaration time.Duration

var mutexLayerGroupUse sync.Mutex

// An LayerGroupUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLayerGroupUse updateLayerGroupUse deleteLayerGroupUse
type LayerGroupUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayerGroupUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLayerGroupUse updateLayerGroupUse
type LayerGroupUseInput struct {
	// The LayerGroupUse to submit or modify
	// in: body
	LayerGroupUse *orm.LayerGroupUseAPI
}

// GetLayerGroupUses
//
// swagger:route GET /layergroupuses layergroupuses getLayerGroupUses
//
// # Get all layergroupuses
//
// Responses:
// default: genericError
//
//	200: layergroupuseDBResponse
func (controller *Controller) GetLayerGroupUses(c *gin.Context) {

	// source slice
	var layergroupuseDBs []orm.LayerGroupUseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayerGroupUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroupUse.GetDB()

	_, err := db.Find(&layergroupuseDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	layergroupuseAPIs := make([]orm.LayerGroupUseAPI, 0)

	// for each layergroupuse, update fields from the database nullable fields
	for idx := range layergroupuseDBs {
		layergroupuseDB := &layergroupuseDBs[idx]
		_ = layergroupuseDB
		var layergroupuseAPI orm.LayerGroupUseAPI

		// insertion point for updating fields
		layergroupuseAPI.ID = layergroupuseDB.ID
		layergroupuseDB.CopyBasicFieldsToLayerGroupUse_WOP(&layergroupuseAPI.LayerGroupUse_WOP)
		layergroupuseAPI.LayerGroupUsePointersEncoding = layergroupuseDB.LayerGroupUsePointersEncoding
		layergroupuseAPIs = append(layergroupuseAPIs, layergroupuseAPI)
	}

	c.JSON(http.StatusOK, layergroupuseAPIs)
}

// PostLayerGroupUse
//
// swagger:route POST /layergroupuses layergroupuses postLayerGroupUse
//
// Creates a layergroupuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLayerGroupUse(c *gin.Context) {

	mutexLayerGroupUse.Lock()
	defer mutexLayerGroupUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLayerGroupUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroupUse.GetDB()

	// Validate input
	var input orm.LayerGroupUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create layergroupuse
	layergroupuseDB := orm.LayerGroupUseDB{}
	layergroupuseDB.LayerGroupUsePointersEncoding = input.LayerGroupUsePointersEncoding
	layergroupuseDB.CopyBasicFieldsFromLayerGroupUse_WOP(&input.LayerGroupUse_WOP)

	_, err = db.Create(&layergroupuseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLayerGroupUse.CheckoutPhaseOneInstance(&layergroupuseDB)
	layergroupuse := backRepo.BackRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr[layergroupuseDB.ID]

	if layergroupuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), layergroupuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, layergroupuseDB)
}

// GetLayerGroupUse
//
// swagger:route GET /layergroupuses/{ID} layergroupuses getLayerGroupUse
//
// Gets the details for a layergroupuse.
//
// Responses:
// default: genericError
//
//	200: layergroupuseDBResponse
func (controller *Controller) GetLayerGroupUse(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayerGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroupUse.GetDB()

	// Get layergroupuseDB in DB
	var layergroupuseDB orm.LayerGroupUseDB
	if _, err := db.First(&layergroupuseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var layergroupuseAPI orm.LayerGroupUseAPI
	layergroupuseAPI.ID = layergroupuseDB.ID
	layergroupuseAPI.LayerGroupUsePointersEncoding = layergroupuseDB.LayerGroupUsePointersEncoding
	layergroupuseDB.CopyBasicFieldsToLayerGroupUse_WOP(&layergroupuseAPI.LayerGroupUse_WOP)

	c.JSON(http.StatusOK, layergroupuseAPI)
}

// UpdateLayerGroupUse
//
// swagger:route PATCH /layergroupuses/{ID} layergroupuses updateLayerGroupUse
//
// # Update a layergroupuse
//
// Responses:
// default: genericError
//
//	200: layergroupuseDBResponse
func (controller *Controller) UpdateLayerGroupUse(c *gin.Context) {

	mutexLayerGroupUse.Lock()
	defer mutexLayerGroupUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLayerGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroupUse.GetDB()

	// Validate input
	var input orm.LayerGroupUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layergroupuseDB orm.LayerGroupUseDB

	// fetch the layergroupuse
	_, err := db.First(&layergroupuseDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layergroupuseDB.CopyBasicFieldsFromLayerGroupUse_WOP(&input.LayerGroupUse_WOP)
	layergroupuseDB.LayerGroupUsePointersEncoding = input.LayerGroupUsePointersEncoding

	db, _ = db.Model(&layergroupuseDB)
	_, err = db.Updates(&layergroupuseDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layergroupuseNew := new(models.LayerGroupUse)
	layergroupuseDB.CopyBasicFieldsToLayerGroupUse(layergroupuseNew)

	// redeem pointers
	layergroupuseDB.DecodePointers(backRepo, layergroupuseNew)

	// get stage instance from DB instance, and call callback function
	layergroupuseOld := backRepo.BackRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr[layergroupuseDB.ID]
	if layergroupuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), layergroupuseOld, layergroupuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layergroupuseDB
	c.JSON(http.StatusOK, layergroupuseDB)
}

// DeleteLayerGroupUse
//
// swagger:route DELETE /layergroupuses/{ID} layergroupuses deleteLayerGroupUse
//
// # Delete a layergroupuse
//
// default: genericError
//
//	200: layergroupuseDBResponse
func (controller *Controller) DeleteLayerGroupUse(c *gin.Context) {

	mutexLayerGroupUse.Lock()
	defer mutexLayerGroupUse.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLayerGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroupUse.GetDB()

	// Get model if exist
	var layergroupuseDB orm.LayerGroupUseDB
	if _, err := db.First(&layergroupuseDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&layergroupuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	layergroupuseDeleted := new(models.LayerGroupUse)
	layergroupuseDB.CopyBasicFieldsToLayerGroupUse(layergroupuseDeleted)

	// get stage instance from DB instance, and call callback function
	layergroupuseStaged := backRepo.BackRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr[layergroupuseDB.ID]
	if layergroupuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), layergroupuseStaged, layergroupuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
