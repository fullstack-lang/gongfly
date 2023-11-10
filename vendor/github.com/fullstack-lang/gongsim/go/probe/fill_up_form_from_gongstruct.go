// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsim/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.DummyAgent:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update DummyAgent Form",
			OnSave: __gong__New__DummyAgentFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Engine:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Engine Form",
			OnSave: __gong__New__EngineFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Event:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Event Form",
			OnSave: __gong__New__EventFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongsimCommand:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongsimCommand Form",
			OnSave: __gong__New__GongsimCommandFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongsimStatus:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongsimStatus Form",
			OnSave: __gong__New__GongsimStatusFormCallback(
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
