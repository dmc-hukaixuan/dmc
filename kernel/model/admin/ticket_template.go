package admin

type TemplateData struct {
	ID          int    `json:"id,omitempty" gorm:"column:id;"`
	Name        string `json:"name,omitempty" gorm:"column:name;"`
	Web         string `json:"web,omitempty" gorm:"column:web;"`
	Mobile      string `json:"mobile,omitempty" gorm:"column:mobile;"`
	Valid_id    string `json:"valid_id,omitempty" gorm:"column:valid_id;"`
	Describe    string `json:"describe,omitempty" gorm:"column:describe;"`
	Icon        string `json:"icon,omitempty" gorm:"column:icon;"`
	Color       string `json:"color,omitempty" gorm:"column:color;"`
	Type        string `json:"type,omitempty" gorm:"column:type;"`
	DisplayType string `json:"display_type,omitempty" gorm:"column:display_type;"`
	Createtime  string `json:"create_time,omitempty" gorm:"column:create_time;"`
	Createby    string `json:"create_by,omitempty" gorm:"column:create_by;"`
	Changetime  string `json:"change_time,omitempty" gorm:"column:change_time;"`
	Changeby    string `json:"change_by,omitempty" gorm:"column:change_by;"`
}

type TemplateField struct {
	ID              int    `json:"id,omitempty" gorm:"column:id;"`
	TemplateID      int    `json:"template_id,omitempty" gorm:"column:template_id;"`
	FieldKey        string `json:"field_key,omitempty" gorm:"column:field_key;"`
	FieldLabel      string `json:"field_label,omitempty" gorm:"column:field_label;"`
	FieldOrder      int    `json:"field_order,omitempty" gorm:"column:field_order;"`
	FieldObject     string `json:"field_object,omitempty" gorm:"column:field_object;"`
	FieldWidth      string `json:"field_width,omitempty" gorm:"column:field_width;"`
	FieldPreference string `json:"field_preference,omitempty" gorm:"column:field_preference;"`
}

type TemplateRole struct {
	ID         int `json:"id,omitempty" gorm:"column:id;"`
	TemplateID int `json:"template_id,omitempty" gorm:"column:template_id;"`
	RoleID     int `json:"role_id,omitempty" gorm:"column:role_id;"`
}
