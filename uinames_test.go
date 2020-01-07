package uinames

import (
	"log"
	"testing"
)

func TestGetName(t *testing.T) {
	name, err := GetName(nil)
	if err != nil {
		t.Errorf("Get Name Error: %v\n", err)
	}
	log.Println(name)
}

/*
func TestGetNameOpt(t *testing.T) {
	name, err := GetName(&Options{Region: "United States"})
	if err != nil {
		t.Errorf("Get Name Error: %v\n", err)
	}
	log.Println(name)
}
*/

func TestGetNameExtra(t *testing.T) {
	name, err := GetNameExtra(&Options{Region: "United States"})
	if err != nil {
		t.Errorf("Get Name Error: %v\n", err)
	}
	log.Println(name)
}

func TestGetNamesOpt(t *testing.T) {
	names, err := GetNamesExtra(&Options{Region: "United States"})
	if err != nil {
		t.Errorf("Get Name Error: %v\n", err)
	}
	log.Println(names)
}
