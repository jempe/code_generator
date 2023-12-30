package generator

type DBData struct {
	Version       int     `json:"version"`        //Database version for migrations
	GoModule      string  `json:"go_module"`      //Name of the go module for imports
	Tables        []Table `json:"tables"`         //List of tables
	SelectedTable Table   `json:"selected_table"` //Selected table
	OpenBraces    string  `json:"open_braces"`    //Open braces because they cannot be used in the template
	CloseBraces   string  `json:"close_braces"`   //Close braces because they cannot be used in the template
}

type Table struct {
	TheItemName    string  `json:"ItemName"`       //CamelCase name of the table
	TheItemsName   string  `json:"ItemsName"`      //CamelCase plural name of the table
	TheitemName    string  `json:"itemName"`       //camelCase name of the table
	TheitemsName   string  `json:"itemsName"`      //camelCase plural name of the table
	The_item_name  string  `json:"item_name"`      //snake_case name of the table
	The_items_name string  `json:"items_name"`     //snake_case plural name of the table
	HasEmbeddings  bool    `json:"has_embeddings"` //If the table has any embedding for semantic search
	Key            Field   `json:"key"`            //Primary key of the table
	Fields         []Field `json:"fields"`         //List of fields of the table except the primary key
	NaturalIndex   int     `json:"-"`              //Index of the table starting from 1
}

type Field struct {
	TheFieldName   string `json:"FieldName"`     //CamelCase name of the field
	The_field_name string `json:"field_name"`    //snake_case name of the field
	EnableFilter   bool   `json:"enable_filter"` //If the field can be used for filtering in the API endpoint
	GoType         string `json:"goType"`        //Go type of the field
	NaturalIndex   int    `json:"-"`             //Index of the field starting from 1
}
