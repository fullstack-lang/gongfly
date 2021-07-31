package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// for the scenario reference
	demoatc_reference "github.com/fullstack-lang/gongfly/go/reference"
	demoatc_engine "github.com/fullstack-lang/gongfly/go/simulation"
	demoatc_visuals "github.com/fullstack-lang/gongfly/go/visuals"

	// demoatc gong stack
	demoatc_controllers "github.com/fullstack-lang/gongfly/go/controllers"
	demoatc_models "github.com/fullstack-lang/gongfly/go/models"
	demoatc_orm "github.com/fullstack-lang/gongfly/go/orm"

	// gongsim stack
	gongsim_controllers "github.com/fullstack-lang/gongsim/go/controllers"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gongsim/go/orm"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"

	// for diagrams
	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"

	// for carto display
	gongleaflet_controllers "github.com/fullstack-lang/gongleaflet/go/controllers"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongleaflet_orm "github.com/fullstack-lang/gongleaflet/go/orm"
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
	db := demoatc_orm.SetupModels(*logDBFlag, ":memory:")
	dbDB, err := db.DB()

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	// add gongsim database
	gongsim_orm.AutoMigrate(db)

	// add gongdoc database
	gongdoc_orm.AutoMigrate(db)

	// add gong database
	gong_orm.AutoMigrate(db)

	// add gongdoc database
	gongleaflet_orm.AutoMigrate(db)

	// attach specific engine callback to the model
	simulation := demoatc_engine.NewSimulation()
	gongsim_models.EngineSingloton.Simulation = simulation

	// attach visual elements
	demoatc_visuals.AttachVisualElementsToModelElements()

	// create the target map
	visualMap := new(gongleaflet_models.VisualMap).Stage()
	visualMap.Lat = demoatc_reference.Scenario1.Lat
	visualMap.Lng = demoatc_reference.Scenario1.Lng
	visualMap.ZoomLevel = demoatc_reference.Scenario1.ZoomLevel

	visualMap.UrlTemplate = "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png"
	visualMap.Attribution = "osm"
	visualMap.MaxZoom = 18

	_true := true
	_false := false
	visualMap.ZoomControl = _false
	visualMap.AttributionControl = _true
	visualMap.ZoomSnap = _true

	// load package to analyse
	modelPkg := &gong_models.ModelPkg{}
	if *diagrams {
		gong_models.Walk("go/models", modelPkg)
		modelPkg.SerializeToStage()
	}

	// create the diagrams
	// prepare the model views
	pkgelt := new(gongdoc_models.Pkgelt)

	// classdiagram can only be fully in memory when they are Unmarshalled
	// for instance, the Name of diagrams or the Name of the Link
	if *diagrams {
		pkgelt.Unmarshall("go/diagrams")
	}
	pkgelt.SerializeToStage()

	demoatc_controllers.RegisterControllers(r)
	gongsim_controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)
	gongleaflet_controllers.RegisterControllers(r)

	//
	// init all back repositories
	//
	demoatc_orm.BackRepo.Init(db)
	gongsim_orm.BackRepo.Init(db)
	gongdoc_orm.BackRepo.Init(db)
	gong_orm.BackRepo.Init(db)
	gongleaflet_orm.BackRepo.Init(db)

	// put all to database
	demoatc_models.Stage.Commit()
	gongsim_models.Stage.Commit()
	gongdoc_models.Stage.Commit()
	gong_models.Stage.Commit()
	gongleaflet_models.Stage.Commit()

	log.Print("Demoatc simulation is ready, waiting for client interactions (play/pause/...)")

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

//go:embed ng/dist/ng
var ng embed.FS

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
