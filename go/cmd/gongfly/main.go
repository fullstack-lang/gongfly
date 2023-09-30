package main

import (
	"flag"
	"log"
	"strconv"

	gongfly_go "github.com/fullstack-lang/gongfly/go"
	gongfly_fullstack "github.com/fullstack-lang/gongfly/go/fullstack"
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_orm "github.com/fullstack-lang/gongfly/go/orm"
	gongfly_probe "github.com/fullstack-lang/gongfly/go/probe"
	gongfly_static "github.com/fullstack-lang/gongfly/go/static"

	"github.com/fullstack-lang/gongfly/go/visuals"

	gongfly_icons "github.com/fullstack-lang/gongfly/go/icons"
	"github.com/fullstack-lang/gongfly/go/reference"
	"github.com/fullstack-lang/gongfly/go/simulation"

	gongfly_visuals "github.com/fullstack-lang/gongfly/go/visuals"

	gongleaflet_go "github.com/fullstack-lang/gongleaflet/go"
	gongleaflet_fullstack "github.com/fullstack-lang/gongleaflet/go/fullstack"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongleaflet_probe "github.com/fullstack-lang/gongleaflet/go/probe"

	gongsim_go "github.com/fullstack-lang/gongsim/go"
	gongsim_fullstack "github.com/fullstack-lang/gongsim/go/fullstack"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_probe "github.com/fullstack-lang/gongsim/go/probe"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongfly: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongfly_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *gongfly_models.StageStruct
	var backRepo *gongfly_orm.BackRepoStruct

	// persistence in a SQLite file on disk in memory
	stage, backRepo = gongfly_fullstack.NewStackInstance(r, "gongfly")

	gongleafletStage, gongleafletBackrepo := gongleaflet_fullstack.NewStackInstance(r, "gongfly")
	gongsimStage, gongsimBackrepo := gongsim_fullstack.NewStackInstance(r, "gongfly")

	simulation := simulation.NewSimulation(stage, gongsimStage, gongleafletStage)

	// the gongsim command orchestrates the simulation engine regarding to the
	// the rest of the stack. It manages when the stage has to be commited to the
	// back repo or when the back repo has to be checked out to the stage
	gongsimCommand := gongsim_models.NewGongSimCommand(gongsimStage, simulation.GetEngine())
	_ = gongsimCommand

	// load all elments
	gongfly_icons.LoadIcons(gongleafletStage)

	// reference.LoadSatellites(stage, simulation.GetEngine())
	reference.LoadCenters(stage, simulation.GetEngine())
	reference.LoadLiners(stage, simulation.GetEngine())

	defaultLayer := new(gongleaflet_models.LayerGroup).Stage(gongleafletStage)
	defaultLayer.Name = "default"
	defaultLayer.DisplayName = "default"
	visuals.LoadLayerGroups(gongleafletStage)
	visuals.LoadLayerGroupsUse(gongleafletStage)
	visuals.LoadMapOptions(gongleafletStage)

	gongfly_visuals.AttachVisualElementsToModelElements(
		stage,
		gongleafletStage, defaultLayer)

	// commits stages
	stage.Commit()
	gongsimStage.Commit()
	gongleafletStage.Commit()

	gongfly_probe.NewProbe(r, gongfly_go.GoModelsDir, gongfly_go.GoDiagramsDir,
		*embeddedDiagrams, "gongfly", stage, backRepo)

	gongleaflet_probe.NewProbe(r, gongleaflet_go.GoModelsDir, gongleaflet_go.GoDiagramsDir,
		*embeddedDiagrams, "gongleaflet", gongleafletStage, gongleafletBackrepo)

	gongsim_probe.NewProbe(r, gongsim_go.GoModelsDir, gongsim_go.GoDiagramsDir,
		*embeddedDiagrams, "gongsim", gongsimStage, gongsimBackrepo)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
