package admin

import (
	"time"

	"gorm.io/gorm"
)

// 在 tag 中添加 omitempty 忽略空值
// 注意这里 hobby,omitempty 合起来是 json tag 值，中间用英文逗号分隔
type ProcessType struct {
	ID         int        `json:"id,omitempty" gorm:"column:id;`
	Name       string     `json:"name,omitempty" binding:"required"`
	Valid      int        `json:"validID,omitempty" binding:"required" gorm:"column:valid_id;"`
	Windows    string     `json:"windows,omitempty" gorm:"column:windows;"`
	CreateTime time.Time  `json:"createTime,omitempty" gorm:"column:create_time;autoCreateTime;"`
	CreateBy   int        `json:"createBy,omitempty" gorm:"column:create_by;"`
	ChangeTime *time.Time `json:"changeTime,omitempty" gorm:"column:change_time;autoCreateTime;"`
	ChangeBy   int        `json:"changeBy,omitempty" gorm:"column:change_by;"`
}

func (pt *ProcessType) TableName() string {
	return "pm_process_type_c"
}
func (pt *ProcessType) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("change_time", time.Now())
	return nil
}

type Process struct {
	ID            int    `json:"ID" gorm:"column:id;"`
	Name          string `json:"Name" gorm:"column:name;"`
	Description   string `json:"Description" gorm:"column:description;"`
	EntityID      string `json:"EntityID" gorm:"column:entity_id;"`
	StateEntityID string `json:"StateEntityID" gorm:"column:state_entity_id;"`
	Layout        string `json:"processConfig" yaml:"Layout" gorm:"column:layout;"`
	Config        string `json:"Config" gorm:"column:config;"`
	ProcessType   int    `json:"ProcessType" gorm:"column:process_type;"`
	CreateTime    string `json:"CreateTime" gorm:"column:create_time;"`
	CreateBy      int    `json:"CreateBy" gorm:"column:create_by;"`
	ChangeTime    string `json:"ChangeTime" gorm:"column:change_time;"`
	ChangeBy      int    `json:"ChangeBy" gorm:"column:change_by;"`
}

func (pt *Process) TableName() string {
	return "pm_process_c"
}

// some opertion in process
//
type ProcessNode struct {
	ID        int    `json:"ID" gorm:"column:id;"`
	Name      string `json:"Name" gorm:"column:name;"`
	ProcessID string `json:"process_id" gorm:"column:process_id;"`
	NodeID    string `json:"nodeID" gorm:"column:node_id;"`
	//StateEntityID string `json:"StateEntityID" gorm:"column:state_entity_id;"`
	Config     string `json:"Config" gorm:"column:config;"`
	CreateTime string `json:"CreateTime" gorm:"column:create_time;"`
	CreateBy   int    `json:"CreateBy" gorm:"column:create_by;"`
	ChangeTime string `json:"ChangeTime" gorm:"column:change_time;"`
	ChangeBy   int    `json:"ChangeBy" gorm:"column:change_by;"`
}

//
//
//
type ProcessTransition struct {
	ID           int    `json:"ID" gorm:"column:id;"`
	Name         string `json:"Name" gorm:"column:name;"`
	ProcessID    string `json:"process_id" gorm:"column:process_id;"`
	TransitionID string `json:"nodeID" gorm:"column:transition_id;"`
	Config       string `json:"Config" gorm:"column:config;"`
	Layout       string `json:"layout" gorm:"column:layout;"`
	CreateTime   string `json:"CreateTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"CreateBy" gorm:"column:create_by;"`
	ChangeTime   string `json:"ChangeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"ChangeBy" gorm:"column:change_by;"`
}

// process transation action struct
// process transation do something
type ProcessTransitionAction struct {
	ID           int    `json:"ID" gorm:"column:id;"`
	Name         string `json:"Name" gorm:"column:name;"`
	ProcessID    string `json:"process_id" gorm:"column:process_id;"`
	TransitionID string `json:"nodeID" gorm:"column:transition_id;"`
	Config       string `json:"Config" gorm:"column:config;"`
	Layout       string `json:"layout" gorm:"column:layout;"`
	CreateTime   string `json:"CreateTime" gorm:"column:create_time;"`
	CreateBy     int    `json:"CreateBy" gorm:"column:create_by;"`
	ChangeTime   string `json:"ChangeTime" gorm:"column:change_time;"`
	ChangeBy     int    `json:"ChangeBy" gorm:"column:change_by;"`
}

