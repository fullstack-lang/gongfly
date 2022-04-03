package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongmarkdown/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gongmarkdown/go")
	{ // insertion point for registrations
		v1.GET("/v1/anotherdummydatas", GetAnotherDummyDatas)
		v1.GET("/v1/anotherdummydatas/:id", GetAnotherDummyData)
		v1.POST("/v1/anotherdummydatas", PostAnotherDummyData)
		v1.PATCH("/v1/anotherdummydatas/:id", UpdateAnotherDummyData)
		v1.PUT("/v1/anotherdummydatas/:id", UpdateAnotherDummyData)
		v1.DELETE("/v1/anotherdummydatas/:id", DeleteAnotherDummyData)

		v1.GET("/v1/cells", GetCells)
		v1.GET("/v1/cells/:id", GetCell)
		v1.POST("/v1/cells", PostCell)
		v1.PATCH("/v1/cells/:id", UpdateCell)
		v1.PUT("/v1/cells/:id", UpdateCell)
		v1.DELETE("/v1/cells/:id", DeleteCell)

		v1.GET("/v1/dummydatas", GetDummyDatas)
		v1.GET("/v1/dummydatas/:id", GetDummyData)
		v1.POST("/v1/dummydatas", PostDummyData)
		v1.PATCH("/v1/dummydatas/:id", UpdateDummyData)
		v1.PUT("/v1/dummydatas/:id", UpdateDummyData)
		v1.DELETE("/v1/dummydatas/:id", DeleteDummyData)

		v1.GET("/v1/elements", GetElements)
		v1.GET("/v1/elements/:id", GetElement)
		v1.POST("/v1/elements", PostElement)
		v1.PATCH("/v1/elements/:id", UpdateElement)
		v1.PUT("/v1/elements/:id", UpdateElement)
		v1.DELETE("/v1/elements/:id", DeleteElement)

		v1.GET("/v1/markdowncontents", GetMarkdownContents)
		v1.GET("/v1/markdowncontents/:id", GetMarkdownContent)
		v1.POST("/v1/markdowncontents", PostMarkdownContent)
		v1.PATCH("/v1/markdowncontents/:id", UpdateMarkdownContent)
		v1.PUT("/v1/markdowncontents/:id", UpdateMarkdownContent)
		v1.DELETE("/v1/markdowncontents/:id", DeleteMarkdownContent)

		v1.GET("/v1/rows", GetRows)
		v1.GET("/v1/rows/:id", GetRow)
		v1.POST("/v1/rows", PostRow)
		v1.PATCH("/v1/rows/:id", UpdateRow)
		v1.PUT("/v1/rows/:id", UpdateRow)
		v1.DELETE("/v1/rows/:id", DeleteRow)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
