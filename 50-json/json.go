package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func demonstrateJsonEncoding[T any](val T) {
	valB, err := json.Marshal(val)
	// valB is a []byte
	fmt.Printf("valB, err := json.Marshal(%#v)\n", val)
	fmt.Printf("valB: %#v, err: %#v\n", valB, err)
	fmt.Printf("string(valB): %s\n", string(valB))
	fmt.Println()
}

func main() {
	fmt.Printf("--- Encoding ---\n\n")
	demonstrateJsonEncoding(true)
	demonstrateJsonEncoding(1)
	demonstrateJsonEncoding(1234567890)
	demonstrateJsonEncoding(3.14)
	demonstrateJsonEncoding("gopher")
	demonstrateJsonEncoding([]string{"apple", "pear"})
	demonstrateJsonEncoding(map[string]int{"five": 5, "four": 4})

	fmt.Println("Encode struct, default field names, hidden privates:")
	type person struct {
		// fields starting with a capital letter (exported fields) will appear in the JSON
		Name    string
		Surname string
		// fields starting with a lowercase letter (hidden fields) will NOT appear in the JSON
		age int
	}
	// the names of the JSON fields default to the same names as the struct's fields
	johnDoe := person{Name: "John", Surname: "Doe", age: 35}
	demonstrateJsonEncoding(&johnDoe)

	fmt.Println("Encode struct, custom field names, hidden privates:")
	type dog struct {
		// to control the name of JSON fields, use tags next to the struct fields
		Name      string `json:"name"`
		Breed     string `json:"breed"`
		weight_kg float64
	}
	fido := dog{Name: "Fido", Breed: "Terrier", weight_kg: 5.2}
	demonstrateJsonEncoding(&fido)

	// Decoding

	fmt.Printf("--- Decoding ---\n\n")
	m1 := map[string]any{"float": 6.28, "strs": []string{"a", "b"}}
	fmt.Printf("m1: %#v\n", m1)
	m1B, _ := json.Marshal(m1)
	fmt.Println("m1B, _ := json.Marshal(m1)")
	fmt.Printf("m1B: %#v\n", m1B)
	fmt.Println()

	// Decode into a map with values of type any
	var m2 map[string]any
	if err := json.Unmarshal(m1B, &m2); err != nil {
		panic(err)
	}
	fmt.Println("json.Unmarshal(m1B, &m2)")
	fmt.Println()

	fmt.Printf("m2: %#v\n", m2)
	fmt.Println("m2[\"float\"]:", m2["float"])
	fmt.Printf("type-of m2[\"float\"]: %T\n", m2["float"])
	fmt.Println()

	fmt.Printf("m2[\"strs\"]: %#v\n", m2["strs"])
	// If we try to index into m2["strs"], we'll get:
	// invalid operation: cannot index m2["chars"] (map index expression of type any)

	// We have to cast the any into a []any
	strs := m2["strs"].([]any)
	// strs has type []any
	fmt.Println("strs := m2[\"strs\"].([]any)")
	fmt.Printf("strs: %#v\n", strs)
	fmt.Printf("strs[0]: %#v\n", strs[0])
	fmt.Printf("type-of strs[0]: %T\n", strs[0])
	fmt.Println()

	str0 := strs[0].(string)
	// str0 has type string
	fmt.Println("str0 := strs[0].(string)")
	fmt.Printf("str0: %#v\n", str0)
	fmt.Println()

	// Decode into a struct
	johnDoeB, _ := json.Marshal(johnDoe)
	var decodedJohnDoe person
	json.Unmarshal(johnDoeB, &decodedJohnDoe)
	fmt.Printf("johnDoe: %#v\n", johnDoe)
	fmt.Println("johnDoeB := json.Marshal(johnDoe)")
	fmt.Printf("string(johnDoeB): %s\n", string(johnDoeB))
	fmt.Println("json.Unmarshal(johnDoeB, &decodedJohnDoe)")
	fmt.Printf("decodedJohnDoe: %#v\n", decodedJohnDoe)
	fmt.Println()

	// Encode directly to an os.Writer like os.Stdout
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("enc := json.NewEncoder(os.stdout)")
	fmt.Printf("enc: %#v\n", enc)
	numstringToInt := map[string]int{"five": 5, "four": 4}
	fmt.Printf("enc.Encode(%#v):\n", numstringToInt)
	enc.Encode(numstringToInt)
}
