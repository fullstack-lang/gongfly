// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Circles           map[*Circle]any
	Circles_mapString map[string]*Circle

	// insertion point for slice of pointers maps
	OnAfterCircleCreateCallback OnAfterCreateInterface[Circle]
	OnAfterCircleUpdateCallback OnAfterUpdateInterface[Circle]
	OnAfterCircleDeleteCallback OnAfterDeleteInterface[Circle]
	OnAfterCircleReadCallback   OnAfterReadInterface[Circle]

	DivIcons           map[*DivIcon]any
	DivIcons_mapString map[string]*DivIcon

	// insertion point for slice of pointers maps
	OnAfterDivIconCreateCallback OnAfterCreateInterface[DivIcon]
	OnAfterDivIconUpdateCallback OnAfterUpdateInterface[DivIcon]
	OnAfterDivIconDeleteCallback OnAfterDeleteInterface[DivIcon]
	OnAfterDivIconReadCallback   OnAfterReadInterface[DivIcon]

	LayerGroups           map[*LayerGroup]any
	LayerGroups_mapString map[string]*LayerGroup

	// insertion point for slice of pointers maps
	OnAfterLayerGroupCreateCallback OnAfterCreateInterface[LayerGroup]
	OnAfterLayerGroupUpdateCallback OnAfterUpdateInterface[LayerGroup]
	OnAfterLayerGroupDeleteCallback OnAfterDeleteInterface[LayerGroup]
	OnAfterLayerGroupReadCallback   OnAfterReadInterface[LayerGroup]

	LayerGroupUses           map[*LayerGroupUse]any
	LayerGroupUses_mapString map[string]*LayerGroupUse

	// insertion point for slice of pointers maps
	OnAfterLayerGroupUseCreateCallback OnAfterCreateInterface[LayerGroupUse]
	OnAfterLayerGroupUseUpdateCallback OnAfterUpdateInterface[LayerGroupUse]
	OnAfterLayerGroupUseDeleteCallback OnAfterDeleteInterface[LayerGroupUse]
	OnAfterLayerGroupUseReadCallback   OnAfterReadInterface[LayerGroupUse]

	MapOptionss           map[*MapOptions]any
	MapOptionss_mapString map[string]*MapOptions

	// insertion point for slice of pointers maps
	MapOptions_LayerGroupUses_reverseMap map[*LayerGroupUse]*MapOptions

	OnAfterMapOptionsCreateCallback OnAfterCreateInterface[MapOptions]
	OnAfterMapOptionsUpdateCallback OnAfterUpdateInterface[MapOptions]
	OnAfterMapOptionsDeleteCallback OnAfterDeleteInterface[MapOptions]
	OnAfterMapOptionsReadCallback   OnAfterReadInterface[MapOptions]

	Markers           map[*Marker]any
	Markers_mapString map[string]*Marker

	// insertion point for slice of pointers maps
	OnAfterMarkerCreateCallback OnAfterCreateInterface[Marker]
	OnAfterMarkerUpdateCallback OnAfterUpdateInterface[Marker]
	OnAfterMarkerDeleteCallback OnAfterDeleteInterface[Marker]
	OnAfterMarkerReadCallback   OnAfterReadInterface[Marker]

	UserClicks           map[*UserClick]any
	UserClicks_mapString map[string]*UserClick

	// insertion point for slice of pointers maps
	OnAfterUserClickCreateCallback OnAfterCreateInterface[UserClick]
	OnAfterUserClickUpdateCallback OnAfterUpdateInterface[UserClick]
	OnAfterUserClickDeleteCallback OnAfterDeleteInterface[UserClick]
	OnAfterUserClickReadCallback   OnAfterReadInterface[UserClick]

	VLines           map[*VLine]any
	VLines_mapString map[string]*VLine

	// insertion point for slice of pointers maps
	OnAfterVLineCreateCallback OnAfterCreateInterface[VLine]
	OnAfterVLineUpdateCallback OnAfterUpdateInterface[VLine]
	OnAfterVLineDeleteCallback OnAfterDeleteInterface[VLine]
	OnAfterVLineReadCallback   OnAfterReadInterface[VLine]

	VisualTracks           map[*VisualTrack]any
	VisualTracks_mapString map[string]*VisualTrack

	// insertion point for slice of pointers maps
	OnAfterVisualTrackCreateCallback OnAfterCreateInterface[VisualTrack]
	OnAfterVisualTrackUpdateCallback OnAfterUpdateInterface[VisualTrack]
	OnAfterVisualTrackDeleteCallback OnAfterDeleteInterface[VisualTrack]
	OnAfterVisualTrackReadCallback   OnAfterReadInterface[VisualTrack]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongleaflet/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitDivIcon(divicon *DivIcon)
	CheckoutDivIcon(divicon *DivIcon)
	CommitLayerGroup(layergroup *LayerGroup)
	CheckoutLayerGroup(layergroup *LayerGroup)
	CommitLayerGroupUse(layergroupuse *LayerGroupUse)
	CheckoutLayerGroupUse(layergroupuse *LayerGroupUse)
	CommitMapOptions(mapoptions *MapOptions)
	CheckoutMapOptions(mapoptions *MapOptions)
	CommitMarker(marker *Marker)
	CheckoutMarker(marker *Marker)
	CommitUserClick(userclick *UserClick)
	CheckoutUserClick(userclick *UserClick)
	CommitVLine(vline *VLine)
	CheckoutVLine(vline *VLine)
	CommitVisualTrack(visualtrack *VisualTrack)
	CheckoutVisualTrack(visualtrack *VisualTrack)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Circles:           make(map[*Circle]any),
		Circles_mapString: make(map[string]*Circle),

		DivIcons:           make(map[*DivIcon]any),
		DivIcons_mapString: make(map[string]*DivIcon),

		LayerGroups:           make(map[*LayerGroup]any),
		LayerGroups_mapString: make(map[string]*LayerGroup),

		LayerGroupUses:           make(map[*LayerGroupUse]any),
		LayerGroupUses_mapString: make(map[string]*LayerGroupUse),

		MapOptionss:           make(map[*MapOptions]any),
		MapOptionss_mapString: make(map[string]*MapOptions),

		Markers:           make(map[*Marker]any),
		Markers_mapString: make(map[string]*Marker),

		UserClicks:           make(map[*UserClick]any),
		UserClicks_mapString: make(map[string]*UserClick),

		VLines:           make(map[*VLine]any),
		VLines_mapString: make(map[string]*VLine),

		VisualTracks:           make(map[*VisualTrack]any),
		VisualTracks_mapString: make(map[string]*VisualTrack),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["DivIcon"] = len(stage.DivIcons)
	stage.Map_GongStructName_InstancesNb["LayerGroup"] = len(stage.LayerGroups)
	stage.Map_GongStructName_InstancesNb["LayerGroupUse"] = len(stage.LayerGroupUses)
	stage.Map_GongStructName_InstancesNb["MapOptions"] = len(stage.MapOptionss)
	stage.Map_GongStructName_InstancesNb["Marker"] = len(stage.Markers)
	stage.Map_GongStructName_InstancesNb["UserClick"] = len(stage.UserClicks)
	stage.Map_GongStructName_InstancesNb["VLine"] = len(stage.VLines)
	stage.Map_GongStructName_InstancesNb["VisualTrack"] = len(stage.VisualTracks)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Circle"] = len(stage.Circles)
	stage.Map_GongStructName_InstancesNb["DivIcon"] = len(stage.DivIcons)
	stage.Map_GongStructName_InstancesNb["LayerGroup"] = len(stage.LayerGroups)
	stage.Map_GongStructName_InstancesNb["LayerGroupUse"] = len(stage.LayerGroupUses)
	stage.Map_GongStructName_InstancesNb["MapOptions"] = len(stage.MapOptionss)
	stage.Map_GongStructName_InstancesNb["Marker"] = len(stage.Markers)
	stage.Map_GongStructName_InstancesNb["UserClick"] = len(stage.UserClicks)
	stage.Map_GongStructName_InstancesNb["VLine"] = len(stage.VLines)
	stage.Map_GongStructName_InstancesNb["VisualTrack"] = len(stage.VisualTracks)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts circle to the model stage
