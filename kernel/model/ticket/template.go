package ticket

// 在 tag 中添加 omitempty 忽略空值
// 注意这里 hobby,omitempty 合起来是 json tag 值，中间用英文逗号分隔
type FieldData struct {
	Name                 string            `json:"name"`
	Default              string            `json:"default"`
	FieldType            string            `json:"type"`
	Label                string            `json:"label"`
	Placeholder          string            `json:"placeholder,omitempty"`
	Display              int               `json:"display"`
	Multiple             int               `json:"multiple"`
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
