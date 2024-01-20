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
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			probe,
			formGroup,
		)
		circle := new(models.Circle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circle, formGroup, probe)
	case "DivIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DivIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DivIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		divicon := new(models.DivIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(divicon, formGroup, probe)
	case "LayerGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LayerGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		layergroup := new(models.LayerGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layergroup, formGroup, probe)
	case "LayerGroupUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LayerGroupUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerGroupUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		layergroupuse := new(models.LayerGroupUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layergroupuse, formGroup, probe)
	case "MapOptions":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MapOptions Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapOptionsFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapoptions := new(models.MapOptions)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapoptions, formGroup, probe)
	case "Marker":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Marker Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkerFormCallback(
			nil,
			probe,
			formGroup,
		)
		marker := new(models.Marker)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(marker, formGroup, probe)
	case "UserClick":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "UserClick Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserClickFormCallback(
			nil,
			probe,
			formGroup,
		)
		userclick := new(models.UserClick)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(userclick, formGroup, probe)
	case "VLine":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "VLine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VLineFormCallback(
			nil,
			probe,
			formGroup,
		)
		vline := new(models.VLine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vline, formGroup, probe)
	case "VisualTrack":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "VisualTrack Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VisualTrackFormCallback(
			nil,
			probe,
			formGroup,
		)
		visualtrack := new(models.VisualTrack)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(visualtrack, formGroup, probe)
	}
	formStage.Commit()
}
