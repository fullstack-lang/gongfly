package gongfly2gongng2charts

import (
	"fmt"
	"time"

	gongfly_models "github.com/fullstack-lang/gongfly/go/models"

	gongng2charts_models "github.com/fullstack-lang/gongng2charts/go/models"
)

// GenerateChartScheduler is the struct of the singloton
// in charge of updating the gantt when there is a new commit from the front or the back
type GenerateChartScheduler struct {
}

var GenerateChartSchedulerSingloton GenerateChartScheduler

// start the scheduler
func init() {
	// go GenerateGanttSchedulerSingloton.checkoutScheduler()
}

func (generateDocumentationScheduler *GenerateChartScheduler) CheckoutScheduler() {

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

			if gongfly_models.Stage.BackRepo != nil {
				newPushFromFront := gongfly_models.Stage.BackRepo.GetLastPushFromFrontNb()
				if lastPushFromFront < newPushFromFront {

					GenerateChartSchedulerSingloton.GenerateDocumentation()
					lastPushFromFront = newPushFromFront
				}
			}
			if gongfly_models.Stage.BackRepo != nil {
				newPushFromBack := gongfly_models.Stage.BackRepo.GetLastCommitFromBackNb()
				if lastPushFromBack < newPushFromBack {

					GenerateChartSchedulerSingloton.GenerateDocumentation()
					lastPushFromBack = newPushFromBack
				}
			}
		}
	}
}

const NbDatasets = 2 // one for lat, one for lng

func (GenerateDocumentationScheduler *GenerateChartScheduler) GenerateDocumentation() {

	// generate documentation
	gongng2charts_models.Stage.Checkout()
	gongng2charts_models.Stage.Reset()
	gongng2charts_models.Stage.Commit()

	var liner *gongfly_models.Liner

	for _liner := range gongfly_models.Stage.Liners {
		liner = _liner
	}

	if liner == nil {
		return
	}

	// Create root element
	chartConfig := (&gongng2charts_models.ChartConfiguration{Name: "LatLng"}).Stage()

	chartConfig.ChartType = gongng2charts_models.LINE

	for idx_dataset := 0; idx_dataset < NbDatasets; idx_dataset = idx_dataset + 1 {
		dataset := (&gongng2charts_models.Dataset{Name: fmt.Sprintf("Dataset %d", 0)}).Stage()
		chartConfig.Datasets = append(chartConfig.Datasets, dataset)

		for idx := 0; idx < len(liner.GetLatPositions()); idx = idx + 1 {
			var datapoint *gongng2charts_models.DataPoint
			if idx_dataset == 0 {
				dataset.Name = "Lat"
				datapoint = (&gongng2charts_models.DataPoint{Name: fmt.Sprintf("Lat %d", idx)}).Stage()
				datapoint.Value = liner.GetLatPositions()[idx]
			}
			if idx_dataset == 1 {
				dataset.Name = "Lng"
				datapoint = (&gongng2charts_models.DataPoint{Name: fmt.Sprintf("Lng %d", idx)}).Stage()
				datapoint.Value = liner.GetLngPositions()[idx]
			}
			dataset.DataPoints = append(dataset.DataPoints, datapoint)
		}

		dataset.Label = fmt.Sprintf("set %d", idx_dataset)
	}

	for idx := 0; idx < len(liner.GetLatPositions()); idx = idx + 1 {
		label := (&gongng2charts_models.Label{Name: fmt.Sprintf("Month %d", idx)}).Stage()
		chartConfig.Labels = append(chartConfig.Labels, label)
	}

	gongng2charts_models.Stage.Commit()
}
