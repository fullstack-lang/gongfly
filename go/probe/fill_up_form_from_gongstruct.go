// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CivilianAirport:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CivilianAirport Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CivilianAirportFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Liner:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Liner Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Message:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Message Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.OpsLine:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "OpsLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpsLineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Radar:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Radar Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RadarFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Satellite:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Satellite Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SatelliteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scenario:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Scenario Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScenarioFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
