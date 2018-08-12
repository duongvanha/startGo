package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func init() {
	fmt.Println("first running")
}

type test struct {
	name string
}

func main() {
	name := "my name ha"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		{{$wisdom := "write code here"}}
		<h1>` + name + `</h1>
		<h2> {{$wisdom}} </h2>
		<ul>
		{{range .}}
    		<li>{{.}}</li>
    	{{end}}
		</ul>

		-----
		<ul>
		{{range $index, $element := .}}
    		<li>{{$index}} - {{ $element}} </li>
    	{{end}}
		</ul>
		-----

		</body>
		</html>
	`)

	nf, _ := os.Create("index.html")

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	tpl, _ := template.ParseFiles("index.html")

	demo, _ := os.Create("demo.html")

	data := map[string]interface{}{
		"a":      "b",
		"c":      "d",
		"e":      "f",
		"g":      "h",
		"array":  []string{"demo", "test"},
		"struct": test{"someName"},
	}

	tpl.Execute(demo, data)

}
