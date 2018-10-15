package airports

import "testing"

func TestLoad(t *testing.T) {
	Load("server/airports.csv")
	airport := Get("LAX")
	if airport.IATA != "LAX" {
		t.Error("LAX not found!")
	}
	airport = Get("BADIATA")
	if airport.IATA != "" {
		t.Error("BADIATA found!")
	}
}

func TestLoadDB(t *testing.T) {
	LoadDB()
	airport := Get("LAX")
	if airport.IATA != "LAX" {
		t.Error("LAX not found!")
	}
	airport = Get("BADIATA")
	if airport.IATA != "" {
		t.Error("BADIATA found!")
	}
}
