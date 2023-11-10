// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
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
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Circle Form",
			OnSave: __gong__New__CircleFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		FillUpForm(circle, formGroup, probe)
	case "DivIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DivIcon Form",
			OnSave: __gong__New__DivIconFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		divicon := new(models.DivIcon)
		FillUpForm(divicon, formGroup, probe)
	case "LayerGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LayerGroup Form",
			OnSave: __gong__New__LayerGroupFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		layergroup := new(models.LayerGroup)
		FillUpForm(layergroup, formGroup, probe)
	case "LayerGroupUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LayerGroupUse Form",
			OnSave: __gong__New__LayerGroupUseFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		layergroupuse := new(models.LayerGroupUse)
		FillUpForm(layergroupuse, formGroup, probe)
	case "MapOptions":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " MapOptions Form",
			OnSave: __gong__New__MapOptionsFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		mapoptions := new(models.MapOptions)
		FillUpForm(mapoptions, formGroup, probe)
	case "Marker":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Marker Form",
			OnSave: __gong__New__MarkerFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		marker := new(models.Marker)
		FillUpForm(marker, formGroup, probe)
	case "UserClick":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " UserClick Form",
			OnSave: __gong__New__UserClickFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		userclick := new(models.UserClick)
		FillUpForm(userclick, formGroup, probe)
	case "VLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " VLine Form",
			OnSave: __gong__New__VLineFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		vline := new(models.VLine)
		FillUpForm(vline, formGroup, probe)
	case "VisualTrack":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " VisualTrack Form",
			OnSave: __gong__New__VisualTrackFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		visualtrack := new(models.VisualTrack)
		FillUpForm(visualtrack, formGroup, probe)
	}
	formStage.Commit()
}
