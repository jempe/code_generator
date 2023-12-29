package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jempe/code_generator/generator"
	"github.com/jempe/code_generator/utils"
)

var output = flag.String("output", "", "path of output file")
var schema = flag.String("schema", "", "path of schema file")
var selectedTable = flag.String("table", "", "name of table")
var overwrite = flag.Bool("overwrite", false, "overwrite files")

func main() {
	flag.Parse()

	var dbData generator.DBData
	var err error

	if *output != "" && utils.FileExists(*output) && !*overwrite {
		fmt.Println("Output file", *output, "already exists, if you want to overwrite it use the -overwrite argument")
		os.Exit(1)
	}

	if !utils.FileExists(*schema) {
		fmt.Println("The schema file", *schema, "doesn't exist")
		os.Exit(1)
	} else {
		dbData, err = generator.ReadFile(*schema)
		panicError(err)
	}

	if *selectedTable == "" {
		var tables []string

		for _, table := range dbData.Tables {
			tables = append(tables, table.The_items_name)
		}

		fmt.Println("The table is required, available tables are:", strings.Join(tables, ", "))
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

	outputCode, err := dbData.ProcessTemplates(*selectedTable, templateFiles...)

	panicError(err)

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
