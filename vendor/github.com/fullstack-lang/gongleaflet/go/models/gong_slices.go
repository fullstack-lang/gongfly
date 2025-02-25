// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Circle
	// insertion point per field

	// Compute reverse map for named struct DivIcon
	// insertion point per field

	// Compute reverse map for named struct LayerGroup
	// insertion point per field

	// Compute reverse map for named struct LayerGroupUse
	// insertion point per field

	// Compute reverse map for named struct MapOptions
	// insertion point per field
	clear(stage.MapOptions_LayerGroupUses_reverseMap)
	stage.MapOptions_LayerGroupUses_reverseMap = make(map[*LayerGroupUse]*MapOptions)
	for mapoptions := range stage.MapOptionss {
		_ = mapoptions
		for _, _layergroupuse := range mapoptions.LayerGroupUses {
			stage.MapOptions_LayerGroupUses_reverseMap[_layergroupuse] = mapoptions
		}
	}

	// Compute reverse map for named struct Marker
	// insertion point per field

	// Compute reverse map for named struct UserClick
	// insertion point per field

	// Compute reverse map for named struct VLine
	// insertion point per field

	// Compute reverse map for named struct VisualTrack
	// insertion point per field

}
