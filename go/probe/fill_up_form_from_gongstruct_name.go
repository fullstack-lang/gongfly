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

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "CivilianAirport":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CivilianAirport Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CivilianAirportFormCallback(
			nil,
			probe,
			formGroup,
		)
		civilianairport := new(models.CivilianAirport)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(civilianairport, formGroup, probe)
	case "Liner":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Liner Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinerFormCallback(
			nil,
			probe,
			formGroup,
		)
		liner := new(models.Liner)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(liner, formGroup, probe)
	case "Message":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Message Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageFormCallback(
			nil,
			probe,
			formGroup,
		)
		message := new(models.Message)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(message, formGroup, probe)
	case "OpsLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "OpsLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpsLineFormCallback(
			nil,
			probe,
			formGroup,
		)
		opsline := new(models.OpsLine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(opsline, formGroup, probe)
	case "Radar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Radar Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RadarFormCallback(
			nil,
			probe,
			formGroup,
		)
		radar := new(models.Radar)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(radar, formGroup, probe)
	case "Satellite":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Satellite Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SatelliteFormCallback(
			nil,
			probe,
			formGroup,
		)
		satellite := new(models.Satellite)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(satellite, formGroup, probe)
	case "Scenario":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Scenario Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScenarioFormCallback(
			nil,
			probe,
			formGroup,
		)
		scenario := new(models.Scenario)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scenario, formGroup, probe)
	}
	formStage.Commit()
}
