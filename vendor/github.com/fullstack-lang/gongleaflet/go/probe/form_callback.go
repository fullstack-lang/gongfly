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
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.probe = probe
	circleFormCallback.circle = circle

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := circleFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	circleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.probe,
	)
	circleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode {
		circleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CircleFormCallback(
				nil,
				circleFormCallback.probe,
			),
		}).Stage(circleFormCallback.probe.formStage)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.probe)
		circleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(circleFormCallback.probe)
}
func __gong__New__DivIconFormCallback(
	divicon *models.DivIcon,
	probe *Probe,
) (diviconFormCallback *DivIconFormCallback) {
	diviconFormCallback = new(DivIconFormCallback)
	diviconFormCallback.probe = probe
	diviconFormCallback.divicon = divicon

	diviconFormCallback.CreationMode = (divicon == nil)

	return
}

type DivIconFormCallback struct {
	divicon *models.DivIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := diviconFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(divicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(divicon_.SVG), formDiv)
		}
	}

	diviconFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DivIcon](
		diviconFormCallback.probe,
	)
	diviconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if diviconFormCallback.CreationMode {
		diviconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DivIconFormCallback(
				nil,
				diviconFormCallback.probe,
			),
		}).Stage(diviconFormCallback.probe.formStage)
		divicon := new(models.DivIcon)
		FillUpForm(divicon, newFormGroup, diviconFormCallback.probe)
		diviconFormCallback.probe.formStage.Commit()
	}

	fillUpTree(diviconFormCallback.probe)
}
func __gong__New__LayerGroupFormCallback(
	layergroup *models.LayerGroup,
	probe *Probe,
) (layergroupFormCallback *LayerGroupFormCallback) {
	layergroupFormCallback = new(LayerGroupFormCallback)
	layergroupFormCallback.probe = probe
	layergroupFormCallback.layergroup = layergroup

	layergroupFormCallback.CreationMode = (layergroup == nil)

	return
}

type LayerGroupFormCallback struct {
	layergroup *models.LayerGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := layergroupFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layergroup_.Name), formDiv)
		case "DisplayName":
			FormDivBasicFieldToField(&(layergroup_.DisplayName), formDiv)
		}
	}

	layergroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LayerGroup](
		layergroupFormCallback.probe,
	)
	layergroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupFormCallback.CreationMode {
		layergroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LayerGroupFormCallback(
				nil,
				layergroupFormCallback.probe,
			),
		}).Stage(layergroupFormCallback.probe.formStage)
		layergroup := new(models.LayerGroup)
		FillUpForm(layergroup, newFormGroup, layergroupFormCallback.probe)
		layergroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(layergroupFormCallback.probe)
}
func __gong__New__LayerGroupUseFormCallback(
	layergroupuse *models.LayerGroupUse,
	probe *Probe,
) (layergroupuseFormCallback *LayerGroupUseFormCallback) {
	layergroupuseFormCallback = new(LayerGroupUseFormCallback)
	layergroupuseFormCallback.probe = probe
	layergroupuseFormCallback.layergroupuse = layergroupuse

	layergroupuseFormCallback.CreationMode = (layergroupuse == nil)

	return
}

type LayerGroupUseFormCallback struct {
	layergroupuse *models.LayerGroupUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := layergroupuseFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	layergroupuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LayerGroupUse](
		layergroupuseFormCallback.probe,
	)
	layergroupuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupuseFormCallback.CreationMode {
		layergroupuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LayerGroupUseFormCallback(
				nil,
				layergroupuseFormCallback.probe,
			),
		}).Stage(layergroupuseFormCallback.probe.formStage)
		layergroupuse := new(models.LayerGroupUse)
		FillUpForm(layergroupuse, newFormGroup, layergroupuseFormCallback.probe)
		layergroupuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(layergroupuseFormCallback.probe)
}
func __gong__New__MapOptionsFormCallback(
	mapoptions *models.MapOptions,
	probe *Probe,
) (mapoptionsFormCallback *MapOptionsFormCallback) {
	mapoptionsFormCallback = new(MapOptionsFormCallback)
	mapoptionsFormCallback.probe = probe
	mapoptionsFormCallback.mapoptions = mapoptions

	mapoptionsFormCallback.CreationMode = (mapoptions == nil)

	return
}

