// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongtable/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_DisplayedColumn_sql sql.NullBool
var dummy_DisplayedColumn_time time.Duration
var dummy_DisplayedColumn_sort sort.Float64Slice

// DisplayedColumnAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model displayedcolumnAPI
type DisplayedColumnAPI struct {
	gorm.Model

	models.DisplayedColumn_WOP

	// encoding of pointers
	DisplayedColumnPointersEncoding DisplayedColumnPointersEncoding
}

// DisplayedColumnPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DisplayedColumnPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// DisplayedColumnDB describes a displayedcolumn in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model displayedcolumnDB
type DisplayedColumnDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field displayedcolumnDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	DisplayedColumnPointersEncoding
}

// DisplayedColumnDBs arrays displayedcolumnDBs
// swagger:response displayedcolumnDBsResponse
type DisplayedColumnDBs []DisplayedColumnDB

// DisplayedColumnDBResponse provides response
// swagger:response displayedcolumnDBResponse
type DisplayedColumnDBResponse struct {
	DisplayedColumnDB
}

// DisplayedColumnWOP is a DisplayedColumn without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DisplayedColumnWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var DisplayedColumn_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoDisplayedColumnStruct struct {
	// stores DisplayedColumnDB according to their gorm ID
	Map_DisplayedColumnDBID_DisplayedColumnDB map[uint]*DisplayedColumnDB

	// stores DisplayedColumnDB ID according to DisplayedColumn address
	Map_DisplayedColumnPtr_DisplayedColumnDBID map[*models.DisplayedColumn]uint

	// stores DisplayedColumn according to their gorm ID
	Map_DisplayedColumnDBID_DisplayedColumnPtr map[uint]*models.DisplayedColumn

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDisplayedColumn.stage
	return
}

func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) GetDB() *gorm.DB {
	return backRepoDisplayedColumn.db
}

// GetDisplayedColumnDBFromDisplayedColumnPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) GetDisplayedColumnDBFromDisplayedColumnPtr(displayedcolumn *models.DisplayedColumn) (displayedcolumnDB *DisplayedColumnDB) {
	id := backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]
	displayedcolumnDB = backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[id]
	return
}

// BackRepoDisplayedColumn.CommitPhaseOne commits all staged instances of DisplayedColumn to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for displayedcolumn := range stage.DisplayedColumns {
		backRepoDisplayedColumn.CommitPhaseOneInstance(displayedcolumn)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, displayedcolumn := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr {
		if _, ok := stage.DisplayedColumns[displayedcolumn]; !ok {
			backRepoDisplayedColumn.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDisplayedColumn.CommitDeleteInstance commits deletion of DisplayedColumn to the BackRepo
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CommitDeleteInstance(id uint) (Error error) {

	displayedcolumn := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[id]

	// displayedcolumn is not staged anymore, remove displayedcolumnDB
	displayedcolumnDB := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[id]
	query := backRepoDisplayedColumn.db.Unscoped().Delete(&displayedcolumnDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID, displayedcolumn)
	delete(backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr, id)
	delete(backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB, id)

	return
}

// BackRepoDisplayedColumn.CommitPhaseOneInstance commits displayedcolumn staged instances of DisplayedColumn to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CommitPhaseOneInstance(displayedcolumn *models.DisplayedColumn) (Error error) {

	// check if the displayedcolumn is not commited yet
	if _, ok := backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]; ok {
		return
	}

	// initiate displayedcolumn
	var displayedcolumnDB DisplayedColumnDB
	displayedcolumnDB.CopyBasicFieldsFromDisplayedColumn(displayedcolumn)

	query := backRepoDisplayedColumn.db.Create(&displayedcolumnDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn] = displayedcolumnDB.ID
	backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID] = displayedcolumn
	backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[displayedcolumnDB.ID] = &displayedcolumnDB

	return
}

// BackRepoDisplayedColumn.CommitPhaseTwo commits all staged instances of DisplayedColumn to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, displayedcolumn := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr {
		backRepoDisplayedColumn.CommitPhaseTwoInstance(backRepo, idx, displayedcolumn)
	}

	return
}

// BackRepoDisplayedColumn.CommitPhaseTwoInstance commits {{structname }} of models.DisplayedColumn to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, displayedcolumn *models.DisplayedColumn) (Error error) {

	// fetch matching displayedcolumnDB
	if displayedcolumnDB, ok := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[idx]; ok {

		displayedcolumnDB.CopyBasicFieldsFromDisplayedColumn(displayedcolumn)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoDisplayedColumn.db.Save(&displayedcolumnDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DisplayedColumn intance %s", displayedcolumn.Name))
		return err
	}

	return
}

