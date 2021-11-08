package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongleaflet/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gongleaflet/go")
	{ // insertion point for registrations
		v1.GET("/v1/visualcenters", GetVisualCenters)
		v1.GET("/v1/visualcenters/:id", GetVisualCenter)
		v1.POST("/v1/visualcenters", PostVisualCenter)
		v1.PATCH("/v1/visualcenters/:id", UpdateVisualCenter)
		v1.PUT("/v1/visualcenters/:id", UpdateVisualCenter)
		v1.DELETE("/v1/visualcenters/:id", DeleteVisualCenter)

		v1.GET("/v1/visualcircles", GetVisualCircles)
		v1.GET("/v1/visualcircles/:id", GetVisualCircle)
		v1.POST("/v1/visualcircles", PostVisualCircle)
		v1.PATCH("/v1/visualcircles/:id", UpdateVisualCircle)
		v1.PUT("/v1/visualcircles/:id", UpdateVisualCircle)
		v1.DELETE("/v1/visualcircles/:id", DeleteVisualCircle)

		v1.GET("/v1/visualicons", GetVisualIcons)
		v1.GET("/v1/visualicons/:id", GetVisualIcon)
		v1.POST("/v1/visualicons", PostVisualIcon)
		v1.PATCH("/v1/visualicons/:id", UpdateVisualIcon)
		v1.PUT("/v1/visualicons/:id", UpdateVisualIcon)
		v1.DELETE("/v1/visualicons/:id", DeleteVisualIcon)

		v1.GET("/v1/visuallayers", GetVisualLayers)
		v1.GET("/v1/visuallayers/:id", GetVisualLayer)
		v1.POST("/v1/visuallayers", PostVisualLayer)
		v1.PATCH("/v1/visuallayers/:id", UpdateVisualLayer)
		v1.PUT("/v1/visuallayers/:id", UpdateVisualLayer)
		v1.DELETE("/v1/visuallayers/:id", DeleteVisualLayer)

		v1.GET("/v1/visuallines", GetVisualLines)
		v1.GET("/v1/visuallines/:id", GetVisualLine)
		v1.POST("/v1/visuallines", PostVisualLine)
		v1.PATCH("/v1/visuallines/:id", UpdateVisualLine)
		v1.PUT("/v1/visuallines/:id", UpdateVisualLine)
		v1.DELETE("/v1/visuallines/:id", DeleteVisualLine)

		v1.GET("/v1/visualmaps", GetVisualMaps)
		v1.GET("/v1/visualmaps/:id", GetVisualMap)
		v1.POST("/v1/visualmaps", PostVisualMap)
		v1.PATCH("/v1/visualmaps/:id", UpdateVisualMap)
		v1.PUT("/v1/visualmaps/:id", UpdateVisualMap)
		v1.DELETE("/v1/visualmaps/:id", DeleteVisualMap)

		v1.GET("/v1/visualtracks", GetVisualTracks)
		v1.GET("/v1/visualtracks/:id", GetVisualTrack)
		v1.POST("/v1/visualtracks", PostVisualTrack)
		v1.PATCH("/v1/visualtracks/:id", UpdateVisualTrack)
		v1.PUT("/v1/visualtracks/:id", UpdateVisualTrack)
		v1.DELETE("/v1/visualtracks/:id", DeleteVisualTrack)

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