func (circle *Circle) Stage(stage *StageStruct) *Circle {
	stage.Circles[circle] = __member
	stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage(stage *StageStruct) *Circle {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
	return circle
}

// UnstageVoid removes circle off the model stage
func (circle *Circle) UnstageVoid(stage *StageStruct) {
	delete(stage.Circles, circle)
	delete(stage.Circles_mapString, circle.Name)
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit(stage *StageStruct) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

func (circle *Circle) CommitVoid(stage *StageStruct) {
	circle.Commit(stage)
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout(stage *StageStruct) *Circle {
	if _, ok := stage.Circles[circle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCircle(circle)
		}
	}
	return circle
}

// for satisfaction of GongStruct interface
func (circle *Circle) GetName() (res string) {
	return circle.Name
}

// Stage puts divicon to the model stage
func (divicon *DivIcon) Stage(stage *StageStruct) *DivIcon {
	stage.DivIcons[divicon] = __member
	stage.DivIcons_mapString[divicon.Name] = divicon

	return divicon
}

// Unstage removes divicon off the model stage
func (divicon *DivIcon) Unstage(stage *StageStruct) *DivIcon {
	delete(stage.DivIcons, divicon)
	delete(stage.DivIcons_mapString, divicon.Name)
	return divicon
}

// UnstageVoid removes divicon off the model stage
func (divicon *DivIcon) UnstageVoid(stage *StageStruct) {
	delete(stage.DivIcons, divicon)
	delete(stage.DivIcons_mapString, divicon.Name)
}

// commit divicon to the back repo (if it is already staged)
func (divicon *DivIcon) Commit(stage *StageStruct) *DivIcon {
	if _, ok := stage.DivIcons[divicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDivIcon(divicon)
		}
	}
	return divicon
}

func (divicon *DivIcon) CommitVoid(stage *StageStruct) {
	divicon.Commit(stage)
}

// Checkout divicon to the back repo (if it is already staged)
func (divicon *DivIcon) Checkout(stage *StageStruct) *DivIcon {
	if _, ok := stage.DivIcons[divicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDivIcon(divicon)
		}
	}
	return divicon
}

// for satisfaction of GongStruct interface
func (divicon *DivIcon) GetName() (res string) {
	return divicon.Name
}

// Stage puts layergroup to the model stage
func (layergroup *LayerGroup) Stage(stage *StageStruct) *LayerGroup {
	stage.LayerGroups[layergroup] = __member
	stage.LayerGroups_mapString[layergroup.Name] = layergroup

	return layergroup
}

// Unstage removes layergroup off the model stage
func (layergroup *LayerGroup) Unstage(stage *StageStruct) *LayerGroup {
	delete(stage.LayerGroups, layergroup)
	delete(stage.LayerGroups_mapString, layergroup.Name)
	return layergroup
}

// UnstageVoid removes layergroup off the model stage
func (layergroup *LayerGroup) UnstageVoid(stage *StageStruct) {
	delete(stage.LayerGroups, layergroup)
	delete(stage.LayerGroups_mapString, layergroup.Name)
}

