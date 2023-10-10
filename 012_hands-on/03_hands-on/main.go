package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

// Create a data structure to pass to a template which contains information about California hotels including Name, Address, City, Zip, Region
type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

// can hold an unlimited number of hotels
type hotels []hotel

// Create a template with the name tpl.gohtml
func init() {
	tpl = template.Must(template.ParseFiles("012_hands-on/03_hands-on/tpl.gohtml"))
}

func main() {

	//Create a slice of hotels. Create at least 3 hotels. Include the information from step 2
	hotels := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
		hotel{
			Name:    "Hotel Denver",
			Address: "42 Sunset Saacashvili",
			City:    "Denver",
			Zip:     "85232",
			Region:  "northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		panic(err)
	}
}
