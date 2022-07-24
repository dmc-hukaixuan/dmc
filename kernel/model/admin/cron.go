package admin

type SchedulerTask struct {
    ID                    int    `json:"id,omitempty" gorm:"column:id;"`
    Ident                 int    `json:"ident,omitempty" gorm:"column:ident;"`
    Name                  string `json:"name,omitempty" binding:"required" gorm:"column:name;"`
    TaskType              string `json:"task_type,omitempty" gorm:"column:task_type;"`
    TaskData              string `json:"task_data,omitempty" gorm:"column:task_data;"`
    Attempts              string `json:"attempts,omitempty" gorm:"column:attempts;"`
    LastExecutionTime     string `json:"config,omitempty" gorm:"<-:false"`
    CronTime              string `json:"cron_time,omitempty" gorm:"column:cron_time;"`
    LastWorkerStatus      string `json:"last_worker_status,omitempty" gorm:"column:last_worker_status;"`
    NextExecutionTime     string `json:"next_execution_time,omitempty" gorm:"column:next_execution_time;"`
    LastWorkerRunningTime string `json:"last_worker_running_time,omitempty" gorm:"column:last_worker_running_time;"`
    Description           string `json:"description,omitempty" gorm:"column:Description;"`
    ValidID               uint   `json:"validid,omitempty" gorm:"column:valid_id;"`
    CreateBy              int    `json:"createBy,omitempty" gorm:"column:create_by;"`
    CreateByName          string `json:"createByName,omitempty" gorm:"<-:false"`
    ChangeTime            string `json:"changeTime,omitempty" gorm:"column:change_time;"`
    ChangeBy              int    `json:"changeBy,omitempty" gorm:"column:change_by;"`
    ChangeByName          string `json:"changeByName,omitempty" gorm:"<-:false"`
}

type ReportSendConfig struct {
    Subject string   `json:"subject"`
    To      []string `json:"to"`
    Cc      []string `json:"cc"`
    Body    string   `json:"body"`
}
