package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gongfly_go "github.com/fullstack-lang/gongfly/go"
	gongfly_fullstack "github.com/fullstack-lang/gongfly/go/fullstack"
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_static "github.com/fullstack-lang/gongfly/go/static"

	gongfly_icons "github.com/fullstack-lang/gongfly/go/icons"
	"github.com/fullstack-lang/gongfly/go/reference"
	"github.com/fullstack-lang/gongfly/go/simulation"

	gongfly_visuals "github.com/fullstack-lang/gongfly/go/visuals"

	gongleaflet_fullstack "github.com/fullstack-lang/gongleaflet/go/fullstack"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	gongsim_go "github.com/fullstack-lang/gongsim/go"
	gongsim_fullstack "github.com/fullstack-lang/gongsim/go/fullstack"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongfly_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gongfly/go/models", "main")
}

func main() {

	log.SetPrefix("gongfly: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongfly_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	gongflyStage := gongfly_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongfly/go/models")
	gongleafletStage := gongleaflet_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongleaflet/go/models")
	gongsimStage := gongsim_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsim/go/models")
	_ = gongsimStage

	simulation := simulation.NewSimulation(gongflyStage, gongsimStage, gongleafletStage)

	// load all elments
	gongfly_icons.LoadIcons(gongleafletStage)
	reference.LoadCenters(gongflyStage, simulation.GetEngine())
	reference.LoadLiners(gongflyStage, simulation.GetEngine())
	reference.LoadSatellites(gongflyStage, simulation.GetEngine())
	reference.LoadScenario(gongflyStage)

	defaultLayer := new(gongleaflet_models.LayerGroup).Stage(gongleafletStage)
	defaultLayer.Name = "default"
	defaultLayer.DisplayName = "default"

	gongfly_visuals.AttachVisualElementsToModelElements(
		gongflyStage,
		gongleafletStage, defaultLayer)

	gongdoc_load.Load(
		"gongfly",
		"github.com/fullstack-lang/gongfly/go/models",
		gongfly_go.GoModelsDir,
		gongfly_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&gongflyStage.Map_GongStructName_InstancesNb)

	gongdoc_load.Load(
		"gongsim",
		"github.com/fullstack-lang/gongsim/go/models",
		gongsim_go.GoModelsDir,
		gongsim_go.GoDiagramsDir,
		r,
		true,
		&gongflyStage.Map_GongStructName_InstancesNb)

	// commits stages
	gongflyStage.Commit()
	gongsimStage.Commit()
	gongleafletStage.Commit()

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
