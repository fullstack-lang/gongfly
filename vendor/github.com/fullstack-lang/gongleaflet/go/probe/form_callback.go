// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongleaflet/go/models"
	"github.com/fullstack-lang/gongleaflet/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CircleFormCallback(
	circle *models.Circle,
	probe *Probe,
	formGroup *table.FormGroup,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.probe = probe
	circleFormCallback.circle = circle
	circleFormCallback.formGroup = formGroup

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.probe.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.probe.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	for _, formDiv := range circleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(circle_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(circle_.Lng), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(circle_.Radius), formDiv)
		case "ColorEnum":
			FormDivEnumStringFieldToField(&(circle_.ColorEnum), formDiv)
		case "DashStyleEnum":
			FormDivEnumStringFieldToField(&(circle_.DashStyleEnum), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(circle_.LayerGroup), circleFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circle_.Unstage(circleFormCallback.probe.stageOfInterest)
	}

	circleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.probe,
	)
	circleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode || circleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(circleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			circleFormCallback.probe,
			newFormGroup,
		)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.probe)
		circleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(circleFormCallback.probe)
}
func __gong__New__DivIconFormCallback(
	divicon *models.DivIcon,
	probe *Probe,
	formGroup *table.FormGroup,
) (diviconFormCallback *DivIconFormCallback) {
	diviconFormCallback = new(DivIconFormCallback)
	diviconFormCallback.probe = probe
	diviconFormCallback.divicon = divicon
	diviconFormCallback.formGroup = formGroup

	diviconFormCallback.CreationMode = (divicon == nil)

	return
}

