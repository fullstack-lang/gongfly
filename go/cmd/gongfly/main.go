package main

import (
	"flag"
	"log"
	"strconv"

	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_stack "github.com/fullstack-lang/gongfly/go/stack"
	gongfly_static "github.com/fullstack-lang/gongfly/go/static"

	"github.com/fullstack-lang/gongfly/go/visuals"

	gongfly_icons "github.com/fullstack-lang/gongfly/go/icons"
	"github.com/fullstack-lang/gongfly/go/reference"
	"github.com/fullstack-lang/gongfly/go/simulation"

	gongfly_visuals "github.com/fullstack-lang/gongfly/go/visuals"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongleaflet_stack "github.com/fullstack-lang/gongleaflet/go/stack"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_stack "github.com/fullstack-lang/gongsim/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

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
	stack := gongfly_stack.NewStack(r, gongfly_models.GongflyStackName.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()
	stage := stack.Stage
	gongleafletStack := gongleaflet_stack.NewStack(r, gongfly_models.GongLeafleatStackName.ToString(), "", "", "", true, true)
	gongsimStack := gongsim_stack.NewStack(r, gongfly_models.GongsimStackName.ToString(), "", "", "", true, true)

	simulation := simulation.NewSimulation(stage, gongsimStack.Stage, gongleafletStack.Stage)

	// the gongsim command orchestrates the simulation engine regarding to the
	// the rest of the stack. It manages when the stage has to be commited to the
	// back repo or when the back repo has to be checked out to the stage
	gongsimCommand := gongsim_models.NewGongSimCommand(gongsimStack.Stage, simulation.GetEngine())
	_ = gongsimCommand

	// load all elments
	gongfly_icons.LoadIcons(gongleafletStack.Stage)

	// reference.LoadSatellites(stage, simulation.GetEngine())
	reference.LoadCenters(stage, simulation.GetEngine())
	reference.LoadLiners(stage, simulation.GetEngine())

	defaultLayer := new(gongleaflet_models.LayerGroup).Stage(gongleafletStack.Stage)
	defaultLayer.Name = "default"
	defaultLayer.DisplayName = "default"
	visuals.LoadLayerGroups(gongleafletStack.Stage)
	visuals.LoadLayerGroupsUse(gongleafletStack.Stage)
	visuals.LoadMapOptions(gongleafletStack.Stage)

	gongfly_visuals.AttachVisualElementsToModelElements(
		stage,
		gongleafletStack.Stage, defaultLayer)

	// commits stages
	stage.Commit()
	gongsimStack.Stage.Commit()
	gongleafletStack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
