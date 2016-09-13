package citiesBase

import (
	"log"
	"testing"
)

func TestSimpleFind(t *testing.T) {
	ciB, err := NewDataBase("test/cities_base.sql")
	if err != nil {
		t.Fatal("Can not open database")
	}
	lett := "А"
	city, res := ciB.FindCityOnLetter(lett, 1)
	log.Print(city)
	if res != cityIsFound {
		t.Fatal("Can not find \"А\" city")
	}

	city, res = ciB.FindCityOnLetter("ы", 1)
	log.Print(city)
	if res != cityDoesNotExist {
		t.Fatal("Can not find \"ы\" city")
	}
}

func TestTempTableUtils(t *testing.T) {
	ciB, err := NewDataBase(":memory:")
	if err != nil {
		t.Fatal("Can not create database")
	}
	exists := ciB.tableExist(1)
	if exists == true {
		t.Fatal("False table exists")
	}

	err = ciB.createTempTable(1)
	exists1 := ciB.tableExist(1)
	exists2 := ciB.tableExist(2)
	if exists1 != true || exists2 != false || err != nil {
		t.Fatal("Can not create table1")
	}

	err = ciB.createTempTable(2)
	exists1 = ciB.tableExist(1)
	exists2 = ciB.tableExist(2)

	if exists1 != true || exists2 != true || err != nil {
		t.Fatal("Can not create table2")
	}

	err = ciB.deleteTempTable(1)
	exists1 = ciB.tableExist(1)
	exists2 = ciB.tableExist(2)
	if exists1 != false || exists2 != true || err != nil {
		t.Fatal("Can not delete table1")
	}

}
