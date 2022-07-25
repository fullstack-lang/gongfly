package main

import (
	"embed"
	"flag"
	"fmt"
	"go/token"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// for the scenario reference
	gongfly "github.com/fullstack-lang/gongfly"
	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/reference"
	gonglfy_engine "github.com/fullstack-lang/gongfly/go/simulation"
	gongfly_visuals "github.com/fullstack-lang/gongfly/go/visuals"

	// demoatc gong stack
	gongfly_controllers "github.com/fullstack-lang/gongfly/go/controllers"
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongly_orm "github.com/fullstack-lang/gongfly/go/orm"

	// gongsim stack
	gongsim_controllers "github.com/fullstack-lang/gongsim/go/controllers"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gongsim/go/orm"
	_ "github.com/fullstack-lang/gongsim/ng"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"
	_ "github.com/fullstack-lang/gong/ng"

	// for diagrams
	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"
	_ "github.com/fullstack-lang/gongdoc/ng"

	// for carto display
	gongleaflet_controllers "github.com/fullstack-lang/gongleaflet/go/controllers"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongleaflet_orm "github.com/fullstack-lang/gongleaflet/go/orm"
	_ "github.com/fullstack-lang/gongleaflet/ng"

	// for document display
	gongmarkdown_controllers "github.com/fullstack-lang/gongmarkdown/go/controllers"
	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"
	gongmarkdown_orm "github.com/fullstack-lang/gongmarkdown/go/orm"
	_ "github.com/fullstack-lang/gongmarkdown/ng"

	// for the scheduler
	"github.com/fullstack-lang/gongfly/go/gongfly2markdown"

	// for graph display
	gongng2charts_controllers "github.com/fullstack-lang/gongng2charts/go/controllers"
	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"
	gongng2charts_orm "github.com/fullstack-lang/gongng2charts/go/orm"
	_ "github.com/fullstack-lang/gongng2charts/ng"

	// for the scheduler
	"github.com/fullstack-lang/gongfly/go/gongfly2gongng2charts"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	diagrams = flag.Bool("diagrams", true, "parse diagrams (takes a few seconds)")
)

func main() {

	log.SetPrefix("gongfly: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM with the
	db := gongly_orm.SetupModels(*logDBFlag, ":memory:")
	dbDB, err := db.DB()

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	// add gongsim database
	gongsim_orm.AutoMigrate(db)
	gongdoc_orm.AutoMigrate(db)
	gong_orm.AutoMigrate(db)
	gongleaflet_orm.AutoMigrate(db)
	gongmarkdown_orm.AutoMigrate(db)
	gongng2charts_orm.AutoMigrate(db)

	// attach specific engine callback to the model
	simulation := gonglfy_engine.NewSimulation()
	gongsim_models.EngineSingloton.Simulation = simulation

	// add sim event
	reference.Sc1_AF_3577_MDM.QueueUpdateEvent(1 * time.Second)

	// add sim event
	reference.Sat1.QueueUpdateEvent(1 * time.Second)
	reference.Sat3.QueueUpdateEvent(1 * time.Second)
	reference.Sat4.QueueUpdateEvent(1 * time.Second)
	reference.Sat2.QueueUpdateEvent(1 * time.Second)
	reference.ISS.QueueUpdateEvent(1 * time.Second)
	reference.Sat5.QueueUpdateEvent(1 * time.Second)
	reference.Sat6.QueueUpdateEvent(1 * time.Second)
	reference.Sat7.QueueUpdateEvent(1 * time.Second)
	reference.Sat8.QueueUpdateEvent(1 * time.Second)

	defaultLayer := new(gongleaflet_models.LayerGroup).Stage()
	defaultLayer.Name = "default"
	defaultLayer.DisplayName = "default"

	// attach visual elements
	gongfly_visuals.AttachVisualElementsToModelElements(defaultLayer)

	if *diagrams {

		// Analyse package
		modelPkg := &gong_models.ModelPkg{}

		// since the source is embedded, one needs to
		// compute the Abstract syntax tree in a special manner
		pkgs := gong_models.ParseEmbedModel(gongfly.GoDir, "go/models")

		gong_models.WalkParser(pkgs, modelPkg)
		modelPkg.SerializeToStage()
		gong_models.Stage.Commit()

		// create the diagrams
		// prepare the model views
		pkgelt := new(gongdoc_models.Pkgelt)

		// first, get all gong struct in the model
		for gongStruct := range gong_models.Stage.GongStructs {

			// let create the gong struct in the gongdoc models
			// and put the numbre of instances
			gongStruct_ := (&gongdoc_models.GongStruct{Name: gongStruct.Name}).Stage()
			nbInstances, ok := models.Stage.Map_GongStructName_InstancesNb[gongStruct.Name]
			if ok {
				gongStruct_.NbInstances = nbInstances
			}
		}

		// classdiagram can only be fully in memory when they are Unmarshalled
		// for instance, the Name of diagrams or the Name of the Link
		fset := new(token.FileSet)
		pkgsParser := gong_models.ParseEmbedModel(gongfly.GoDir, "go/diagrams")
		if len(pkgsParser) != 1 {
			log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
		}
		if pkgParser, ok := pkgsParser["diagrams"]; ok {
			pkgelt.Unmarshall(modelPkg, pkgParser, fset, "go/diagrams")
		}
		pkgelt.SerializeToStage()
	}

	gongfly_controllers.RegisterControllers(r)
	gongsim_controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)
	gongleaflet_controllers.RegisterControllers(r)
	gongmarkdown_controllers.RegisterControllers(r)
	gongng2charts_controllers.RegisterControllers(r)

	// put all to database
	gongfly_models.Stage.Commit()
	gongsim_models.Stage.Commit()
	gongdoc_models.Stage.Commit()
	gong_models.Stage.Commit()
	gongleaflet_models.Stage.Commit()
	gongmarkdown_models.Stage.Commit()
	gongng2charts_models.Stage.Commit()

	go gongfly2markdown.GenerateDocumentationSchedulerSingloton.CheckoutScheduler()
	go gongfly2gongng2charts.GenerateChartSchedulerSingloton.CheckoutScheduler()

	log.Print("Demoatc simulation is ready, waiting for client interactions (play/pause/...)")

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongfly.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