// commit layergroup to the back repo (if it is already staged)
func (layergroup *LayerGroup) Commit(stage *StageStruct) *LayerGroup {
	if _, ok := stage.LayerGroups[layergroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayerGroup(layergroup)
		}
	}
	return layergroup
}

func (layergroup *LayerGroup) CommitVoid(stage *StageStruct) {
	layergroup.Commit(stage)
}

// Checkout layergroup to the back repo (if it is already staged)
func (layergroup *LayerGroup) Checkout(stage *StageStruct) *LayerGroup {
	if _, ok := stage.LayerGroups[layergroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLayerGroup(layergroup)
		}
	}
	return layergroup
}

// for satisfaction of GongStruct interface
func (layergroup *LayerGroup) GetName() (res string) {
	return layergroup.Name
}

// Stage puts layergroupuse to the model stage
func (layergroupuse *LayerGroupUse) Stage(stage *StageStruct) *LayerGroupUse {
	stage.LayerGroupUses[layergroupuse] = __member
	stage.LayerGroupUses_mapString[layergroupuse.Name] = layergroupuse

	return layergroupuse
}

// Unstage removes layergroupuse off the model stage
func (layergroupuse *LayerGroupUse) Unstage(stage *StageStruct) *LayerGroupUse {
	delete(stage.LayerGroupUses, layergroupuse)
	delete(stage.LayerGroupUses_mapString, layergroupuse.Name)
	return layergroupuse
}

// UnstageVoid removes layergroupuse off the model stage
func (layergroupuse *LayerGroupUse) UnstageVoid(stage *StageStruct) {
	delete(stage.LayerGroupUses, layergroupuse)
	delete(stage.LayerGroupUses_mapString, layergroupuse.Name)
}

// commit layergroupuse to the back repo (if it is already staged)
func (layergroupuse *LayerGroupUse) Commit(stage *StageStruct) *LayerGroupUse {
	if _, ok := stage.LayerGroupUses[layergroupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayerGroupUse(layergroupuse)
		}
	}
	return layergroupuse
}

func (layergroupuse *LayerGroupUse) CommitVoid(stage *StageStruct) {
	layergroupuse.Commit(stage)
}

// Checkout layergroupuse to the back repo (if it is already staged)
func (layergroupuse *LayerGroupUse) Checkout(stage *StageStruct) *LayerGroupUse {
	if _, ok := stage.LayerGroupUses[layergroupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLayerGroupUse(layergroupuse)
		}
	}
	return layergroupuse
}

// for satisfaction of GongStruct interface
func (layergroupuse *LayerGroupUse) GetName() (res string) {
	return layergroupuse.Name
}

// Stage puts mapoptions to the model stage
func (mapoptions *MapOptions) Stage(stage *StageStruct) *MapOptions {
	stage.MapOptionss[mapoptions] = __member
	stage.MapOptionss_mapString[mapoptions.Name] = mapoptions

	return mapoptions
}

// Unstage removes mapoptions off the model stage
func (mapoptions *MapOptions) Unstage(stage *StageStruct) *MapOptions {
	delete(stage.MapOptionss, mapoptions)
	delete(stage.MapOptionss_mapString, mapoptions.Name)
	return mapoptions
}

// UnstageVoid removes mapoptions off the model stage
func (mapoptions *MapOptions) UnstageVoid(stage *StageStruct) {
	delete(stage.MapOptionss, mapoptions)
	delete(stage.MapOptionss_mapString, mapoptions.Name)
}

// commit mapoptions to the back repo (if it is already staged)
func (mapoptions *MapOptions) Commit(stage *StageStruct) *MapOptions {
	if _, ok := stage.MapOptionss[mapoptions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMapOptions(mapoptions)
		}
	}
	return mapoptions
}

func (mapoptions *MapOptions) CommitVoid(stage *StageStruct) {
	mapoptions.Commit(stage)
}

// Checkout mapoptions to the back repo (if it is already staged)
func (mapoptions *MapOptions) Checkout(stage *StageStruct) *MapOptions {
	if _, ok := stage.MapOptionss[mapoptions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMapOptions(mapoptions)
		}
	}
	return mapoptions
}

// for satisfaction of GongStruct interface
func (mapoptions *MapOptions) GetName() (res string) {
	return mapoptions.Name
}

// Stage puts marker to the model stage
func (marker *Marker) Stage(stage *StageStruct) *Marker {
	stage.Markers[marker] = __member
	stage.Markers_mapString[marker.Name] = marker

	return marker
}

// Unstage removes marker off the model stage
func (marker *Marker) Unstage(stage *StageStruct) *Marker {
	delete(stage.Markers, marker)
	delete(stage.Markers_mapString, marker.Name)
	return marker
}

// UnstageVoid removes marker off the model stage
func (marker *Marker) UnstageVoid(stage *StageStruct) {
	delete(stage.Markers, marker)
	delete(stage.Markers_mapString, marker.Name)
}

// commit marker to the back repo (if it is already staged)
func (marker *Marker) Commit(stage *StageStruct) *Marker {
	if _, ok := stage.Markers[marker]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMarker(marker)
		}
	}
	return marker
}

func (marker *Marker) CommitVoid(stage *StageStruct) {
	marker.Commit(stage)
}

// Checkout marker to the back repo (if it is already staged)
func (marker *Marker) Checkout(stage *StageStruct) *Marker {
	if _, ok := stage.Markers[marker]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMarker(marker)
		}
	}
	return marker
}

// for satisfaction of GongStruct interface
func (marker *Marker) GetName() (res string) {
	return marker.Name
}

