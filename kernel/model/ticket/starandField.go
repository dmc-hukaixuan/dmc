package ticket

type TicketPriroity struct {
	ID         int    `json:"id" gorm:"column:id;"`
	Name       string `json:"name" gorm:"column:name;"`
	ValidID    int    `json:"validID" gorm:"column:valid_id;"`
	Color      string `json:"color" gorm:"column:color;"`
	Icon       string `json:"icon" gorm:"column:icon;"`
	Comment    string `json:"comment" gorm:"column:comment;"`
	CreateTime string `json:"createTime" gorm:"column:create_time;"`
	CreateBy   int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy   int    `json:"changeBy" gorm:"column:change_by;"`
}

type TicketType struct {
	ID         int    `json:"id" gorm:"column:id;"`
	Name       string `json:"name" gorm:"column:name;"`
	ValidID    int    `json:"validID" gorm:"column:valid_id;"`
	Color      string `json:"color" gorm:"column:color;"`
	Icon       string `json:"icon" gorm:"column:icon;"`
	Comment    string `json:"comment" gorm:"column:comment;"`
	CreateTime string `json:"createTime" gorm:"column:create_time;"`
	CreateBy   int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy   int    `json:"changeBy" gorm:"column:change_by;"`
}

type TicketState struct {
	ID          int    `json:"id" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	ValidID     int    `json:"validID" gorm:"column:valid_id;"`
	Color       string `json:"color" gorm:"column:color;"`
	StateTypeID int    `json:"typeID" gorm:"column:type_id;"`
	Icon        string `json:"icon" gorm:"column:icon;"`
	Comment     string `json:"comment" gorm:"column:comment;"`
	CreateTime  string `json:"createTime" gorm:"column:create_time;"`
	CreateBy    int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime  string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy    int    `json:"changeBy" gorm:"column:change_by;"`
}

type Service struct {
	ID         int    `json:"id" gorm:"column:id;"`
	Name       string `json:"name" gorm:"column:name;"`
	ValidID    int    `json:"validID" gorm:"column:valid_id;"`
	Comment    string `json:"comment" gorm:"column:comment;"`
	CreateTime string `json:"createTime" gorm:"column:create_time;"`
	CreateBy   int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy   int    `json:"changeBy" gorm:"column:change_by;"`
}

type SLA struct {
	ID         int    `json:"id" gorm:"column:id;"`
	Name       string `json:"name" gorm:"column:name;"`
	ValidID    int    `json:"validID" gorm:"column:valid_id;"`
	Comment    string `json:"comment" gorm:"column:comment;"`
	CreateTime string `json:"createTime" gorm:"column:create_time;"`
	CreateBy   int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy   int    `json:"changeBy" gorm:"column:change_by;"`
}
