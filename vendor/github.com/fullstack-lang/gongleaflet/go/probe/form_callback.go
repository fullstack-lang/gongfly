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
	playground *Playground,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.playground = playground
	circleFormCallback.circle = circle

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.playground.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.playground.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	// get the formGroup
	formGroup := circleFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(circle_.LayerGroup), circleFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	circleFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.playground,
	)
	circleFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode {
		circleFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CircleFormCallback(
				nil,
				circleFormCallback.playground,
			),
		}).Stage(circleFormCallback.playground.formStage)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.playground)
		circleFormCallback.playground.formStage.Commit()
	}

	fillUpTree(circleFormCallback.playground)
}
func __gong__New__DivIconFormCallback(
	divicon *models.DivIcon,
	playground *Playground,
) (diviconFormCallback *DivIconFormCallback) {
	diviconFormCallback = new(DivIconFormCallback)
	diviconFormCallback.playground = playground
	diviconFormCallback.divicon = divicon

	diviconFormCallback.CreationMode = (divicon == nil)

	return
}

type DivIconFormCallback struct {
	divicon *models.DivIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (diviconFormCallback *DivIconFormCallback) OnSave() {

	log.Println("DivIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diviconFormCallback.playground.formStage.Checkout()

	if diviconFormCallback.divicon == nil {
		diviconFormCallback.divicon = new(models.DivIcon).Stage(diviconFormCallback.playground.stageOfInterest)
	}
	divicon_ := diviconFormCallback.divicon
	_ = divicon_

	// get the formGroup
	formGroup := diviconFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(divicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(divicon_.SVG), formDiv)
		}
	}

	diviconFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DivIcon](
		diviconFormCallback.playground,
	)
	diviconFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if diviconFormCallback.CreationMode {
		diviconFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DivIconFormCallback(
				nil,
				diviconFormCallback.playground,
			),
		}).Stage(diviconFormCallback.playground.formStage)
		divicon := new(models.DivIcon)
		FillUpForm(divicon, newFormGroup, diviconFormCallback.playground)
		diviconFormCallback.playground.formStage.Commit()
	}

	fillUpTree(diviconFormCallback.playground)
}
func __gong__New__LayerGroupFormCallback(
	layergroup *models.LayerGroup,
	playground *Playground,
) (layergroupFormCallback *LayerGroupFormCallback) {
	layergroupFormCallback = new(LayerGroupFormCallback)
	layergroupFormCallback.playground = playground
	layergroupFormCallback.layergroup = layergroup

	layergroupFormCallback.CreationMode = (layergroup == nil)

	return
}

type LayerGroupFormCallback struct {
	layergroup *models.LayerGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (layergroupFormCallback *LayerGroupFormCallback) OnSave() {

	log.Println("LayerGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layergroupFormCallback.playground.formStage.Checkout()

	if layergroupFormCallback.layergroup == nil {
		layergroupFormCallback.layergroup = new(models.LayerGroup).Stage(layergroupFormCallback.playground.stageOfInterest)
	}
	layergroup_ := layergroupFormCallback.layergroup
	_ = layergroup_

	// get the formGroup
	formGroup := layergroupFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layergroup_.Name), formDiv)
		case "DisplayName":
			FormDivBasicFieldToField(&(layergroup_.DisplayName), formDiv)
		}
	}

	layergroupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LayerGroup](
		layergroupFormCallback.playground,
	)
	layergroupFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupFormCallback.CreationMode {
		layergroupFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LayerGroupFormCallback(
				nil,
				layergroupFormCallback.playground,
			),
		}).Stage(layergroupFormCallback.playground.formStage)
		layergroup := new(models.LayerGroup)
		FillUpForm(layergroup, newFormGroup, layergroupFormCallback.playground)
		layergroupFormCallback.playground.formStage.Commit()
	}

	fillUpTree(layergroupFormCallback.playground)
}
func __gong__New__LayerGroupUseFormCallback(
	layergroupuse *models.LayerGroupUse,
	playground *Playground,
) (layergroupuseFormCallback *LayerGroupUseFormCallback) {
	layergroupuseFormCallback = new(LayerGroupUseFormCallback)
	layergroupuseFormCallback.playground = playground
	layergroupuseFormCallback.layergroupuse = layergroupuse

	layergroupuseFormCallback.CreationMode = (layergroupuse == nil)

	return
}

