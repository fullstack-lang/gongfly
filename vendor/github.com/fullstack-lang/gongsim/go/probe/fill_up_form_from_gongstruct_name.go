// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsim/go/models"
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
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "DummyAgent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DummyAgent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DummyAgentFormCallback(
			nil,
			probe,
			formGroup,
		)
		dummyagent := new(models.DummyAgent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dummyagent, formGroup, probe)
	case "Engine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Engine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EngineFormCallback(
			nil,
			probe,
			formGroup,
		)
		engine := new(models.Engine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(engine, formGroup, probe)
	case "Event":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Event Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EventFormCallback(
			nil,
			probe,
			formGroup,
		)
		event := new(models.Event)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(event, formGroup, probe)
	case "GongsimCommand":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GongsimCommand Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongsimCommandFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongsimcommand := new(models.GongsimCommand)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongsimcommand, formGroup, probe)
	case "GongsimStatus":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GongsimStatus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongsimStatusFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongsimstatus := new(models.GongsimStatus)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongsimstatus, formGroup, probe)
	}
	formStage.Commit()
}
