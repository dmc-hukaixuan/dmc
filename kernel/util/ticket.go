package util

import (
    "database/sql/driver"
    "fmt"

    //pool "service-cool/kernel/system/dao"
    "dmc/global"
    workingtime "dmc/kernel/util/time"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type LocalTime struct {
    time.Time
}

type MoveAge struct {
    L1FisrtMoveTime        string
    L1WaitingTime          int
    L2FirstMoveL3Time      string
    L2FirstMoveL3Flag      int
    L2WaitingTime          int
    L2MoveL3WorkingTime    int
    L1FirstMoveWorkingTime int
    MoveCount              int
    SolutionOfferTime      string
    FadadaInvoke           string
    L2StayTimeFlag         int
    FirstRespsoneTime      string
    SolveTeam              string
    CloseTime              string
    PendingWorkingTime     int
    L2StayWorkingTime      int
    CalendarAge            string
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
    //格式化秒
    seconds := t.Unix()
    return []byte(strconv.FormatInt(seconds, 10)), nil
}
func (t LocalTime) Value() (driver.Value, error) {
    var zeroTime time.Time
    if t.Time.UnixNano() == zeroTime.UnixNano() {
        return nil, nil
    }
    return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
    value, ok := v.(time.Time)
    fmt.Println("value ", value)
    if ok {
        *t = LocalTime{Time: value}
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}

type TicketHistory struct {
    //Id            int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;"`
    Name          string    `json:"name" gorm:"type:varchar(200);column:name;not null;"`
    HistoryType   string    `json:"history_type" gorm:"type:varchar(200);"`
    HistoryTypeID int       `json:"history_type_id" gorm:"type:SMALLINT(6);"`
    TicketID      int64     `json:"ticket_id" gorm:"type:BIGINT;"`
    ArticletID    int64     `josn:"article_id" gorm:"column:article_id;;foreignKey:ID"`
    TypeID        int       `json:"type_id" gorm:"column:type_id;"`
    QueueID       int       `json:"queue_id" gorm:"column:queue_id;"`
    QueueName     string    `json:"queue_name" gorm:"column:queue_name;"`
    OwnerID       int       `json:"owner_id" gorm:"column:owner_id;"`
    PirorityID    int       `json:"priority_id" gorm:"column:priority_id;"`
    StateID       int       `json: "state_id" gorm:"column:state_id;"`
    TransitionID  string    `json:"transition_id"`
    DetailAction  string    `json:"detail_action"`
    Customer      string    `json:"customer"`
    Source        string    `json:"source"`
    CreateBy      int       `json:"create_by"`
    CreateUser    string    `json:"create_user"`
    CreateTime    time.Time `json:"create_time" gorm:"column:create_time;DEFAULT:NULL"`
    ChangeBy      int       `json:"change_by"`
}

/*
   方法说明：

       查询指定工单历史

   入参

       ticketid  工单 ID
       userid   查询人是谁

   返回结果

       TicketHistory 数组，该工单的一条条记录

*/
func TicketHistoryGet(ticketid int64, userid int, historytype []string, starttime, endtime string) []TicketHistory {
    // connect db object

    var lines []TicketHistory

    // sql string TicketHistoryGet ticket_history_2022_02
    SelctSQL := `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type, q.name AS queue_name,
                sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                sh.transition_id, sh.detail_action, sh.customer, sh.source
                FROM ticket_history sh LEFT JOIN  ticket_history_type ht ON ht.id = sh.history_type_id LEFT JOIN queue q ON q.id = sh.queue_id WHERE
                sh.ticket_id = ?
                ORDER BY sh.create_time`
    //fmt.Println("SelctSQL ", starttime, endtime)
    if len(historytype) > 0 && starttime == "" {
        Hisotry := strings.Join(historytype, ",")
        //fmt.Println("historytype ", starttime, endtime)
        // fmt.Println("ss", Hisotry)
        SelctSQL = `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type, q.name AS queue_name,
                    sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                    sh.transition_id, sh.detail_action, sh.customer, sh.source
                    FROM ticket_history sh LEFT JOIN  ticket_history_type ht ON ht.id = sh.history_type_id LEFT JOIN queue q ON q.id = sh.queue_id WHERE
                    sh.ticket_id = ? AND sh.history_type_id IN ( ` + Hisotry +
            ` ) ORDER BY sh.create_time`
    }

    if len(historytype) > 0 && starttime != "" && endtime != "" {
        Hisotry := strings.Join(historytype, ",")
        // fmt.Println("ss", Hisotry)
        SelctSQL = `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type, q.name AS queue_name,
                    sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                    sh.transition_id, sh.detail_action, sh.customer, sh.source
                    FROM ticket_history sh LEFT JOIN  ticket_history_type ht ON ht.id = sh.history_type_id LEFT JOIN queue q ON q.id = sh.queue_id WHERE
                    sh.ticket_id = ? AND sh.create_time > '` + starttime + `' AND sh.create_time < '` + endtime + `' AND sh.history_type_id IN ( ` + Hisotry +
            ` ) ORDER BY sh.create_time`
        //fmt.Println("starttime 11 ", starttime, endtime)
    }
    // ask database and fetch result
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&lines)
    // fmt.Println("lines ", lines)
    return lines
}

func TicketHistoryCustom(ticketid int64, historytype int, nameLike string) []TicketHistory {
    //db := pool.GetDB()
    var lines []TicketHistory
    var SelctSQL string

    if nameLike == "" {
        // fmt.Println("ss", Hisotry)
        SelctSQL = `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type,
                    sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                    sh.transition_id, sh.detail_action, sh.customer, sh.source, u.login AS create_user
                    FROM ticket_history sh LEFT JOIN  ticket_history_type ht ON ht.id = sh.history_type_id LEFT JOIN users u ON u.id = sh.create_by WHERE
                    sh.ticket_id = ? AND sh.history_type_id = ? ORDER BY sh.create_time`
        // ask database and fetch result
        global.GVA_DB_REPORT.Raw(SelctSQL, ticketid, historytype).Scan(&lines)
    } else {
        // fmt.Println("ss", Hisotry)
        SelctSQL = `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type,
                    sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                    sh.transition_id, sh.detail_action, sh.customer, sh.source, u.login AS create_user
                    FROM ticket_history sh LEFT JOIN ticket_history_type ht ON ht.id = sh.history_type_id LEFT JOIN users u ON u.id = sh.create_by  WHERE
                    sh.ticket_id = ? AND sh.history_type_id = ? AND sh.name LIKE '%` + nameLike + `%' ORDER BY sh.create_time`
        // ask database and fetch result

        global.GVA_DB_REPORT.Raw(SelctSQL, ticketid, historytype).Scan(&lines)
    }
    // fmt.Println("lines ", lines)
    return lines
}

/*
   方法说明：

       计算工单字段间的 working time

   入参

       ticketid  工单 ID

   返回结果

       在表中插入数据
*/
func TicketFieldUpdateWorkingTime(ticketid int64, calendarID int, roles string, starttime, endtime string) (rolestay1 map[string]int, moveage MoveAge) {

    type TicketEeveyUpdateWorkingTime struct {
        ID           int64     `json:"id" gorm:"TYPE:bigint(20);NOT NULL;PRIMARY_KEY;INDEX"`
        TicketID     int64     `json:"ticket_id" gorm:"TYPE: bigint(20);NOT NULL;INDEX"`
        Object       string    `json:"object" gorm:"TYPE: VARCHAR(200);"`
        ObjectValue  string    `json:"object_value" gorm:"TYPE: VARCHAR(200);"`
        StartTime    time.Time `json:"start_time" gorm:"TYPE:DATETIME;DEFAULT:NULL;"`
        LeaveTime    time.Time `json:"leave_time" gorm:"TYPE:DATETIME;DEFAULT:NULL;"`
        StartValue   string    `json:"start_value" gorm:"TYPE: VARCHAR(200);"`
        LeaveValue   string    `json:"leave_value" gorm:"TYPE: VARCHAR(200);"`
        WorkingAge   int       `json:"working_age" gorm:"TYPE: int(11);"`
        CalendarAge  float64   `json:"calendar_age" gorm:"TYPE: int(11);"`
        PengdingTime int       `json:"pengding_time" gorm:"TYPE: int(11);"`
        Comments     string    `json:"comments" gorm:"TYPE: VARCHAR(200);"`
    }

    var unFinshField = make(map[string]TicketEeveyUpdateWorkingTime)
    /*
       1, new ticket , 16: move, 26:setpendingtime, 27 : stateupdate, 28: TicketDynamicFieldUpdate
    */
    // var balance = [5]string{"1", "27", "26", "16", "28"}
    var balance = [5]string{"1", "16", "26", "27", "28"}
    // fetch ticket history
    historyLines := TicketHistoryGet(ticketid, 1, balance[:], starttime, endtime)
    if len(historyLines) == 0 {
        return
    }
    //historyLines := TicketHistoryGet(ticketid, 1, []int{})
    // get ticket create during state queue start time
    firstLine := historyLines[0]

    ticketStartTime := firstLine.CreateTime
    var ticketEndTime time.Time

    firstStateWorkingTime := TicketEeveyUpdateWorkingTime{
        TicketID:    ticketid,
        Object:      "StateUpdate",
        ObjectValue: strconv.Itoa(firstLine.StateID),
        StartTime:   firstLine.CreateTime,
        StartValue:  strconv.Itoa(firstLine.StateID),
    }
    firstMoveWorkingTime := TicketEeveyUpdateWorkingTime{
        TicketID:    ticketid,
        Object:      "Move",
        ObjectValue: firstLine.QueueName,
        StartTime:   firstLine.CreateTime,
        StartValue:  firstLine.QueueName,
    }

    moveage.SolveTeam = firstLine.QueueName

    // insert data to db ticket_field_update_workingtime_2
    //result := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Create(&firstStateWorkingTime)
    unFinshField["StateUpdate"] = firstStateWorkingTime
    // if result.Error != nil {
    // 	fmt.Printf("insert update field failed, err : %s\n", result.Error)
    // }
    //_ = global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Create(&firstMoveWorkingTime)
    unFinshField["Move"] = firstMoveWorkingTime

    // get dropdown dynamic field
    pengdingtimeflap := 0
    // DSS L2 : 25, user calendar: 日历三: id is 14
    // DSS L1 : 24, user calendar: 日历一: id is 12
    rolesCandendar := map[string]int{
        "DSS L2":                14,
        "25":                    14, // DSS L2
        "DSS L1":                12,
        "24":                    12, // DSS L1
        "POS L1":                12,
        "22":                    12, // DSS L1
        "POS L2":                13,
        "23":                    13, // DSS L1
        "POS L3":                13,
        "13":                    13, // DSS L1
        "DMS L1":                2,
        "9":                     2, // DSS L1
        "DMS L2":                2,
        "10":                    2, // DSS L1
        "DMS L3":                2,
        "11":                    2, // DSS L1
        "D-Flow L1":             15,
        "D-Flow L2":             16,
        "Datalake Support Team": calendarID,
        "Fadada support team":   14,
        "PDC BFS Support Team":  calendarID,
        "PAG C@P Support Team":  calendarID,
    }

    //  get dss l1 count
    rolestay := make(map[string]int)
LOOP:
    for _, value := range historyLines {
        // eg : %%misc%%10%%junk%%7
        arr := strings.Split(value.Name, "%%")
        var (
            object   string
            valueNew string
            valueOld string
        )
        if value.HistoryType == "StateUpdate" {
            object = "StateUpdate"

            if arr[1] != "" {
                valueOld = arr[1]
            }
            if arr[2] != "" {
                valueNew = arr[2]
            }

        } else if value.HistoryType == "Move" {
            object = "Move"

            moveage.MoveCount++
            if arr[3] != "" {
                valueOld = arr[3]
            }
            if arr[1] != "" {
                valueNew = arr[1]
            }
            // 根据团队的不同，工作日日历不同
            if rolesCandendar[valueOld] > 0 {
                calendarID = rolesCandendar[valueOld]
            }
            l3newmacth, _ := regexp.MatchString("L3", valueNew)
            fadadamacth, _ := regexp.MatchString("Fadada", valueNew)
            l2oldmacth, _ := regexp.MatchString("L2", valueOld)
            l2newmacth, _ := regexp.MatchString("L2", valueNew)
            l1newmacth, _ := regexp.MatchString("L1", valueNew)
            l1oldmacth, _ := regexp.MatchString("L1", valueOld)
            // L1 第一次转移时间
            if l2newmacth && l1oldmacth && moveage.L1FisrtMoveTime == "" {
                if moveage.L2FirstMoveL3Flag == 0 {
                    moveage.SolveTeam = valueNew
                }

                moveage.L1FisrtMoveTime = value.CreateTime.Format("2006-01-02 15:04:05")
            }
            if l2oldmacth && l3newmacth && moveage.L2FirstMoveL3Flag == 0 {
                if moveage.FadadaInvoke == "" {
                    moveage.SolveTeam = valueNew
                }

                moveage.L2FirstMoveL3Time = value.CreateTime.Format("2006-01-02 15:04:05")
                // moveage.L2WaitingTime = TicketPendingTimeGet(ticketid, calendarID, moveage.L2StayStartTime, value.CreateTime)
                moveage.L2FirstMoveL3Flag = 1
            }

            l3oldmacth, _ := regexp.MatchString("L3", valueOld)
            // 由L3处理，则以L3最后一次转回给L2的时间为准
            if l3oldmacth && l2newmacth {
                moveage.SolutionOfferTime = value.CreateTime.Format("2006-01-02 15:04:05")
            }

            // L2处理的，以L2转回给L1的时间为准
            if l1oldmacth && l2newmacth {
                moveage.SolutionOfferTime = value.CreateTime.Format("2006-01-02 15:04:05")
            }
            if fadadamacth {
                // moveage.SolveTeam = valueNew
                moveage.FadadaInvoke = "Y"
            }
            // 当转出 l2 并且转入不是 l1 时，作为在 l2 角色的停留的结束时间
            // 计算在 l2 的停留时长
            if l2oldmacth && !l1newmacth && moveage.L2StayTimeFlag == 0 {
                moveage.L2StayTimeFlag = 1
                //moveage.L2StayWorkingTime = workingtime.WorkingTime(moveage.L2StayStartTime, value.CreateTime, calendarID)
            }
            //	fmt.Println("valueNew:", valueNew, ", calendarID: ",calendarID, "queue id :", value.QueueID)
        } else if value.HistoryType == "TicketDynamicFieldUpdate" {
            l3newmacth, _ := regexp.MatchString("IncidentStartTime", value.Name)
            if l3newmacth {
                response := workingtime.WorkingTime(ticketStartTime, value.CreateTime, rolesCandendar[value.QueueName])
                moveage.FirstRespsoneTime = fmt.Sprintf("%.2f", float64(response)/float64(60))
            } else {
                continue LOOP
            }
            object = arr[2]
            if arr[6] != "" {
                valueOld = arr[6]
            }
            if arr[4] != "" {
                valueNew = arr[4]
            }
        } else if value.HistoryType == "SetPendingTime" {
            object = "SetPendingTime"
            // fmt.Println("SetPendingTime -^^^^^^^^^^^^^^^^^-", pengdingtimeflap, arr[1])
            if _, ok := unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)]; ok {

                valueNew = value.CreateTime.Format("2006-01-02 15:04:05")
                valueOld = unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)].StartTime.Format("2006-01-02 15:04:05")
                // fmt.Println("  start time 888888^^^^^^^^^^^", arr[1], unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)].LeaveTime)
                if arr[1] == "00-00-00 00:00" {
                    unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)] = TicketEeveyUpdateWorkingTime{
                        LeaveTime: value.CreateTime,
                        StartTime: unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)].StartTime,
                    }
                    // 这里需要判断上一次的挂起时间是不是结束了，如果没有结束则可能用户把挂起时间重置了，但是工单的开始挂起时间不能i需改
                } else if unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)].LeaveTime.Year() > 1 {
                    unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap+1)] = TicketEeveyUpdateWorkingTime{
                        StartTime: value.CreateTime,
                    }
                    pengdingtimeflap++
                }

            } else {
                moveage.SolutionOfferTime = value.CreateTime.Format("2006-01-02 15:04:05")
                valueNew = arr[1]
                valueOld = ""
                //fmt.Println("strconv.Itoa(pengdingtimeflap) ", strconv.Itoa(pengdingtimeflap))
                unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)] = TicketEeveyUpdateWorkingTime{
                    StartTime: value.CreateTime,
                }
            }
        } else {
            continue LOOP
        }

        /*
           如果事件改变的值已经存在，并且结束事件没有值，则表示已经发生过该字段的更新
               将字段更新的时间写入字段值的结束时间中。
           如果事件发生改变的字段值，已经存在，并且结束时间存在，
               将上次字段更新的值写入字段更新的开始值，将上次字段更新值的结束时间，作为本字段的
               开始时间，将字段发生更新的时间记为结束时间
           如果事件发生改变的值，不存在，
               将字段更新的时间，写入新的记录中，并记录字段更新的开始时间，开始值
        */
        start_time_ok := unFinshField[object].StartTime.Year()
        leave_time_ok := unFinshField[object].LeaveTime.Year()
        if _, ok := unFinshField[object]; !ok {
            // 如果事件改变的值已经存在，并且结束事件没有值，则
            firstWorkingTime := TicketEeveyUpdateWorkingTime{
                TicketID:    ticketid,
                Object:      object,
                ObjectValue: valueNew,
                StartTime:   value.CreateTime,
                StartValue:  valueNew,
            }
            //fmt.Println("unFinshField[object].StartTime---------- ", object, value.CreateTime, valueOld, valueNew )
            // resultok := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Create(&firstWorkingTime)
            // if resultok.Error != nil {
            // 	fmt.Println("add working error ", resultok.Error)
            // }
            unFinshField[object] = firstWorkingTime
        } else if _, ok := unFinshField[object]; ok && start_time_ok != 1 && leave_time_ok != 1 {

            // calculate working time
            workingAge := workingtime.WorkingTime(unFinshField[object].LeaveTime, value.CreateTime, calendarID)
            // fmt.Println("rtTime  ",  object, unFinshField[object].LeaveTime, value.CreateTime, valueOld ,valueNew, workingAge, calendarID ,value.QueueID)
            var pengdingtime int
            var pengdingTimeStartTime time.Time
            // check if the ticket is pending start during field update
            if _, ok := unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)]; ok && object != "SetPendingTime" {
                for i := 0; i <= pengdingtimeflap; i++ {
                    //	fmt.Println("ss ", unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.After(unFinshField[object].LeaveTime), unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Before(value.CreateTime), value.CreateTime.Sub(unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime))
                    //	fmt.Println("start", object, unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, unFinshField[object].LeaveTime, value.CreateTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime)
                    // 挂起时间
                    // fmt.Println("++",unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, unFinshField[object].LeaveTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime, value.CreateTime)
                    // 在该事件发生之前工单已经挂起，并且工单挂起事件还未结束，挂起时间： 该事件发生事件 至 现在的时间
                    if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.Before(unFinshField[object].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() <= 1 && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        pengdingtime = workingtime.WorkingTime(unFinshField[object].LeaveTime, value.CreateTime, calendarID) + pengdingtime
                        // 该事件发生之前工单已经挂起，并且工单挂起的事件已经结束
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.Before(unFinshField[object].LeaveTime) && value.CreateTime.Before(unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        //} else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.Before(unFinshField[object].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Before(value.CreateTime) {
                        pengdingtime = workingtime.WorkingTime(unFinshField[object].LeaveTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime, calendarID) + pengdingtime
                        // 该事件发生得时候工单还未挂起，并且在该事件结束前，工单已经结束挂起，则挂起时间: 工单挂起时间 至 工单挂起结束时间
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.After(unFinshField[object].LeaveTime) && (unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Before(value.CreateTime) || unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime == value.CreateTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() != 1 {
                        //fmt.Println("++++++++++++++++++++++++++++")
                        pengdingtime = workingtime.WorkingTime(unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime, calendarID) + pengdingtime
                        // 该事件发生得时候工单还未挂起，并且工单挂起事件没有结束，计算挂起事件：工单挂起时间 至 该事件创建时间
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.After(unFinshField[object].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() >= 1 && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        pengdingtime = workingtime.WorkingTime(unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, value.CreateTime, calendarID) + pengdingtime
                    }
                    pengdingTimeStartTime = unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime
                }
            }

            if pengdingTimeStartTime.Year() > 1 {
                //workingAge = workingtime.WorkingTime(unFinshField[object].LeaveTime, pengdingTimeStartTime, calendarID)
                //fmt.Println("pengd ----", unFinshField[object].LeaveTime, pengdingTimeStartTime)
            }

            //fmt.Println("pengd ----", object,valueOld, valueNew,pengdingTimeStartTime, workingAge,pengdingtime, calendarID )
            firstWorkingTime := TicketEeveyUpdateWorkingTime{
                TicketID:     ticketid,
                Object:       object,
                ObjectValue:  valueNew,
                StartTime:    unFinshField[object].LeaveTime,
                LeaveTime:    value.CreateTime,
                StartValue:   valueOld,
                LeaveValue:   valueNew,
                WorkingAge:   workingAge - pengdingtime,
                CalendarAge:  value.CreateTime.Sub(unFinshField[object].StartTime).Seconds(),
                PengdingTime: pengdingtime,
                Comments:     "",
            }

            if value.HistoryType == "Move" {
                //	fmt.Println("--------------", valueNew, valueOld, workingAge, pengdingtime, unFinshField[object].LeaveTime, value.CreateTime)
                if moveage.MoveCount == 1 {
                    moveage.L1FirstMoveWorkingTime = workingAge - pengdingtime
                    moveage.L1WaitingTime = pengdingtime
                    moveage.MoveCount++
                }

                if moveage.L2FirstMoveL3Flag == 1 {
                    moveage.L2MoveL3WorkingTime = workingAge - pengdingtime
                    moveage.L2WaitingTime = pengdingtime
                    moveage.L2StayWorkingTime = rolestay[valueOld] + workingAge - pengdingtime
                    moveage.L2FirstMoveL3Flag++
                    // fmt.Println("L2FirstMoveL3Flag ", )
                }
                rolestay[valueOld] = rolestay[valueOld] + workingAge - pengdingtime
                //fmt.Println("valueNew Move ---------:",valueOld, valueNew, workingAge, unFinshField[object].LeaveTime, value.CreateTime, rolestay[valueOld], calendarID)
            }
            // insert data to db
            // resultok := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Create(&firstWorkingTime)
            // if resultok.Error != nil {
            // 	fmt.Println("update working error ", resultok.Error)
            // }
            unFinshField[object] = firstWorkingTime
        } else if _, ok := unFinshField[object]; ok && start_time_ok != 1 && leave_time_ok <= 1 {
            // calculate working time

            workingAge := workingtime.WorkingTime(unFinshField[object].StartTime, value.CreateTime, calendarID)

            var pengdingtime int
            //var pengdingTimeStartTime time.Time
            // check if the ticket is pending start during field update
            if _, ok := unFinshField["SetPendingTime"+strconv.Itoa(pengdingtimeflap)]; ok {
                for i := 0; i <= pengdingtimeflap; i++ {
                    // && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime
                    // 在该事件发生之前工单已经挂起，并且工单挂起事件还未结束，挂起时间： 该事件发生事件 至 现在的时间
                    if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.Before(unFinshField[object].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() <= 1 && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        pengdingtime = workingtime.WorkingTime(unFinshField[object].LeaveTime, value.CreateTime, calendarID) + pengdingtime
                        // 该事件发生之前工单已经挂起，并且工单挂起的事件已经结束
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.Before(unFinshField[object].LeaveTime) && value.CreateTime.Before(unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        pengdingtime = workingtime.WorkingTime(unFinshField[object].LeaveTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime, calendarID) + pengdingtime
                        // 该事件发生得时候工单还未挂起，并且在该事件结束前，工单已经结束挂起，则挂起时间: 工单挂起时间 至 工单挂起结束时间
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.After(unFinshField[object].LeaveTime) && (unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Before(value.CreateTime) || unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime == value.CreateTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() != 1 {
                        pengdingtime = workingtime.WorkingTime(unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime, calendarID) + pengdingtime
                        // 该事件发生得时候工单还未挂起，并且工单挂起事件没有结束，计算挂起事件：工单挂起时间 至 该事件创建时间
                    } else if unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime.After(unFinshField[object].LeaveTime) && unFinshField["SetPendingTime"+strconv.Itoa(i)].LeaveTime.Year() <= 1 && unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime != value.CreateTime {
                        pengdingtime = workingtime.WorkingTime(unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime, value.CreateTime, calendarID) + pengdingtime
                    }
                    // pengdingTimeStartTime = unFinshField["SetPendingTime"+strconv.Itoa(i)].StartTime
                    //fmt.Println("pengdingTimeStartTime 12 ", pengdingtime,  pengdingTimeStartTime)
                }
            }

            // update already exists data
            firstWorkingTime := TicketEeveyUpdateWorkingTime{
                StartTime:    unFinshField[object].StartTime,
                LeaveTime:    value.CreateTime,
                StartValue:   valueOld,
                LeaveValue:   valueNew,
                WorkingAge:   workingAge - pengdingtime,
                CalendarAge:  value.CreateTime.Sub(unFinshField[object].StartTime).Seconds(),
                PengdingTime: pengdingtime,
                Comments:     "",
            }
            if value.HistoryType == "Move" {
                //fmt.Println("--------------", valueNew, valueOld, workingAge, pengdingtime, unFinshField[object].LeaveTime, value.CreateTime)
                if moveage.MoveCount == 1 {
                    moveage.L1FirstMoveWorkingTime = workingAge - pengdingtime
                    moveage.L1WaitingTime = pengdingtime
                    moveage.MoveCount++
                }
                if moveage.L2FirstMoveL3Flag == 1 {
                    moveage.L2MoveL3WorkingTime = workingAge - pengdingtime
                    moveage.L2WaitingTime = pengdingtime
                    moveage.L2StayWorkingTime = rolestay[valueOld] + workingAge - pengdingtime
                    moveage.L2FirstMoveL3Flag++
                    // fmt.Println("L2FirstMoveL3Flag ", )
                }

                //fmt.Println("valueOld", valueOld, valueNew, workingAge, pengdingtime)
                rolestay[valueOld] = rolestay[valueOld] + workingAge - pengdingtime
                //fmt.Println("valueNew 22:", workingAge - pengdingtime,  value.QueueID, rolestay[valueOld])
            }
            // update database
            // resultok := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Where("id = ?", unFinshField[object].ID).Updates(&firstWorkingTime)
            // if resultok.Error != nil {
            // 	fmt.Println("update working error ", resultok.Error)
            // }
            unFinshField[object] = firstWorkingTime
        }

        // 如果工单关闭，则计算工单关闭时工单所在的队列，的时长
        // 15 ：首次电话解决 , 34 ： 用户无回应
        // 2，closed successful
        if value.HistoryType == "StateUpdate" && (value.StateID == 15 || value.StateID == 34 || value.StateID == 2 || value.StateID == 2) {
            moveage.CloseTime = value.CreateTime.Format("2006-01-02 15:04:05")
            lastCalander := rolesCandendar[unFinshField["Move"].LeaveValue]
            //fmt.Println(" :lastCalander:", lastCalander, unFinshField["Move"].StartValue, rolesCandendar[unFinshField["Move"].LeaveValue])
            if lastCalander == 0 {
                lastCalander = rolesCandendar[unFinshField["Move"].StartValue]
                if lastCalander == 0 {
                    lastCalander = calendarID
                }
            }

            var pengdingtime int
            var workingAge int
            if unFinshField["Move"].LeaveTime.Year() == 1 {
                pengdingtime = TicketPendingTimeGet(ticketid, lastCalander, unFinshField["Move"].StartTime, value.CreateTime)
                workingAge = workingtime.WorkingTime(unFinshField["Move"].StartTime, value.CreateTime, lastCalander) - pengdingtime
            } else {

                pengdingtime = TicketPendingTimeGet(ticketid, lastCalander, unFinshField["Move"].LeaveTime, value.CreateTime)

                workingAge = workingtime.WorkingTime(unFinshField["Move"].LeaveTime, value.CreateTime, lastCalander) - pengdingtime
            }
            if moveage.MoveCount == 0 {
                moveage.L1WaitingTime = pengdingtime
            }
            // calculate working time
            if unFinshField["Move"].LeaveTime.Year() == 1 {
                firstWorkingTime := TicketEeveyUpdateWorkingTime{
                    StartTime:    unFinshField["Move"].StartTime,
                    LeaveTime:    value.CreateTime,
                    StartValue:   unFinshField["Move"].StartValue,
                    LeaveValue:   unFinshField["Move"].StartValue,
                    WorkingAge:   workingAge,
                    CalendarAge:  value.CreateTime.Sub(unFinshField["Move"].StartTime).Seconds(),
                    PengdingTime: pengdingtime,
                    Comments:     "",
                }

                // update database
                // resultok := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Where("id = ?", unFinshField["Move"].ID).Updates(&firstWorkingTime)
                // if resultok.Error != nil {
                // 	fmt.Println("update working error ", resultok.Error)
                // }
                unFinshField["Move"] = firstWorkingTime
                //rolestay["stayTime"] = workingAge
            }
            // fmt.Println("pengdingtime ",pengdingtime)
            //fmt.Println("valueNew--", object,unFinshField["Move"].LeaveValue,valueNew,  unFinshField["Move"].LeaveTime, workingAge - pengdingtime, pengdingTimeStartTime)
            // update already exists data
            firstWorkingTime := TicketEeveyUpdateWorkingTime{
                TicketID:     ticketid,
                Object:       object,
                ObjectValue:  valueNew,
                StartTime:    unFinshField["Move"].LeaveTime,
                LeaveTime:    value.CreateTime,
                StartValue:   unFinshField["Move"].LeaveValue,
                LeaveValue:   valueNew,
                WorkingAge:   workingAge,
                CalendarAge:  value.CreateTime.Sub(unFinshField["Move"].LeaveTime).Seconds(),
                PengdingTime: pengdingtime,
                Comments:     "",
            }

            rolestay[unFinshField["Move"].LeaveValue] = rolestay[unFinshField["Move"].LeaveValue] + workingAge
            //fmt.Println("valueNew:", workingAge - pengdingtime,  value.QueueID, rolestay[valueOld])
            // update database
            // resultok := global.GVA_DB_REPORT.Table("ticket_field_update_workingtime").Create(&firstWorkingTime)
            // if resultok.Error != nil {
            // 	fmt.Println("update working error ", resultok.Error)
            // }
            unFinshField[object] = firstWorkingTime
        }
        ticketEndTime = value.CreateTime
    }

    moveage.CalendarAge = fmt.Sprintf("%.2f", ticketEndTime.Sub(ticketStartTime).Minutes())
    // fmt.Println("ticketEndTime ", ticketStartTime, ticketEndTime,moveage.CalendarAge)
    var lines []TicketHistory
    // 获取挂起时间
    SelctSQL := `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by,
                sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                q.name AS queue_name, sh.transition_id, sh.detail_action, sh.customer, sh.source
                FROM ticket_history sh left join ticket_history_type tht ON tht.id = sh.history_type_id
                LEFT JOIN queue q ON q.id = sh.queue_id
                WHERE sh.ticket_id = ? AND tht.id = sh.history_type_id AND sh.history_type_id = 26
                ORDER BY sh.create_time`
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&lines)
    var (
        pendingEndTime     time.Time
        pendingStartTime   time.Time
        pengdingWorkingAge int
    )
    for _, v := range lines {
        // eg : %%misc%%10%%junk%%7
        arr := strings.Split(v.Name, "%%")
        if arr[1] == "00-00-00 00:00" {
            pendingEndTime = v.CreateTime
            cal := rolesCandendar[v.QueueName]
            if cal == 0 {
                cal = 12
            }
            pengdingWorkingAge = workingtime.WorkingTime(pendingStartTime, pendingEndTime, cal) + pengdingWorkingAge
            pendingStartTime = v.CreateTime
        } else if pendingStartTime.Year() == 1 {
            pendingStartTime = v.CreateTime
        }
    }

    moveage.PendingWorkingTime = pengdingWorkingAge
    return rolestay, moveage

    //fmt.Println("rolestay ", rolestay )
}

func TicketFirstRespsoneTime(ticketid int64, createtime time.Time) string {
    rolesCandendar := map[string]int{
        "DSS L2":                14,
        "25":                    14, // DSS L2
        "DSS L1":                12,
        "24":                    12, // DSS L1
        "POS L1":                12,
        "22":                    12, // DSS L1
        "POS L2":                13,
        "23":                    13, // DSS L1
        "POS L3":                13,
        "13":                    13, // DSS L1
        "DMS L1":                2,
        "9":                     2, // DSS L1
        "DMS L2":                2,
        "10":                    2, // DSS L1
        "DMS L3":                2,
        "11":                    2, // DSS L1
        "D-Flow L1":             15,
        "D-Flow L2":             16,
        "Datalake Support Team": 12,
        "Fadada support team":   14,
        "PDC BFS Support Team":  12,
        "PAG C@P Support Team":  12,
    }
    var lines []TicketHistory
    SelctSQL := `SELECT sh.name as name, sh.create_time as create_time, q.name AS queue_name
                FROM ticket_history sh LEFT JOIN queue q ON q.id = sh.queue_id WHERE sh.ticket_id = ? AND sh.name LIKE '%IncidentStartTime%' `
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&lines)

    response := workingtime.WorkingTime(createtime, lines[0].CreateTime, rolesCandendar[lines[0].QueueName])
    responsetime := fmt.Sprintf("%.2f", float64(response)/float64(60))
    return responsetime
}

/*
   获取某段时间内的所有挂起事件，及挂起的 working time
*/
func TicketPendingTimeGet(ticketid int64, calendarID int, startTime, endTime time.Time) int {
    //fmt.Println(" TicketPendingTimeGet  ----", startTime, endTime)
    var lines []TicketHistory
    // db := pool.GetDB()

    endUnix := endTime.Unix() + 60
    timeEnd := time.Unix(endUnix, 0)

    // sql string TicketHistoryGet ticket_history_2022_02
    SelctSQL := `SELECT sh.name as name, sh.article_id as article_id, sh.create_time as create_time, sh.create_by, ht.name as history_type,
                sh.queue_id as queue_id, sh.owner_id as owner_id, sh.priority_id as priority_id, sh.state_id as state_id, sh.history_type_id, sh.type_id,
                sh.transition_id, sh.detail_action, sh.customer, sh.source
                FROM ticket_history sh, ticket_history_type ht WHERE
                sh.ticket_id = ? AND ht.id = sh.history_type_id AND sh.history_type_id = 26 AND sh.create_time > ? AND sh.create_time <= ?
                ORDER BY sh.create_time`
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid, startTime, timeEnd).Scan(&lines)
    var (
        pendingEndTime     time.Time
        pendingStartTime   time.Time
        pengdingWorkingAge int
    )
    for _, v := range lines {

        // eg : %%misc%%10%%junk%%7
        arr := strings.Split(v.Name, "%%")
        fmt.Println("arr ", arr)
        if arr[1] == "00-00-00 00:00" {
            pendingEndTime = v.CreateTime
            pengdingWorkingAge = workingtime.WorkingTime(pendingStartTime, pendingEndTime, calendarID) + pengdingWorkingAge
            pendingStartTime = v.CreateTime
        } else if pendingStartTime.Year() == 1 {
            pendingStartTime = v.CreateTime
        }
    }

    return pengdingWorkingAge
}

/*
   获取g单所有动态字段的值
*/
func TicketynamicFieldValueGet(ticketid int64) map[string]string {
    type DynamicFieldValue struct {
        Label     string `json:"label"`
        Name      string `json:"name"`
        ValueText string `json:"value_text"`
    }
    var dfv []DynamicFieldValue
    //db := pool.GetDB()
    // sql string TicketHistoryGet ticket_history_2022_02
    SelctSQL := `SELECT dfv.id, dfv.field_id, df.name as name, df.label as label, max(case when dfv.value_text IS NOT NULL then dfv.value_text ELSE dfv.value_date END) as value_text
                FROM dynamic_field_value dfv LEFT JOIN dynamic_field df ON df.id = dfv.field_id WHERE dfv.object_id = ?
                GROUP BY dfv.id`
    // ask database and fetch result
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&dfv)
    dynamic_field := make(map[string]string)
    for _, v := range dfv {
        dynamic_field[v.Name] = v.ValueText
    }
    return dynamic_field
}

