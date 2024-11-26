# Code Generator
Generate code to avoid repetitive tasks

## Struct Variables:

**StructName**		name

**StructPluralName**	name + "s"

**TableName**		table

**TableItem**		item

**KeyType**			type of field where key equals true

**KeyName**			name of field where key equals true

**KeyFieldName**		field_name or name (lowercase) of field where key equals true

**StructVar**		name (camelcase)

**OpenBraces**		{{

**CloseBraces**		}}

## Field Variables:

**Name**			name

**FieldName**		field_name or name (lowercase)

**Key**			true if field is tablet key

**Type**			type

**Maxlength**		maxlength

**Minlength**		minlength

**DBExtras**		db_extras

**DBDefault**		db_default

**Reference**		reference

**ReferenceType**		"" or type of the foreign table ID if type equals reference

**ValidateFunction**	validate_function or Valid + name + Default 

## Usage

To generate code using this tool, you need to provide a schema file, specify the table, and include at least one template file. Here is the basic usage:

```sh
jempe_code_generator -schema=schema.json -table=users templates/user.tmpl
```

## Example
```sh
jempe_code_generator -schema=schema.json -table=users templates/user.tmpl
```

## Options

- `-output`: Path of the output file. If not provided, the generated code will be printed to the standard output.
- `-schema`: Path of the schema file. This file should contain the database schema.
- `-table`: Name of the table for which you want to generate code.
- `-overwrite`: If set to true, existing files will be overwritten.
- `-template_files`: List of template files to be used for code generation.

## Example with Output File
```sh
jempe_code_generator -schema=schema.json -table=users -output=generated_code.go templates/user.tmpl
```

This command will generate the code based on the provided schema and template files and save it to `generated_code.go`.

## Notes

- Ensure the schema file exists and is correctly formatted.
- Provide at least one template file to generate the code.
- If the output file already exists and you don't use the -overwrite flag, the tool will not overwrite the existing file.

