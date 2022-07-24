package report

type Stats struct {
    ID           int    `json:"id" gorm:"column:create_by;"`
    Name         string `json:"name" gorm:"column:create_by;"`
    ScaleType    string `json:"scaleType" gorm:"column:scale_type;"`
    Config       string `json:"config" gorm:"column:config;"`
    ValidID      string `json:"validId" gorm:"column:valid_id;"`
    Comments     string `json:"comments" gorm:"column:comments;"`
    CreateBy     int    `json:"createBy,omitempty" gorm:"column:create_by;"`
    CreateByName string `json:"createByName,omitempty" gorm:"<-:false"`
    ChangeTime   string `json:"changeTime,omitempty" gorm:"column:change_time;"`
    ChangeBy     int    `json:"changeBy,omitempty" gorm:"column:change_by;"`
    ChangeByName string `json:"changeByName,omitempty" gorm:"<-:false"`
}

type StatsDynamicConfig struct {
    UseAsXvalue      string                 `json:"useAsXvalue"`
    UseAsValueSeries string                 `json:"useAsValueSeries"`
    UseAsRestriction map[string]interface{} `json:"useAsRestriction"`
}
