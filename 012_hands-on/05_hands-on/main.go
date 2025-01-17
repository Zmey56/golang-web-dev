package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

// Create a data structure to pass to a template which contains information about
// restaurant's menu including Breakfast, Lunch, and Dinner items
type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

// can hold an unlimited number of meals
type menu []meal

// Create a template with the name tpl.gohtml
func init() {
	tpl = template.Must(template.ParseFiles("012_hands-on/05_hands-on/tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name:    "Oatmeal",
					Descrip: "yum yum",
					Price:   4.95,
				},
				item{
					Name:    "Cheerios",
					Descrip: "American eating food traditional now",
					Price:   3.95,
				},
				item{
					Name:    "Juice Orange",
					Descrip: "Delicious drinking in throat squeezed fresh",
					Price:   2.95,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:    "Hamburger",
					Descrip: "Delicous good eating for you",
					Price:   6.95,
				},
				item{
					Name:    "Cheese Melted Sandwich",
					Descrip: "Make cheese bread melt grease hot",
					Price:   3.95,
				},
				item{
					Name:    "French Fries",
					Descrip: "French eat potatoe fingers",
					Price:   2.95,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Item: []item{
				item{
					Name:    "Pasta Bolognese",
					Descrip: "From Italy delicious eating",
					Price:   7.95,
				},
				item{
					Name:    "Steak",
					Descrip: "Dead cow grilled bloody",
					Price:   13.95,
				},
				item{
					Name:    "Bistro Potatoe",
					Descrip: "Bistro bar wood American bacon",
					Price:   6.95,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