type Transations struct {
	ID               int    `json:"id"`
	Condition        string `json:"condition"`
	ConditionName    string `json:"conditionName"`
	ConditionLinking string `json:"conditionLinking"`
	SourceP          string `json:"sourceP"`
	TargetP          string `json:"targetP"`
	TransitionID     string `json:"transitionID"`
	Count            string `json:"Count"`
}

// process node struct detail
type NodeDetail struct {
	ID           int      `json:"id"`
	TemplateList []string `json:"templateList"`
	Name         string   `json:"name"`
}

type Transition struct {
	// ConditionName string `json:"conditionName"`
	// Conditions    string `json:"conditions"`
	// ID            int    `json:"id"`
	// EntityID      string `json:"entityID"`
	Name     string `json:"name"`
	Config   string `json:"config"`
	ID       int    `json:"id"`
	EntityID string `json:"entityID"`
}

type TransitionAction struct {
	Name     string `json:"name"`
	Config   string `json:"config"`
	ID       int    `json:"id"`
	EntityID string `json:"entityID"`
}

type Nodeinfo struct {
	Name         string   `json:"name"`
	NodeID       string   `json:"nodeID"`
	ID           int      `json:"id"`
	TemplateList []string `json:"templateList"`
}

type ActivityNote struct {
	ID       string `json:"id"`
	Label    string `json:"Label"`
	NodeType string `json:"nodeType"`
	Style    string `json:"style"`
	Left     int    `json:"x"`
	Right    int    `json:"y"`
}

type ProcessTypeList struct {
	ID      int    `gorm:"column:id;"`
	Name    string `gorm:"column:name;"`
	ValidID string `gorm:"column:valid_id;"`
	Windows string `gorm:"column:windows;"`
	Count   int    `gorm:"column:count_process;"`
}

type Path struct {
	StartNode string                            `json:"startNode"`
	EndNode   string                            `json:"endNode"`
	Path      map[string]map[string]interface{} `json:"path"`
}

type PorcessLayout struct {
	FlowData    FlowData    `json:"flowData"`
	ProcessData ProcessData `json:"processData"`
}

type FlowData struct {
	Edges      []Edges      `json:"edges"`
	NodeLayout []NodeLayout `json:"nodes"`
}

// type ProcessData struct {
// 	Nodes
// 	Edges
// }
type EndPoint struct {
	X  int    `json:"x"`
	Y  int    `json:"y"`
	ID string `json:"id"`
}
type StarrPoint struct {
	X  int    `json:"x "`
	Y  int    `json:"y "`
	ID string `json:"id "`
}
type Style struct {
}
type Edges struct {
	EndPoint   EndPoint   `json:"endPoint "`
	ID         string     `json:"id"`
	Label      string     `json:"label "`
	Source     string     `json:"source "`
	StarrPoint StarrPoint `json:"starrPoint "`
	Style      Style      `json:"style "`
	Target     string     `json:"target "`
	Type       string     `json:"type "`
}

type ActivityTemplate struct {
	ID         int    `json:"id " gorm:"column:id;"`
	ProcessID  string `json:"processID" gorm:"column:process_id;"`
	NodeID     string `json:"nodeID" gorm:"column:node_id;"`
	TemplateID int    `json:"templateID" gorm:"column:ticket_template_id;"`
}
type Activity struct {
	Name         string `json:"Name"`
	TemplateList []int  `json:"templateList"`
}
type NodeLayout struct {
	ID    string `json:"id "`
	Label string `json:" label ,omitempty"`
	Style Style  `json:"style "`
	Type  string `json:"type "`
	X     string `json:"x "`
	Y     string `json:"y "`
}

type ProcessData struct {
	ID               int                         `json:"id"`
	ProcessID        string                      `json:"processID"`
	Name             string                      `json:"name"`
	Comments         string                      `json:"commnes"`
	ProcessState     string                      `json:"processState"`
	ProcessTypeID    int                         `json:"processTypeID"`
	ValidID          int                         `json:"valid_id"`
	Nodes            map[string]Activity         `json:"nodes"`
	Transition       map[string]Transition       `json:"transition"`
	TransitionAction map[string]TransitionAction `json:"transitionAction"`
}
