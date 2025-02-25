// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DivIcon:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DivIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DivIconFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LayerGroup:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LayerGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerGroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LayerGroupUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LayerGroupUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerGroupUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapOptions:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MapOptions Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapOptionsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Marker:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Marker Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UserClick:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "UserClick Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserClickFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VLine:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "VLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VLineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VisualTrack:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "VisualTrack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VisualTrackFormCallback(
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
