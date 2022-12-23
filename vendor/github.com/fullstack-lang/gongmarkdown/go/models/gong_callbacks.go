package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AnotherDummyData:
		if stage.OnAfterAnotherDummyDataCreateCallback != nil {
			stage.OnAfterAnotherDummyDataCreateCallback.OnAfterCreate(stage, target)
		}
	case *Cell:
		if stage.OnAfterCellCreateCallback != nil {
			stage.OnAfterCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *DummyData:
		if stage.OnAfterDummyDataCreateCallback != nil {
			stage.OnAfterDummyDataCreateCallback.OnAfterCreate(stage, target)
		}
	case *Element:
		if stage.OnAfterElementCreateCallback != nil {
			stage.OnAfterElementCreateCallback.OnAfterCreate(stage, target)
		}
	case *MarkdownContent:
		if stage.OnAfterMarkdownContentCreateCallback != nil {
			stage.OnAfterMarkdownContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Row:
		if stage.OnAfterRowCreateCallback != nil {
			stage.OnAfterRowCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AnotherDummyData:
		newTarget := any(new).(*AnotherDummyData)
		if stage.OnAfterAnotherDummyDataUpdateCallback != nil {
			stage.OnAfterAnotherDummyDataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Cell:
		newTarget := any(new).(*Cell)
		if stage.OnAfterCellUpdateCallback != nil {
			stage.OnAfterCellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DummyData:
		newTarget := any(new).(*DummyData)
		if stage.OnAfterDummyDataUpdateCallback != nil {
			stage.OnAfterDummyDataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Element:
		newTarget := any(new).(*Element)
		if stage.OnAfterElementUpdateCallback != nil {
			stage.OnAfterElementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MarkdownContent:
		newTarget := any(new).(*MarkdownContent)
		if stage.OnAfterMarkdownContentUpdateCallback != nil {
			stage.OnAfterMarkdownContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Row:
		newTarget := any(new).(*Row)
		if stage.OnAfterRowUpdateCallback != nil {
			stage.OnAfterRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AnotherDummyData:
		if stage.OnAfterAnotherDummyDataDeleteCallback != nil {
			staged := any(staged).(*AnotherDummyData)
			stage.OnAfterAnotherDummyDataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Cell:
		if stage.OnAfterCellDeleteCallback != nil {
			staged := any(staged).(*Cell)
			stage.OnAfterCellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DummyData:
		if stage.OnAfterDummyDataDeleteCallback != nil {
			staged := any(staged).(*DummyData)
			stage.OnAfterDummyDataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Element:
		if stage.OnAfterElementDeleteCallback != nil {
			staged := any(staged).(*Element)
			stage.OnAfterElementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MarkdownContent:
		if stage.OnAfterMarkdownContentDeleteCallback != nil {
			staged := any(staged).(*MarkdownContent)
			stage.OnAfterMarkdownContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Row:
		if stage.OnAfterRowDeleteCallback != nil {
			staged := any(staged).(*Row)
			stage.OnAfterRowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AnotherDummyData:
		if stage.OnAfterAnotherDummyDataReadCallback != nil {
			stage.OnAfterAnotherDummyDataReadCallback.OnAfterRead(stage, target)
		}
	case *Cell:
		if stage.OnAfterCellReadCallback != nil {
			stage.OnAfterCellReadCallback.OnAfterRead(stage, target)
		}
	case *DummyData:
		if stage.OnAfterDummyDataReadCallback != nil {
			stage.OnAfterDummyDataReadCallback.OnAfterRead(stage, target)
		}
	case *Element:
		if stage.OnAfterElementReadCallback != nil {
			stage.OnAfterElementReadCallback.OnAfterRead(stage, target)
		}
	case *MarkdownContent:
		if stage.OnAfterMarkdownContentReadCallback != nil {
			stage.OnAfterMarkdownContentReadCallback.OnAfterRead(stage, target)
		}
	case *Row:
		if stage.OnAfterRowReadCallback != nil {
			stage.OnAfterRowReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AnotherDummyData:
		stage.OnAfterAnotherDummyDataUpdateCallback = any(callback).(OnAfterUpdateInterface[AnotherDummyData])
	
	case *Cell:
		stage.OnAfterCellUpdateCallback = any(callback).(OnAfterUpdateInterface[Cell])
	
	case *DummyData:
		stage.OnAfterDummyDataUpdateCallback = any(callback).(OnAfterUpdateInterface[DummyData])
	
	case *Element:
		stage.OnAfterElementUpdateCallback = any(callback).(OnAfterUpdateInterface[Element])
	
	case *MarkdownContent:
		stage.OnAfterMarkdownContentUpdateCallback = any(callback).(OnAfterUpdateInterface[MarkdownContent])
	
	case *Row:
		stage.OnAfterRowUpdateCallback = any(callback).(OnAfterUpdateInterface[Row])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AnotherDummyData:
		stage.OnAfterAnotherDummyDataCreateCallback = any(callback).(OnAfterCreateInterface[AnotherDummyData])
	
	case *Cell:
		stage.OnAfterCellCreateCallback = any(callback).(OnAfterCreateInterface[Cell])
	
	case *DummyData:
		stage.OnAfterDummyDataCreateCallback = any(callback).(OnAfterCreateInterface[DummyData])
	
	case *Element:
		stage.OnAfterElementCreateCallback = any(callback).(OnAfterCreateInterface[Element])
	
	case *MarkdownContent:
		stage.OnAfterMarkdownContentCreateCallback = any(callback).(OnAfterCreateInterface[MarkdownContent])
	
	case *Row:
		stage.OnAfterRowCreateCallback = any(callback).(OnAfterCreateInterface[Row])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AnotherDummyData:
		stage.OnAfterAnotherDummyDataDeleteCallback = any(callback).(OnAfterDeleteInterface[AnotherDummyData])
	
	case *Cell:
		stage.OnAfterCellDeleteCallback = any(callback).(OnAfterDeleteInterface[Cell])
	
	case *DummyData:
		stage.OnAfterDummyDataDeleteCallback = any(callback).(OnAfterDeleteInterface[DummyData])
	
	case *Element:
		stage.OnAfterElementDeleteCallback = any(callback).(OnAfterDeleteInterface[Element])
	
	case *MarkdownContent:
		stage.OnAfterMarkdownContentDeleteCallback = any(callback).(OnAfterDeleteInterface[MarkdownContent])
	
	case *Row:
		stage.OnAfterRowDeleteCallback = any(callback).(OnAfterDeleteInterface[Row])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AnotherDummyData:
		stage.OnAfterAnotherDummyDataReadCallback = any(callback).(OnAfterReadInterface[AnotherDummyData])
	
	case *Cell:
		stage.OnAfterCellReadCallback = any(callback).(OnAfterReadInterface[Cell])
	
	case *DummyData:
		stage.OnAfterDummyDataReadCallback = any(callback).(OnAfterReadInterface[DummyData])
	
	case *Element:
		stage.OnAfterElementReadCallback = any(callback).(OnAfterReadInterface[Element])
	
	case *MarkdownContent:
		stage.OnAfterMarkdownContentReadCallback = any(callback).(OnAfterReadInterface[MarkdownContent])
	
	case *Row:
		stage.OnAfterRowReadCallback = any(callback).(OnAfterReadInterface[Row])
	
	}
}
