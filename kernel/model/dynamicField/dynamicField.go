package dynamicfield

type DynamicFieldValue struct {
	ID        int64       `json:"id" gorm:"column:id;"`
	FieldID   int         `json:"field_id" gorm:"column:field_id;"`
	ObjectID  int64       `json:"object_id" gorm:"column:object_id;"`
	ValueText interface{} `json:"value_text" gorm:"column:value_text;"`
	ValueDate interface{} `json:"value_date" gorm:"column:value_date;"`
	ValueInt  interface{} `json:"value_int" gorm:"column:value_int;"`
}