// Stage puts userclick to the model stage
func (userclick *UserClick) Stage(stage *StageStruct) *UserClick {
	stage.UserClicks[userclick] = __member
	stage.UserClicks_mapString[userclick.Name] = userclick

	return userclick
}

// Unstage removes userclick off the model stage
func (userclick *UserClick) Unstage(stage *StageStruct) *UserClick {
	delete(stage.UserClicks, userclick)
	delete(stage.UserClicks_mapString, userclick.Name)
	return userclick
}

// UnstageVoid removes userclick off the model stage
func (userclick *UserClick) UnstageVoid(stage *StageStruct) {
	delete(stage.UserClicks, userclick)
	delete(stage.UserClicks_mapString, userclick.Name)
}

// commit userclick to the back repo (if it is already staged)
func (userclick *UserClick) Commit(stage *StageStruct) *UserClick {
	if _, ok := stage.UserClicks[userclick]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUserClick(userclick)
		}
	}
	return userclick
}

func (userclick *UserClick) CommitVoid(stage *StageStruct) {
	userclick.Commit(stage)
}

// Checkout userclick to the back repo (if it is already staged)
func (userclick *UserClick) Checkout(stage *StageStruct) *UserClick {
	if _, ok := stage.UserClicks[userclick]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUserClick(userclick)
		}
	}
	return userclick
}

// for satisfaction of GongStruct interface
func (userclick *UserClick) GetName() (res string) {
	return userclick.Name
}

// Stage puts vline to the model stage
func (vline *VLine) Stage(stage *StageStruct) *VLine {
	stage.VLines[vline] = __member
	stage.VLines_mapString[vline.Name] = vline

	return vline
}

// Unstage removes vline off the model stage
func (vline *VLine) Unstage(stage *StageStruct) *VLine {
	delete(stage.VLines, vline)
	delete(stage.VLines_mapString, vline.Name)
	return vline
}

// UnstageVoid removes vline off the model stage
func (vline *VLine) UnstageVoid(stage *StageStruct) {
	delete(stage.VLines, vline)
	delete(stage.VLines_mapString, vline.Name)
}

// commit vline to the back repo (if it is already staged)
func (vline *VLine) Commit(stage *StageStruct) *VLine {
	if _, ok := stage.VLines[vline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVLine(vline)
		}
	}
	return vline
}

func (vline *VLine) CommitVoid(stage *StageStruct) {
	vline.Commit(stage)
}

// Checkout vline to the back repo (if it is already staged)
func (vline *VLine) Checkout(stage *StageStruct) *VLine {
	if _, ok := stage.VLines[vline]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVLine(vline)
		}
	}
	return vline
}

// for satisfaction of GongStruct interface
func (vline *VLine) GetName() (res string) {
	return vline.Name
}

// Stage puts visualtrack to the model stage
func (visualtrack *VisualTrack) Stage(stage *StageStruct) *VisualTrack {
	stage.VisualTracks[visualtrack] = __member
	stage.VisualTracks_mapString[visualtrack.Name] = visualtrack

	return visualtrack
}

// Unstage removes visualtrack off the model stage
func (visualtrack *VisualTrack) Unstage(stage *StageStruct) *VisualTrack {
	delete(stage.VisualTracks, visualtrack)
	delete(stage.VisualTracks_mapString, visualtrack.Name)
	return visualtrack
}

// UnstageVoid removes visualtrack off the model stage
func (visualtrack *VisualTrack) UnstageVoid(stage *StageStruct) {
	delete(stage.VisualTracks, visualtrack)
	delete(stage.VisualTracks_mapString, visualtrack.Name)
}

// commit visualtrack to the back repo (if it is already staged)
func (visualtrack *VisualTrack) Commit(stage *StageStruct) *VisualTrack {
	if _, ok := stage.VisualTracks[visualtrack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVisualTrack(visualtrack)
		}
	}
	return visualtrack
}

func (visualtrack *VisualTrack) CommitVoid(stage *StageStruct) {
	visualtrack.Commit(stage)
}

// Checkout visualtrack to the back repo (if it is already staged)
func (visualtrack *VisualTrack) Checkout(stage *StageStruct) *VisualTrack {
	if _, ok := stage.VisualTracks[visualtrack]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVisualTrack(visualtrack)
		}
	}
	return visualtrack
}

