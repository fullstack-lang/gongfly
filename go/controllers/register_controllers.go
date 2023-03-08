package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongfly/go")
	{ // insertion point for registrations
		v1.GET("/v1/civilianairports", GetController().GetCivilianAirports)
		v1.GET("/v1/civilianairports/:id", GetController().GetCivilianAirport)
		v1.POST("/v1/civilianairports", GetController().PostCivilianAirport)
		v1.PATCH("/v1/civilianairports/:id", GetController().UpdateCivilianAirport)
		v1.PUT("/v1/civilianairports/:id", GetController().UpdateCivilianAirport)
		v1.DELETE("/v1/civilianairports/:id", GetController().DeleteCivilianAirport)

		v1.GET("/v1/liners", GetController().GetLiners)
		v1.GET("/v1/liners/:id", GetController().GetLiner)
		v1.POST("/v1/liners", GetController().PostLiner)
		v1.PATCH("/v1/liners/:id", GetController().UpdateLiner)
		v1.PUT("/v1/liners/:id", GetController().UpdateLiner)
		v1.DELETE("/v1/liners/:id", GetController().DeleteLiner)

		v1.GET("/v1/messages", GetController().GetMessages)
		v1.GET("/v1/messages/:id", GetController().GetMessage)
		v1.POST("/v1/messages", GetController().PostMessage)
		v1.PATCH("/v1/messages/:id", GetController().UpdateMessage)
		v1.PUT("/v1/messages/:id", GetController().UpdateMessage)
		v1.DELETE("/v1/messages/:id", GetController().DeleteMessage)

		v1.GET("/v1/opslines", GetController().GetOpsLines)
		v1.GET("/v1/opslines/:id", GetController().GetOpsLine)
		v1.POST("/v1/opslines", GetController().PostOpsLine)
		v1.PATCH("/v1/opslines/:id", GetController().UpdateOpsLine)
		v1.PUT("/v1/opslines/:id", GetController().UpdateOpsLine)
		v1.DELETE("/v1/opslines/:id", GetController().DeleteOpsLine)

		v1.GET("/v1/radars", GetController().GetRadars)
		v1.GET("/v1/radars/:id", GetController().GetRadar)
		v1.POST("/v1/radars", GetController().PostRadar)
		v1.PATCH("/v1/radars/:id", GetController().UpdateRadar)
		v1.PUT("/v1/radars/:id", GetController().UpdateRadar)
		v1.DELETE("/v1/radars/:id", GetController().DeleteRadar)

		v1.GET("/v1/satellites", GetController().GetSatellites)
		v1.GET("/v1/satellites/:id", GetController().GetSatellite)
		v1.POST("/v1/satellites", GetController().PostSatellite)
		v1.PATCH("/v1/satellites/:id", GetController().UpdateSatellite)
		v1.PUT("/v1/satellites/:id", GetController().UpdateSatellite)
		v1.DELETE("/v1/satellites/:id", GetController().DeleteSatellite)

		v1.GET("/v1/scenarios", GetController().GetScenarios)
		v1.GET("/v1/scenarios/:id", GetController().GetScenario)
		v1.POST("/v1/scenarios", GetController().PostScenario)
		v1.PATCH("/v1/scenarios/:id", GetController().UpdateScenario)
		v1.PUT("/v1/scenarios/:id", GetController().UpdateScenario)
		v1.DELETE("/v1/scenarios/:id", GetController().DeleteScenario)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func(controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
