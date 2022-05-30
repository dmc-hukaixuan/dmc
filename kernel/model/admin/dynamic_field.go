package admin

import (
	"time"
	//"gorm.io/gorm"
)

type DynamicField struct {
	ID            int                `json:"id,omitempty" gorm:"column:id;"`
	InternalField uint               `json:"internalField,omitempty" binding:"required" gorm:"column:internal_field;"`
	Name          string             `json:"name,omitempty" binding:"required" gorm:"column:name;"`
	Label         string             `json:"label,omitempty" gorm:"column:label;"`
	FieldType     string             `json:"field_type,omitempty" gorm:"column:field_type;"`
	ObjectType    string             `json:"object_type,omitempty" gorm:"column:object_type;"`
	Config        DynamicFieldConfig `json:"config,omitempty" gorm:"column:config;"`
	ValidID       uint               `json:"valid_id,omitempty" gorm:"column:valid_id;"`
	CreateTime    *time.Time         `json:"createTime,omitempty" gorm:"column:create_time;autoCreateTime;"`
	CreateBy      int                `json:"createBy,omitempty" gorm:"column:create_by;"`
	ChangeTime    *time.Time         `json:"changeTime,omitempty" gorm:"column:change_time;autoCreateTime;"`
	ChangeBy      int                `json:"changeBy,omitempty" gorm:"column:change_by;"`
}

type DynamicFieldConfig struct {
	DefaultValue     string            `mytag:"DefaultValue" json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Formula          string            `mytag:"Formula" json:"Formula,omitempty" yaml:"formula,omitempty"`
	HintType         int               `mytag:"HintType" json:"HintType,omitempty" `
	HintContent      string            `mytag:"HintContent" json:"HintContent,omitempty" `
	PossibleComments map[string]string `mytag:"PossibleComments" json:"PossibleComments" yaml:"PossibleComments,omitempty"`
	PossibleValues   map[string]string `mytag:"PossibleValues" json:"PossibleValues" yaml:"PossibleValues,omitempty"`
	Multiple         int               `mytag:"Multiple" json:"Multiple,omitempty" `
	Regex            string            `mytag:"Regex" json:"regex,omitempty" `
	Rows             string            `mytag:"Rows" json:"Rows,omitempty" `
	Columns          string            `mytag:"Columns" json:"columns,omitempty" `
	RegexHint        int               `mytag:"RegexHint" json:"RegexHint,omitempty" `
	TreeView         int               `mytag:"TreView" json:"TreView,omitempty" `
}
