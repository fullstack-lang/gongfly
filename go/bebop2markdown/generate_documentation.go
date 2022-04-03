package bebop2markdown

import (
	bebop_models "bebop/go/models"
	"sort"
	"strings"
	"time"

	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"
)

// GenerateDocumentationScheduler is the struct of the singloton
// in charge of updating the gantt when there is a new commit from the front or the back
type GenerateDocumentationScheduler struct {
}

var GenerateDocumentationSchedulerSingloton GenerateDocumentationScheduler

// start the scheduler
func init() {
	// go GenerateGanttSchedulerSingloton.checkoutScheduler()
}

func (generateDocumentationScheduler *GenerateDocumentationScheduler) CheckoutScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	var CheckoutSchedulerPeriod = time.NewTicker(100 * time.Millisecond)

	lastPushFromFront := uint(0)
	lastPushFromBack := uint(0)
	for {
		select {
		case t := <-CheckoutSchedulerPeriod.C:

			_ = t

			if bebop_models.Stage.BackRepo != nil {
				newPushFromFront := bebop_models.Stage.BackRepo.GetLastPushFromFrontNb()
				if lastPushFromFront < newPushFromFront {

					GenerateDocumentationSchedulerSingloton.GenerateDocumentation()
					lastPushFromFront = newPushFromFront
				}
			}
			if bebop_models.Stage.BackRepo != nil {
				newPushFromBack := bebop_models.Stage.BackRepo.GetLastCommitFromBackNb()
				if lastPushFromBack < newPushFromBack {

					GenerateDocumentationSchedulerSingloton.GenerateDocumentation()
					lastPushFromBack = newPushFromBack
				}
			}
		}
	}
}

func (GenerateDocumentationScheduler *GenerateDocumentationScheduler) GenerateDocumentation() {

	// generate documentation
	gongmarkdown_models.Stage.Checkout()
	gongmarkdown_models.Stage.Reset()
	gongmarkdown_models.Stage.Commit()

	// Create root element
	markdownContent := (&gongmarkdown_models.MarkdownContent{Name: "Dummy"}).Stage()
	markdownContent.Name = "Dummy"

	root := (&gongmarkdown_models.Element{Name: "Root"}).Stage()
	root.Type = gongmarkdown_models.PARAGRAPH
	markdownContent.Root = root

	// generate documentation
	groupTable := []bebop_models.GongStructInterface{}

	for group := range bebop_models.Stage.Groups {
		groupTable = append(groupTable, group)
	}

	sort.Slice(groupTable[:], func(i, j int) bool {
		return groupTable[i].GetName() < groupTable[j].GetName()
	})

	// only display Name
	groupFieldsToDisplay := []FieldHeadOfTable{
		{
			FieldName: "Name",
			TableHead: "Groupe",
		},
		{
			FieldName: "Centres",
			TableHead: "Liste ordonée des centres",
		},
	}

	root.SubElements = append(root.SubElements,
		GongStructSlice2MarkdownTable(groupTable, groupFieldsToDisplay))

	// table of deploiements
	deploiementsTable := []bebop_models.GongStructInterface{}

	for deploiement := range bebop_models.Stage.Deploiements {
		deploiementsTable = append(deploiementsTable, deploiement)
	}

	sort.Slice(deploiementsTable[:], func(i, j int) bool {
		return deploiementsTable[i].GetName() < deploiementsTable[j].GetName()
	})

	// only display Name
	deploiementFieldsToDisplay := []FieldHeadOfTable{
		{
			FieldName: "Name",
			TableHead: "Déploiement",
		},
		{
			FieldName: "Commentaire",
			TableHead: "Description du déploiement",
		},
		{
			FieldName: "DatesAuto",
			TableHead: "Dates calculées",
		},
		{
			FieldName: "Csds",
			TableHead: "Liste des CSD",
		},
	}

	root.SubElements = append(root.SubElements,
		GongStructSlice2MarkdownTable(deploiementsTable, deploiementFieldsToDisplay))

	markdownContent.UpdateContent()

	gongmarkdown_models.Stage.Commit()
}

type FieldHeadOfTable struct {

	// FieldName of field to display
	FieldName string

	// name of the displayed
	TableHead string
}

func GongStructSlice2MarkdownTable(gongStructInstances []bebop_models.GongStructInterface, fieldsToDisplay []FieldHeadOfTable) (element *gongmarkdown_models.Element) {

	if len(gongStructInstances) == 0 {
		return
	}

	//
	element = (&gongmarkdown_models.Element{Name: "Table "}).Stage()
	element.Type = gongmarkdown_models.TABLE

	// fill up title cells
	titleRow := (&gongmarkdown_models.Row{Name: "Title Row "}).Stage()
	element.Rows = append(element.Rows, titleRow)

	// fields := gongStructInstances[0].GetFields()
	for _, field := range fieldsToDisplay {
		titleRow.Cells = append(titleRow.Cells, (&gongmarkdown_models.Cell{Name: field.TableHead}).Stage())
	}

	for _, gongStructInstance := range gongStructInstances {
		// fill up title cells
		row := (&gongmarkdown_models.Row{Name: "Date Row of " + gongStructInstance.GetFieldStringValue("Name")}).Stage()
		element.Rows = append(element.Rows, row)

		for _, field := range fieldsToDisplay {

			stringValue := gongStructInstance.GetFieldStringValue(field.FieldName)

			// replace new lines with <br>
			stringValue = strings.ReplaceAll(stringValue, "\n", "<br>")

			cell := (&gongmarkdown_models.Cell{Name: stringValue}).Stage()
			row.Cells = append(row.Cells, cell)

		}
	}

	return
}
