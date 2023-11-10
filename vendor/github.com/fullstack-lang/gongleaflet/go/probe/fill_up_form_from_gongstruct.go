// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Circle Form",
			OnSave: __gong__New__CircleFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DivIcon:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update DivIcon Form",
			OnSave: __gong__New__DivIconFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LayerGroup:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update LayerGroup Form",
			OnSave: __gong__New__LayerGroupFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LayerGroupUse:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update LayerGroupUse Form",
			OnSave: __gong__New__LayerGroupUseFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapOptions:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update MapOptions Form",
			OnSave: __gong__New__MapOptionsFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Marker:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Marker Form",
			OnSave: __gong__New__MarkerFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UserClick:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update UserClick Form",
			OnSave: __gong__New__UserClickFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VLine:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update VLine Form",
			OnSave: __gong__New__VLineFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VisualTrack:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update VisualTrack Form",
			OnSave: __gong__New__VisualTrackFormCallback(
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
