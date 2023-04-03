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

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongleaflet/go")
	{ // insertion point for registrations
		v1.GET("/v1/circles", GetController().GetCircles)
		v1.GET("/v1/circles/:id", GetController().GetCircle)
		v1.POST("/v1/circles", GetController().PostCircle)
		v1.PATCH("/v1/circles/:id", GetController().UpdateCircle)
		v1.PUT("/v1/circles/:id", GetController().UpdateCircle)
		v1.DELETE("/v1/circles/:id", GetController().DeleteCircle)

		v1.GET("/v1/divicons", GetController().GetDivIcons)
		v1.GET("/v1/divicons/:id", GetController().GetDivIcon)
		v1.POST("/v1/divicons", GetController().PostDivIcon)
		v1.PATCH("/v1/divicons/:id", GetController().UpdateDivIcon)
		v1.PUT("/v1/divicons/:id", GetController().UpdateDivIcon)
		v1.DELETE("/v1/divicons/:id", GetController().DeleteDivIcon)

		v1.GET("/v1/layergroups", GetController().GetLayerGroups)
		v1.GET("/v1/layergroups/:id", GetController().GetLayerGroup)
		v1.POST("/v1/layergroups", GetController().PostLayerGroup)
		v1.PATCH("/v1/layergroups/:id", GetController().UpdateLayerGroup)
		v1.PUT("/v1/layergroups/:id", GetController().UpdateLayerGroup)
		v1.DELETE("/v1/layergroups/:id", GetController().DeleteLayerGroup)

		v1.GET("/v1/layergroupuses", GetController().GetLayerGroupUses)
		v1.GET("/v1/layergroupuses/:id", GetController().GetLayerGroupUse)
		v1.POST("/v1/layergroupuses", GetController().PostLayerGroupUse)
		v1.PATCH("/v1/layergroupuses/:id", GetController().UpdateLayerGroupUse)
		v1.PUT("/v1/layergroupuses/:id", GetController().UpdateLayerGroupUse)
		v1.DELETE("/v1/layergroupuses/:id", GetController().DeleteLayerGroupUse)

		v1.GET("/v1/mapoptionss", GetController().GetMapOptionss)
		v1.GET("/v1/mapoptionss/:id", GetController().GetMapOptions)
		v1.POST("/v1/mapoptionss", GetController().PostMapOptions)
		v1.PATCH("/v1/mapoptionss/:id", GetController().UpdateMapOptions)
		v1.PUT("/v1/mapoptionss/:id", GetController().UpdateMapOptions)
		v1.DELETE("/v1/mapoptionss/:id", GetController().DeleteMapOptions)

		v1.GET("/v1/markers", GetController().GetMarkers)
		v1.GET("/v1/markers/:id", GetController().GetMarker)
		v1.POST("/v1/markers", GetController().PostMarker)
		v1.PATCH("/v1/markers/:id", GetController().UpdateMarker)
		v1.PUT("/v1/markers/:id", GetController().UpdateMarker)
		v1.DELETE("/v1/markers/:id", GetController().DeleteMarker)

		v1.GET("/v1/userclicks", GetController().GetUserClicks)
		v1.GET("/v1/userclicks/:id", GetController().GetUserClick)
		v1.POST("/v1/userclicks", GetController().PostUserClick)
		v1.PATCH("/v1/userclicks/:id", GetController().UpdateUserClick)
		v1.PUT("/v1/userclicks/:id", GetController().UpdateUserClick)
		v1.DELETE("/v1/userclicks/:id", GetController().DeleteUserClick)

		v1.GET("/v1/vlines", GetController().GetVLines)
		v1.GET("/v1/vlines/:id", GetController().GetVLine)
		v1.POST("/v1/vlines", GetController().PostVLine)
		v1.PATCH("/v1/vlines/:id", GetController().UpdateVLine)
		v1.PUT("/v1/vlines/:id", GetController().UpdateVLine)
		v1.DELETE("/v1/vlines/:id", GetController().DeleteVLine)

		v1.GET("/v1/visualtracks", GetController().GetVisualTracks)
		v1.GET("/v1/visualtracks/:id", GetController().GetVisualTrack)
		v1.POST("/v1/visualtracks", GetController().PostVisualTrack)
		v1.PATCH("/v1/visualtracks/:id", GetController().UpdateVisualTrack)
		v1.PUT("/v1/visualtracks/:id", GetController().UpdateVisualTrack)
		v1.DELETE("/v1/visualtracks/:id", GetController().DeleteVisualTrack)

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
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
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
