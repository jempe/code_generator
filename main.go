package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jempe/code_generator/generator"
	"github.com/jempe/code_generator/utils"
)

var output = flag.String("output", "", "path of output file")
var schema = flag.String("schema", "", "path of schema file")
var selectedStruct = flag.String("struct", "", "name of selected struct")
var overwrite = flag.Bool("overwrite", false, "overwrite files")

func main() {
	flag.Parse()

	if *output != "" && utils.FileExists(*output) && !*overwrite {
		fmt.Println("Output file", *output, "already exists, if you want to overwrite it use the -overwrite argument")
		os.Exit(1)
	}

	if !utils.FileExists(*schema) {
		fmt.Println("The schema file", *schema, "doesn't exist")
		os.Exit(1)
	}

	if *selectedStruct == "" {
		fmt.Println("The struct is required")
		os.Exit(1)
	}

	templateFiles := flag.Args()

	if len(templateFiles) == 0 {
		fmt.Println("You need at least one template file")
		os.Exit(1)
	}

	for _, templateFile := range templateFiles {
		if !utils.FileExists(templateFile) {
			fmt.Println("template:", templateFile, "doesn't exist")
			os.Exit(1)
		}
	}

	dbData, err := generator.ReadFile(*schema)
	panicError(err)

	outputCode := dbData.ProcessTemplates(*selectedStruct, templateFiles...)

	if *output == "" {
		fmt.Println(outputCode)
	} else {
		err = ioutil.WriteFile(*output, []byte(outputCode), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("code generated at ", *output)
		}
	}
}

func panicError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
