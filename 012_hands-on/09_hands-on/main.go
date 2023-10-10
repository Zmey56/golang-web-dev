package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
// Create struct for table in table.csv
type Price struct {
	Date     string `csv:"date"`
	Open     string `csv:"open"`
	High     string `csv:"high"`
	Close    string `csv:"close"`
	Low      string `csv:"low"`
	Volume   string `csv:"volume"`
	AdjClose string `csv:"adj_close"`
}

type PriceTwoFields struct {
	Date     string `csv:"date"`
	AdjClose string `csv:"adj_close"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("012_hands-on/09_hands-on/tpl.gohtml"))
}

func main() {

	http.HandleFunc("/", ParseCSVToWeb)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func ParseCSVToWeb(response http.ResponseWriter, r *http.Request) {

	lines, err := ReadCSV("012_hands-on/09_hands-on/table.csv")
	if err != nil {
		panic(err)
	}

	data := []PriceTwoFields{}

	for _, line := range lines {
		tmp := PriceTwoFields{
			Date:     line[0],
			AdjClose: line[6],
		}
		data = append(data, tmp)
	}

	//Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
	err = tpl.Execute(response, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
