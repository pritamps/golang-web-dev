package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// TWO STEPS -- Parse and Execute

	// STEP 1 - PARSE
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Creates a writer interface
	nf, err := os.Create("index.html")

	if err != nil {
		log.Println("Couldn't create file :( ")
	}
	defer nf.Close()

	// STEP 2 - EXECUTE
	// Executes the template TO the writer `nf`
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