// BackRepoDisplayedColumn.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CheckoutPhaseOne() (Error error) {

	displayedcolumnDBArray := make([]DisplayedColumnDB, 0)
	query := backRepoDisplayedColumn.db.Find(&displayedcolumnDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	displayedcolumnInstancesToBeRemovedFromTheStage := make(map[*models.DisplayedColumn]any)
	for key, value := range backRepoDisplayedColumn.stage.DisplayedColumns {
		displayedcolumnInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, displayedcolumnDB := range displayedcolumnDBArray {
		backRepoDisplayedColumn.CheckoutPhaseOneInstance(&displayedcolumnDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		displayedcolumn, ok := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
		if ok {
			delete(displayedcolumnInstancesToBeRemovedFromTheStage, displayedcolumn)
		}
	}

	// remove from stage and back repo's 3 maps all displayedcolumns that are not in the checkout
	for displayedcolumn := range displayedcolumnInstancesToBeRemovedFromTheStage {
		displayedcolumn.Unstage(backRepoDisplayedColumn.GetStage())

		// remove instance from the back repo 3 maps
		displayedcolumnID := backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]
		delete(backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID, displayedcolumn)
		delete(backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB, displayedcolumnID)
		delete(backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr, displayedcolumnID)
	}

	return
}

// CheckoutPhaseOneInstance takes a displayedcolumnDB that has been found in the DB, updates the backRepo and stages the
// models version of the displayedcolumnDB
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CheckoutPhaseOneInstance(displayedcolumnDB *DisplayedColumnDB) (Error error) {

	displayedcolumn, ok := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
	if !ok {
		displayedcolumn = new(models.DisplayedColumn)

		backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID] = displayedcolumn
		backRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn] = displayedcolumnDB.ID

		// append model store with the new element
		displayedcolumn.Name = displayedcolumnDB.Name_Data.String
		displayedcolumn.Stage(backRepoDisplayedColumn.GetStage())
	}
	displayedcolumnDB.CopyBasicFieldsToDisplayedColumn(displayedcolumn)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	displayedcolumn.Stage(backRepoDisplayedColumn.GetStage())

	// preserve pointer to displayedcolumnDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DisplayedColumnDBID_DisplayedColumnDB)[displayedcolumnDB hold variable pointers
	displayedcolumnDB_Data := *displayedcolumnDB
	preservedPtrToDisplayedColumn := &displayedcolumnDB_Data
	backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[displayedcolumnDB.ID] = preservedPtrToDisplayedColumn

	return
}

// BackRepoDisplayedColumn.CheckoutPhaseTwo Checkouts all staged instances of DisplayedColumn to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, displayedcolumnDB := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {
		backRepoDisplayedColumn.CheckoutPhaseTwoInstance(backRepo, displayedcolumnDB)
	}
	return
}

// BackRepoDisplayedColumn.CheckoutPhaseTwoInstance Checkouts staged instances of DisplayedColumn to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, displayedcolumnDB *DisplayedColumnDB) (Error error) {

	displayedcolumn := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr[displayedcolumnDB.ID]
	_ = displayedcolumn // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitDisplayedColumn allows commit of a single displayedcolumn (if already staged)
