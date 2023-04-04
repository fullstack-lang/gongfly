package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gongfly_go "github.com/fullstack-lang/gongfly/go"
	gongfly_fullstack "github.com/fullstack-lang/gongfly/go/fullstack"
	"github.com/fullstack-lang/gongfly/go/icons"
	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongfly_static "github.com/fullstack-lang/gongfly/go/static"

	gongleaflet_fullstack "github.com/fullstack-lang/gongleaflet/go/fullstack"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
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

	gongleafletStage := gongleaflet_fullstack.NewStackInstance(r, "github.com/fullstack-lang/leaflet/go/models")
	icons.LoadIcons(gongleafletStage)

	if *unmarshallFromCode != "" {
		gongflyStage.Checkout()
		gongflyStage.Reset()
		gongflyStage.Commit()
		err := gongfly_models.ParseAstFile(gongflyStage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		gongflyStage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		gongflyStage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		gongflyStage.OnInitCommitFromFrontCallback = hook
	}

	gongdoc_load.Load(
		"gongfly",
		"github.com/fullstack-lang/gongfly/go/models",
		gongfly_go.GoModelsDir,
		gongfly_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&gongflyStage.Map_GongStructName_InstancesNb)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
