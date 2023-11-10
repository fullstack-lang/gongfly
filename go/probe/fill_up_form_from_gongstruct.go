// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CivilianAirport Form",
			OnSave: __gong__New__CivilianAirportFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Liner:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Liner Form",
			OnSave: __gong__New__LinerFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Message:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Message Form",
			OnSave: __gong__New__MessageFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.OpsLine:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update OpsLine Form",
			OnSave: __gong__New__OpsLineFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Radar:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Radar Form",
			OnSave: __gong__New__RadarFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Satellite:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Satellite Form",
			OnSave: __gong__New__SatelliteFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scenario:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Scenario Form",
			OnSave: __gong__New__ScenarioFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
