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

	"github.com/fullstack-lang/gongng2charts/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Label_sql sql.NullBool
var dummy_Label_time time.Duration
var dummy_Label_sort sort.Float64Slice

// LabelAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model labelAPI
type LabelAPI struct {
	gorm.Model

	models.Label

	// encoding of pointers
	LabelPointersEnconding
}

// LabelPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LabelPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field ChartConfiguration{}.Labels []*Label
	ChartConfiguration_LabelsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	ChartConfiguration_LabelsDBID_Index sql.NullInt64
}

// LabelDB describes a label in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model labelDB
type LabelDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field labelDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	LabelPointersEnconding
}

// LabelDBs arrays labelDBs
// swagger:response labelDBsResponse
type LabelDBs []LabelDB

// LabelDBResponse provides response
// swagger:response labelDBResponse
type LabelDBResponse struct {
	LabelDB
}

// LabelWOP is a Label without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LabelWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Label_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoLabelStruct struct {
	// stores LabelDB according to their gorm ID
	Map_LabelDBID_LabelDB *map[uint]*LabelDB

	// stores LabelDB ID according to Label address
	Map_LabelPtr_LabelDBID *map[*models.Label]uint

	// stores Label according to their gorm ID
	Map_LabelDBID_LabelPtr *map[uint]*models.Label

	db *gorm.DB
}

func (backRepoLabel *BackRepoLabelStruct) GetDB() *gorm.DB {
	return backRepoLabel.db
}

// GetLabelDBFromLabelPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLabel *BackRepoLabelStruct) GetLabelDBFromLabelPtr(label *models.Label) (labelDB *LabelDB) {
	id := (*backRepoLabel.Map_LabelPtr_LabelDBID)[label]
	labelDB = (*backRepoLabel.Map_LabelDBID_LabelDB)[id]
	return
}

// BackRepoLabel.Init set up the BackRepo of the Label
func (backRepoLabel *BackRepoLabelStruct) Init(db *gorm.DB) (Error error) {

	if backRepoLabel.Map_LabelDBID_LabelPtr != nil {
		err := errors.New("In Init, backRepoLabel.Map_LabelDBID_LabelPtr should be nil")
		return err
	}

	if backRepoLabel.Map_LabelDBID_LabelDB != nil {
		err := errors.New("In Init, backRepoLabel.Map_LabelDBID_LabelDB should be nil")
		return err
	}

	if backRepoLabel.Map_LabelPtr_LabelDBID != nil {
		err := errors.New("In Init, backRepoLabel.Map_LabelPtr_LabelDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Label, 0)
	backRepoLabel.Map_LabelDBID_LabelPtr = &tmp

	tmpDB := make(map[uint]*LabelDB, 0)
	backRepoLabel.Map_LabelDBID_LabelDB = &tmpDB

	tmpID := make(map[*models.Label]uint, 0)
	backRepoLabel.Map_LabelPtr_LabelDBID = &tmpID

	backRepoLabel.db = db
	return
}

// BackRepoLabel.CommitPhaseOne commits all staged instances of Label to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLabel *BackRepoLabelStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for label := range stage.Labels {
		backRepoLabel.CommitPhaseOneInstance(label)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, label := range *backRepoLabel.Map_LabelDBID_LabelPtr {
		if _, ok := stage.Labels[label]; !ok {
			backRepoLabel.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLabel.CommitDeleteInstance commits deletion of Label to the BackRepo
func (backRepoLabel *BackRepoLabelStruct) CommitDeleteInstance(id uint) (Error error) {

	label := (*backRepoLabel.Map_LabelDBID_LabelPtr)[id]

	// label is not staged anymore, remove labelDB
	labelDB := (*backRepoLabel.Map_LabelDBID_LabelDB)[id]
	query := backRepoLabel.db.Unscoped().Delete(&labelDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoLabel.Map_LabelPtr_LabelDBID), label)
	delete((*backRepoLabel.Map_LabelDBID_LabelPtr), id)
	delete((*backRepoLabel.Map_LabelDBID_LabelDB), id)

	return
}

// BackRepoLabel.CommitPhaseOneInstance commits label staged instances of Label to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLabel *BackRepoLabelStruct) CommitPhaseOneInstance(label *models.Label) (Error error) {

	// check if the label is not commited yet
	if _, ok := (*backRepoLabel.Map_LabelPtr_LabelDBID)[label]; ok {
		return
	}

	// initiate label
	var labelDB LabelDB
	labelDB.CopyBasicFieldsFromLabel(label)

	query := backRepoLabel.db.Create(&labelDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoLabel.Map_LabelPtr_LabelDBID)[label] = labelDB.ID
	(*backRepoLabel.Map_LabelDBID_LabelPtr)[labelDB.ID] = label
	(*backRepoLabel.Map_LabelDBID_LabelDB)[labelDB.ID] = &labelDB

	return
}

// BackRepoLabel.CommitPhaseTwo commits all staged instances of Label to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLabel *BackRepoLabelStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, label := range *backRepoLabel.Map_LabelDBID_LabelPtr {
		backRepoLabel.CommitPhaseTwoInstance(backRepo, idx, label)
	}

	return
}