type LayerGroupUseFormCallback struct {
	layergroupuse *models.LayerGroupUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (layergroupuseFormCallback *LayerGroupUseFormCallback) OnSave() {

	log.Println("LayerGroupUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layergroupuseFormCallback.playground.formStage.Checkout()

	if layergroupuseFormCallback.layergroupuse == nil {
		layergroupuseFormCallback.layergroupuse = new(models.LayerGroupUse).Stage(layergroupuseFormCallback.playground.stageOfInterest)
	}
	layergroupuse_ := layergroupuseFormCallback.layergroupuse
	_ = layergroupuse_

	// get the formGroup
	formGroup := layergroupuseFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layergroupuse_.Name), formDiv)
		case "Display":
			FormDivBasicFieldToField(&(layergroupuse_.Display), formDiv)
		case "LayerGroup":
			FormDivSelectFieldToField(&(layergroupuse_.LayerGroup), layergroupuseFormCallback.playground.stageOfInterest, formDiv)
		case "MapOptions:LayerGroupUses":
			// we need to retrieve the field owner before the change
			var pastMapOptionsOwner *models.MapOptions
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "MapOptions"
			rf.Fieldname = "LayerGroupUses"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				layergroupuseFormCallback.playground.stageOfInterest,
				layergroupuseFormCallback.playground.backRepoOfInterest,
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
				for _mapoptions := range *models.GetGongstructInstancesSet[models.MapOptions](layergroupuseFormCallback.playground.stageOfInterest) {

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

	layergroupuseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LayerGroupUse](
		layergroupuseFormCallback.playground,
	)
	layergroupuseFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if layergroupuseFormCallback.CreationMode {
		layergroupuseFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LayerGroupUseFormCallback(
				nil,
				layergroupuseFormCallback.playground,
			),
		}).Stage(layergroupuseFormCallback.playground.formStage)
		layergroupuse := new(models.LayerGroupUse)
		FillUpForm(layergroupuse, newFormGroup, layergroupuseFormCallback.playground)
		layergroupuseFormCallback.playground.formStage.Commit()
	}

	fillUpTree(layergroupuseFormCallback.playground)
}
func __gong__New__MapOptionsFormCallback(
	mapoptions *models.MapOptions,
	playground *Playground,
) (mapoptionsFormCallback *MapOptionsFormCallback) {
	mapoptionsFormCallback = new(MapOptionsFormCallback)
	mapoptionsFormCallback.playground = playground
	mapoptionsFormCallback.mapoptions = mapoptions

	mapoptionsFormCallback.CreationMode = (mapoptions == nil)

	return
}

type MapOptionsFormCallback struct {
	mapoptions *models.MapOptions

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (mapoptionsFormCallback *MapOptionsFormCallback) OnSave() {

	log.Println("MapOptionsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapoptionsFormCallback.playground.formStage.Checkout()

	if mapoptionsFormCallback.mapoptions == nil {
		mapoptionsFormCallback.mapoptions = new(models.MapOptions).Stage(mapoptionsFormCallback.playground.stageOfInterest)
	}
	mapoptions_ := mapoptionsFormCallback.mapoptions
	_ = mapoptions_

	// get the formGroup
	formGroup := mapoptionsFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	mapoptionsFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.MapOptions](
		mapoptionsFormCallback.playground,
	)
	mapoptionsFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if mapoptionsFormCallback.CreationMode {
		mapoptionsFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MapOptionsFormCallback(
				nil,
				mapoptionsFormCallback.playground,
			),
		}).Stage(mapoptionsFormCallback.playground.formStage)
		mapoptions := new(models.MapOptions)
		FillUpForm(mapoptions, newFormGroup, mapoptionsFormCallback.playground)
		mapoptionsFormCallback.playground.formStage.Commit()
	}

	fillUpTree(mapoptionsFormCallback.playground)
}
func __gong__New__MarkerFormCallback(
	marker *models.Marker,
	playground *Playground,
) (markerFormCallback *MarkerFormCallback) {
	markerFormCallback = new(MarkerFormCallback)
	markerFormCallback.playground = playground
	markerFormCallback.marker = marker

	markerFormCallback.CreationMode = (marker == nil)

	return
}

type MarkerFormCallback struct {
	marker *models.Marker

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (markerFormCallback *MarkerFormCallback) OnSave() {

	log.Println("MarkerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	markerFormCallback.playground.formStage.Checkout()

	if markerFormCallback.marker == nil {
		markerFormCallback.marker = new(models.Marker).Stage(markerFormCallback.playground.stageOfInterest)
	}
	marker_ := markerFormCallback.marker
	_ = marker_

	// get the formGroup
	formGroup := markerFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(marker_.LayerGroup), markerFormCallback.playground.stageOfInterest, formDiv)
		case "DivIcon":
			FormDivSelectFieldToField(&(marker_.DivIcon), markerFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	markerFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Marker](
		markerFormCallback.playground,
	)
	markerFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if markerFormCallback.CreationMode {
		markerFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MarkerFormCallback(
				nil,
				markerFormCallback.playground,
			),
		}).Stage(markerFormCallback.playground.formStage)
		marker := new(models.Marker)
		FillUpForm(marker, newFormGroup, markerFormCallback.playground)
		markerFormCallback.playground.formStage.Commit()
	}

	fillUpTree(markerFormCallback.playground)
}
func __gong__New__UserClickFormCallback(
	userclick *models.UserClick,
	playground *Playground,
) (userclickFormCallback *UserClickFormCallback) {
	userclickFormCallback = new(UserClickFormCallback)
	userclickFormCallback.playground = playground
	userclickFormCallback.userclick = userclick

	userclickFormCallback.CreationMode = (userclick == nil)

	return
}

