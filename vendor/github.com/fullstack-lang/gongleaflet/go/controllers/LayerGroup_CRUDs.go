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
var __LayerGroup__dummysDeclaration__ models.LayerGroup
var __LayerGroup_time__dummyDeclaration time.Duration

var mutexLayerGroup sync.Mutex

// An LayerGroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLayerGroup updateLayerGroup deleteLayerGroup
type LayerGroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayerGroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLayerGroup updateLayerGroup
type LayerGroupInput struct {
	// The LayerGroup to submit or modify
	// in: body
	LayerGroup *orm.LayerGroupAPI
}

// GetLayerGroups
//
// swagger:route GET /layergroups layergroups getLayerGroups
//
// # Get all layergroups
//
// Responses:
// default: genericError
//
//	200: layergroupDBResponse
func (controller *Controller) GetLayerGroups(c *gin.Context) {

	// source slice
	var layergroupDBs []orm.LayerGroupDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayerGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroup.GetDB()

	query := db.Find(&layergroupDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	layergroupAPIs := make([]orm.LayerGroupAPI, 0)

	// for each layergroup, update fields from the database nullable fields
	for idx := range layergroupDBs {
		layergroupDB := &layergroupDBs[idx]
		_ = layergroupDB
		var layergroupAPI orm.LayerGroupAPI

		// insertion point for updating fields
		layergroupAPI.ID = layergroupDB.ID
		layergroupDB.CopyBasicFieldsToLayerGroup_WOP(&layergroupAPI.LayerGroup_WOP)
		layergroupAPI.LayerGroupPointersEncoding = layergroupDB.LayerGroupPointersEncoding
		layergroupAPIs = append(layergroupAPIs, layergroupAPI)
	}

	c.JSON(http.StatusOK, layergroupAPIs)
}

// PostLayerGroup
//
// swagger:route POST /layergroups layergroups postLayerGroup
//
// Creates a layergroup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLayerGroup(c *gin.Context) {

	mutexLayerGroup.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLayerGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroup.GetDB()

	// Validate input
	var input orm.LayerGroupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create layergroup
	layergroupDB := orm.LayerGroupDB{}
	layergroupDB.LayerGroupPointersEncoding = input.LayerGroupPointersEncoding
	layergroupDB.CopyBasicFieldsFromLayerGroup_WOP(&input.LayerGroup_WOP)

	query := db.Create(&layergroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLayerGroup.CheckoutPhaseOneInstance(&layergroupDB)
	layergroup := backRepo.BackRepoLayerGroup.Map_LayerGroupDBID_LayerGroupPtr[layergroupDB.ID]

	if layergroup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), layergroup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, layergroupDB)

	mutexLayerGroup.Unlock()
}

// GetLayerGroup
//
// swagger:route GET /layergroups/{ID} layergroups getLayerGroup
//
// Gets the details for a layergroup.
//
// Responses:
// default: genericError
//
//	200: layergroupDBResponse
func (controller *Controller) GetLayerGroup(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayerGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroup.GetDB()

	// Get layergroupDB in DB
	var layergroupDB orm.LayerGroupDB
	if err := db.First(&layergroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var layergroupAPI orm.LayerGroupAPI
	layergroupAPI.ID = layergroupDB.ID
	layergroupAPI.LayerGroupPointersEncoding = layergroupDB.LayerGroupPointersEncoding
	layergroupDB.CopyBasicFieldsToLayerGroup_WOP(&layergroupAPI.LayerGroup_WOP)

	c.JSON(http.StatusOK, layergroupAPI)
}

// UpdateLayerGroup
//
// swagger:route PATCH /layergroups/{ID} layergroups updateLayerGroup
//
// # Update a layergroup
//
// Responses:
// default: genericError
//
//	200: layergroupDBResponse
func (controller *Controller) UpdateLayerGroup(c *gin.Context) {

	mutexLayerGroup.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLayerGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroup.GetDB()

	// Validate input
	var input orm.LayerGroupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layergroupDB orm.LayerGroupDB

	// fetch the layergroup
	query := db.First(&layergroupDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layergroupDB.CopyBasicFieldsFromLayerGroup_WOP(&input.LayerGroup_WOP)
	layergroupDB.LayerGroupPointersEncoding = input.LayerGroupPointersEncoding

	query = db.Model(&layergroupDB).Updates(layergroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layergroupNew := new(models.LayerGroup)
	layergroupDB.CopyBasicFieldsToLayerGroup(layergroupNew)

	// redeem pointers
	layergroupDB.DecodePointers(backRepo, layergroupNew)

	// get stage instance from DB instance, and call callback function
	layergroupOld := backRepo.BackRepoLayerGroup.Map_LayerGroupDBID_LayerGroupPtr[layergroupDB.ID]
	if layergroupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), layergroupOld, layergroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layergroupDB
	c.JSON(http.StatusOK, layergroupDB)

	mutexLayerGroup.Unlock()
}

// DeleteLayerGroup
//
// swagger:route DELETE /layergroups/{ID} layergroups deleteLayerGroup
//
// # Delete a layergroup
//
// default: genericError
//
//	200: layergroupDBResponse
func (controller *Controller) DeleteLayerGroup(c *gin.Context) {

	mutexLayerGroup.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLayerGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLayerGroup.GetDB()

	// Get model if exist
	var layergroupDB orm.LayerGroupDB
	if err := db.First(&layergroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&layergroupDB)

	// get an instance (not staged) from DB instance, and call callback function
	layergroupDeleted := new(models.LayerGroup)
	layergroupDB.CopyBasicFieldsToLayerGroup(layergroupDeleted)

	// get stage instance from DB instance, and call callback function
	layergroupStaged := backRepo.BackRepoLayerGroup.Map_LayerGroupDBID_LayerGroupPtr[layergroupDB.ID]
	if layergroupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), layergroupStaged, layergroupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLayerGroup.Unlock()
}
