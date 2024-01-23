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
var __Scenario__dummysDeclaration__ models.Scenario
var __Scenario_time__dummyDeclaration time.Duration

var mutexScenario sync.Mutex

// An ScenarioID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScenario updateScenario deleteScenario
type ScenarioID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ScenarioInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScenario updateScenario
type ScenarioInput struct {
	// The Scenario to submit or modify
	// in: body
	Scenario *orm.ScenarioAPI
}

// GetScenarios
//
// swagger:route GET /scenarios scenarios getScenarios
//
// # Get all scenarios
//
// Responses:
// default: genericError
//
//	200: scenarioDBResponse
func (controller *Controller) GetScenarios(c *gin.Context) {

	// source slice
	var scenarioDBs []orm.ScenarioDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScenarios", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScenario.GetDB()

	query := db.Find(&scenarioDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	scenarioAPIs := make([]orm.ScenarioAPI, 0)

	// for each scenario, update fields from the database nullable fields
	for idx := range scenarioDBs {
		scenarioDB := &scenarioDBs[idx]
		_ = scenarioDB
		var scenarioAPI orm.ScenarioAPI

		// insertion point for updating fields
		scenarioAPI.ID = scenarioDB.ID
		scenarioDB.CopyBasicFieldsToScenario_WOP(&scenarioAPI.Scenario_WOP)
		scenarioAPI.ScenarioPointersEncoding = scenarioDB.ScenarioPointersEncoding
		scenarioAPIs = append(scenarioAPIs, scenarioAPI)
	}

	c.JSON(http.StatusOK, scenarioAPIs)
}

// PostScenario
//
// swagger:route POST /scenarios scenarios postScenario
//
// Creates a scenario
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScenario(c *gin.Context) {

	mutexScenario.Lock()
	defer mutexScenario.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScenarios", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScenario.GetDB()

	// Validate input
	var input orm.ScenarioAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create scenario
	scenarioDB := orm.ScenarioDB{}
	scenarioDB.ScenarioPointersEncoding = input.ScenarioPointersEncoding
	scenarioDB.CopyBasicFieldsFromScenario_WOP(&input.Scenario_WOP)

	query := db.Create(&scenarioDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScenario.CheckoutPhaseOneInstance(&scenarioDB)
	scenario := backRepo.BackRepoScenario.Map_ScenarioDBID_ScenarioPtr[scenarioDB.ID]

	if scenario != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), scenario)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, scenarioDB)
}

// GetScenario
//
// swagger:route GET /scenarios/{ID} scenarios getScenario
//
// Gets the details for a scenario.
//
// Responses:
// default: genericError
//
//	200: scenarioDBResponse
func (controller *Controller) GetScenario(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScenario", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScenario.GetDB()

	// Get scenarioDB in DB
	var scenarioDB orm.ScenarioDB
	if err := db.First(&scenarioDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var scenarioAPI orm.ScenarioAPI
	scenarioAPI.ID = scenarioDB.ID
	scenarioAPI.ScenarioPointersEncoding = scenarioDB.ScenarioPointersEncoding
	scenarioDB.CopyBasicFieldsToScenario_WOP(&scenarioAPI.Scenario_WOP)

	c.JSON(http.StatusOK, scenarioAPI)
}

// UpdateScenario
//
// swagger:route PATCH /scenarios/{ID} scenarios updateScenario
//
// # Update a scenario
//
// Responses:
// default: genericError
//
//	200: scenarioDBResponse
func (controller *Controller) UpdateScenario(c *gin.Context) {

	mutexScenario.Lock()
	defer mutexScenario.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScenario", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScenario.GetDB()

	// Validate input
	var input orm.ScenarioAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var scenarioDB orm.ScenarioDB

	// fetch the scenario
	query := db.First(&scenarioDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	scenarioDB.CopyBasicFieldsFromScenario_WOP(&input.Scenario_WOP)
	scenarioDB.ScenarioPointersEncoding = input.ScenarioPointersEncoding

	query = db.Model(&scenarioDB).Updates(scenarioDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	scenarioNew := new(models.Scenario)
	scenarioDB.CopyBasicFieldsToScenario(scenarioNew)

	// redeem pointers
	scenarioDB.DecodePointers(backRepo, scenarioNew)

	// get stage instance from DB instance, and call callback function
	scenarioOld := backRepo.BackRepoScenario.Map_ScenarioDBID_ScenarioPtr[scenarioDB.ID]
	if scenarioOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), scenarioOld, scenarioNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the scenarioDB
	c.JSON(http.StatusOK, scenarioDB)
}

// DeleteScenario
//
// swagger:route DELETE /scenarios/{ID} scenarios deleteScenario
//
// # Delete a scenario
//
// default: genericError
//
//	200: scenarioDBResponse
func (controller *Controller) DeleteScenario(c *gin.Context) {

	mutexScenario.Lock()
	defer mutexScenario.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScenario", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongfly/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScenario.GetDB()

	// Get model if exist
	var scenarioDB orm.ScenarioDB
	if err := db.First(&scenarioDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&scenarioDB)

	// get an instance (not staged) from DB instance, and call callback function
	scenarioDeleted := new(models.Scenario)
	scenarioDB.CopyBasicFieldsToScenario(scenarioDeleted)

	// get stage instance from DB instance, and call callback function
	scenarioStaged := backRepo.BackRepoScenario.Map_ScenarioDBID_ScenarioPtr[scenarioDB.ID]
	if scenarioStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), scenarioStaged, scenarioDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