type DivIconFormCallback struct {
	divicon *models.DivIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diviconFormCallback *DivIconFormCallback) OnSave() {

	log.Println("DivIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diviconFormCallback.probe.formStage.Checkout()

	if diviconFormCallback.divicon == nil {
		diviconFormCallback.divicon = new(models.DivIcon).Stage(diviconFormCallback.probe.stageOfInterest)
	}
	divicon_ := diviconFormCallback.divicon
	_ = divicon_

	for _, formDiv := range diviconFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(divicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(divicon_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if diviconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		divicon_.Unstage(diviconFormCallback.probe.stageOfInterest)
	}

	diviconFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DivIcon](
		diviconFormCallback.probe,
	)
	diviconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if diviconFormCallback.CreationMode || diviconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diviconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(diviconFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DivIconFormCallback(
			nil,
			diviconFormCallback.probe,
			newFormGroup,
		)
		divicon := new(models.DivIcon)
		FillUpForm(divicon, newFormGroup, diviconFormCallback.probe)
		diviconFormCallback.probe.formStage.Commit()
	}

	fillUpTree(diviconFormCallback.probe)
}
func __gong__New__LayerGroupFormCallback(
	layergroup *models.LayerGroup,
	probe *Probe,
	formGroup *table.FormGroup,
) (layergroupFormCallback *LayerGroupFormCallback) {
	layergroupFormCallback = new(LayerGroupFormCallback)
	layergroupFormCallback.probe = probe
	layergroupFormCallback.layergroup = layergroup
	layergroupFormCallback.formGroup = formGroup

	layergroupFormCallback.CreationMode = (layergroup == nil)

	return
}

type LayerGroupFormCallback struct {
	layergroup *models.LayerGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (layergroupFormCallback *LayerGroupFormCallback) OnSave() {

	log.Println("LayerGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layergroupFormCallback.probe.formStage.Checkout()

	if layergroupFormCallback.layergroup == nil {
		layergroupFormCallback.layergroup = new(models.LayerGroup).Stage(layergroupFormCallback.probe.stageOfInterest)
	}
	layergroup_ := layergroupFormCallback.layergroup
	_ = layergroup_

	for _, formDiv := range layergroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layergroup_.Name), formDiv)
		case "DisplayName":
			FormDivBasicFieldToField(&(layergroup_.DisplayName), formDiv)
		}
	}

	// manage the suppress operation
	if layergroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layergroup_.Unstage(layergroupFormCallback.probe.stageOfInterest)
	}

	layergroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LayerGroup](
		layergroupFormCallback.probe,
	)
	layergroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupFormCallback.CreationMode || layergroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layergroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(layergroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LayerGroupFormCallback(
			nil,
			layergroupFormCallback.probe,
			newFormGroup,
		)
		layergroup := new(models.LayerGroup)
		FillUpForm(layergroup, newFormGroup, layergroupFormCallback.probe)
		layergroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(layergroupFormCallback.probe)
}
func __gong__New__LayerGroupUseFormCallback(
	layergroupuse *models.LayerGroupUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (layergroupuseFormCallback *LayerGroupUseFormCallback) {
	layergroupuseFormCallback = new(LayerGroupUseFormCallback)
	layergroupuseFormCallback.probe = probe
	layergroupuseFormCallback.layergroupuse = layergroupuse
	layergroupuseFormCallback.formGroup = formGroup

	layergroupuseFormCallback.CreationMode = (layergroupuse == nil)

	return
}

type LayerGroupUseFormCallback struct {
	layergroupuse *models.LayerGroupUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (layergroupuseFormCallback *LayerGroupUseFormCallback) OnSave() {

	log.Println("LayerGroupUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layergroupuseFormCallback.probe.formStage.Checkout()

	if layergroupuseFormCallback.layergroupuse == nil {
		layergroupuseFormCallback.layergroupuse = new(models.LayerGroupUse).Stage(layergroupuseFormCallback.probe.stageOfInterest)
	}
	layergroupuse_ := layergroupuseFormCallback.layergroupuse
	_ = layergroupuse_

	for _, formDiv := range layergroupuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layergroupuse_.Name), formDiv)
		case "IsDisplayed":
			FormDivBasicFieldToField(&(layergroupuse_.IsDisplayed), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(layergroupuse_.LayerGroup), layergroupuseFormCallback.probe.stageOfInterest, formDiv)
		case "MapOptions:LayerGroupUses":
			// we need to retrieve the field owner before the change
			var pastMapOptionsOwner *models.MapOptions
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "MapOptions"
			rf.Fieldname = "LayerGroupUses"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				layergroupuseFormCallback.probe.stageOfInterest,
				layergroupuseFormCallback.probe.backRepoOfInterest,
				layergroupuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastMapOptionsOwner = reverseFieldOwner.(*models.MapOptions)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMapOptionsOwner != nil {
					idx := slices.Index(pastMapOptionsOwner.LayerGroupUses, layergroupuse_)
					pastMapOptionsOwner.LayerGroupUses = slices.Delete(pastMapOptionsOwner.LayerGroupUses, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _mapoptions := range *models.GetGongstructInstancesSet[models.MapOptions](layergroupuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _mapoptions.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMapOptionsOwner := _mapoptions // we have a match
						if pastMapOptionsOwner != nil {
							if newMapOptionsOwner != pastMapOptionsOwner {
								idx := slices.Index(pastMapOptionsOwner.LayerGroupUses, layergroupuse_)
								pastMapOptionsOwner.LayerGroupUses = slices.Delete(pastMapOptionsOwner.LayerGroupUses, idx, idx+1)
								newMapOptionsOwner.LayerGroupUses = append(newMapOptionsOwner.LayerGroupUses, layergroupuse_)
							}
						} else {
							newMapOptionsOwner.LayerGroupUses = append(newMapOptionsOwner.LayerGroupUses, layergroupuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if layergroupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layergroupuse_.Unstage(layergroupuseFormCallback.probe.stageOfInterest)
	}

	layergroupuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LayerGroupUse](
		layergroupuseFormCallback.probe,
	)
	layergroupuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupuseFormCallback.CreationMode || layergroupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layergroupuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(layergroupuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LayerGroupUseFormCallback(
			nil,
			layergroupuseFormCallback.probe,
			newFormGroup,
		)
		layergroupuse := new(models.LayerGroupUse)
		FillUpForm(layergroupuse, newFormGroup, layergroupuseFormCallback.probe)
		layergroupuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(layergroupuseFormCallback.probe)
}
func __gong__New__MapOptionsFormCallback(
	mapoptions *models.MapOptions,
	probe *Probe,
	formGroup *table.FormGroup,
) (mapoptionsFormCallback *MapOptionsFormCallback) {
	mapoptionsFormCallback = new(MapOptionsFormCallback)
	mapoptionsFormCallback.probe = probe
	mapoptionsFormCallback.mapoptions = mapoptions
	mapoptionsFormCallback.formGroup = formGroup

	mapoptionsFormCallback.CreationMode = (mapoptions == nil)

	return
}

type MapOptionsFormCallback struct {
	mapoptions *models.MapOptions

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mapoptionsFormCallback *MapOptionsFormCallback) OnSave() {

	log.Println("MapOptionsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapoptionsFormCallback.probe.formStage.Checkout()

	if mapoptionsFormCallback.mapoptions == nil {
		mapoptionsFormCallback.mapoptions = new(models.MapOptions).Stage(mapoptionsFormCallback.probe.stageOfInterest)
	}
	mapoptions_ := mapoptionsFormCallback.mapoptions
	_ = mapoptions_

	for _, formDiv := range mapoptionsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(mapoptions_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(mapoptions_.Lng), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(mapoptions_.Name), formDiv)
		case "ZoomLevel":
			FormDivBasicFieldToField(&(mapoptions_.ZoomLevel), formDiv)
		case "UrlTemplate":
			FormDivBasicFieldToField(&(mapoptions_.UrlTemplate), formDiv)
		case "Attribution":
			FormDivBasicFieldToField(&(mapoptions_.Attribution), formDiv)
		case "MaxZoom":
			FormDivBasicFieldToField(&(mapoptions_.MaxZoom), formDiv)
		case "ZoomControl":
			FormDivBasicFieldToField(&(mapoptions_.ZoomControl), formDiv)
		case "AttributionControl":
			FormDivBasicFieldToField(&(mapoptions_.AttributionControl), formDiv)
		case "ZoomSnap":
			FormDivBasicFieldToField(&(mapoptions_.ZoomSnap), formDiv)
		}
	}

	// manage the suppress operation
	if mapoptionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapoptions_.Unstage(mapoptionsFormCallback.probe.stageOfInterest)
	}

	mapoptionsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MapOptions](
		mapoptionsFormCallback.probe,
	)
	mapoptionsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mapoptionsFormCallback.CreationMode || mapoptionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapoptionsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mapoptionsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MapOptionsFormCallback(
			nil,
			mapoptionsFormCallback.probe,
			newFormGroup,
		)
		mapoptions := new(models.MapOptions)
		FillUpForm(mapoptions, newFormGroup, mapoptionsFormCallback.probe)
		mapoptionsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mapoptionsFormCallback.probe)
}
func __gong__New__MarkerFormCallback(
	marker *models.Marker,
	probe *Probe,
	formGroup *table.FormGroup,
) (markerFormCallback *MarkerFormCallback) {
	markerFormCallback = new(MarkerFormCallback)
	markerFormCallback.probe = probe
	markerFormCallback.marker = marker
	markerFormCallback.formGroup = formGroup

	markerFormCallback.CreationMode = (marker == nil)

	return
}

type MarkerFormCallback struct {
	marker *models.Marker

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (markerFormCallback *MarkerFormCallback) OnSave() {

	log.Println("MarkerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	markerFormCallback.probe.formStage.Checkout()

	if markerFormCallback.marker == nil {
		markerFormCallback.marker = new(models.Marker).Stage(markerFormCallback.probe.stageOfInterest)
	}
	marker_ := markerFormCallback.marker
	_ = marker_

	for _, formDiv := range markerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(marker_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(marker_.Lng), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(marker_.Name), formDiv)
		case "ColorEnum":
			FormDivEnumStringFieldToField(&(marker_.ColorEnum), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(marker_.LayerGroup), markerFormCallback.probe.stageOfInterest, formDiv)
		case "DivIcon":
			FormDivSelectFieldToField(&(marker_.DivIcon), markerFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if markerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		marker_.Unstage(markerFormCallback.probe.stageOfInterest)
	}

	markerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Marker](
		markerFormCallback.probe,
	)
	markerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if markerFormCallback.CreationMode || markerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		markerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(markerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MarkerFormCallback(
			nil,
			markerFormCallback.probe,
			newFormGroup,
		)
		marker := new(models.Marker)
		FillUpForm(marker, newFormGroup, markerFormCallback.probe)
		markerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(markerFormCallback.probe)
}
func __gong__New__UserClickFormCallback(
	userclick *models.UserClick,
	probe *Probe,
	formGroup *table.FormGroup,
) (userclickFormCallback *UserClickFormCallback) {
	userclickFormCallback = new(UserClickFormCallback)
	userclickFormCallback.probe = probe
	userclickFormCallback.userclick = userclick
	userclickFormCallback.formGroup = formGroup

	userclickFormCallback.CreationMode = (userclick == nil)

	return
}

type UserClickFormCallback struct {
	userclick *models.UserClick

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (userclickFormCallback *UserClickFormCallback) OnSave() {

	log.Println("UserClickFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	userclickFormCallback.probe.formStage.Checkout()

	if userclickFormCallback.userclick == nil {
		userclickFormCallback.userclick = new(models.UserClick).Stage(userclickFormCallback.probe.stageOfInterest)
	}
	userclick_ := userclickFormCallback.userclick
	_ = userclick_

	for _, formDiv := range userclickFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(userclick_.Name), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(userclick_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(userclick_.Lng), formDiv)
		case "TimeOfClick":
			FormDivBasicFieldToField(&(userclick_.TimeOfClick), formDiv)
		}
	}

	// manage the suppress operation
	if userclickFormCallback.formGroup.HasSuppressButtonBeenPressed {
		userclick_.Unstage(userclickFormCallback.probe.stageOfInterest)
	}

	userclickFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.UserClick](
		userclickFormCallback.probe,
	)
	userclickFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if userclickFormCallback.CreationMode || userclickFormCallback.formGroup.HasSuppressButtonBeenPressed {
		userclickFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(userclickFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UserClickFormCallback(
			nil,
			userclickFormCallback.probe,
			newFormGroup,
		)
		userclick := new(models.UserClick)
		FillUpForm(userclick, newFormGroup, userclickFormCallback.probe)
		userclickFormCallback.probe.formStage.Commit()
	}

	fillUpTree(userclickFormCallback.probe)
}
func __gong__New__VLineFormCallback(
	vline *models.VLine,
	probe *Probe,
	formGroup *table.FormGroup,
) (vlineFormCallback *VLineFormCallback) {
	vlineFormCallback = new(VLineFormCallback)
	vlineFormCallback.probe = probe
	vlineFormCallback.vline = vline
	vlineFormCallback.formGroup = formGroup

	vlineFormCallback.CreationMode = (vline == nil)

	return
}

type VLineFormCallback struct {
	vline *models.VLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (vlineFormCallback *VLineFormCallback) OnSave() {

	log.Println("VLineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	vlineFormCallback.probe.formStage.Checkout()

	if vlineFormCallback.vline == nil {
		vlineFormCallback.vline = new(models.VLine).Stage(vlineFormCallback.probe.stageOfInterest)
	}
	vline_ := vlineFormCallback.vline
	_ = vline_

	for _, formDiv := range vlineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "StartLat":
			FormDivBasicFieldToField(&(vline_.StartLat), formDiv)
		case "StartLng":
			FormDivBasicFieldToField(&(vline_.StartLng), formDiv)
		case "EndLat":
			FormDivBasicFieldToField(&(vline_.EndLat), formDiv)
		case "EndLng":
			FormDivBasicFieldToField(&(vline_.EndLng), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(vline_.Name), formDiv)
		case "ColorEnum":
			FormDivEnumStringFieldToField(&(vline_.ColorEnum), formDiv)
		case "DashStyleEnum":
			FormDivEnumStringFieldToField(&(vline_.DashStyleEnum), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(vline_.LayerGroup), vlineFormCallback.probe.stageOfInterest, formDiv)
		case "IsTransmitting":
			FormDivEnumStringFieldToField(&(vline_.IsTransmitting), formDiv)
		case "Message":
			FormDivBasicFieldToField(&(vline_.Message), formDiv)
		case "IsTransmittingBackward":
			FormDivEnumStringFieldToField(&(vline_.IsTransmittingBackward), formDiv)
		case "MessageBackward":
			FormDivBasicFieldToField(&(vline_.MessageBackward), formDiv)
		}
	}

	// manage the suppress operation
	if vlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		vline_.Unstage(vlineFormCallback.probe.stageOfInterest)
	}

	vlineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VLine](
		vlineFormCallback.probe,
	)
	vlineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if vlineFormCallback.CreationMode || vlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		vlineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(vlineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VLineFormCallback(
			nil,
			vlineFormCallback.probe,
			newFormGroup,
		)
		vline := new(models.VLine)
		FillUpForm(vline, newFormGroup, vlineFormCallback.probe)
		vlineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(vlineFormCallback.probe)
}
func __gong__New__VisualTrackFormCallback(
	visualtrack *models.VisualTrack,
	probe *Probe,
	formGroup *table.FormGroup,
) (visualtrackFormCallback *VisualTrackFormCallback) {
	visualtrackFormCallback = new(VisualTrackFormCallback)
	visualtrackFormCallback.probe = probe
	visualtrackFormCallback.visualtrack = visualtrack
	visualtrackFormCallback.formGroup = formGroup

	visualtrackFormCallback.CreationMode = (visualtrack == nil)

	return
}

type VisualTrackFormCallback struct {
	visualtrack *models.VisualTrack

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (visualtrackFormCallback *VisualTrackFormCallback) OnSave() {

	log.Println("VisualTrackFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	visualtrackFormCallback.probe.formStage.Checkout()

	if visualtrackFormCallback.visualtrack == nil {
		visualtrackFormCallback.visualtrack = new(models.VisualTrack).Stage(visualtrackFormCallback.probe.stageOfInterest)
	}
	visualtrack_ := visualtrackFormCallback.visualtrack
	_ = visualtrack_

	for _, formDiv := range visualtrackFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Lat":
			FormDivBasicFieldToField(&(visualtrack_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(visualtrack_.Lng), formDiv)
		case "Heading":
			FormDivBasicFieldToField(&(visualtrack_.Heading), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(visualtrack_.Level), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(visualtrack_.Speed), formDiv)
		case "VerticalSpeed":
			FormDivBasicFieldToField(&(visualtrack_.VerticalSpeed), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(visualtrack_.Name), formDiv)
		case "ColorEnum":
			FormDivEnumStringFieldToField(&(visualtrack_.ColorEnum), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(visualtrack_.LayerGroup), visualtrackFormCallback.probe.stageOfInterest, formDiv)
		case "DivIcon":
			FormDivSelectFieldToField(&(visualtrack_.DivIcon), visualtrackFormCallback.probe.stageOfInterest, formDiv)
		case "DisplayTrackHistory":
			FormDivBasicFieldToField(&(visualtrack_.DisplayTrackHistory), formDiv)
		case "DisplayLevelAndSpeed":
			FormDivBasicFieldToField(&(visualtrack_.DisplayLevelAndSpeed), formDiv)
		}
	}

	// manage the suppress operation
	if visualtrackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		visualtrack_.Unstage(visualtrackFormCallback.probe.stageOfInterest)
	}

	visualtrackFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VisualTrack](
		visualtrackFormCallback.probe,
	)
	visualtrackFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if visualtrackFormCallback.CreationMode || visualtrackFormCallback.formGroup.HasSuppressButtonBeenPressed {
		visualtrackFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(visualtrackFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VisualTrackFormCallback(
			nil,
			visualtrackFormCallback.probe,
			newFormGroup,
		)
		visualtrack := new(models.VisualTrack)
		FillUpForm(visualtrack, newFormGroup, visualtrackFormCallback.probe)
		visualtrackFormCallback.probe.formStage.Commit()
	}

	fillUpTree(visualtrackFormCallback.probe)
}
