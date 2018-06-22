package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type AdvancedData struct {
	Name    string
	KVPairs map[string]int
	Points  []Point
}

const templ = `
The circle has centered at X={{.X}}, Y={{.Y}}
`

const biggerTempl = `
This is a very advanced data structure called {{.Name}}.

It contains several important datapoint suchs as
{{range $k, $v := .KVPairs}}{{$k}}={{$v}}
{{end}}
and also secret Cartesian coordiates such as
{{range .Points}}X={{.X}},Y={{.Y}}
{{end}}
`

func main() {
	fmt.Println("hej")
	p := Point{1, 2}
	c := Circle{p, 10}
	ad := AdvancedData{"FBILevelData",
		map[string]int{"Password1": 123, "Password2": 321},
		[]Point{p, {3, 4}},
	}
	fmt.Println(p)
	fmt.Println(c)
	fmt.Printf("%v, %v, %v", c.X, c.Y, c.Radius)

	report, err := template.New("report").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	biggerReport, err := template.New("biggerReport").Parse(biggerTempl)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(os.Stdout, c)
	biggerReport.Execute(os.Stdout, ad)
}
