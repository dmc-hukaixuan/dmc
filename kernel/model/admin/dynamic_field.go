package admin

import (
	"time"
	//"gorm.io/gorm"
)

type DynamicField struct {
	ID            int        `json:"id,omitempty" gorm:"column:id;"`
	InternalField uint       `json:"internalField,omitempty" binding:"required" gorm:"column:internal_field;"`
	Name          string     `json:"name,omitempty" binding:"required" gorm:"column:name;"`
	Label         string     `json:"label,omitempty" gorm:"column:label;"`
	FieldType     string     `json:"field_type,omitempty" gorm:"column:field_type;"`
	ObjectType    string     `json:"object_type,omitempty" gorm:"column:object_type;"`
	Config        string     `json:"config,omitempty" gorm:"column:config;"`
	ValidID       uint       `json:"valid_id,omitempty" gorm:"column:valid_id;"`
	CreateTime    *time.Time `json:"createTime,omitempty" gorm:"column:create_time;autoCreateTime;"`
	CreateBy      int        `json:"createBy,omitempty" gorm:"column:create_by;"`
	ChangeTime    *time.Time `json:"changeTime,omitempty" gorm:"column:change_time;autoCreateTime;"`
	ChangeBy      int        `json:"changeBy,omitempty" gorm:"column:change_by;"`
}
