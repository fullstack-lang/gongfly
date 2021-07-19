package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongfly/go/orm"
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
		v1.GET("/v1/civilianairports", GetCivilianAirports)
		v1.GET("/v1/civilianairports/:id", GetCivilianAirport)
		v1.POST("/v1/civilianairports", PostCivilianAirport)
		v1.PATCH("/v1/civilianairports/:id", UpdateCivilianAirport)
		v1.PUT("/v1/civilianairports/:id", UpdateCivilianAirport)
		v1.DELETE("/v1/civilianairports/:id", DeleteCivilianAirport)

		v1.GET("/v1/liners", GetLiners)
		v1.GET("/v1/liners/:id", GetLiner)
		v1.POST("/v1/liners", PostLiner)
		v1.PATCH("/v1/liners/:id", UpdateLiner)
		v1.PUT("/v1/liners/:id", UpdateLiner)
		v1.DELETE("/v1/liners/:id", DeleteLiner)

		v1.GET("/v1/messages", GetMessages)
		v1.GET("/v1/messages/:id", GetMessage)
		v1.POST("/v1/messages", PostMessage)
		v1.PATCH("/v1/messages/:id", UpdateMessage)
		v1.PUT("/v1/messages/:id", UpdateMessage)
		v1.DELETE("/v1/messages/:id", DeleteMessage)

		v1.GET("/v1/opslines", GetOpsLines)
		v1.GET("/v1/opslines/:id", GetOpsLine)
		v1.POST("/v1/opslines", PostOpsLine)
		v1.PATCH("/v1/opslines/:id", UpdateOpsLine)
		v1.PUT("/v1/opslines/:id", UpdateOpsLine)
		v1.DELETE("/v1/opslines/:id", DeleteOpsLine)

		v1.GET("/v1/orders", GetOrders)
		v1.GET("/v1/orders/:id", GetOrder)
		v1.POST("/v1/orders", PostOrder)
		v1.PATCH("/v1/orders/:id", UpdateOrder)
		v1.PUT("/v1/orders/:id", UpdateOrder)
		v1.DELETE("/v1/orders/:id", DeleteOrder)

		v1.GET("/v1/radars", GetRadars)
		v1.GET("/v1/radars/:id", GetRadar)
		v1.POST("/v1/radars", PostRadar)
		v1.PATCH("/v1/radars/:id", UpdateRadar)
		v1.PUT("/v1/radars/:id", UpdateRadar)
		v1.DELETE("/v1/radars/:id", DeleteRadar)

		v1.GET("/v1/reports", GetReports)
		v1.GET("/v1/reports/:id", GetReport)
		v1.POST("/v1/reports", PostReport)
		v1.PATCH("/v1/reports/:id", UpdateReport)
		v1.PUT("/v1/reports/:id", UpdateReport)
		v1.DELETE("/v1/reports/:id", DeleteReport)

		v1.GET("/v1/scenarios", GetScenarios)
		v1.GET("/v1/scenarios/:id", GetScenario)
		v1.POST("/v1/scenarios", PostScenario)
		v1.PATCH("/v1/scenarios/:id", UpdateScenario)
		v1.PUT("/v1/scenarios/:id", UpdateScenario)
		v1.DELETE("/v1/scenarios/:id", DeleteScenario)

		v1.GET("/commitnb", GetLastCommitNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
