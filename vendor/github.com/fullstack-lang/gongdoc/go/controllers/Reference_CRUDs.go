// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Reference__dummysDeclaration__ models.Reference
var __Reference_time__dummyDeclaration time.Duration

// An ReferenceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getReference updateReference deleteReference
type ReferenceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ReferenceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postReference updateReference
type ReferenceInput struct {
	// The Reference to submit or modify
	// in: body
	Reference *orm.ReferenceAPI
}

// GetReferences
//
// swagger:route GET /references references getReferences
//
// # Get all references
//
// Responses:
// default: genericError
//
//	200: referenceDBResponse
func GetReferences(c *gin.Context) {
	db := orm.BackRepo.BackRepoReference.GetDB()

	// source slice
	var referenceDBs []orm.ReferenceDB
	query := db.Find(&referenceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	referenceAPIs := make([]orm.ReferenceAPI, 0)

	// for each reference, update fields from the database nullable fields
	for idx := range referenceDBs {
		referenceDB := &referenceDBs[idx]
		_ = referenceDB
		var referenceAPI orm.ReferenceAPI

		// insertion point for updating fields
		referenceAPI.ID = referenceDB.ID
		referenceDB.CopyBasicFieldsToReference(&referenceAPI.Reference)
		referenceAPI.ReferencePointersEnconding = referenceDB.ReferencePointersEnconding
		referenceAPIs = append(referenceAPIs, referenceAPI)
	}

	c.JSON(http.StatusOK, referenceAPIs)
}

// PostReference
//
// swagger:route POST /references references postReference
//
// Creates a reference
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostReference(c *gin.Context) {
	db := orm.BackRepo.BackRepoReference.GetDB()

	// Validate input
	var input orm.ReferenceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reference
	referenceDB := orm.ReferenceDB{}
	referenceDB.ReferencePointersEnconding = input.ReferencePointersEnconding
	referenceDB.CopyBasicFieldsFromReference(&input.Reference)

	query := db.Create(&referenceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoReference.CheckoutPhaseOneInstance(&referenceDB)
	reference := (*orm.BackRepo.BackRepoReference.Map_ReferenceDBID_ReferencePtr)[referenceDB.ID]

	if reference != nil {
		models.AfterCreateFromFront(&models.Stage, reference)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, referenceDB)
}

// GetReference
//
// swagger:route GET /references/{ID} references getReference
//
// Gets the details for a reference.
//
// Responses:
// default: genericError
//
//	200: referenceDBResponse
func GetReference(c *gin.Context) {
	db := orm.BackRepo.BackRepoReference.GetDB()

	// Get referenceDB in DB
	var referenceDB orm.ReferenceDB
	if err := db.First(&referenceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var referenceAPI orm.ReferenceAPI
	referenceAPI.ID = referenceDB.ID
	referenceAPI.ReferencePointersEnconding = referenceDB.ReferencePointersEnconding
	referenceDB.CopyBasicFieldsToReference(&referenceAPI.Reference)

	c.JSON(http.StatusOK, referenceAPI)
}

// UpdateReference
//
// swagger:route PATCH /references/{ID} references updateReference
//
// # Update a reference
//
// Responses:
// default: genericError
//
//	200: referenceDBResponse
func UpdateReference(c *gin.Context) {
	db := orm.BackRepo.BackRepoReference.GetDB()

	// Get model if exist
	var referenceDB orm.ReferenceDB

	// fetch the reference
	query := db.First(&referenceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ReferenceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	referenceDB.CopyBasicFieldsFromReference(&input.Reference)
	referenceDB.ReferencePointersEnconding = input.ReferencePointersEnconding

	query = db.Model(&referenceDB).Updates(referenceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	referenceNew := new(models.Reference)
	referenceDB.CopyBasicFieldsToReference(referenceNew)

	// get stage instance from DB instance, and call callback function
	referenceOld := (*orm.BackRepo.BackRepoReference.Map_ReferenceDBID_ReferencePtr)[referenceDB.ID]
	if referenceOld != nil {
		models.AfterUpdateFromFront(&models.Stage, referenceOld, referenceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the referenceDB
	c.JSON(http.StatusOK, referenceDB)
}

// DeleteReference
//
// swagger:route DELETE /references/{ID} references deleteReference
//
// # Delete a reference
//
// default: genericError
//
//	200: referenceDBResponse
func DeleteReference(c *gin.Context) {
	db := orm.BackRepo.BackRepoReference.GetDB()

	// Get model if exist
	var referenceDB orm.ReferenceDB
	if err := db.First(&referenceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&referenceDB)

	// get an instance (not staged) from DB instance, and call callback function
	referenceDeleted := new(models.Reference)
	referenceDB.CopyBasicFieldsToReference(referenceDeleted)

	// get stage instance from DB instance, and call callback function
	referenceStaged := (*orm.BackRepo.BackRepoReference.Map_ReferenceDBID_ReferencePtr)[referenceDB.ID]
	if referenceStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, referenceStaged, referenceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}