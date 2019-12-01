package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

// Record is struct for reading data into
type Record struct {
	Date                                     time.Time
	Open, High, Low, Close, Volume, AdjClose float64
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := parseCsv("table.csv")
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}

func parseCsv(filename string) []Record {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// read into a var
	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]Record, 0, len(rows))
	// loop through csv file
	for _, row := range rows {

		// parse the data into correct formats
		_date, _ := time.Parse("2006-01-02", row[0])
		_open, _ := strconv.ParseFloat(row[1], 64)
		_high, _ := strconv.ParseFloat(row[2], 64)
		_low, _ := strconv.ParseFloat(row[3], 64)
		_close, _ := strconv.ParseFloat(row[4], 64)
		_volume, _ := strconv.ParseFloat(row[5], 64)
		_adjClose, _ := strconv.ParseFloat(row[6], 64)

		data = append(data,

			Record{
				Date:     _date,
				Open:     _open,
				High:     _high,
				Low:      _low,
				Close:    _close,
				Volume:   _volume,
				AdjClose: _adjClose,
			},
		)
	}

	return data
}
