package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jempe/code_generator/generator"
	"github.com/jempe/code_generator/utils"
)

var output = flag.String("output", "", "path of output file")
var schema = flag.String("schema", "", "path of schema file")
var selectedStruct = flag.String("struct", "", "name of selected struct")
var dbType = flag.String("db", "", "type of db to generate code")
var overwrite = flag.Bool("overwrite", false, "overwrite files")

func main() {
	flag.Parse()

	if !(*dbType == "postgre" || *dbType == "boltdb") {
		fmt.Println("The DB Type", *dbType, "is not supported, the available options are: postgre, boltdb")
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("Output file is required")
		os.Exit(1)
	} else if utils.FileExists(*output) && !*overwrite {
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

	fmt.Println(outputCode)
}

func panicError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
