package admin

type TemplateData struct {
	ID              int                     `json:"id,omitempty" gorm:"column:id;"`
	Name            string                  `json:"name,omitempty" gorm:"column:name;"`
	Web             string                  `json:"web,omitempty" gorm:"column:web;"`
	Mobile          string                  `json:"mobile,omitempty" gorm:"column:mobile;"`
	ValidID         int                     `json:"valid_id" gorm:"column:valid_id;"`
	Describes       string                  `json:"describes" gorm:"column:describes;"`
	Icon            string                  `json:"icon" gorm:"column:icon;"`
	Color           string                  `json:"color" gorm:"column:color;"`
	Type            string                  `json:"type,omitempty" gorm:"column:type;"`
	DisplayType     string                  `json:"display_type" gorm:"column:display_type;"`
	FieldOrder      []string                `json:"fieldorder,omitempty"`
	FieldData       map[string]FieldData    `json:"fieldData,omitempty"`
	FilterCondition *map[string]interface{} `json:"fieldOrder,omitempty"`
	Roles           []int                   `json:"roles,omitempty"`
	Createtime      string                  `json:"create_time,omitempty" gorm:"column:create_time;"`
	Createby        int                     `json:"create_by,omitempty" gorm:"column:create_by;"`
	CreatebyName    string                  `json:"create_by_name,omitempty" gorm:"<-:false"`
	Changetime      string                  `json:"change_time,omitempty" gorm:"column:change_time;"`
	ChangebyName    string                  `json:"change_by_name,omitempty" gorm:"<-:false"`
	Changeby        int                     `json:"change_by,omitempty" gorm:"column:change_by;"`
}

type TemplateField struct {
	ID              int    `json:"id,omitempty" gorm:"column:id;"`
	TemplateID      int    `json:"template_id,omitempty" gorm:"column:template_id;"`
	FieldKey        string `json:"field_key,omitempty" gorm:"column:field_key;"`
	Display         int    `json:"display,omitempty" gorm:"display"`
	FieldLabel      string `json:"field_label,omitempty" gorm:"column:field_label;"`
	FieldType       string `json:"field_type,omitempty" gorm:"column:field_type;"`
	FieldOrder      int    `json:"field_order,omitempty" gorm:"column:field_order;"`
	FieldObject     string `json:"field_object,omitempty" gorm:"column:field_object;"`
	FieldWidth      int    `json:"field_width,omitempty" gorm:"column:field_width;"`
	FieldPreference string `json:"field_preference,omitempty" gorm:"column:field_preference;"`
}

/*
	template role list
*/
type TemplateRole struct {
	ID         int `json:"id,omitempty" gorm:"column:id;"`
	TemplateID int `json:"template_id,omitempty" gorm:"column:template_id;"`
	RoleID     int `json:"role_id,omitempty" gorm:"column:role_id;"`
}

type FieldData struct {
	Name                 string            `json:"name,omitempty"`
	Default              interface{}       `json:"default"`
	FieldType            string            `json:"type,omitempty"`
	Label                string            `json:"label,omitempty"`
	Placeholder          string            `json:"placeholder,omitempty"`
	Display              int               `json:"display,omitempty"`
	Multiple             int               `json:"multiple,omitempty"`
	Impacts              []string          `json:"impacts,omitempty"`
	DependsOn            []string          `json:"dependsOn,omitempty"`
	PromptCode           int               `json:"promptCode,omitempty"`
	PromptMessage        string            `json:"PromptMessage,omitempty"`
	AutoComplete         bool              `json:"autoComplete,omitempty"`
	Options              map[string]string `json:"options,omitempty"`
	OptionsType          string            `json:"optionsType,omitempty"`
	OptionsList          []string          `json:"optionsList,omitempty"`
	OptionsValueComments map[string]string `json:"optionsValueComments,omitempty"`
	HintMessage          string            `json:"hint,omitempty"`
	HintType             int               `json:"hintType,omitempty"`
	Width                int               `json:"width,omitempty"`
	RegexError           string            `json:"regexError,omitempty"`
	Regex                string            `json:"regex,omitempty"`
}

type TemplateIDList struct {
	TemplateID []int `json:"template"`
}
