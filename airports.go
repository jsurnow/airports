package airports

import (
	"database/sql"
	"encoding/csv"
	_ "github.com/lib/pq"
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

func LoadDB() {
	db, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable host=postgres")
	if (err != nil) {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT iata, name FROM airport")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var iata string
		var name string
		if err := rows.Scan(&iata, &name); err != nil {
			log.Fatal(err)
		}
		airport := new(Airport)
		airport.IATA = iata
		airport.Name = name
		airports[airport.IATA] = *airport
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func Get(iata string) Airport {
	return airports[iata]
}
