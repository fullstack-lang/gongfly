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
var __DivIcon__dummysDeclaration__ models.DivIcon
var __DivIcon_time__dummyDeclaration time.Duration

var mutexDivIcon sync.Mutex

// An DivIconID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDivIcon updateDivIcon deleteDivIcon
type DivIconID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DivIconInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDivIcon updateDivIcon
type DivIconInput struct {
	// The DivIcon to submit or modify
	// in: body
	DivIcon *orm.DivIconAPI
}

// GetDivIcons
//
// swagger:route GET /divicons divicons getDivIcons
//
// # Get all divicons
//
// Responses:
// default: genericError
//
//	200: diviconDBResponse
func (controller *Controller) GetDivIcons(c *gin.Context) {

	// source slice
	var diviconDBs []orm.DivIconDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDivIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDivIcon.GetDB()

	query := db.Find(&diviconDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	diviconAPIs := make([]orm.DivIconAPI, 0)

	// for each divicon, update fields from the database nullable fields
	for idx := range diviconDBs {
		diviconDB := &diviconDBs[idx]
		_ = diviconDB
		var diviconAPI orm.DivIconAPI

		// insertion point for updating fields
		diviconAPI.ID = diviconDB.ID
		diviconDB.CopyBasicFieldsToDivIcon_WOP(&diviconAPI.DivIcon_WOP)
		diviconAPI.DivIconPointersEncoding = diviconDB.DivIconPointersEncoding
		diviconAPIs = append(diviconAPIs, diviconAPI)
	}

	c.JSON(http.StatusOK, diviconAPIs)
}

// PostDivIcon
//
// swagger:route POST /divicons divicons postDivIcon
//
// Creates a divicon
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDivIcon(c *gin.Context) {

	mutexDivIcon.Lock()
	defer mutexDivIcon.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDivIcons", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDivIcon.GetDB()

	// Validate input
	var input orm.DivIconAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create divicon
	diviconDB := orm.DivIconDB{}
	diviconDB.DivIconPointersEncoding = input.DivIconPointersEncoding
	diviconDB.CopyBasicFieldsFromDivIcon_WOP(&input.DivIcon_WOP)

	query := db.Create(&diviconDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDivIcon.CheckoutPhaseOneInstance(&diviconDB)
	divicon := backRepo.BackRepoDivIcon.Map_DivIconDBID_DivIconPtr[diviconDB.ID]

	if divicon != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), divicon)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, diviconDB)
}

// GetDivIcon
//
// swagger:route GET /divicons/{ID} divicons getDivIcon
//
// Gets the details for a divicon.
//
// Responses:
// default: genericError
//
//	200: diviconDBResponse
func (controller *Controller) GetDivIcon(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDivIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDivIcon.GetDB()

	// Get diviconDB in DB
	var diviconDB orm.DivIconDB
	if err := db.First(&diviconDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var diviconAPI orm.DivIconAPI
	diviconAPI.ID = diviconDB.ID
	diviconAPI.DivIconPointersEncoding = diviconDB.DivIconPointersEncoding
	diviconDB.CopyBasicFieldsToDivIcon_WOP(&diviconAPI.DivIcon_WOP)

	c.JSON(http.StatusOK, diviconAPI)
}

// UpdateDivIcon
//
// swagger:route PATCH /divicons/{ID} divicons updateDivIcon
//
// # Update a divicon
//
// Responses:
// default: genericError
//
//	200: diviconDBResponse
func (controller *Controller) UpdateDivIcon(c *gin.Context) {

	mutexDivIcon.Lock()
	defer mutexDivIcon.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDivIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDivIcon.GetDB()

	// Validate input
	var input orm.DivIconAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var diviconDB orm.DivIconDB

	// fetch the divicon
	query := db.First(&diviconDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	diviconDB.CopyBasicFieldsFromDivIcon_WOP(&input.DivIcon_WOP)
	diviconDB.DivIconPointersEncoding = input.DivIconPointersEncoding

	query = db.Model(&diviconDB).Updates(diviconDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	diviconNew := new(models.DivIcon)
	diviconDB.CopyBasicFieldsToDivIcon(diviconNew)

	// redeem pointers
	diviconDB.DecodePointers(backRepo, diviconNew)

	// get stage instance from DB instance, and call callback function
	diviconOld := backRepo.BackRepoDivIcon.Map_DivIconDBID_DivIconPtr[diviconDB.ID]
	if diviconOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), diviconOld, diviconNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the diviconDB
	c.JSON(http.StatusOK, diviconDB)
}

// DeleteDivIcon
//
// swagger:route DELETE /divicons/{ID} divicons deleteDivIcon
//
// # Delete a divicon
//
// default: genericError
//
//	200: diviconDBResponse
func (controller *Controller) DeleteDivIcon(c *gin.Context) {

	mutexDivIcon.Lock()
	defer mutexDivIcon.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDivIcon", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDivIcon.GetDB()

	// Get model if exist
	var diviconDB orm.DivIconDB
	if err := db.First(&diviconDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&diviconDB)

	// get an instance (not staged) from DB instance, and call callback function
	diviconDeleted := new(models.DivIcon)
	diviconDB.CopyBasicFieldsToDivIcon(diviconDeleted)

	// get stage instance from DB instance, and call callback function
	diviconStaged := backRepo.BackRepoDivIcon.Map_DivIconDBID_DivIconPtr[diviconDB.ID]
	if diviconStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), diviconStaged, diviconDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