func (backRepo *BackRepoStruct) CommitDisplayedColumn(displayedcolumn *models.DisplayedColumn) {
	backRepo.BackRepoDisplayedColumn.CommitPhaseOneInstance(displayedcolumn)
	if id, ok := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]; ok {
		backRepo.BackRepoDisplayedColumn.CommitPhaseTwoInstance(backRepo, id, displayedcolumn)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDisplayedColumn allows checkout of a single displayedcolumn (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDisplayedColumn(displayedcolumn *models.DisplayedColumn) {
	// check if the displayedcolumn is staged
	if _, ok := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]; ok {

		if id, ok := backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnPtr_DisplayedColumnDBID[displayedcolumn]; ok {
			var displayedcolumnDB DisplayedColumnDB
			displayedcolumnDB.ID = id

			if err := backRepo.BackRepoDisplayedColumn.db.First(&displayedcolumnDB, id).Error; err != nil {
				log.Fatalln("CheckoutDisplayedColumn : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDisplayedColumn.CheckoutPhaseOneInstance(&displayedcolumnDB)
			backRepo.BackRepoDisplayedColumn.CheckoutPhaseTwoInstance(backRepo, &displayedcolumnDB)
		}
	}
}

// CopyBasicFieldsFromDisplayedColumn
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsFromDisplayedColumn(displayedcolumn *models.DisplayedColumn) {
	// insertion point for fields commit

	displayedcolumnDB.Name_Data.String = displayedcolumn.Name
	displayedcolumnDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDisplayedColumn_WOP
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsFromDisplayedColumn_WOP(displayedcolumn *models.DisplayedColumn_WOP) {
	// insertion point for fields commit

	displayedcolumnDB.Name_Data.String = displayedcolumn.Name
	displayedcolumnDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDisplayedColumnWOP
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsFromDisplayedColumnWOP(displayedcolumn *DisplayedColumnWOP) {
	// insertion point for fields commit

	displayedcolumnDB.Name_Data.String = displayedcolumn.Name
	displayedcolumnDB.Name_Data.Valid = true
}

// CopyBasicFieldsToDisplayedColumn
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsToDisplayedColumn(displayedcolumn *models.DisplayedColumn) {
	// insertion point for checkout of basic fields (back repo to stage)
	displayedcolumn.Name = displayedcolumnDB.Name_Data.String
}

// CopyBasicFieldsToDisplayedColumn_WOP
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsToDisplayedColumn_WOP(displayedcolumn *models.DisplayedColumn_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	displayedcolumn.Name = displayedcolumnDB.Name_Data.String
}

// CopyBasicFieldsToDisplayedColumnWOP
func (displayedcolumnDB *DisplayedColumnDB) CopyBasicFieldsToDisplayedColumnWOP(displayedcolumn *DisplayedColumnWOP) {
	displayedcolumn.ID = int(displayedcolumnDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	displayedcolumn.Name = displayedcolumnDB.Name_Data.String
}

// Backup generates a json file from a slice of all DisplayedColumnDB instances in the backrepo
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DisplayedColumnDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DisplayedColumnDB, 0)
	for _, displayedcolumnDB := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {
		forBackup = append(forBackup, displayedcolumnDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json DisplayedColumn ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json DisplayedColumn file", err.Error())
	}
}

// Backup generates a json file from a slice of all DisplayedColumnDB instances in the backrepo
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DisplayedColumnDB, 0)
	for _, displayedcolumnDB := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {
		forBackup = append(forBackup, displayedcolumnDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DisplayedColumn")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DisplayedColumn_Fields, -1)
	for _, displayedcolumnDB := range forBackup {

		var displayedcolumnWOP DisplayedColumnWOP
		displayedcolumnDB.CopyBasicFieldsToDisplayedColumnWOP(&displayedcolumnWOP)

		row := sh.AddRow()
		row.WriteStruct(&displayedcolumnWOP, -1)
	}
}

// RestoreXL from the "DisplayedColumn" sheet all DisplayedColumnDB instances
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDisplayedColumnid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DisplayedColumn"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDisplayedColumn.rowVisitorDisplayedColumn)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) rowVisitorDisplayedColumn(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var displayedcolumnWOP DisplayedColumnWOP
		row.ReadStruct(&displayedcolumnWOP)

		// add the unmarshalled struct to the stage
		displayedcolumnDB := new(DisplayedColumnDB)
		displayedcolumnDB.CopyBasicFieldsFromDisplayedColumnWOP(&displayedcolumnWOP)

		displayedcolumnDB_ID_atBackupTime := displayedcolumnDB.ID
		displayedcolumnDB.ID = 0
		query := backRepoDisplayedColumn.db.Create(displayedcolumnDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[displayedcolumnDB.ID] = displayedcolumnDB
		BackRepoDisplayedColumnid_atBckpTime_newID[displayedcolumnDB_ID_atBackupTime] = displayedcolumnDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DisplayedColumnDB.json" in dirPath that stores an array
// of DisplayedColumnDB and stores it in the database
// the map BackRepoDisplayedColumnid_atBckpTime_newID is updated accordingly
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDisplayedColumnid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DisplayedColumnDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json DisplayedColumn file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DisplayedColumnDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DisplayedColumnDBID_DisplayedColumnDB
	for _, displayedcolumnDB := range forRestore {

		displayedcolumnDB_ID_atBackupTime := displayedcolumnDB.ID
		displayedcolumnDB.ID = 0
		query := backRepoDisplayedColumn.db.Create(displayedcolumnDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[displayedcolumnDB.ID] = displayedcolumnDB
		BackRepoDisplayedColumnid_atBckpTime_newID[displayedcolumnDB_ID_atBackupTime] = displayedcolumnDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json DisplayedColumn file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DisplayedColumn>id_atBckpTime_newID
// to compute new index
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) RestorePhaseTwo() {

	for _, displayedcolumnDB := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {

		// next line of code is to avert unused variable compilation error
		_ = displayedcolumnDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoDisplayedColumn.db.Model(displayedcolumnDB).Updates(*displayedcolumnDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDisplayedColumn.ResetReversePointers commits all staged instances of DisplayedColumn to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, displayedcolumn := range backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnPtr {
		backRepoDisplayedColumn.ResetReversePointersInstance(backRepo, idx, displayedcolumn)
	}

	return
}

func (backRepoDisplayedColumn *BackRepoDisplayedColumnStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.DisplayedColumn) (Error error) {

	// fetch matching displayedcolumnDB
	if displayedcolumnDB, ok := backRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB[idx]; ok {
		_ = displayedcolumnDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDisplayedColumnid_atBckpTime_newID map[uint]uint