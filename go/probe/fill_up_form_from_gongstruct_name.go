// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongfly/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
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
				playground,
			),
		}).Stage(formStage)
		civilianairport := new(models.CivilianAirport)
		FillUpForm(civilianairport, formGroup, playground)
	case "Liner":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Liner Form",
			OnSave: __gong__New__LinerFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		liner := new(models.Liner)
		FillUpForm(liner, formGroup, playground)
	case "Message":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Message Form",
			OnSave: __gong__New__MessageFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		message := new(models.Message)
		FillUpForm(message, formGroup, playground)
	case "OpsLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " OpsLine Form",
			OnSave: __gong__New__OpsLineFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		opsline := new(models.OpsLine)
		FillUpForm(opsline, formGroup, playground)
	case "Radar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Radar Form",
			OnSave: __gong__New__RadarFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		radar := new(models.Radar)
		FillUpForm(radar, formGroup, playground)
	case "Satellite":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Satellite Form",
			OnSave: __gong__New__SatelliteFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		satellite := new(models.Satellite)
		FillUpForm(satellite, formGroup, playground)
	case "Scenario":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Scenario Form",
			OnSave: __gong__New__ScenarioFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		scenario := new(models.Scenario)
		FillUpForm(scenario, formGroup, playground)
	}
	formStage.Commit()
}
