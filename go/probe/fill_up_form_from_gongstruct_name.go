// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "CivilianAirport":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CivilianAirport Form",
			OnSave: __gong__New__CivilianAirportFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		civilianairport := new(models.CivilianAirport)
		FillUpForm(civilianairport, formGroup, probe)
	case "Liner":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Liner Form",
			OnSave: __gong__New__LinerFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		liner := new(models.Liner)
		FillUpForm(liner, formGroup, probe)
	case "Message":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Message Form",
			OnSave: __gong__New__MessageFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		message := new(models.Message)
		FillUpForm(message, formGroup, probe)
	case "OpsLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " OpsLine Form",
			OnSave: __gong__New__OpsLineFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		opsline := new(models.OpsLine)
		FillUpForm(opsline, formGroup, probe)
	case "Radar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Radar Form",
			OnSave: __gong__New__RadarFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		radar := new(models.Radar)
		FillUpForm(radar, formGroup, probe)
	case "Satellite":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Satellite Form",
			OnSave: __gong__New__SatelliteFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		satellite := new(models.Satellite)
		FillUpForm(satellite, formGroup, probe)
	case "Scenario":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Scenario Form",
			OnSave: __gong__New__ScenarioFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		scenario := new(models.Scenario)
		FillUpForm(scenario, formGroup, probe)
	}
	formStage.Commit()
}
