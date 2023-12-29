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
	TheItemsName   string  `json:"ItemsName"`
	TheitemName    string  `json:"itemName"`
	TheitemsName   string  `json:"itemsName"`
	The_item_name  string  `json:"item_name"`
	The_items_name string  `json:"items_name"`
	HasEmbeddings  bool    `json:"has_embeddings"`
	Key            Field   `json:"key"`
	Fields         []Field `json:"fields"`
}

type Field struct {
	TheFieldName    string `json:"FieldName"`
	The_field_name  string `json:"field_name"`
	EnableFilter    bool   `json:"enable_filter"`
	FilterValidator string `json:"filter_validator"`
	GoType          string `json:"goType"`
}
