package ticket

import "time"

type TicketBaseData struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:api路径"`                 // 工单的 ID
	TicketNumber      string    `json:"tn" gorm:"column:tn; comment:api中文描述"`                             // 工单的单号
	Title             string    `json:"title" gorm:"comment: ticket title"`                               // 工单的标题
	Queue             string    `json:"queue" gorm:"-"`                                                   // 工单角色
	QueueID           int       `json:"queue_id" gorm:"comment:api组"`                                     // 工单角色 ID
	LockID            int       `json:"ticket_lock_id" gorm:"column:ticket_lock_id;comment:api组"`         // 工单是否锁定 id
	Lock              string    `json:"ticket_lock" gorm:"-"`                                             // 锁定的
	TypeID            int       `json:"type_id" gorm:"column:type_id;comment:api组"`                       // 类型 ID
	Type              string    `json:"type" gorm:"-"`                                                    // 工单类型
	ServiceID         int       `json:"service_id" gorm:"column:service_id;comment:api组"`                 // 工单服务 id
	Service           string    `json:"service" gorm:"-"`                                                 // 工单服务
	SLAID             int       `json:"sla_id" gorm:"default:null;comment:api组"`                          // 工单 SLAID
	SLA               string    `json:"sla" gorm:"-"`                                                     // 工单 SLA
	User              string    `json:"user" gorm:"-"`                                                    // 工单指定处理人 ID
	UserID            int       `json:"user_id" gorm:"comment:api组"`                                      // 工单指定处理人
	ResponsibleUserID int       `json:"responsible_user_id" gorm:"comment:api组"`                          // 工单负责人 ID
	ResponsibleUser   string    `json:"responsible_user" gorm:"-"`                                        // 工单负责人
	PriorityID        int       `json:"ticket_priority_id" gorm:"column:ticket_priority_id;comment:api组"` // 工单优先级 ID
	Priority          string    `json:"ticket_priority" gorm:"-"`                                         // 工单优先级
	StateID           int       `json:"ticket_state_id" gorm:"column:ticket_state_id;comment:api组"`       // 工单状态 ID
	State             string    `json:"ticket_state" gorm:"-"`                                            // 工单状态
	CustomerID        string    `json:"customer_id" gorm:"comment:api组"`                                  // 工单的客户
	CustomerUserID    string    `json:"customer_user_id" gorm:"comment:api组"`                             // 工单的客户用户
	CreateCustomer    string    `json:"create_customer" gorm:"comment:api组"`                              // 工单是否由用户创建，如果是用户创建，那么这里就是用户的登录名，如果不是用户创建这里则为空
	Timeout           int       `json:"timeout" gorm:"comment:api组"`                                      // 工单锁定超时时间
	UntilTime         int       `json:"until_time" gorm:"comment:api组"`                                   // 工单挂起时间，如果工单的状态是挂起，则这里的值为挂起的目标时间
	ISEscalationd     int       `json:"is_escalationd" gorm:"comment:api组"`                               // 工单是否升级过
	WorkingTimeAge    int       `json:"working_time_age" gorm:"comment:api组"`                             // 工单的时长
	ProcessID         string    `json:"process_id" gorm:"comment:api组"`                                   // 工单所属的流程
	ActivityID        string    `json:"activity_id" gorm:"comment:api组"`                                  // 工单的环节
	ArchiveFlag       int       `json:"archive_flag" gorm:"comment:api组"`                                 // 工单是否转存
	CreateTime        time.Time `json:"create_time" gorm:"autoCreateTime;"`                               // 工单的创建时间
	CreateBy          int       `json:"create_by" gorm:"comment:api组"`
	CreateByUserName  string    `json:"create_by_login" gorm:"comment:api组"` // 工单的创建人
	ChangeTime        time.Time `json:"change_time" gorm:"autoUpdateTime;"`  // 工单的修改时间
	ChangeBy          int       `json:"change_by" gorm:"comment:工单的修改人"`     // 工单的修改人
}

type TicketHistory struct {
	ID            int64  `json:"id"`
	HistoryTypeID int    `json:"history_type_id"`
	Object        string `json:"object"`
	TicketID      int64  `json:"ticket_id"`
	StartValue    string `json:"start_value"`
	LeaveValue    string `json:"leave_value"`
	OwnerID       int    `json:"owner_id"`
	QueueID       int    `json:"queue_id"`
	WorkingTime   int    `json:"working_time"`
	CalendarTime  int    `json:"calendar_time"`
	ArticleID     int64  `json:"article_id"`
	TemplateID    int    `json:"template_id"`
	TemplateName  string `json:"template_name"`
	Source        string `json:"source"`
	CreateBy      int    `json:"create_by"`
	CreateTime    string `json:"create_time"`
}

type TicketHistoryFieldData struct {
	TicketID        int64
	OwnerID         int
	QueueID         int
	ArticleID       int64
	TemplateID      int
	FieldOrder      []string
	TemplateName    string
	Source          string
	TicketFiledList map[string]interface{}
	UserID          int
}
