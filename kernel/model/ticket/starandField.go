package ticket

type TicketPriority struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Color        string `json:"color" gorm:"column:color;"`
	Icon         string `json:"icon" gorm:"column:icon;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type TicketType struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Color        string `json:"color" gorm:"column:color;"`
	Icon         string `json:"icon" gorm:"column:icon;"`
	TNStart      string `json:"tnstart" gorm:"column:tnstart;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

/*
	ticket state
*/
type TicketState struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Color        string `json:"color" gorm:"column:color;"`
	StateTypeID  int    `json:"typeID" gorm:"column:type_id;"`
	Icon         string `json:"icon" gorm:"column:icon;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type TicketSource struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Color        string `json:"color" gorm:"column:color;"`
	Icon         string `json:"icon" gorm:"column:icon;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type TicketStateType struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type Service struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type SLA struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}

type SLAService struct {
	ID      int `json:"id" gorm:"column:id;"`
	SLA     int `json:"sla" gorm:"column:sla;"`
	Service int `json:"service" gorm:"column:service;"`
}

type WorkingCalender struct {
	WorkingHours    map[string][]string `yaml:"WorkingHours"`
	ExtraWorkingDay []map[string]string `yaml:"ExtraWorkingDay"`
	VacationDays    []map[string]string `yaml:"VacationDays"`
}

type WorkingTimeCalender struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	TimeZone     string `json:"time_zone" gorm:"column:time_zone;"`
	WeekDayStart string `json:"weekdaystart" gorm:"column:week_day_start;"`
	WorkingTime  string `json:"workingtime" gorm:"column:working_time;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createby" gorm:"column:create_by;"`
	CreateByName string `json:"create_by_name" gorm:"<-:false"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeby" gorm:"column:change_by;"`
	ChangeByName string `json:"change_by_name,omitempty" gorm:"<-:false"`
}

type SLACalender struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Weekday     string `db:"week_day_start"`
	WorkingTime string `db:"working_time"`
}

type District struct {
	ID           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	ValidID      int    `json:"validID" gorm:"column:valid_id;"`
	Comment      string `json:"comment" gorm:"column:comment;"`
	CreateTime   string `json:"createTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by;"`
	CreateByName string `gorm:"<-:false" json:"create_by_name"`
	ChangeTime   string `json:"changeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"changeBy" gorm:"column:change_by;"`
	ChangeByName string `gorm:"<-:false" json:"change_by_name,omitempty"`
}