// BackRepoLabel.CommitPhaseTwoInstance commits {{structname }} of models.Label to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLabel *BackRepoLabelStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, label *models.Label) (Error error) {

	// fetch matching labelDB
	if labelDB, ok := (*backRepoLabel.Map_LabelDBID_LabelDB)[idx]; ok {

		labelDB.CopyBasicFieldsFromLabel(label)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoLabel.db.Save(&labelDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Label intance %s", label.Name))
		return err
	}

	return
}

// BackRepoLabel.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoLabel *BackRepoLabelStruct) CheckoutPhaseOne() (Error error) {

	labelDBArray := make([]LabelDB, 0)
	query := backRepoLabel.db.Find(&labelDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	labelInstancesToBeRemovedFromTheStage := make(map[*models.Label]any)
	for key, value := range models.Stage.Labels {
		labelInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, labelDB := range labelDBArray {
		backRepoLabel.CheckoutPhaseOneInstance(&labelDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		label, ok := (*backRepoLabel.Map_LabelDBID_LabelPtr)[labelDB.ID]
		if ok {
			delete(labelInstancesToBeRemovedFromTheStage, label)
		}
	}

	// remove from stage and back repo's 3 maps all labels that are not in the checkout
	for label := range labelInstancesToBeRemovedFromTheStage {
		label.Unstage()

		// remove instance from the back repo 3 maps
		labelID := (*backRepoLabel.Map_LabelPtr_LabelDBID)[label]
		delete((*backRepoLabel.Map_LabelPtr_LabelDBID), label)
		delete((*backRepoLabel.Map_LabelDBID_LabelDB), labelID)
		delete((*backRepoLabel.Map_LabelDBID_LabelPtr), labelID)
	}

	return
}

// CheckoutPhaseOneInstance takes a labelDB that has been found in the DB, updates the backRepo and stages the
// models version of the labelDB
func (backRepoLabel *BackRepoLabelStruct) CheckoutPhaseOneInstance(labelDB *LabelDB) (Error error) {

	label, ok := (*backRepoLabel.Map_LabelDBID_LabelPtr)[labelDB.ID]
	if !ok {
		label = new(models.Label)

		(*backRepoLabel.Map_LabelDBID_LabelPtr)[labelDB.ID] = label
		(*backRepoLabel.Map_LabelPtr_LabelDBID)[label] = labelDB.ID

		// append model store with the new element
		label.Name = labelDB.Name_Data.String
		label.Stage()
	}
	labelDB.CopyBasicFieldsToLabel(label)

	// preserve pointer to labelDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LabelDBID_LabelDB)[labelDB hold variable pointers
	labelDB_Data := *labelDB
	preservedPtrToLabel := &labelDB_Data
	(*backRepoLabel.Map_LabelDBID_LabelDB)[labelDB.ID] = preservedPtrToLabel

	return
}

// BackRepoLabel.CheckoutPhaseTwo Checkouts all staged instances of Label to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLabel *BackRepoLabelStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, labelDB := range *backRepoLabel.Map_LabelDBID_LabelDB {
		backRepoLabel.CheckoutPhaseTwoInstance(backRepo, labelDB)
	}
	return
}

// BackRepoLabel.CheckoutPhaseTwoInstance Checkouts staged instances of Label to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLabel *BackRepoLabelStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, labelDB *LabelDB) (Error error) {

	label := (*backRepoLabel.Map_LabelDBID_LabelPtr)[labelDB.ID]
	_ = label // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitLabel allows commit of a single label (if already staged)
func (backRepo *BackRepoStruct) CommitLabel(label *models.Label) {
	backRepo.BackRepoLabel.CommitPhaseOneInstance(label)
	if id, ok := (*backRepo.BackRepoLabel.Map_LabelPtr_LabelDBID)[label]; ok {
		backRepo.BackRepoLabel.CommitPhaseTwoInstance(backRepo, id, label)
	}
}

