package main

import (
	"embed"
	"flag"
	"fmt"
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

	gongfly_models "github.com/fullstack-lang/gongfly/go/models"

	// gongsim stack
	gongsim_fullstack "github.com/fullstack-lang/gongsim/go/fullstack"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	gongmarkdown_fullstack "github.com/fullstack-lang/gongmarkdown/go/fullstack"
	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"

	gongng2charts_fullstack "github.com/fullstack-lang/gongng2charts/go/fullstack"
	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"

	// for diagrams
	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	// for carto display
	gongleaflet_fullstack "github.com/fullstack-lang/gongleaflet/go/fullstack"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"

	// for document display

	// for the scheduler
	"github.com/fullstack-lang/gongfly/go/gongfly2markdown"

	// for the scheduler
	"github.com/fullstack-lang/gongfly/go/gongfly2gongng2charts"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
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

	gongsim_fullstack.Init(r)
	gongleaflet_fullstack.Init(r)
	gongmarkdown_fullstack.Init(r)
	gongng2charts_fullstack.Init(r)

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

	defaultLayer := new(gongleaflet_models.LayerGroup).Stage(&gongleaflet_models.Stage)
	defaultLayer.Name = "default"
	defaultLayer.DisplayName = "default"

	// attach visual elements
	gongfly_visuals.AttachVisualElementsToModelElements(defaultLayer)

	if *diagrams {
		gongdoc_load.Load(
			"gongfly",
			"github.com/fullstack-lang/gongfly/go/models",
			gongfly.GoDir,
			r,
			*embeddedDiagrams,
			&models.Stage.Map_GongStructName_InstancesNb)
	}

	// put all to database
	gongfly_models.Stage.Commit()
	gongsim_models.Stage.Commit()
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
