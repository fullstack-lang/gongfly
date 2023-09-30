// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsim/go/models"
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
	case "DummyAgent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DummyAgent Form",
			OnSave: __gong__New__DummyAgentFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		dummyagent := new(models.DummyAgent)
		FillUpForm(dummyagent, formGroup, playground)
	case "Engine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Engine Form",
			OnSave: __gong__New__EngineFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		engine := new(models.Engine)
		FillUpForm(engine, formGroup, playground)
	case "Event":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Event Form",
			OnSave: __gong__New__EventFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		event := new(models.Event)
		FillUpForm(event, formGroup, playground)
	case "GongsimCommand":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongsimCommand Form",
			OnSave: __gong__New__GongsimCommandFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongsimcommand := new(models.GongsimCommand)
		FillUpForm(gongsimcommand, formGroup, playground)
	case "GongsimStatus":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongsimStatus Form",
			OnSave: __gong__New__GongsimStatusFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongsimstatus := new(models.GongsimStatus)
		FillUpForm(gongsimstatus, formGroup, playground)
	}
	formStage.Commit()
}
