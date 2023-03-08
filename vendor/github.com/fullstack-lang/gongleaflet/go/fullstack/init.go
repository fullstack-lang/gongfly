package fullstack

import (
	// gongleaflet stack for model analysis

	gongleaflet_controllers "github.com/fullstack-lang/gongleaflet/go/controllers"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongleaflet_orm "github.com/fullstack-lang/gongleaflet/go/orm"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gongleaflet/ng/projects"
)

func Init(r *gin.Engine, filenames ...string) {

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	db_inMemory := gongleaflet_orm.SetupModels(&gongleaflet_models.Stage, false, filenames[0])

	// since gongleafletsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	gongleaflet_controllers.RegisterControllers(r)
}