// for satisfaction of GongStruct interface
func (visualtrack *VisualTrack) GetName() (res string) {
	return visualtrack.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCircle(Circle *Circle)
	CreateORMDivIcon(DivIcon *DivIcon)
	CreateORMLayerGroup(LayerGroup *LayerGroup)
	CreateORMLayerGroupUse(LayerGroupUse *LayerGroupUse)
	CreateORMMapOptions(MapOptions *MapOptions)
	CreateORMMarker(Marker *Marker)
	CreateORMUserClick(UserClick *UserClick)
	CreateORMVLine(VLine *VLine)
	CreateORMVisualTrack(VisualTrack *VisualTrack)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCircle(Circle *Circle)
	DeleteORMDivIcon(DivIcon *DivIcon)
	DeleteORMLayerGroup(LayerGroup *LayerGroup)
	DeleteORMLayerGroupUse(LayerGroupUse *LayerGroupUse)
	DeleteORMMapOptions(MapOptions *MapOptions)
	DeleteORMMarker(Marker *Marker)
	DeleteORMUserClick(UserClick *UserClick)
	DeleteORMVLine(VLine *VLine)
	DeleteORMVisualTrack(VisualTrack *VisualTrack)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Circles = make(map[*Circle]any)
	stage.Circles_mapString = make(map[string]*Circle)

	stage.DivIcons = make(map[*DivIcon]any)
	stage.DivIcons_mapString = make(map[string]*DivIcon)

	stage.LayerGroups = make(map[*LayerGroup]any)
	stage.LayerGroups_mapString = make(map[string]*LayerGroup)

	stage.LayerGroupUses = make(map[*LayerGroupUse]any)
	stage.LayerGroupUses_mapString = make(map[string]*LayerGroupUse)

	stage.MapOptionss = make(map[*MapOptions]any)
	stage.MapOptionss_mapString = make(map[string]*MapOptions)

	stage.Markers = make(map[*Marker]any)
	stage.Markers_mapString = make(map[string]*Marker)

	stage.UserClicks = make(map[*UserClick]any)
	stage.UserClicks_mapString = make(map[string]*UserClick)

	stage.VLines = make(map[*VLine]any)
	stage.VLines_mapString = make(map[string]*VLine)

	stage.VisualTracks = make(map[*VisualTrack]any)
	stage.VisualTracks_mapString = make(map[string]*VisualTrack)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.DivIcons = nil
	stage.DivIcons_mapString = nil

	stage.LayerGroups = nil
	stage.LayerGroups_mapString = nil

	stage.LayerGroupUses = nil
	stage.LayerGroupUses_mapString = nil

	stage.MapOptionss = nil
	stage.MapOptionss_mapString = nil

	stage.Markers = nil
	stage.Markers_mapString = nil

	stage.UserClicks = nil
	stage.UserClicks_mapString = nil

	stage.VLines = nil
	stage.VLines_mapString = nil

	stage.VisualTracks = nil
	stage.VisualTracks_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for circle := range stage.Circles {
		circle.Unstage(stage)
	}

	for divicon := range stage.DivIcons {
		divicon.Unstage(stage)
	}

	for layergroup := range stage.LayerGroups {
		layergroup.Unstage(stage)
	}

	for layergroupuse := range stage.LayerGroupUses {
		layergroupuse.Unstage(stage)
	}

	for mapoptions := range stage.MapOptionss {
		mapoptions.Unstage(stage)
	}

	for marker := range stage.Markers {
		marker.Unstage(stage)
	}

	for userclick := range stage.UserClicks {
		userclick.Unstage(stage)
	}

	for vline := range stage.VLines {
		vline.Unstage(stage)
	}

	for visualtrack := range stage.VisualTracks {
		visualtrack.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {

}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Circle]any:
		return any(&stage.Circles).(*Type)
	case map[*DivIcon]any:
		return any(&stage.DivIcons).(*Type)
	case map[*LayerGroup]any:
		return any(&stage.LayerGroups).(*Type)
	case map[*LayerGroupUse]any:
		return any(&stage.LayerGroupUses).(*Type)
	case map[*MapOptions]any:
		return any(&stage.MapOptionss).(*Type)
	case map[*Marker]any:
		return any(&stage.Markers).(*Type)
	case map[*UserClick]any:
		return any(&stage.UserClicks).(*Type)
	case map[*VLine]any:
		return any(&stage.VLines).(*Type)
	case map[*VisualTrack]any:
		return any(&stage.VisualTracks).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Circle:
		return any(&stage.Circles_mapString).(*Type)
	case map[string]*DivIcon:
		return any(&stage.DivIcons_mapString).(*Type)
	case map[string]*LayerGroup:
		return any(&stage.LayerGroups_mapString).(*Type)
	case map[string]*LayerGroupUse:
		return any(&stage.LayerGroupUses_mapString).(*Type)
	case map[string]*MapOptions:
		return any(&stage.MapOptionss_mapString).(*Type)
	case map[string]*Marker:
		return any(&stage.Markers_mapString).(*Type)
	case map[string]*UserClick:
		return any(&stage.UserClicks_mapString).(*Type)
	case map[string]*VLine:
		return any(&stage.VLines_mapString).(*Type)
	case map[string]*VisualTrack:
		return any(&stage.VisualTracks_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Circle:
		return any(&stage.Circles).(*map[*Type]any)
	case DivIcon:
		return any(&stage.DivIcons).(*map[*Type]any)
	case LayerGroup:
		return any(&stage.LayerGroups).(*map[*Type]any)
	case LayerGroupUse:
		return any(&stage.LayerGroupUses).(*map[*Type]any)
	case MapOptions:
		return any(&stage.MapOptionss).(*map[*Type]any)
	case Marker:
		return any(&stage.Markers).(*map[*Type]any)
	case UserClick:
		return any(&stage.UserClicks).(*map[*Type]any)
	case VLine:
		return any(&stage.VLines).(*map[*Type]any)
	case VisualTrack:
		return any(&stage.VisualTracks).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Circle:
		return any(&stage.Circles).(*map[Type]any)
	case *DivIcon:
		return any(&stage.DivIcons).(*map[Type]any)
	case *LayerGroup:
		return any(&stage.LayerGroups).(*map[Type]any)
	case *LayerGroupUse:
		return any(&stage.LayerGroupUses).(*map[Type]any)
	case *MapOptions:
		return any(&stage.MapOptionss).(*map[Type]any)
	case *Marker:
		return any(&stage.Markers).(*map[Type]any)
	case *UserClick:
		return any(&stage.UserClicks).(*map[Type]any)
	case *VLine:
		return any(&stage.VLines).(*map[Type]any)
	case *VisualTrack:
		return any(&stage.VisualTracks).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Circle:
		return any(&stage.Circles_mapString).(*map[string]*Type)
	case DivIcon:
		return any(&stage.DivIcons_mapString).(*map[string]*Type)
	case LayerGroup:
		return any(&stage.LayerGroups_mapString).(*map[string]*Type)
	case LayerGroupUse:
		return any(&stage.LayerGroupUses_mapString).(*map[string]*Type)
	case MapOptions:
		return any(&stage.MapOptionss_mapString).(*map[string]*Type)
	case Marker:
		return any(&stage.Markers_mapString).(*map[string]*Type)
	case UserClick:
		return any(&stage.UserClicks_mapString).(*map[string]*Type)
	case VLine:
		return any(&stage.VLines_mapString).(*map[string]*Type)
	case VisualTrack:
		return any(&stage.VisualTracks_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Circle:
		return any(&Circle{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroup with the name of the field
			LayerGroup: &LayerGroup{Name: "LayerGroup"},
		}).(*Type)
	case DivIcon:
		return any(&DivIcon{
			// Initialisation of associations
		}).(*Type)
	case LayerGroup:
		return any(&LayerGroup{
			// Initialisation of associations
		}).(*Type)
	case LayerGroupUse:
		return any(&LayerGroupUse{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroup with the name of the field
			LayerGroup: &LayerGroup{Name: "LayerGroup"},
		}).(*Type)
	case MapOptions:
		return any(&MapOptions{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroupUse with the name of the field
			LayerGroupUses: []*LayerGroupUse{{Name: "LayerGroupUses"}},
		}).(*Type)
	case Marker:
		return any(&Marker{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroup with the name of the field
			LayerGroup: &LayerGroup{Name: "LayerGroup"},
			// field is initialized with an instance of DivIcon with the name of the field
			DivIcon: &DivIcon{Name: "DivIcon"},
		}).(*Type)
	case UserClick:
		return any(&UserClick{
			// Initialisation of associations
		}).(*Type)
	case VLine:
		return any(&VLine{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroup with the name of the field
			LayerGroup: &LayerGroup{Name: "LayerGroup"},
		}).(*Type)
	case VisualTrack:
		return any(&VisualTrack{
			// Initialisation of associations
			// field is initialized with an instance of LayerGroup with the name of the field
			LayerGroup: &LayerGroup{Name: "LayerGroup"},
			// field is initialized with an instance of DivIcon with the name of the field
			DivIcon: &DivIcon{Name: "DivIcon"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroup":
			res := make(map[*LayerGroup][]*Circle)
			for circle := range stage.Circles {
				if circle.LayerGroup != nil {
					layergroup_ := circle.LayerGroup
					var circles []*Circle
					_, ok := res[layergroup_]
					if ok {
						circles = res[layergroup_]
					} else {
						circles = make([]*Circle, 0)
					}
					circles = append(circles, circle)
					res[layergroup_] = circles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DivIcon
	case DivIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LayerGroup
	case LayerGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LayerGroupUse
	case LayerGroupUse:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroup":
			res := make(map[*LayerGroup][]*LayerGroupUse)
			for layergroupuse := range stage.LayerGroupUses {
				if layergroupuse.LayerGroup != nil {
					layergroup_ := layergroupuse.LayerGroup
					var layergroupuses []*LayerGroupUse
					_, ok := res[layergroup_]
					if ok {
						layergroupuses = res[layergroup_]
					} else {
						layergroupuses = make([]*LayerGroupUse, 0)
					}
					layergroupuses = append(layergroupuses, layergroupuse)
					res[layergroup_] = layergroupuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MapOptions
	case MapOptions:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Marker
	case Marker:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroup":
			res := make(map[*LayerGroup][]*Marker)
			for marker := range stage.Markers {
				if marker.LayerGroup != nil {
					layergroup_ := marker.LayerGroup
					var markers []*Marker
					_, ok := res[layergroup_]
					if ok {
						markers = res[layergroup_]
					} else {
						markers = make([]*Marker, 0)
					}
					markers = append(markers, marker)
					res[layergroup_] = markers
				}
			}
			return any(res).(map[*End][]*Start)
		case "DivIcon":
			res := make(map[*DivIcon][]*Marker)
			for marker := range stage.Markers {
				if marker.DivIcon != nil {
					divicon_ := marker.DivIcon
					var markers []*Marker
					_, ok := res[divicon_]
					if ok {
						markers = res[divicon_]
					} else {
						markers = make([]*Marker, 0)
					}
					markers = append(markers, marker)
					res[divicon_] = markers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of UserClick
	case UserClick:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VLine
	case VLine:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroup":
			res := make(map[*LayerGroup][]*VLine)
			for vline := range stage.VLines {
				if vline.LayerGroup != nil {
					layergroup_ := vline.LayerGroup
					var vlines []*VLine
					_, ok := res[layergroup_]
					if ok {
						vlines = res[layergroup_]
					} else {
						vlines = make([]*VLine, 0)
					}
					vlines = append(vlines, vline)
					res[layergroup_] = vlines
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VisualTrack
	case VisualTrack:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroup":
			res := make(map[*LayerGroup][]*VisualTrack)
			for visualtrack := range stage.VisualTracks {
				if visualtrack.LayerGroup != nil {
					layergroup_ := visualtrack.LayerGroup
					var visualtracks []*VisualTrack
					_, ok := res[layergroup_]
					if ok {
						visualtracks = res[layergroup_]
					} else {
						visualtracks = make([]*VisualTrack, 0)
					}
					visualtracks = append(visualtracks, visualtrack)
					res[layergroup_] = visualtracks
				}
			}
			return any(res).(map[*End][]*Start)
		case "DivIcon":
			res := make(map[*DivIcon][]*VisualTrack)
			for visualtrack := range stage.VisualTracks {
				if visualtrack.DivIcon != nil {
					divicon_ := visualtrack.DivIcon
					var visualtracks []*VisualTrack
					_, ok := res[divicon_]
					if ok {
						visualtracks = res[divicon_]
					} else {
						visualtracks = make([]*VisualTrack, 0)
					}
					visualtracks = append(visualtracks, visualtrack)
					res[divicon_] = visualtracks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Circle
	case Circle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DivIcon
	case DivIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LayerGroup
	case LayerGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LayerGroupUse
	case LayerGroupUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapOptions
	case MapOptions:
		switch fieldname {
		// insertion point for per direct association field
		case "LayerGroupUses":
			res := make(map[*LayerGroupUse]*MapOptions)
			for mapoptions := range stage.MapOptionss {
				for _, layergroupuse_ := range mapoptions.LayerGroupUses {
					res[layergroupuse_] = mapoptions
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Marker
	case Marker:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UserClick
	case UserClick:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VLine
	case VLine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VisualTrack
	case VisualTrack:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Circle:
		res = "Circle"
	case DivIcon:
		res = "DivIcon"
	case LayerGroup:
		res = "LayerGroup"
	case LayerGroupUse:
		res = "LayerGroupUse"
	case MapOptions:
		res = "MapOptions"
	case Marker:
		res = "Marker"
	case UserClick:
		res = "UserClick"
	case VLine:
		res = "VLine"
	case VisualTrack:
		res = "VisualTrack"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Circle:
		res = "Circle"
	case *DivIcon:
		res = "DivIcon"
	case *LayerGroup:
		res = "LayerGroup"
	case *LayerGroupUse:
		res = "LayerGroupUse"
	case *MapOptions:
		res = "MapOptions"
	case *Marker:
		res = "Marker"
	case *UserClick:
		res = "UserClick"
	case *VLine:
		res = "VLine"
	case *VisualTrack:
		res = "VisualTrack"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Circle:
		res = []string{"Lat", "Lng", "Name", "Radius", "ColorEnum", "DashStyleEnum", "LayerGroup"}
	case DivIcon:
		res = []string{"Name", "SVG"}
	case LayerGroup:
		res = []string{"Name", "DisplayName"}
	case LayerGroupUse:
		res = []string{"Name", "IsDisplayed", "LayerGroup"}
	case MapOptions:
		res = []string{"Lat", "Lng", "Name", "ZoomLevel", "UrlTemplate", "Attribution", "MaxZoom", "ZoomControl", "AttributionControl", "ZoomSnap", "LayerGroupUses"}
	case Marker:
		res = []string{"Lat", "Lng", "Name", "ColorEnum", "LayerGroup", "DivIcon"}
	case UserClick:
		res = []string{"Name", "Lat", "Lng", "TimeOfClick"}
	case VLine:
		res = []string{"StartLat", "StartLng", "EndLat", "EndLng", "Name", "ColorEnum", "DashStyleEnum", "LayerGroup", "IsTransmitting", "Message", "IsTransmittingBackward", "MessageBackward"}
	case VisualTrack:
		res = []string{"Lat", "Lng", "Heading", "Level", "Speed", "VerticalSpeed", "Name", "ColorEnum", "LayerGroup", "DivIcon", "DisplayTrackHistory", "DisplayLevelAndSpeed"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Circle:
		var rf ReverseField
		_ = rf
	case DivIcon:
		var rf ReverseField
		_ = rf
	case LayerGroup:
		var rf ReverseField
		_ = rf
	case LayerGroupUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "MapOptions"
		rf.Fieldname = "LayerGroupUses"
		res = append(res, rf)
	case MapOptions:
		var rf ReverseField
		_ = rf
	case Marker:
		var rf ReverseField
		_ = rf
	case UserClick:
		var rf ReverseField
		_ = rf
	case VLine:
		var rf ReverseField
		_ = rf
	case VisualTrack:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Circle:
		res = []string{"Lat", "Lng", "Name", "Radius", "ColorEnum", "DashStyleEnum", "LayerGroup"}
	case *DivIcon:
		res = []string{"Name", "SVG"}
	case *LayerGroup:
		res = []string{"Name", "DisplayName"}
	case *LayerGroupUse:
		res = []string{"Name", "IsDisplayed", "LayerGroup"}
	case *MapOptions:
		res = []string{"Lat", "Lng", "Name", "ZoomLevel", "UrlTemplate", "Attribution", "MaxZoom", "ZoomControl", "AttributionControl", "ZoomSnap", "LayerGroupUses"}
	case *Marker:
		res = []string{"Lat", "Lng", "Name", "ColorEnum", "LayerGroup", "DivIcon"}
	case *UserClick:
		res = []string{"Name", "Lat", "Lng", "TimeOfClick"}
	case *VLine:
		res = []string{"StartLat", "StartLng", "EndLat", "EndLng", "Name", "ColorEnum", "DashStyleEnum", "LayerGroup", "IsTransmitting", "Message", "IsTransmittingBackward", "MessageBackward"}
	case *VisualTrack:
		res = []string{"Lat", "Lng", "Heading", "Level", "Speed", "VerticalSpeed", "Name", "ColorEnum", "LayerGroup", "DivIcon", "DisplayTrackHistory", "DisplayLevelAndSpeed"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Circle:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "Radius":
			res = fmt.Sprintf("%f", inferedInstance.Radius)
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "DashStyleEnum":
			enum := inferedInstance.DashStyleEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		}
	case *DivIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SVG":
			res = inferedInstance.SVG
		}
	case *LayerGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DisplayName":
			res = inferedInstance.DisplayName
		}
	case *LayerGroupUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		}
	case *MapOptions:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "ZoomLevel":
			res = fmt.Sprintf("%f", inferedInstance.ZoomLevel)
		case "UrlTemplate":
			res = inferedInstance.UrlTemplate
		case "Attribution":
			res = inferedInstance.Attribution
		case "MaxZoom":
			res = fmt.Sprintf("%d", inferedInstance.MaxZoom)
		case "ZoomControl":
			res = fmt.Sprintf("%t", inferedInstance.ZoomControl)
		case "AttributionControl":
			res = fmt.Sprintf("%t", inferedInstance.AttributionControl)
		case "ZoomSnap":
			res = fmt.Sprintf("%d", inferedInstance.ZoomSnap)
		case "LayerGroupUses":
			for idx, __instance__ := range inferedInstance.LayerGroupUses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Marker:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "DivIcon":
			if inferedInstance.DivIcon != nil {
				res = inferedInstance.DivIcon.Name
			}
		}
	case *UserClick:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "TimeOfClick":
			res = inferedInstance.TimeOfClick.String()
		}
	case *VLine:
		switch fieldName {
		// string value of fields
		case "StartLat":
			res = fmt.Sprintf("%f", inferedInstance.StartLat)
		case "StartLng":
			res = fmt.Sprintf("%f", inferedInstance.StartLng)
		case "EndLat":
			res = fmt.Sprintf("%f", inferedInstance.EndLat)
		case "EndLng":
			res = fmt.Sprintf("%f", inferedInstance.EndLng)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "DashStyleEnum":
			enum := inferedInstance.DashStyleEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "IsTransmitting":
			enum := inferedInstance.IsTransmitting
			res = enum.ToCodeString()
		case "Message":
			res = inferedInstance.Message
		case "IsTransmittingBackward":
			enum := inferedInstance.IsTransmittingBackward
			res = enum.ToCodeString()
		case "MessageBackward":
			res = inferedInstance.MessageBackward
		}
	case *VisualTrack:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "DivIcon":
			if inferedInstance.DivIcon != nil {
				res = inferedInstance.DivIcon.Name
			}
		case "DisplayTrackHistory":
			res = fmt.Sprintf("%t", inferedInstance.DisplayTrackHistory)
		case "DisplayLevelAndSpeed":
			res = fmt.Sprintf("%t", inferedInstance.DisplayLevelAndSpeed)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Circle:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "Radius":
			res = fmt.Sprintf("%f", inferedInstance.Radius)
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "DashStyleEnum":
			enum := inferedInstance.DashStyleEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		}
	case DivIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SVG":
			res = inferedInstance.SVG
		}
	case LayerGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DisplayName":
			res = inferedInstance.DisplayName
		}
	case LayerGroupUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsDisplayed":
			res = fmt.Sprintf("%t", inferedInstance.IsDisplayed)
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		}
	case MapOptions:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "ZoomLevel":
			res = fmt.Sprintf("%f", inferedInstance.ZoomLevel)
		case "UrlTemplate":
			res = inferedInstance.UrlTemplate
		case "Attribution":
			res = inferedInstance.Attribution
		case "MaxZoom":
			res = fmt.Sprintf("%d", inferedInstance.MaxZoom)
		case "ZoomControl":
			res = fmt.Sprintf("%t", inferedInstance.ZoomControl)
		case "AttributionControl":
			res = fmt.Sprintf("%t", inferedInstance.AttributionControl)
		case "ZoomSnap":
			res = fmt.Sprintf("%d", inferedInstance.ZoomSnap)
		case "LayerGroupUses":
			for idx, __instance__ := range inferedInstance.LayerGroupUses {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Marker:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "DivIcon":
			if inferedInstance.DivIcon != nil {
				res = inferedInstance.DivIcon.Name
			}
		}
	case UserClick:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "TimeOfClick":
			res = inferedInstance.TimeOfClick.String()
		}
	case VLine:
		switch fieldName {
		// string value of fields
		case "StartLat":
			res = fmt.Sprintf("%f", inferedInstance.StartLat)
		case "StartLng":
			res = fmt.Sprintf("%f", inferedInstance.StartLng)
		case "EndLat":
			res = fmt.Sprintf("%f", inferedInstance.EndLat)
		case "EndLng":
			res = fmt.Sprintf("%f", inferedInstance.EndLng)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "DashStyleEnum":
			enum := inferedInstance.DashStyleEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "IsTransmitting":
			enum := inferedInstance.IsTransmitting
			res = enum.ToCodeString()
		case "Message":
			res = inferedInstance.Message
		case "IsTransmittingBackward":
			enum := inferedInstance.IsTransmittingBackward
			res = enum.ToCodeString()
		case "MessageBackward":
			res = inferedInstance.MessageBackward
		}
	case VisualTrack:
		switch fieldName {
		// string value of fields
		case "Lat":
			res = fmt.Sprintf("%f", inferedInstance.Lat)
		case "Lng":
			res = fmt.Sprintf("%f", inferedInstance.Lng)
		case "Heading":
			res = fmt.Sprintf("%f", inferedInstance.Heading)
		case "Level":
			res = fmt.Sprintf("%f", inferedInstance.Level)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "VerticalSpeed":
			res = fmt.Sprintf("%f", inferedInstance.VerticalSpeed)
		case "Name":
			res = inferedInstance.Name
		case "ColorEnum":
			enum := inferedInstance.ColorEnum
			res = enum.ToCodeString()
		case "LayerGroup":
			if inferedInstance.LayerGroup != nil {
				res = inferedInstance.LayerGroup.Name
			}
		case "DivIcon":
			if inferedInstance.DivIcon != nil {
				res = inferedInstance.DivIcon.Name
			}
		case "DisplayTrackHistory":
			res = fmt.Sprintf("%t", inferedInstance.DisplayTrackHistory)
		case "DisplayLevelAndSpeed":
			res = fmt.Sprintf("%t", inferedInstance.DisplayLevelAndSpeed)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