// CommitLabel allows checkout of a single label (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLabel(label *models.Label) {
	// check if the label is staged
	if _, ok := (*backRepo.BackRepoLabel.Map_LabelPtr_LabelDBID)[label]; ok {

		if id, ok := (*backRepo.BackRepoLabel.Map_LabelPtr_LabelDBID)[label]; ok {
			var labelDB LabelDB
			labelDB.ID = id

			if err := backRepo.BackRepoLabel.db.First(&labelDB, id).Error; err != nil {
				log.Panicln("CheckoutLabel : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLabel.CheckoutPhaseOneInstance(&labelDB)
			backRepo.BackRepoLabel.CheckoutPhaseTwoInstance(backRepo, &labelDB)
		}
	}
}

// CopyBasicFieldsFromLabel
func (labelDB *LabelDB) CopyBasicFieldsFromLabel(label *models.Label) {
	// insertion point for fields commit

	labelDB.Name_Data.String = label.Name
	labelDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromLabelWOP
func (labelDB *LabelDB) CopyBasicFieldsFromLabelWOP(label *LabelWOP) {
	// insertion point for fields commit

	labelDB.Name_Data.String = label.Name
	labelDB.Name_Data.Valid = true
}

// CopyBasicFieldsToLabel
func (labelDB *LabelDB) CopyBasicFieldsToLabel(label *models.Label) {
	// insertion point for checkout of basic fields (back repo to stage)
	label.Name = labelDB.Name_Data.String
}

// CopyBasicFieldsToLabelWOP
func (labelDB *LabelDB) CopyBasicFieldsToLabelWOP(label *LabelWOP) {
	label.ID = int(labelDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	label.Name = labelDB.Name_Data.String
}

// Backup generates a json file from a slice of all LabelDB instances in the backrepo
func (backRepoLabel *BackRepoLabelStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LabelDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LabelDB, 0)
	for _, labelDB := range *backRepoLabel.Map_LabelDBID_LabelDB {
		forBackup = append(forBackup, labelDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Label ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Label file", err.Error())
	}
}

// Backup generates a json file from a slice of all LabelDB instances in the backrepo
func (backRepoLabel *BackRepoLabelStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LabelDB, 0)
	for _, labelDB := range *backRepoLabel.Map_LabelDBID_LabelDB {
		forBackup = append(forBackup, labelDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Label")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Label_Fields, -1)
	for _, labelDB := range forBackup {

		var labelWOP LabelWOP
		labelDB.CopyBasicFieldsToLabelWOP(&labelWOP)

		row := sh.AddRow()
		row.WriteStruct(&labelWOP, -1)
	}
}

// RestoreXL from the "Label" sheet all LabelDB instances
func (backRepoLabel *BackRepoLabelStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLabelid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Label"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLabel.rowVisitorLabel)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLabel *BackRepoLabelStruct) rowVisitorLabel(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var labelWOP LabelWOP
		row.ReadStruct(&labelWOP)

		// add the unmarshalled struct to the stage
		labelDB := new(LabelDB)
		labelDB.CopyBasicFieldsFromLabelWOP(&labelWOP)

		labelDB_ID_atBackupTime := labelDB.ID
		labelDB.ID = 0
		query := backRepoLabel.db.Create(labelDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLabel.Map_LabelDBID_LabelDB)[labelDB.ID] = labelDB
		BackRepoLabelid_atBckpTime_newID[labelDB_ID_atBackupTime] = labelDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LabelDB.json" in dirPath that stores an array
// of LabelDB and stores it in the database
// the map BackRepoLabelid_atBckpTime_newID is updated accordingly
func (backRepoLabel *BackRepoLabelStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLabelid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LabelDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Label file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LabelDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LabelDBID_LabelDB
	for _, labelDB := range forRestore {

		labelDB_ID_atBackupTime := labelDB.ID
		labelDB.ID = 0
		query := backRepoLabel.db.Create(labelDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLabel.Map_LabelDBID_LabelDB)[labelDB.ID] = labelDB
		BackRepoLabelid_atBckpTime_newID[labelDB_ID_atBackupTime] = labelDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Label file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Label>id_atBckpTime_newID
// to compute new index
func (backRepoLabel *BackRepoLabelStruct) RestorePhaseTwo() {

	for _, labelDB := range *backRepoLabel.Map_LabelDBID_LabelDB {

		// next line of code is to avert unused variable compilation error
		_ = labelDB

		// insertion point for reindexing pointers encoding
		// This reindex label.Labels
		if labelDB.ChartConfiguration_LabelsDBID.Int64 != 0 {
			labelDB.ChartConfiguration_LabelsDBID.Int64 =
				int64(BackRepoChartConfigurationid_atBckpTime_newID[uint(labelDB.ChartConfiguration_LabelsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLabel.db.Model(labelDB).Updates(*labelDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLabelid_atBckpTime_newID map[uint]uint
