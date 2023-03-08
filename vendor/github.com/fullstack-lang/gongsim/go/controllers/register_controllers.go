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
	v1 := r.Group("/api/github.com/fullstack-lang/gongsim/go")
	{ // insertion point for registrations
		v1.GET("/v1/dummyagents", GetController().GetDummyAgents)
		v1.GET("/v1/dummyagents/:id", GetController().GetDummyAgent)
		v1.POST("/v1/dummyagents", GetController().PostDummyAgent)
		v1.PATCH("/v1/dummyagents/:id", GetController().UpdateDummyAgent)
		v1.PUT("/v1/dummyagents/:id", GetController().UpdateDummyAgent)
		v1.DELETE("/v1/dummyagents/:id", GetController().DeleteDummyAgent)

		v1.GET("/v1/engines", GetController().GetEngines)
		v1.GET("/v1/engines/:id", GetController().GetEngine)
		v1.POST("/v1/engines", GetController().PostEngine)
		v1.PATCH("/v1/engines/:id", GetController().UpdateEngine)
		v1.PUT("/v1/engines/:id", GetController().UpdateEngine)
		v1.DELETE("/v1/engines/:id", GetController().DeleteEngine)

		v1.GET("/v1/events", GetController().GetEvents)
		v1.GET("/v1/events/:id", GetController().GetEvent)
		v1.POST("/v1/events", GetController().PostEvent)
		v1.PATCH("/v1/events/:id", GetController().UpdateEvent)
		v1.PUT("/v1/events/:id", GetController().UpdateEvent)
		v1.DELETE("/v1/events/:id", GetController().DeleteEvent)

		v1.GET("/v1/gongsimcommands", GetController().GetGongsimCommands)
		v1.GET("/v1/gongsimcommands/:id", GetController().GetGongsimCommand)
		v1.POST("/v1/gongsimcommands", GetController().PostGongsimCommand)
		v1.PATCH("/v1/gongsimcommands/:id", GetController().UpdateGongsimCommand)
		v1.PUT("/v1/gongsimcommands/:id", GetController().UpdateGongsimCommand)
		v1.DELETE("/v1/gongsimcommands/:id", GetController().DeleteGongsimCommand)

		v1.GET("/v1/gongsimstatuss", GetController().GetGongsimStatuss)
		v1.GET("/v1/gongsimstatuss/:id", GetController().GetGongsimStatus)
		v1.POST("/v1/gongsimstatuss", GetController().PostGongsimStatus)
		v1.PATCH("/v1/gongsimstatuss/:id", GetController().UpdateGongsimStatus)
		v1.PUT("/v1/gongsimstatuss/:id", GetController().UpdateGongsimStatus)
		v1.DELETE("/v1/gongsimstatuss/:id", GetController().DeleteGongsimStatus)

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