type MapOptionsFormCallback struct {
	mapoptions *models.MapOptions

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := mapoptionsFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	mapoptionsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MapOptions](
		mapoptionsFormCallback.probe,
	)
	mapoptionsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mapoptionsFormCallback.CreationMode {
		mapoptionsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MapOptionsFormCallback(
				nil,
				mapoptionsFormCallback.probe,
			),
		}).Stage(mapoptionsFormCallback.probe.formStage)
		mapoptions := new(models.MapOptions)
		FillUpForm(mapoptions, newFormGroup, mapoptionsFormCallback.probe)
		mapoptionsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mapoptionsFormCallback.probe)
}
func __gong__New__MarkerFormCallback(
	marker *models.Marker,
	probe *Probe,
) (markerFormCallback *MarkerFormCallback) {
	markerFormCallback = new(MarkerFormCallback)
	markerFormCallback.probe = probe
	markerFormCallback.marker = marker

	markerFormCallback.CreationMode = (marker == nil)

	return
}

type MarkerFormCallback struct {
	marker *models.Marker

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := markerFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	markerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Marker](
		markerFormCallback.probe,
	)
	markerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if markerFormCallback.CreationMode {
		markerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MarkerFormCallback(
				nil,
				markerFormCallback.probe,
			),
		}).Stage(markerFormCallback.probe.formStage)
		marker := new(models.Marker)
		FillUpForm(marker, newFormGroup, markerFormCallback.probe)
		markerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(markerFormCallback.probe)
}
func __gong__New__UserClickFormCallback(
	userclick *models.UserClick,
	probe *Probe,
) (userclickFormCallback *UserClickFormCallback) {
	userclickFormCallback = new(UserClickFormCallback)
	userclickFormCallback.probe = probe
	userclickFormCallback.userclick = userclick

	userclickFormCallback.CreationMode = (userclick == nil)

	return
}

type UserClickFormCallback struct {
	userclick *models.UserClick

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := userclickFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	userclickFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.UserClick](
		userclickFormCallback.probe,
	)
	userclickFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if userclickFormCallback.CreationMode {
		userclickFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__UserClickFormCallback(
				nil,
				userclickFormCallback.probe,
			),
		}).Stage(userclickFormCallback.probe.formStage)
		userclick := new(models.UserClick)
		FillUpForm(userclick, newFormGroup, userclickFormCallback.probe)
		userclickFormCallback.probe.formStage.Commit()
	}

	fillUpTree(userclickFormCallback.probe)
}
func __gong__New__VLineFormCallback(
	vline *models.VLine,
	probe *Probe,
) (vlineFormCallback *VLineFormCallback) {
	vlineFormCallback = new(VLineFormCallback)
	vlineFormCallback.probe = probe
	vlineFormCallback.vline = vline

	vlineFormCallback.CreationMode = (vline == nil)

	return
}

type VLineFormCallback struct {
	vline *models.VLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := vlineFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	vlineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VLine](
		vlineFormCallback.probe,
	)
	vlineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if vlineFormCallback.CreationMode {
		vlineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__VLineFormCallback(
				nil,
				vlineFormCallback.probe,
			),
		}).Stage(vlineFormCallback.probe.formStage)
		vline := new(models.VLine)
		FillUpForm(vline, newFormGroup, vlineFormCallback.probe)
		vlineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(vlineFormCallback.probe)
}
func __gong__New__VisualTrackFormCallback(
	visualtrack *models.VisualTrack,
	probe *Probe,
) (visualtrackFormCallback *VisualTrackFormCallback) {
	visualtrackFormCallback = new(VisualTrackFormCallback)
	visualtrackFormCallback.probe = probe
	visualtrackFormCallback.visualtrack = visualtrack

	visualtrackFormCallback.CreationMode = (visualtrack == nil)

	return
}

type VisualTrackFormCallback struct {
	visualtrack *models.VisualTrack

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
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

	// get the formGroup
	formGroup := visualtrackFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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

	visualtrackFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VisualTrack](
		visualtrackFormCallback.probe,
	)
	visualtrackFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if visualtrackFormCallback.CreationMode {
		visualtrackFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__VisualTrackFormCallback(
				nil,
				visualtrackFormCallback.probe,
			),
		}).Stage(visualtrackFormCallback.probe.formStage)
		visualtrack := new(models.VisualTrack)
		FillUpForm(visualtrack, newFormGroup, visualtrackFormCallback.probe)
		visualtrackFormCallback.probe.formStage.Commit()
	}

	fillUpTree(visualtrackFormCallback.probe)
}
