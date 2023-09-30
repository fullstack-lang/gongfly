// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
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
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Circle Form",
			OnSave: __gong__New__CircleFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		FillUpForm(circle, formGroup, playground)
	case "DivIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DivIcon Form",
			OnSave: __gong__New__DivIconFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		divicon := new(models.DivIcon)
		FillUpForm(divicon, formGroup, playground)
	case "LayerGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LayerGroup Form",
			OnSave: __gong__New__LayerGroupFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		layergroup := new(models.LayerGroup)
		FillUpForm(layergroup, formGroup, playground)
	case "LayerGroupUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LayerGroupUse Form",
			OnSave: __gong__New__LayerGroupUseFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		layergroupuse := new(models.LayerGroupUse)
		FillUpForm(layergroupuse, formGroup, playground)
	case "MapOptions":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " MapOptions Form",
			OnSave: __gong__New__MapOptionsFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		mapoptions := new(models.MapOptions)
		FillUpForm(mapoptions, formGroup, playground)
	case "Marker":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Marker Form",
			OnSave: __gong__New__MarkerFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		marker := new(models.Marker)
		FillUpForm(marker, formGroup, playground)
	case "UserClick":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " UserClick Form",
			OnSave: __gong__New__UserClickFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		userclick := new(models.UserClick)
		FillUpForm(userclick, formGroup, playground)
	case "VLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " VLine Form",
			OnSave: __gong__New__VLineFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		vline := new(models.VLine)
		FillUpForm(vline, formGroup, playground)
	case "VisualTrack":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " VisualTrack Form",
			OnSave: __gong__New__VisualTrackFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		visualtrack := new(models.VisualTrack)
		FillUpForm(visualtrack, formGroup, playground)
	}
	formStage.Commit()
}
