package citiesBase

import (
	"log"
	"strconv"

	"github.com/mxk/go-sqlite/sqlite3"
)

const (
	cityIsFound = iota
	// Can not find any city which starts from this letter
	cityDoesNotExist = iota
	// All cities which start from this letter were already used
	allCitiesWereUsed = iota
)

// DataBase sturct stores connection to database
type DataBase struct {
	conn *sqlite3.Conn
}

// NewDataBase Creates new connection to database
func NewDataBase(dbPath string) (*DataBase, error) {
	var err error
	db := new(DataBase)
	if db.conn == nil {
		db.conn, err = sqlite3.Open(dbPath)
	}
	return db, err
}

// FindCityOnLetter Looks for a city name which starts from specified letter and hasn't been used yet
func (db DataBase) FindCityOnLetter(letter string, chatID int64) (string, int) {
	if db.tableExist(chatID) == false {
		db.createTempTable(chatID)
	}

	s, err := db.conn.Prepare("SELECT name FROM cities as t1 LEFT JOIN id_? WHERE name LIKE '?%';")
	if err != nil {
		log.Panicf("Can not prepare request: %s", err)
	}

	for err := s.Query(chatID, letter); err == nil; err = s.Next() {
		var city string
		s.Scan(&city)
		return city, cityIsFound
	}
	return "", cityDoesNotExist
}

func (db DataBase) tableExist(userID int64) bool {
	log.Printf("Check if temp table for user %d exist", userID)
	sql := "SELECT * FROM sqlite_temp_master WHERE type='table' AND NAME='id_" +
		strconv.FormatInt(userID, 10) + "';"

	s, err := db.conn.Query(sql)
	if err == nil {
		s.Close()
		return true
	}

	return false
}

func (db DataBase) deleteTempTable(userID int64) error {
	log.Printf("Delete temp table for user %d", userID)
	sql := "DROP TABLE id_" + strconv.FormatInt(userID, 10) + ";"
	err := db.conn.Exec(sql)
	if err != nil {
		log.Printf("Can not delete temp table for user %d, err %s\n%s",
			userID, err.Error(), sql)
	}
	return err
}

func (db DataBase) createTempTable(userID int64) error {
	log.Printf("Create temp table for user %d", userID)

	sql := "CREATE TEMP TABLE id_" + strconv.FormatInt(userID, 10) + "(name TEXT);"
	err := db.conn.Exec(sql)
	if err != nil {
		log.Printf("Can not create temp table for user %d, err %s\n%s",
			userID, err.Error(), sql)
	}
	return err
}