type UserClickFormCallback struct {
	userclick *models.UserClick

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (userclickFormCallback *UserClickFormCallback) OnSave() {

	log.Println("UserClickFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	userclickFormCallback.playground.formStage.Checkout()

	if userclickFormCallback.userclick == nil {
		userclickFormCallback.userclick = new(models.UserClick).Stage(userclickFormCallback.playground.stageOfInterest)
	}
	userclick_ := userclickFormCallback.userclick
	_ = userclick_

	// get the formGroup
	formGroup := userclickFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(userclick_.Name), formDiv)
		case "Lat":
			FormDivBasicFieldToField(&(userclick_.Lat), formDiv)
		case "Lng":
			FormDivBasicFieldToField(&(userclick_.Lng), formDiv)
		}
	}

	userclickFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.UserClick](
		userclickFormCallback.playground,
	)
	userclickFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if userclickFormCallback.CreationMode {
		userclickFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__UserClickFormCallback(
				nil,
				userclickFormCallback.playground,
			),
		}).Stage(userclickFormCallback.playground.formStage)
		userclick := new(models.UserClick)
		FillUpForm(userclick, newFormGroup, userclickFormCallback.playground)
		userclickFormCallback.playground.formStage.Commit()
	}

	fillUpTree(userclickFormCallback.playground)
}
func __gong__New__VLineFormCallback(
	vline *models.VLine,
	playground *Playground,
) (vlineFormCallback *VLineFormCallback) {
	vlineFormCallback = new(VLineFormCallback)
	vlineFormCallback.playground = playground
	vlineFormCallback.vline = vline

	vlineFormCallback.CreationMode = (vline == nil)

	return
}

type VLineFormCallback struct {
	vline *models.VLine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (vlineFormCallback *VLineFormCallback) OnSave() {

	log.Println("VLineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	vlineFormCallback.playground.formStage.Checkout()

	if vlineFormCallback.vline == nil {
		vlineFormCallback.vline = new(models.VLine).Stage(vlineFormCallback.playground.stageOfInterest)
	}
	vline_ := vlineFormCallback.vline
	_ = vline_

	// get the formGroup
	formGroup := vlineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(vline_.LayerGroup), vlineFormCallback.playground.stageOfInterest, formDiv)
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

	vlineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.VLine](
		vlineFormCallback.playground,
	)
	vlineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if vlineFormCallback.CreationMode {
		vlineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__VLineFormCallback(
				nil,
				vlineFormCallback.playground,
			),
		}).Stage(vlineFormCallback.playground.formStage)
		vline := new(models.VLine)
		FillUpForm(vline, newFormGroup, vlineFormCallback.playground)
		vlineFormCallback.playground.formStage.Commit()
	}

	fillUpTree(vlineFormCallback.playground)
}
func __gong__New__VisualTrackFormCallback(
	visualtrack *models.VisualTrack,
	playground *Playground,
) (visualtrackFormCallback *VisualTrackFormCallback) {
	visualtrackFormCallback = new(VisualTrackFormCallback)
	visualtrackFormCallback.playground = playground
	visualtrackFormCallback.visualtrack = visualtrack

	visualtrackFormCallback.CreationMode = (visualtrack == nil)

	return
}

type VisualTrackFormCallback struct {
	visualtrack *models.VisualTrack

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (visualtrackFormCallback *VisualTrackFormCallback) OnSave() {

	log.Println("VisualTrackFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	visualtrackFormCallback.playground.formStage.Checkout()

	if visualtrackFormCallback.visualtrack == nil {
		visualtrackFormCallback.visualtrack = new(models.VisualTrack).Stage(visualtrackFormCallback.playground.stageOfInterest)
	}
	visualtrack_ := visualtrackFormCallback.visualtrack
	_ = visualtrack_

	// get the formGroup
	formGroup := visualtrackFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(visualtrack_.LayerGroup), visualtrackFormCallback.playground.stageOfInterest, formDiv)
		case "DivIcon":
			FormDivSelectFieldToField(&(visualtrack_.DivIcon), visualtrackFormCallback.playground.stageOfInterest, formDiv)
		case "DisplayTrackHistory":
			FormDivBasicFieldToField(&(visualtrack_.DisplayTrackHistory), formDiv)
		case "DisplayLevelAndSpeed":
			FormDivBasicFieldToField(&(visualtrack_.DisplayLevelAndSpeed), formDiv)
		}
	}

	visualtrackFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.VisualTrack](
		visualtrackFormCallback.playground,
	)
	visualtrackFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if visualtrackFormCallback.CreationMode {
		visualtrackFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__VisualTrackFormCallback(
				nil,
				visualtrackFormCallback.playground,
			),
		}).Stage(visualtrackFormCallback.playground.formStage)
		visualtrack := new(models.VisualTrack)
		FillUpForm(visualtrack, newFormGroup, visualtrackFormCallback.playground)
		visualtrackFormCallback.playground.formStage.Commit()
	}

	fillUpTree(visualtrackFormCallback.playground)
}
