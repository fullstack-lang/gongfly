// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Radar__dummysDeclaration__ models.Radar
var __Radar_time__dummyDeclaration time.Duration

// An RadarID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRadar updateRadar deleteRadar
type RadarID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RadarInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRadar updateRadar
type RadarInput struct {
	// The Radar to submit or modify
	// in: body
	Radar *orm.RadarAPI
}

// GetRadars
//
// swagger:route GET /radars radars getRadars
//
// Get all radars
//
// Responses:
//    default: genericError
//        200: radarDBsResponse
func GetRadars(c *gin.Context) {
	db := orm.BackRepo.BackRepoRadar.GetDB()

	// source slice
	var radarDBs []orm.RadarDB
	query := db.Find(&radarDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	radarAPIs := make([]orm.RadarAPI, 0)

	// for each radar, update fields from the database nullable fields
	for idx := range radarDBs {
		radarDB := &radarDBs[idx]
		_ = radarDB
		var radarAPI orm.RadarAPI

		// insertion point for updating fields
		radarAPI.ID = radarDB.ID
		radarDB.CopyBasicFieldsToRadar(&radarAPI.Radar)
		radarAPI.RadarPointersEnconding = radarDB.RadarPointersEnconding
		radarAPIs = append(radarAPIs, radarAPI)
	}

	c.JSON(http.StatusOK, radarAPIs)
}

// PostRadar
//
// swagger:route POST /radars radars postRadar
//
// Creates a radar
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: radarDBResponse
func PostRadar(c *gin.Context) {
	db := orm.BackRepo.BackRepoRadar.GetDB()

	// Validate input
	var input orm.RadarAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create radar
	radarDB := orm.RadarDB{}
	radarDB.RadarPointersEnconding = input.RadarPointersEnconding
	radarDB.CopyBasicFieldsFromRadar(&input.Radar)

	query := db.Create(&radarDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, radarDB)
}

// GetRadar
//
// swagger:route GET /radars/{ID} radars getRadar
//
// Gets the details for a radar.
//
// Responses:
//    default: genericError
//        200: radarDBResponse
func GetRadar(c *gin.Context) {
	db := orm.BackRepo.BackRepoRadar.GetDB()

	// Get radarDB in DB
	var radarDB orm.RadarDB
	if err := db.First(&radarDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var radarAPI orm.RadarAPI
	radarAPI.ID = radarDB.ID
	radarAPI.RadarPointersEnconding = radarDB.RadarPointersEnconding
	radarDB.CopyBasicFieldsToRadar(&radarAPI.Radar)

	c.JSON(http.StatusOK, radarAPI)
}

// UpdateRadar
//
// swagger:route PATCH /radars/{ID} radars updateRadar
//
// Update a radar
//
// Responses:
//    default: genericError
//        200: radarDBResponse
func UpdateRadar(c *gin.Context) {
	db := orm.BackRepo.BackRepoRadar.GetDB()

	// Get model if exist
	var radarDB orm.RadarDB

	// fetch the radar
	query := db.First(&radarDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.RadarAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	radarDB.CopyBasicFieldsFromRadar(&input.Radar)
	radarDB.RadarPointersEnconding = input.RadarPointersEnconding

	query = db.Model(&radarDB).Updates(radarDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the radarDB
	c.JSON(http.StatusOK, radarDB)
}

// DeleteRadar
//
// swagger:route DELETE /radars/{ID} radars deleteRadar
//
// Delete a radar
//
// Responses:
//    default: genericError
func DeleteRadar(c *gin.Context) {
	db := orm.BackRepo.BackRepoRadar.GetDB()

	// Get model if exist
	var radarDB orm.RadarDB
	if err := db.First(&radarDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&radarDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
