// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongfly/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	civilianairportDBs map[uint]*CivilianAirportDB

	nextIDCivilianAirportDB uint

	linerDBs map[uint]*LinerDB

	nextIDLinerDB uint

	messageDBs map[uint]*MessageDB

	nextIDMessageDB uint

	opslineDBs map[uint]*OpsLineDB

	nextIDOpsLineDB uint

	radarDBs map[uint]*RadarDB

	nextIDRadarDB uint

	satelliteDBs map[uint]*SatelliteDB

	nextIDSatelliteDB uint

	scenarioDBs map[uint]*ScenarioDB

	nextIDScenarioDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		civilianairportDBs: make(map[uint]*CivilianAirportDB),

		linerDBs: make(map[uint]*LinerDB),

		messageDBs: make(map[uint]*MessageDB),

		opslineDBs: make(map[uint]*OpsLineDB),

		radarDBs: make(map[uint]*RadarDB),

		satelliteDBs: make(map[uint]*SatelliteDB),

		scenarioDBs: make(map[uint]*ScenarioDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *CivilianAirportDB:
		db.nextIDCivilianAirportDB++
		v.ID = db.nextIDCivilianAirportDB
		db.civilianairportDBs[v.ID] = v
	case *LinerDB:
		db.nextIDLinerDB++
		v.ID = db.nextIDLinerDB
		db.linerDBs[v.ID] = v
	case *MessageDB:
		db.nextIDMessageDB++
		v.ID = db.nextIDMessageDB
		db.messageDBs[v.ID] = v
	case *OpsLineDB:
		db.nextIDOpsLineDB++
		v.ID = db.nextIDOpsLineDB
		db.opslineDBs[v.ID] = v
	case *RadarDB:
		db.nextIDRadarDB++
		v.ID = db.nextIDRadarDB
		db.radarDBs[v.ID] = v
	case *SatelliteDB:
		db.nextIDSatelliteDB++
		v.ID = db.nextIDSatelliteDB
		db.satelliteDBs[v.ID] = v
	case *ScenarioDB:
		db.nextIDScenarioDB++
		v.ID = db.nextIDScenarioDB
		db.scenarioDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CivilianAirportDB:
		delete(db.civilianairportDBs, v.ID)
	case *LinerDB:
		delete(db.linerDBs, v.ID)
	case *MessageDB:
		delete(db.messageDBs, v.ID)
	case *OpsLineDB:
		delete(db.opslineDBs, v.ID)
	case *RadarDB:
		delete(db.radarDBs, v.ID)
	case *SatelliteDB:
		delete(db.satelliteDBs, v.ID)
	case *ScenarioDB:
		delete(db.scenarioDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CivilianAirportDB:
		db.civilianairportDBs[v.ID] = v
		return db, nil
	case *LinerDB:
		db.linerDBs[v.ID] = v
		return db, nil
	case *MessageDB:
		db.messageDBs[v.ID] = v
		return db, nil
	case *OpsLineDB:
		db.opslineDBs[v.ID] = v
		return db, nil
	case *RadarDB:
		db.radarDBs[v.ID] = v
		return db, nil
	case *SatelliteDB:
		db.satelliteDBs[v.ID] = v
		return db, nil
	case *ScenarioDB:
		db.scenarioDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *CivilianAirportDB:
		if existing, ok := db.civilianairportDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *LinerDB:
		if existing, ok := db.linerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *MessageDB:
		if existing, ok := db.messageDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *OpsLineDB:
		if existing, ok := db.opslineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *RadarDB:
		if existing, ok := db.radarDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *SatelliteDB:
		if existing, ok := db.satelliteDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	case *ScenarioDB:
		if existing, ok := db.scenarioDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongfly/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]CivilianAirportDB:
        *ptr = make([]CivilianAirportDB, 0, len(db.civilianairportDBs))
        for _, v := range db.civilianairportDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]LinerDB:
        *ptr = make([]LinerDB, 0, len(db.linerDBs))
        for _, v := range db.linerDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]MessageDB:
        *ptr = make([]MessageDB, 0, len(db.messageDBs))
        for _, v := range db.messageDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]OpsLineDB:
        *ptr = make([]OpsLineDB, 0, len(db.opslineDBs))
        for _, v := range db.opslineDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]RadarDB:
        *ptr = make([]RadarDB, 0, len(db.radarDBs))
        for _, v := range db.radarDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SatelliteDB:
        *ptr = make([]SatelliteDB, 0, len(db.satelliteDBs))
        for _, v := range db.satelliteDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ScenarioDB:
        *ptr = make([]ScenarioDB, 0, len(db.scenarioDBs))
        for _, v := range db.scenarioDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("github.com/fullstack-lang/gongfly/go, Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *CivilianAirportDB:
		tmp, ok := db.civilianairportDBs[uint(i)]

		civilianairportDB, _ := instanceDB.(*CivilianAirportDB)
		*civilianairportDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *LinerDB:
		tmp, ok := db.linerDBs[uint(i)]

		linerDB, _ := instanceDB.(*LinerDB)
		*linerDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *MessageDB:
		tmp, ok := db.messageDBs[uint(i)]

		messageDB, _ := instanceDB.(*MessageDB)
		*messageDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *OpsLineDB:
		tmp, ok := db.opslineDBs[uint(i)]

		opslineDB, _ := instanceDB.(*OpsLineDB)
		*opslineDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *RadarDB:
		tmp, ok := db.radarDBs[uint(i)]

		radarDB, _ := instanceDB.(*RadarDB)
		*radarDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SatelliteDB:
		tmp, ok := db.satelliteDBs[uint(i)]

		satelliteDB, _ := instanceDB.(*SatelliteDB)
		*satelliteDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ScenarioDB:
		tmp, ok := db.scenarioDBs[uint(i)]

		scenarioDB, _ := instanceDB.(*ScenarioDB)
		*scenarioDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongfly/go, Unkown type")
	}
	
	return db, nil
}

