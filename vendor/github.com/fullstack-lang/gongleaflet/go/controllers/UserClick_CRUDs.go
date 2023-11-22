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
var __UserClick__dummysDeclaration__ models.UserClick
var __UserClick_time__dummyDeclaration time.Duration

var mutexUserClick sync.Mutex

// An UserClickID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUserClick updateUserClick deleteUserClick
type UserClickID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UserClickInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUserClick updateUserClick
type UserClickInput struct {
	// The UserClick to submit or modify
	// in: body
	UserClick *orm.UserClickAPI
}

// GetUserClicks
//
// swagger:route GET /userclicks userclicks getUserClicks
//
// # Get all userclicks
//
// Responses:
// default: genericError
//
//	200: userclickDBResponse
func (controller *Controller) GetUserClicks(c *gin.Context) {

	// source slice
	var userclickDBs []orm.UserClickDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUserClicks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserClick.GetDB()

	query := db.Find(&userclickDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	userclickAPIs := make([]orm.UserClickAPI, 0)

	// for each userclick, update fields from the database nullable fields
	for idx := range userclickDBs {
		userclickDB := &userclickDBs[idx]
		_ = userclickDB
		var userclickAPI orm.UserClickAPI

		// insertion point for updating fields
		userclickAPI.ID = userclickDB.ID
		userclickDB.CopyBasicFieldsToUserClick_WOP(&userclickAPI.UserClick_WOP)
		userclickAPI.UserClickPointersEncoding = userclickDB.UserClickPointersEncoding
		userclickAPIs = append(userclickAPIs, userclickAPI)
	}

	c.JSON(http.StatusOK, userclickAPIs)
}

// PostUserClick
//
// swagger:route POST /userclicks userclicks postUserClick
//
// Creates a userclick
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUserClick(c *gin.Context) {

	mutexUserClick.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUserClicks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserClick.GetDB()

	// Validate input
	var input orm.UserClickAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create userclick
	userclickDB := orm.UserClickDB{}
	userclickDB.UserClickPointersEncoding = input.UserClickPointersEncoding
	userclickDB.CopyBasicFieldsFromUserClick_WOP(&input.UserClick_WOP)

	query := db.Create(&userclickDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUserClick.CheckoutPhaseOneInstance(&userclickDB)
	userclick := backRepo.BackRepoUserClick.Map_UserClickDBID_UserClickPtr[userclickDB.ID]

	if userclick != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), userclick)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, userclickDB)

	mutexUserClick.Unlock()
}

// GetUserClick
//
// swagger:route GET /userclicks/{ID} userclicks getUserClick
//
// Gets the details for a userclick.
//
// Responses:
// default: genericError
//
//	200: userclickDBResponse
func (controller *Controller) GetUserClick(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUserClick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserClick.GetDB()

	// Get userclickDB in DB
	var userclickDB orm.UserClickDB
	if err := db.First(&userclickDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var userclickAPI orm.UserClickAPI
	userclickAPI.ID = userclickDB.ID
	userclickAPI.UserClickPointersEncoding = userclickDB.UserClickPointersEncoding
	userclickDB.CopyBasicFieldsToUserClick_WOP(&userclickAPI.UserClick_WOP)

	c.JSON(http.StatusOK, userclickAPI)
}

// UpdateUserClick
//
// swagger:route PATCH /userclicks/{ID} userclicks updateUserClick
//
// # Update a userclick
//
// Responses:
// default: genericError
//
//	200: userclickDBResponse
func (controller *Controller) UpdateUserClick(c *gin.Context) {

	mutexUserClick.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUserClick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserClick.GetDB()

	// Validate input
	var input orm.UserClickAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var userclickDB orm.UserClickDB

	// fetch the userclick
	query := db.First(&userclickDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	userclickDB.CopyBasicFieldsFromUserClick_WOP(&input.UserClick_WOP)
	userclickDB.UserClickPointersEncoding = input.UserClickPointersEncoding

	query = db.Model(&userclickDB).Updates(userclickDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	userclickNew := new(models.UserClick)
	userclickDB.CopyBasicFieldsToUserClick(userclickNew)

	// redeem pointers
	userclickDB.DecodePointers(backRepo, userclickNew)

	// get stage instance from DB instance, and call callback function
	userclickOld := backRepo.BackRepoUserClick.Map_UserClickDBID_UserClickPtr[userclickDB.ID]
	if userclickOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), userclickOld, userclickNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the userclickDB
	c.JSON(http.StatusOK, userclickDB)

	mutexUserClick.Unlock()
}

// DeleteUserClick
//
// swagger:route DELETE /userclicks/{ID} userclicks deleteUserClick
//
// # Delete a userclick
//
// default: genericError
//
//	200: userclickDBResponse
func (controller *Controller) DeleteUserClick(c *gin.Context) {

	mutexUserClick.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUserClick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongleaflet/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserClick.GetDB()

	// Get model if exist
	var userclickDB orm.UserClickDB
	if err := db.First(&userclickDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&userclickDB)

	// get an instance (not staged) from DB instance, and call callback function
	userclickDeleted := new(models.UserClick)
	userclickDB.CopyBasicFieldsToUserClick(userclickDeleted)

	// get stage instance from DB instance, and call callback function
	userclickStaged := backRepo.BackRepoUserClick.Map_UserClickDBID_UserClickPtr[userclickDB.ID]
	if userclickStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), userclickStaged, userclickDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexUserClick.Unlock()
}