type Article struct {
    Body       string `json :"body"`
    AFrom      string `json :"afrom"`
    CreateTime string `json :"create_time"`
    SenderType string `json :"sender_type"`
}

/*
   获取工单关联的信件
*/
func TicketArticleList(ticketid int64) (articleTree string) {
    //db := pool.GetDB()
    var article []Article
    // sql string TicketHistoryGet ticket_history_2022_02
    SelctSQL := `SELECT adm.a_body AS body, ast.name AS sender_type, adm.a_from AS afrom, adm.create_time as create_time FROM article_data_mime adm
                LEFT JOIN article act ON act.id = adm.article_id
                LEFT JOIN article_sender_type ast ON ast.id = act.article_sender_type_id
                WHERE act.ticket_id = ? ORDER BY adm.create_time`
    // ask database and fetch result
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&article)
    for _, v := range article {
        articleTree += "\n-->" + `||` + v.SenderType + `||` + v.AFrom + `||` + v.CreateTime + `||<--------------\n` + v.Body
    }
    return articleTree
}

/*
   获取工单链接的工单及单号
*/
func TicketLinkList(ticketid int64) (linktn string, linkCount int) {
    type linklist struct {
        Tn string `json :"tn"`
    }
    var tnlist []linklist

    // sql string TicketHistoryGet ticket_history_2022_02
    SelctSQL := `SELECT t.tn as tn FROM link_relation lr LEFT JOIN ticket t ON t.id = lr.target_key WHERE lr.source_key = ?`
    // ask database and fetch result
    global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&tnlist)

    for _, v := range tnlist {
        linktn += "," + v.Tn
        linkCount++
    }
    return linktn, linkCount
}

// /*
// 	创建工单函数
// */
// func TicketCreate(TicketBaseData m_ticket.TicketBaseData) (err error, TicketID int64) {
// 	// get db object
// 	// db := pool.GetDB()
// 	numberbuild := number.TicketNumber()
// 	// check ticket number if not null, generate ticket number
// 	if TicketBaseData.TicketNumber == "" {
// 		TicketBaseData.TicketNumber = numberbuild.TicketNumberBuild()
// 	}
// 	fmt.Println("TicketBaseData ", TicketBaseData)
// 	if !errors.Is(global.GVA_DB_REPORT.Table("sc_ticket").Create(&TicketBaseData).Error, gorm.ErrRecordNotFound) {
// 		return errors.New("Ticket create failed"), 0
// 	}

// 	return err, TicketBaseData.ID
// }

/*
   获取工单数据
*/
// func TicketGet(ticketid int64) m_ticket.TicketBaseData {
// 	SelctSQL := `SELECT t.tn as tn FROM link_relation lr LEFT JOIN ticket t ON t.id = lr.target_key WHERE lr.source_key = ?`
// 	// ask database and fetch result
// 	global.GVA_DB_REPORT.Raw(SelctSQL, ticketid).Scan(&tnlist)
// }
