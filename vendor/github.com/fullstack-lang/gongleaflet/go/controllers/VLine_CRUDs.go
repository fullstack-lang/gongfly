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
var __VLine__dummysDeclaration__ models.VLine
var __VLine_time__dummyDeclaration time.Duration

var mutexVLine sync.Mutex

// An VLineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVLine updateVLine deleteVLine
type VLineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VLineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVLine updateVLine
type VLineInput struct {
	// The VLine to submit or modify
	// in: body
	VLine *orm.VLineAPI
}

// GetVLines
//
// swagger:route GET /vlines vlines getVLines
//
// # Get all vlines
//
// Responses:
// default: genericError
//
//	200: vlineDBResponse
func (controller *Controller) GetVLines(c *gin.Context) {

	// source slice
	var vlineDBs []orm.VLineDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVLine.GetDB()

	query := db.Find(&vlineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	vlineAPIs := make([]orm.VLineAPI, 0)

	// for each vline, update fields from the database nullable fields
	for idx := range vlineDBs {
		vlineDB := &vlineDBs[idx]
		_ = vlineDB
		var vlineAPI orm.VLineAPI

		// insertion point for updating fields
		vlineAPI.ID = vlineDB.ID
		vlineDB.CopyBasicFieldsToVLine_WOP(&vlineAPI.VLine_WOP)
		vlineAPI.VLinePointersEncoding = vlineDB.VLinePointersEncoding
		vlineAPIs = append(vlineAPIs, vlineAPI)
	}

	c.JSON(http.StatusOK, vlineAPIs)
}

// PostVLine
//
// swagger:route POST /vlines vlines postVLine
//
// Creates a vline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVLine(c *gin.Context) {

	mutexVLine.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVLines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVLine.GetDB()

	// Validate input
	var input orm.VLineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vline
	vlineDB := orm.VLineDB{}
	vlineDB.VLinePointersEncoding = input.VLinePointersEncoding
	vlineDB.CopyBasicFieldsFromVLine_WOP(&input.VLine_WOP)

	query := db.Create(&vlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVLine.CheckoutPhaseOneInstance(&vlineDB)
	vline := backRepo.BackRepoVLine.Map_VLineDBID_VLinePtr[vlineDB.ID]

	if vline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), vline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, vlineDB)

	mutexVLine.Unlock()
}

// GetVLine
//
// swagger:route GET /vlines/{ID} vlines getVLine
//
// Gets the details for a vline.
//
// Responses:
// default: genericError
//
//	200: vlineDBResponse
func (controller *Controller) GetVLine(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVLine.GetDB()

	// Get vlineDB in DB
	var vlineDB orm.VLineDB
	if err := db.First(&vlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var vlineAPI orm.VLineAPI
	vlineAPI.ID = vlineDB.ID
	vlineAPI.VLinePointersEncoding = vlineDB.VLinePointersEncoding
	vlineDB.CopyBasicFieldsToVLine_WOP(&vlineAPI.VLine_WOP)

	c.JSON(http.StatusOK, vlineAPI)
}

// UpdateVLine
//
// swagger:route PATCH /vlines/{ID} vlines updateVLine
//
// # Update a vline
//
// Responses:
// default: genericError
//
//	200: vlineDBResponse
func (controller *Controller) UpdateVLine(c *gin.Context) {

	mutexVLine.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVLine.GetDB()

	// Validate input
	var input orm.VLineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var vlineDB orm.VLineDB

	// fetch the vline
	query := db.First(&vlineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	vlineDB.CopyBasicFieldsFromVLine_WOP(&input.VLine_WOP)
	vlineDB.VLinePointersEncoding = input.VLinePointersEncoding

	query = db.Model(&vlineDB).Updates(vlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	vlineNew := new(models.VLine)
	vlineDB.CopyBasicFieldsToVLine(vlineNew)

	// get stage instance from DB instance, and call callback function
	vlineOld := backRepo.BackRepoVLine.Map_VLineDBID_VLinePtr[vlineDB.ID]
	if vlineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), vlineOld, vlineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the vlineDB
	c.JSON(http.StatusOK, vlineDB)

	mutexVLine.Unlock()
}

// DeleteVLine
//
// swagger:route DELETE /vlines/{ID} vlines deleteVLine
//
// # Delete a vline
//
// default: genericError
//
//	200: vlineDBResponse
func (controller *Controller) DeleteVLine(c *gin.Context) {

	mutexVLine.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVLine", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVLine.GetDB()

	// Get model if exist
	var vlineDB orm.VLineDB
	if err := db.First(&vlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&vlineDB)

	// get an instance (not staged) from DB instance, and call callback function
	vlineDeleted := new(models.VLine)
	vlineDB.CopyBasicFieldsToVLine(vlineDeleted)

	// get stage instance from DB instance, and call callback function
	vlineStaged := backRepo.BackRepoVLine.Map_VLineDBID_VLinePtr[vlineDB.ID]
	if vlineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), vlineStaged, vlineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexVLine.Unlock()
}
