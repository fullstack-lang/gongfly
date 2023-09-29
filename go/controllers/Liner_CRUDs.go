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
var __Liner__dummysDeclaration__ models.Liner
var __Liner_time__dummyDeclaration time.Duration

var mutexLiner sync.Mutex

// An LinerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLiner updateLiner deleteLiner
type LinerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLiner updateLiner
type LinerInput struct {
	// The Liner to submit or modify
	// in: body
	Liner *orm.LinerAPI
}

// GetLiners
//
// swagger:route GET /liners liners getLiners
//
// # Get all liners
//
// Responses:
// default: genericError
//
//	200: linerDBResponse
func (controller *Controller) GetLiners(c *gin.Context) {

	// source slice
	var linerDBs []orm.LinerDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLiners", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLiner.GetDB()

	query := db.Find(&linerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linerAPIs := make([]orm.LinerAPI, 0)

	// for each liner, update fields from the database nullable fields
	for idx := range linerDBs {
		linerDB := &linerDBs[idx]
		_ = linerDB
		var linerAPI orm.LinerAPI

		// insertion point for updating fields
		linerAPI.ID = linerDB.ID
		linerDB.CopyBasicFieldsToLiner(&linerAPI.Liner)
		linerAPI.LinerPointersEnconding = linerDB.LinerPointersEnconding
		linerAPIs = append(linerAPIs, linerAPI)
	}

	c.JSON(http.StatusOK, linerAPIs)
}

// PostLiner
//
// swagger:route POST /liners liners postLiner
//
// Creates a liner
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLiner(c *gin.Context) {

	mutexLiner.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLiners", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLiner.GetDB()

	// Validate input
	var input orm.LinerAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create liner
	linerDB := orm.LinerDB{}
	linerDB.LinerPointersEnconding = input.LinerPointersEnconding
	linerDB.CopyBasicFieldsFromLiner(&input.Liner)

	query := db.Create(&linerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLiner.CheckoutPhaseOneInstance(&linerDB)
	liner := backRepo.BackRepoLiner.Map_LinerDBID_LinerPtr[linerDB.ID]

	if liner != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), liner)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, linerDB)

	mutexLiner.Unlock()
}

// GetLiner
//
// swagger:route GET /liners/{ID} liners getLiner
//
// Gets the details for a liner.
//
// Responses:
// default: genericError
//
//	200: linerDBResponse
func (controller *Controller) GetLiner(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLiner", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLiner.GetDB()

	// Get linerDB in DB
	var linerDB orm.LinerDB
	if err := db.First(&linerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var linerAPI orm.LinerAPI
	linerAPI.ID = linerDB.ID
	linerAPI.LinerPointersEnconding = linerDB.LinerPointersEnconding
	linerDB.CopyBasicFieldsToLiner(&linerAPI.Liner)

	c.JSON(http.StatusOK, linerAPI)
}

// UpdateLiner
//
// swagger:route PATCH /liners/{ID} liners updateLiner
//
// # Update a liner
//
// Responses:
// default: genericError
//
//	200: linerDBResponse
func (controller *Controller) UpdateLiner(c *gin.Context) {

	mutexLiner.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLiner", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLiner.GetDB()

	// Validate input
	var input orm.LinerAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var linerDB orm.LinerDB

	// fetch the liner
	query := db.First(&linerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	linerDB.CopyBasicFieldsFromLiner(&input.Liner)
	linerDB.LinerPointersEnconding = input.LinerPointersEnconding

	query = db.Model(&linerDB).Updates(linerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linerNew := new(models.Liner)
	linerDB.CopyBasicFieldsToLiner(linerNew)

	// get stage instance from DB instance, and call callback function
	linerOld := backRepo.BackRepoLiner.Map_LinerDBID_LinerPtr[linerDB.ID]
	if linerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), linerOld, linerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linerDB
	c.JSON(http.StatusOK, linerDB)

	mutexLiner.Unlock()
}

// DeleteLiner
//
// swagger:route DELETE /liners/{ID} liners deleteLiner
//
// # Delete a liner
//
// default: genericError
//
//	200: linerDBResponse
func (controller *Controller) DeleteLiner(c *gin.Context) {

	mutexLiner.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLiner", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLiner.GetDB()

	// Get model if exist
	var linerDB orm.LinerDB
	if err := db.First(&linerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&linerDB)

	// get an instance (not staged) from DB instance, and call callback function
	linerDeleted := new(models.Liner)
	linerDB.CopyBasicFieldsToLiner(linerDeleted)

	// get stage instance from DB instance, and call callback function
	linerStaged := backRepo.BackRepoLiner.Map_LinerDBID_LinerPtr[linerDB.ID]
	if linerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), linerStaged, linerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLiner.Unlock()
}
