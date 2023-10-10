package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	langs := []string{
		"Go",
		"Python",
		"Ruby",
		"Javascript",
	}
	t1.Execute(os.Stdout, "hello")
	t1.Execute(os.Stdout, 51)
	t1.Execute(os.Stdout, langs)
	fmt.Println()

	t2 := template.Must(template.New("t2").Parse("Full name: {{.Name}} {{.Surname}}\n"))

	type Person struct {
		// The struct's attributes must be uppercase to be accessible from the template.
		Name, Surname string
	}
	janeStruct := Person{"Jane", "Doe"}
	t2.Execute(os.Stdout, janeStruct)

	// We can use lowercase keys for the map, but the template must call them as lowercase.
	johnMap := map[string]string{"Name": "John", "Surname": "Doe"}
	t2.Execute(os.Stdout, johnMap)
	fmt.Println()

	t3 := template.Must(template.New("t3").Parse("{{.}} is {{if . -}} truthy {{else -}} falsey {{end}}\n"))
	t3.Execute(os.Stdout, 0)
	t3.Execute(os.Stdout, -1)
	t3.Execute(os.Stdout, 5)
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, " ")
	t3.Execute(os.Stdout, ' ')
	t3.Execute(os.Stdout, "hello")
	fmt.Println()

	t4 := template.Must(template.New("t4").Parse("Range: {{range .}}{{.}} {{end}}\n"))
	t4.Execute(os.Stdout, langs)
}
