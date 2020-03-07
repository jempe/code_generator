# code_generator
Generate code to avoid repetitive tasks

##Struct Variables:

StructName		name
StructPluralName	name + "s"
TableName		table
TableItem		item
KeyType			type of field where key equals true
KeyName			name of field where key equals true
KeyFieldName		field_name or name (lowercase) of field where key equals true
StructVar		name (camelcase)
OpenBraces		{{
CloseBraces		}}

##Field Variables:

Name			name
FieldName		field_name or name (lowercase)
Key			true if field is tablet key
Type			type
Maxlength		maxlength
Minlength		minlength
DBExtras		db_extras
DBDefault		db_default
Reference		reference
ReferenceType		"" or type of the foreign table ID if type equals reference
ValidateFunction	validate_function or Valid + name + Default 
