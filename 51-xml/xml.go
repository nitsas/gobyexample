package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	// The XML element's name will be "plant".
	XMLName xml.Name `xml:"plant"`
	// Id will live as an attribute on the root XML element.
	Id     int      `xml:"id,attr"`
	Name   string   `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("%#v", p)
}

func main() {
	coffee := Plant{
		Id:     27,
		Name:   "Coffee",
		Origin: []string{"Ethiopia", "Brazil"},
	}
	// fmt.Println() calls coffee.String() automatically
	fmt.Println("coffee:", &coffee)

	// The 2nd argument "" means how to indent the whole element.
	// The 3rd argument "  " means to indent sub-elements.
	coffeeXML, err := xml.MarshalIndent(&coffee, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(`coffeeXML := xml.MarshalIndent(coffee, "", "  ")`)
	fmt.Println(string(coffeeXML))
	fmt.Println()

	var coffeeParsed Plant
	xml.Unmarshal(coffeeXML, &coffeeParsed)
	fmt.Println(`xml.Unmarshal(coffeeXML, &coffeeParsed)`)
	fmt.Println("coffeeParsed:", &coffeeParsed)
	fmt.Println()

	tomato := Plant{
		Id:     28,
		Name:   "Tomato",
		Origin: []string{"Mexico", "California"},
	}
	fmt.Println("tomato:", &tomato)

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		// Tag `xml:"plants>plant"` means nest the elements of this array
		// under a "plants" parent element. The default would NOT nest at all.
		Plants []*Plant `xml:"plants>plant"`
	}
	nesting := &Nesting{
		Plants: []*Plant{&coffee, &tomato},
	}
	fmt.Printf("nesting: %#v\n", nesting)
	nestingXML, err := xml.MarshalIndent(nesting, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(`nestingXML := xml.MarshalIndent(nesting, "", "  ")`)
	fmt.Println(string(nestingXML))
}
