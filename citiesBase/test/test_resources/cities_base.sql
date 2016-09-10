package citiesBase

import (
	"fmt"
	"github.com/mxk/go-sqlite/sqlite3"
	"go/constant"
)

const (
	cityIsFound = iota
	// Can not find any city which starts from this letter
	cityDoesNotExist = iota
	// All cities which start from this letter were already used
	allCitiesWereUsed = iota
)

type dataBase struct {
	conn *sqlite3.Conn
}

func NewDataBase(dbPath string) (error, *dataBase) {
	var err error
	db := new(dataBase)
	if db.conn == nil {
		db.conn, err = sqlite3.Open(dbPath)
	}
	return err, &db
}

func (db dataBase) checkCitiesOnLetter(letter string) (string, int) {
	sql := "SELECT name FROM cities WHERE name LIKE '" + letter + "%';"
	//row := make(sqlite3.RowMap)
	for s, err := db.conn.Query(sql); err == nil; err = s.Next() {
		var city string
		s.Scan(&city)
		return city, cityIsFound
	}
	return "", cityDoesNotExist
}

func FindNextCity(startLetter string, sessionId uint64) string {
	err, db := NewDataBase("cities.sql")
	if err {
		return "aaa"
	}
	city, _ := db.checkCitiesOnLetter("а")
	return city
}
