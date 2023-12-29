package generator

type DBData struct {
	Version       int     `json:"version"`
	GoModule      string  `json:"go_module"`
	Tables        []Table `json:"tables"`
	SelectedTable Table   `json:"selected_table"`
	OpenBraces    string  `json:"open_braces"`
	CloseBraces   string  `json:"close_braces"`
}

type Table struct {
	TheItemName    string  `json:"ItemName"`
	The_item_name  string  `json:"item_name"`
	TheItemsName   string  `json:"ItemsName"`
	The_items_name string  `json:"items_name"`
	TheitemName    string  `json:"itemName"`
	TheitemsName   string  `json:"itemsName"`
	HasEmbeddings  bool    `json:"has_embeddings"`
	Key            Field   `json:"key"`
	Fields         []Field `json:"fields"`
}

type Field struct {
	FieldName  string `json:"FieldName"`
	field_name string `json:"field_name"`
	goType     string `json:"goType"`
}
