package airports

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

type Airport struct {
	IATA string
	Name string
}

var airports = make(map[string]Airport)

func Load(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		airport := new(Airport)
		airport.IATA = record[13]
		airport.Name = record[3]
		airports[airport.IATA] = *airport
	}
}

func Get(iata string) Airport  {
	return airports[iata]
}
