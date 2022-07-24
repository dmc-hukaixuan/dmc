package stats

import (
    "dmc/global"
    "dmc/kernel/util"
    "fmt"
    "strconv"
    "strings"
    "sync"
    "time"

    //"reflect"
    "github.com/xuri/excelize/v2"
)

type TicketPorsche struct {
    ID                     int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:api路径"`                 // 工单的 ID
    TicketNumber           string `json:"tn" gorm:"column:tn; comment:api中文描述"`                             // 工单的单号
    Title                  string `json:"title" gorm:"comment: ticket title"`                               // 工单的标题
    Queue                  string `json:"queue" gorm:"-"`                                                   // 工单角色
    QueueID                string `json:"queue_id" gorm:"comment:api组"`                                     // 工单角色 ID
    LockID                 string `json:"ticket_lock_id" gorm:"column:ticket_lock_id;comment:api组"`         // 工单是否锁定 id
    Lock                   string `json:"ticket_lock" gorm:"-"`                                             // 锁定的
    TypeID                 string `json:"type_id" gorm:"column:type_id;comment:api组"`                       // 类型 ID
    Type                   string `json:"type" gorm:"-"`                                                    // 工单类型
    ServiceID              string `json:"service_id" gorm:"column:service_id;comment:api组"`                 // 工单服务 id
    Service                string `json:"service" gorm:"-"`                                                 // 工单服务
    SLAID                  int    `json:"sla_id" gorm:"default:null;comment:api组"`                          // 工单 SLAID
    SLA                    string `json:"sla" gorm:"-"`                                                     // 工单 SLA
    User                   string `json:"user" gorm:"-"`                                                    // 工单指定处理人 ID
    UserID                 string `json:"user_id" gorm:"comment:api组"`                                      // 工单指定处理人
    ResponsibleUserID      string `json:"responsible_user_id" gorm:"comment:api组"`                          // 工单负责人 ID
    ResponsibleUser        string `json:"responsible_user" gorm:"-"`                                        // 工单负责人
    PriorityID             string `json:"ticket_priority_id" gorm:"column:ticket_priority_id;comment:api组"` // 工单优先级 ID
    Priority               string `json:"ticket_priority" gorm:"-"`                                         // 工单优先级
    StateID                string `json:"ticket_state_id" gorm:"column:ticket_state_id;comment:api组"`       // 工单状态 ID
    State                  string `json:"ticket_state" gorm:"-"`                                            // 工单状态
    CustomerID             string `json:"customer_id" gorm:"comment:api组"`                                  // 工单的客户
    CustomerUserID         string `json:"customer_user_id" gorm:"comment:api组"`                             // 工单的客户用户
    CreateCustomer         string `json:"create_customer" gorm:"comment:api组"`                              // 工单是否由用户创建，如果是用户创建，那么这里就是用户的登录名，如果不是用户创建这里则为空
    EscalationTime         int    `json:"escalation_time" gorm:"comment:api组"`
    EscalationUpdateTime   int    `json:"escalation_update_time" gorm:"comment:api组"`
    EscalationResponseTime int    `json:"escalation_response_time" gorm:"comment:api组"`
    EscalationSolutionTime int    `json:"escalation_solution_time" gorm:"comment:api组"`
    Timeout                string `json:"timeout" gorm:"comment:api组"`        // 工单锁定超时时间
    UntilTime              string `json:"until_time" gorm:"comment:api组"`     // 工单的环节
    ArchiveFlag            string `json:"archive_flag" gorm:"comment:api组"`   // 工单是否转存
    CreateTime             string `json:"create_time" gorm:"autoCreateTime;"` // 工单的创建时间
    CreateBy               string `json:"create_by" gorm:"comment:api组"`      // 工单的创建人
    ChangeTime             string `json:"change_time" gorm:"autoUpdateTime;"` // 工单的修改时间
    ChangeBy               string `json:"change_by" gorm:"comment:工单的修改人"`    // 工单的修改人
}

type TicketList struct{}

func (*TicketList) GenerateDynamicStatsRun(roles string, searchType string, starttime, stoptime string) {

    // mail.SendMail()

    // return
    //var tp []TicketPorsche
    // 创建一个 excel 写入数据
    // td := TicketDetailData(145778, "DSS")
    // fmt.Println("td, ", td.Jjfksj, "Yhfksj ", td.Yhfksj)
    // return
    // 查询 关闭事件在本月的所有工单
    // 126259  121483   121496
    // pos 122465, 127702 , 134416 , 119848
    // DSS 127881

    // rolecostTime, moveInfo := system.TicketFieldUpdateWorkingTime(147143, 12, "POS", "", "")

    // fmt.Println("test ", rolecostTime, moveInfo)
    // fmt.Println("L1FisrtMoveTime", moveInfo.SolveTeam, moveInfo.L2FirstMoveL3Time)

    // return
    //fmt.Println("PendingWorkingTime ", moveInfo.PendingWorkingTime, " L2MoveL3WorkingTime ", moveInfo.L2MoveL3WorkingTime," L2WaitingTime.",  moveInfo.L2WaitingTime )
    //fmt.Println("L2MoveL3WorkingTime ---- ", moveInfo.FirstRespsoneTime, rolecostTime)
    // fmt.Println("L2MoveL3WorkingTime ---- ", moveInfo.L2MoveL3WorkingTime, "L1WaitingTime: ", moveInfo.L1WaitingTime, ",L1FisrtMoveTime:", moveInfo.L1FisrtMoveTime,",L1FirstMoveWorkingTime :", moveInfo.L1FirstMoveWorkingTime)
    // "POS", "DSS", "D-Flow","Fadada support team"
    // porscheReprot := []string{"POS", "DSS", "D-Flow"}
    // starttime := "2022-07-01"
    // endtime := "2022-07-16"
    // starttime := "2022-05-01"
    // endtime := "2022-06-01"
    // "create", "close"
    // searchGroup := []string{"create"}
    // for _, searchType := range porscheReprot {
    //     mail.SendMail(searchType)
    // }
    // return
    // for _, v := range porscheReprot {
    //     for _, searchType := range searchGroup {
    //         //	mail.SendMail(v)
    //         TicketReportDataSyncGet(v, starttime, endtime, searchType)
    //         //mail.SendMail(v)
    //     }
    // }
    TicketReportDataSyncGet(roles, starttime, stoptime, searchType)
}

var wg sync.WaitGroup

func TicketReportDataSyncGet(roles, starttime, endtime, searchType string) {
    // get sla mapping calendar
    type SLACalender struct {
        SLAIQ      int `json:"id"  gorm:"column:id;" `
        CalendarID int `json:"create_by" gorm:"column:calendar_name;"`
    }
    var slaCalender []SLACalender
    global.GVA_DB_REPORT.Raw("select id, calendar_name from sla").Scan(&slaCalender)
    slaCalender1 := make(map[int]int)
    for _, sla := range slaCalender {
        slaCalender1[sla.SLAIQ] = sla.CalendarID
    }
    slatime := map[int]map[string]int{
        20: {
            "DSS L1":    30 * 60,
            "DSS L2":    60 * 60,
            "POS L1":    30 * 60,
            "POS L2":    2 * 60 * 60,
            "POS L3":    2 * 60 * 60,
            "D-Flow L1": 30 * 60,
            "D-Flow L2": 2 * 60 * 60,
        },
        21: {
            "DSS L1":    60 * 60,
            "DSS L2":    90 * 60,
            "POS L1":    60 * 60,
            "POS L2":    4 * 60 * 60,
            "POS L3":    4 * 60 * 60,
            "D-Flow L1": 60 * 60,
            "D-Flow L2": 6 * 60 * 60,
        },
        22: {
            "DSS L1":    2 * 60 * 60, // 2 working hours
            "DSS L2":    1210 * 60,   // 1210 minutes
            "POS L1":    2 * 60 * 60, //
            "POS L2":    8 * 60 * 60,
            "POS L3":    2 * 8 * 60 * 60,
            "D-Flow L1": 2 * 60 * 60,
            "D-Flow L2": 2 * 9 * 60 * 60,
        },
        23: {
            "DSS L1":    4 * 60 * 60,     // 4 working hours
            "DSS L2":    3250 * 60,       // 3250 minutes
            "POS L1":    4 * 60 * 60,     // 4 working hours
            "POS L2":    3 * 8 * 60 * 60, // 3 working days
            "POS L3":    5 * 8 * 60 * 60, // 5 working days
            "D-Flow L1": 4 * 60 * 60,
            "D-Flow L2": 5 * 9 * 60 * 60,
        },
        35: {
            "DSS L1":    30 * 60,
            "DSS L2":    60 * 60,
            "POS L1":    30 * 60,
            "POS L2":    2 * 60 * 60,
            "POS L3":    2 * 60 * 60,
            "D-Flow L1": 30 * 60,
            "D-Flow L2": 4 * 60 * 60,
        },
        36: {
            "DSS L1":    60 * 60,
            "DSS L2":    90 * 60,
            "POS L1":    60 * 60,
            "POS L2":    4 * 60 * 60,
            "POS L3":    4 * 60 * 60,
            "D-Flow L1": 60 * 60,
            "D-Flow L2": 6 * 60 * 60,
        },
        37: {
            "DSS L1":    2 * 60 * 60, // 2 working hours
            "DSS L2":    1210 * 60,   // 1210 minutes
            "POS L1":    2 * 60 * 60, //
            "POS L2":    8 * 60 * 60,
            "POS L3":    2 * 8 * 60 * 60,
            "D-Flow L1": 2 * 60 * 60,
            "D-Flow L2": 2 * 9 * 60 * 60,
        },
        38: {
            "DSS L1":    4 * 60 * 60,     // 4 working hours
            "DSS L2":    3250 * 60,       // 3250 minutes
            "POS L1":    4 * 60 * 60,     // 4 working hours
            "POS L2":    3 * 8 * 60 * 60, // 3 working days
            "POS L3":    5 * 8 * 60 * 60, // 5 working days
            "D-Flow L1": 4 * 60 * 60,
            "D-Flow L2": 5 * 9 * 60 * 60,
        },
    }
    forwordTime := map[int]map[string]int{
        20: { // P1(240min)
            "SLASoluationTime": 240,
            "DSS L1":           10 * 60,
            "DSS L2":           30 * 60,
            "POS L1":           10 * 60,
            "POS L2":           30 * 60,
            "POS L3":           10 * 60,
            "D-Flow L1":        10 * 60,
            "D-Flow L2":        10 * 60,
        },
        21: { // P2(480min)
            "SLASoluationTime": 480,
            "DSS L1":           10 * 60,
            "DSS L2":           60 * 60,
            "POS L1":           10 * 60,
            "POS L2":           60 * 60,
            "POS L3":           4 * 60 * 60,
            "D-Flow L1":        10 * 60,
            "D-Flow L2":        60 * 60,
        },
        22: { // P3(2160min)
            "SLASoluationTime": 2160,
            "DSS L1":           30 * 60,  // 2 working hours
            "DSS L2":           240 * 60, // 1210 minutes
            "POS L1":           30 * 60,  //
            "POS L2":           240 * 60,
            "POS L3":           2 * 8 * 60 * 60,
            "D-Flow L1":        30 * 60,
            "D-Flow L2":        240 * 60,
        },
        23: { // P4(5760min)
            "SLASoluationTime": 5760,
            "DSS L1":           30 * 60,         // 4 working hours
            "DSS L2":           480 * 60,        // 3250 minutes
            "POS L1":           30 * 60,         // 4 working hours
            "POS L2":           480 * 60,        // 3 working days
            "POS L3":           5 * 8 * 60 * 60, // 5 working days
            "D-Flow L1":        30 * 60,
            "D-Flow L2":        540 * 60,
        },
        35: { // P1(240min)-d-flow
            "SLASoluationTime": 240,
            "DSS L1":           10 * 60,
            "DSS L2":           10 * 60,
            "POS L1":           10 * 60,
            "POS L2":           10 * 60,
            "POS L3":           10 * 60,
            "D-Flow L1":        10 * 60,
            "D-Flow L2":        10 * 60,
        },
        36: { // P2(480min)-d-flow
            "SLASoluationTime": 480,
            "DSS L1":           10 * 60,
            "DSS L2":           60 * 60,
            "POS L1":           10 * 60,
            "POS L2":           60 * 60,
            "POS L3":           4 * 60 * 60,
            "D-Flow L1":        10 * 60,
            "D-Flow L2":        60 * 60,
        },
        37: { // P3(2160min)-d-flow
            "SLASoluationTime": 2160,
            "DSS L1":           30 * 60,  // 2 working hours
            "DSS L2":           240 * 60, // 1210 minutes
            "POS L1":           30 * 60,  //
            "POS L2":           240 * 60,
            "POS L3":           2 * 8 * 60 * 60,
            "D-Flow L1":        30 * 60,
            "D-Flow L2":        240 * 60,
        },
        38: { // P4(5760min)-d-flow
            "SLASoluationTime": 5760,
            "DSS L1":           30 * 60,         // 4 working hours
            "DSS L2":           480 * 60,        // 3250 minutes
            "POS L1":           30 * 60,         // 4 working hours
            "POS L2":           480 * 60,        // 3 working days
            "POS L3":           5 * 8 * 60 * 60, // 5 working days
            "D-Flow L1":        30 * 60,
            "D-Flow L2":        540 * 60,
        },
    }
    // 获取相关队列的所有数据
    ticketList := selectTicket(roles, starttime, endtime, searchType)
    rows1 := [][]string{}
    rows2 := [][]string{}
    rows3 := [][]string{}

    count := 1

    // 创建有5 个缓冲的通道，数据类型是  *TicketReprot
    ch := make(chan TicketReprot, 10)
    wg.Add(8)

    for i := 0; i < 8; i++ {
        go func() {
            defer wg.Done()
            for ticket := range ch {
                row3 := []string{}
                row1 := []string{}
                row2 := []string{}

                rolecostTime, moveInfo := util.TicketFieldUpdateWorkingTime(ticket.ID, slaCalender1[ticket.SLAID], roles, "", "")
                ///fmt.Println(" rolecostTime ---- ", ticket.ID, " moveInfo:",moveInfo.L1FisrtMoveTime, moveInfo.L1WaitingTime, ", MoveCount ", moveInfo.MoveCount)
                if moveInfo.MoveCount == 0 {
                    moveInfo.L1FisrtMoveTime = ""
                }

                var L1Reach string
                var L2Reach string
                var L3Reach string
                var L1ForwordReach string
                var L2ForwordReach string
                // 判断是否达标
                if slatime[ticket.SLAID][roles+" L1"] > rolecostTime[roles+" L1"] {
                    L1Reach = "Y"
                } else {
                    L1Reach = "N"
                }
                if slatime[ticket.SLAID][roles+" L2"] > rolecostTime[roles+" L2"] {
                    L2Reach = "Y"
                } else {
                    L2Reach = "N"
                }
                if slatime[ticket.SLAID][roles+" L3"] > rolecostTime[roles+" L3"] {
                    L3Reach = "Y"
                } else {
                    L3Reach = "N"
                }
                if forwordTime[ticket.SLAID][roles+" L1"] > moveInfo.L1FirstMoveWorkingTime {
                    L1ForwordReach = "Y"
                } else {
                    L1ForwordReach = "N"
                }
                if forwordTime[ticket.SLAID][roles+" L2"] > moveInfo.L2MoveL3WorkingTime {
                    L2ForwordReach = "Y"
                } else {
                    L2ForwordReach = "N"
                }
                var L1WaitingTime string
                var L2WaitingTime string
                var L1roles string
                var L3roles string
                var L2roles string

                var L1FirstmoveWorkingTime string
                //var L2StayWorkingTime string
                //
                if moveInfo.L1WaitingTime > 0 {
                    L1WaitingTime = fmt.Sprintf("%.2f", float64(moveInfo.L1WaitingTime)/float64(60))
                }
                if moveInfo.L2WaitingTime > 0 {
                    L2WaitingTime = fmt.Sprintf("%.2f", float64(moveInfo.L2WaitingTime)/float64(60))
                }
                if rolecostTime[roles+" L1"] > 0 {
                    L1roles = fmt.Sprintf("%.2f", float64(rolecostTime[roles+" L1"])/float64(60))
                    //fmt.Println("L1roles", L1roles)
                }
                if rolecostTime[roles+" L2"] > 0 {
                    L2roles = fmt.Sprintf("%.2f", float64(rolecostTime[roles+" L2"])/float64(60))
                    //fmt.Println("L2roles", L2roles)
                }
                if rolecostTime[roles+" L3"] > 0 {
                    L3roles = fmt.Sprintf("%.2f", float64(rolecostTime[roles+" L3"])/float64(60))
                    //fmt.Println("L3roles", L3roles)
                }

                if moveInfo.L1FirstMoveWorkingTime > 0 {
                    L1FirstmoveWorkingTime = fmt.Sprintf("%.2f", float64(moveInfo.L1FirstMoveWorkingTime)/float64(60))
                }
                L2MoveL3WorkingTime := fmt.Sprintf("%.2f", float64(moveInfo.L2MoveL3WorkingTime)/float64(60))
                // 第一张明细表
                // 获取工单的额外信息
                tdi := TicketDetailData(ticket.ID, roles)
                pendingWorkingTime := fmt.Sprintf("%.2f", float64(moveInfo.PendingWorkingTime)/float64(60))
                var Solutime int
                for _, v := range rolecostTime {
                    Solutime += v
                }
                Soluation := fmt.Sprintf("%.2f", float64(moveInfo.PendingWorkingTime)/float64(60))
                SolutimeMinu := fmt.Sprintf("%.2f", float64(forwordTime[ticket.SLAID]["SLASoluationTime"])-float64(Solutime)/float64(60))
                var actualCostTime string
                if tdi.Yhfksj != "" && moveInfo.SolutionOfferTime != "" {
                    rolecostTime1, _ := util.TicketFieldUpdateWorkingTime(ticket.ID, slaCalender1[ticket.SLAID], roles, tdi.Yhfksj, moveInfo.SolutionOfferTime)
                    var actualTime int
                    for _, v := range rolecostTime1 {
                        actualTime += v
                    }
                    actualCostTime = fmt.Sprintf("%.2f", float64(forwordTime[ticket.SLAID]["SLASoluationTime"])-float64(Solutime)/float64(60))
                }
                ticketOfID := strconv.FormatInt(ticket.ID, 10)
                CreateTimeStr := ticket.CreateTime.Format("2006-01-02 15:04:05")
                CloseTimeStr := ticket.CloseTime.Format("2006-01-02 15:04:05")
                //fmt.Println("tdi", tdi)
                row1 = []string{
                    ticketOfID, ticket.TicketNumber, tdi.Jxs, ticket.CreateCustomer, ticket.CreateUserName, CreateTimeStr, CloseTimeStr, moveInfo.FadadaInvoke, pendingWorkingTime, ticket.SLAName, moveInfo.CalendarAge,
                    Soluation, SolutimeMinu, ticket.TicketClassify, ticket.TwoTicketClassify, tdi.Jxdkh, tdi.Jxdgx, tdi.Jxdgxry, tdi.Yxfw, tdi.Yxfwgx, tdi.Yxfwgxry,
                    tdi.Slakh, tdi.Slagx, tdi.SlaGXRY, ticket.StateName, ticket.Moudle, ticket.MoudleClassfiy, ticket.Title, tdi.Xjs, ticket.RoleName, ticket.OwnerName, tdi.Yhfksj, tdi.Jjfksj,
                    actualCostTime, tdi.Scyysj, moveInfo.FirstRespsoneTime, tdi.Wtxt, tdi.Wtyy, ticket.TicketSource, tdi.Xtmc, tdi.Czhj, tdi.Zdyx, ticket.SLAHM, tdi.LinkCount, tdi.LinkList,
                }
                // row1 = []string{
                // 	ticketOfID, ticket.TicketNumber, tdi.Jxs, ticket.CreateCustomer, ticket.CreateUserName, CreateTimeStr, CloseTimeStr, moveInfo.FadadaInvoke, pendingWorkingTime, ticket.SLAName, moveInfo.CalendarAge,
                // 	Soluation, SolutimeMinu, ticket.TicketClassify, ticket.TwoTicketClassify, tdi.Jxdkh, tdi.Yxfw,
                // 	ticket.SLAName, ticket.StateName, ticket.Moudle, ticket.MoudleClassfiy, ticket.Title, tdi.Xjs, ticket.RoleName, ticket.OwnerName, tdi.Yhfksj, tdi.Jjfksj,
                // 	actualCostTime, tdi.Scyysj, moveInfo.FirstRespsoneTime, tdi.Wtxt, tdi.Wtyy, ticket.TicketSource, tdi.Xtmc, tdi.Czhj, tdi.Zdyx, ticket.SLAHM,
                // }
                // if moveInfo.L2StayWorkingTime> 0 {
                // 	L2StayWorkingTime = fmt.Sprintf("%.2f", float64(moveInfo.L2StayWorkingTime)/float64(60))
                // }
                ticket.SolveTeam = ""
                if ticket.StateName == "closed successful" || ticket.StateName == "首次电话解决" || ticket.StateName == "用户无回应" {
                    ticket.SolveTeam = moveInfo.SolveTeam
                    CloseTimeStr = moveInfo.CloseTime
                } else {
                    CloseTimeStr = ""
                    ticket.SolveTeam = ticket.RoleName
                }
                // "转移时长L2（分钟）", "L2转L3 SLA是否达标",
                if roles == "D-Flow" {
                    // "工单 ID","工单号", "创建时间", "关闭时间", "服务水平协议", "工单来源", "类型", "模块", "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）","L1转L2 SLA是否达标", "SLA豁免", "解决团队"
                    row2 = append(row2, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.SLAName, ticket.TicketSource, ticket.TypeName, ticket.Moudle)
                    row2 = append(row2, moveInfo.L1FisrtMoveTime, L1WaitingTime, L1FirstmoveWorkingTime, L1ForwordReach)
                    // "工单 ID","工单号", "创建时间", "关闭时间", "服务水平协议", "工单来源", "类型", "模块", "L1队列总时长", "L1队列总时长是否达标", "L2队列总时长", "L2队列总时长是否达标", "SLA豁免", "解决团队"
                    row3 = append(row3, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.SLAName, ticket.TicketSource, "类型", ticket.Moudle, L1roles, L1Reach, L2roles, L2Reach)

                    // row 1
                    // "工单 ID", "工单号", "所属经销商", "创建人（用户）", "创建人（支持）", "创建时间", "关闭时间", "解决团队", "挂起时长（分钟）", "SLA目标时间（分钟）", "总时长（分钟）",
                    // "解决时间（分钟）", "解决时间差（分钟）", "类型", "紧急度 - 客户", "影响范围 - 客户", "服务水平协议 - 客户", "状态", "模块", "标题", "信件树", "队列", "服务人员/所有者",
                    // "用户反馈时间", "解决方案提供时间", "实际处理时长", "首次响应时间", "首次响应时长（分钟）", "问题系统", "工单来源", "系统名称", "SLA豁免", "link的工单数", "link的工单号",
                    row1 = []string{
                        ticketOfID, ticket.TicketNumber, tdi.Jxs, ticket.CreateCustomer, ticket.CreateUserName, CreateTimeStr, CloseTimeStr, pendingWorkingTime, ticket.SLAName, moveInfo.CalendarAge,
                        Soluation, SolutimeMinu, ticket.TypeName, tdi.Jxdkh, tdi.Yxfw, ticket.SLAName, ticket.StateName, ticket.Moudle, ticket.Title, tdi.Xjs, ticket.RoleName, ticket.OwnerName,
                        tdi.Yhfksj, tdi.Jjfksj, actualCostTime, tdi.Scyysj, moveInfo.FirstRespsoneTime, tdi.Wtxt, ticket.TicketSource, tdi.Xtmc, ticket.SLAHM, tdi.LinkCount, tdi.LinkList,
                    }
                    // row1 = []string{
                    // 	ticketOfID, ticket.TicketNumber, tdi.Jxs, ticket.CreateCustomer, ticket.CreateUserName, CreateTimeStr, CloseTimeStr, pendingWorkingTime, ticket.SLAName, moveInfo.CalendarAge,
                    // 	Soluation, SolutimeMinu, ticket.TypeName, tdi.Jxdkh, tdi.Yxfw, ticket.SLAName, ticket.StateName, ticket.Moudle, ticket.Title, tdi.Xjs, ticket.RoleName, ticket.OwnerName,
                    // 	tdi.Yhfksj, tdi.Jjfksj, actualCostTime, tdi.Scyysj, moveInfo.FirstRespsoneTime, tdi.Wtxt, ticket.TicketSource, tdi.Xtmc, ticket.SLAHM,
                    // }
                } else if roles == "POS" {
                    // "工单 ID", "工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议",
                    // "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）", "L1转L2 SLA是否达标", "L2第一次转L3的时间", "等待总时长L2（分钟）", "转移时长L2（分钟）", "L2转L3 SLA是否达标", "SLA豁免", "解决团队"
                    row2 = append(row2, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.TicketSource, ticket.TicketClassify, ticket.TwoTicketClassify, ticket.Moudle, ticket.MoudleClassfiy, ticket.SLAName)
                    row2 = append(row2, moveInfo.L1FisrtMoveTime, L1WaitingTime, L1FirstmoveWorkingTime, L1ForwordReach, moveInfo.L2FirstMoveL3Time, L2WaitingTime, L2MoveL3WorkingTime, L2ForwordReach)
                    // "工单 ID","工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类",  "服务水平协议", "L1队列总时长", "L1队列总时长是否达标" , "L2队列总时长", "L2队列总时长是否达标", "L3队列总时长", "L3队列总时长是否达标", "SLA豁免", "解决团队"
                    row3 = append(row3, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.TicketSource, ticket.TicketClassify, ticket.TwoTicketClassify, ticket.Moudle, ticket.MoudleClassfiy, ticket.SLAName, L1roles, L1Reach, L2roles, L2Reach)
                    // "L3队列总时长是否达标"
                    FadadaCost := fmt.Sprintf("%.2f", float64(rolecostTime["Fadada support team"])/float64(60))
                    row3 = append(row3, L3roles, L3Reach, FadadaCost)
                } else {
                    FadadaCost := fmt.Sprintf("%.2f", float64(rolecostTime["Fadada support team"])/float64(60))
                    // "工单 ID","工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类",  "服务水平协议", "L1队列总时长", "L1队列总时长是否达标" , "L2队列总时长", "L2队列总时长是否达标", "Fadada 总时长" "SLA豁免", "解决团队"
                    row3 = append(row3, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.TicketSource, ticket.TicketClassify, ticket.TwoTicketClassify, ticket.Moudle, ticket.MoudleClassfiy, ticket.SLAName, L1roles, L1Reach, L2roles, L2Reach, FadadaCost)
                    // "工单 ID","工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议", "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）","L1转L2 SLA是否达标", "SLA豁免", "解决团队"
                    row2 = append(row2, ticketOfID, ticket.TicketNumber, CreateTimeStr, CloseTimeStr, ticket.TicketSource, ticket.TicketClassify, ticket.TwoTicketClassify, ticket.Moudle, ticket.MoudleClassfiy, ticket.SLAName)
                    row2 = append(row2, moveInfo.L1FisrtMoveTime, L1WaitingTime, L1FirstmoveWorkingTime, L1ForwordReach)
                }

                row2 = append(row2, ticket.SLAHM, ticket.SolveTeam)
                rows2 = append(rows2, row2)
                row3 = append(row3, ticket.SLAHM, ticket.SolveTeam)
                rows3 = append(rows3, row3)

                if roles == "Fadada support team" {
                    costTimeStr := fmt.Sprintf("%.2f", float64(float64(rolecostTime[roles])/float64(60)))
                    // "工单 ID", "工单号", "工单标题", "所属经销商", "系统名称", "优先级", "创建时间", "状态", "指定处理人", "角色", "关闭时间", "总时长（分钟）",
                    row1 = []string{
                        ticketOfID, ticket.TicketNumber, ticket.Title, tdi.Jxs, tdi.Xtmc, ticket.PriorityName, CreateTimeStr, ticket.StateName, ticket.OwnerName, ticket.RoleName, CloseTimeStr, costTimeStr,
                    }
                    rows1 = append(rows1, row1)
                } else {
                    row1 = append(row1, ticket.SolveTeam)
                    rows1 = append(rows1, row1)
                }
            }
        }()
    }

    ticketCount := len(ticketList)
    for _, ticket := range ticketList {
        fmt.Println("count  :", ticketCount, count, ticket.ID)
        ch <- ticket
        count++

    }
    close(ch)
    wg.Wait()
    headerNameArray3 := []string{"工单 ID", "工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议", "L1队列总时长", "L1队列总时长是否达标", "L2队列总时长", "L2队列总时长是否达标", "L3队列总时长", "L3队列总时长是否达标", "FaDaDa 队列总时长", "SLA豁免", "解决团队"}

    headerNameArray2 := []string{"工单 ID", "工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议", "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）", "L1转L2 SLA是否达标", "L2第一次转L3的时间", "等待总时长L2（分钟）", "转移时长L2（分钟）", "L2转L3 SLA是否达标", "SLA豁免", "解决团队"}
    headerNameArray1 := []string{
        "工单 ID", "工单号", "所属经销商", "创建人（用户）", "创建人（支持）", "创建时间", "关闭时间", "是否转给法大大", "挂起时长（分钟）", "SLA目标时间（分钟）", "总时长（分钟）",
        "解决时间（分钟）", "解决时间差（分钟）", "问题分类1", "问题分类2", "紧急度 - 客户", "紧急度 - 更新", "紧急度 - 更新人员", "影响范围 - 客户", "影响范围 - 更新", "影响范围 - 更新人员",
        "服务水平协议 - 客户", "服务水平协议 - 更新", "服务水平协议 - 更新人员", "状态", "模块", "模块分类", "标题", "信件树", "队列", "服务人员/所有者", "用户反馈时间", "解决方案提供时间",
        "实际处理时长", "首次响应时间", "首次响应时长（分钟）", "问题系统", "问题原因", "工单来源", "系统名称", "操作环境", "重大影响", "SLA豁免", "link的工单数", "link的工单号", "解决团队",
    }
    // headerNameArray1 := []string{
    // 	"工单 ID", "工单号", "所属经销商", "创建人（用户）", "创建人（支持）", "创建时间", "关闭时间", "是否转给法大大", "挂起时长（分钟）", "SLA目标时间（分钟）", "总时长（分钟）",
    // 	"解决时间（分钟）", "解决时间差（分钟）", "问题分类1", "问题分类2", "紧急度 - 客户", "影响范围 - 客户",
    // 	"服务水平协议 - 客户", "状态", "模块", "模块分类", "标题", "信件树", "队列", "服务人员/所有者", "用户反馈时间", "解决方案提供时间",
    // 	"实际处理时长", "首次响应时间", "首次响应时长（分钟）", "问题系统", "问题原因", "工单来源", "系统名称", "操作环境", "重大影响", "SLA豁免", "解决团队",
    // }
    if roles == "DSS" {
        headerNameArray2 = []string{"工单 ID", "工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议", "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）", "L1转L2 SLA是否达标", "SLA豁免", "解决团队"}
        headerNameArray3 = []string{"工单 ID", "工单号", "创建时间", "关闭时间", "工单来源", "问题分类1", "问题分类2", "模块", "模块分类", "服务水平协议", "L1队列总时长", "L1队列总时长是否达标", "L2队列总时长", "L2队列总时长是否达标", "FaDaDa 队列总时长", "SLA豁免", "解决团队"}
    } else if roles == "D-Flow" {
        headerNameArray1 = []string{
            "工单 ID", "工单号", "所属经销商", "创建人（用户）", "创建人（支持）", "创建时间", "关闭时间", "挂起时长（分钟）", "SLA目标时间（分钟）", "总时长（分钟）",
            "解决时间（分钟）", "解决时间差（分钟）", "类型", "紧急度 - 客户", "影响范围 - 客户", "服务水平协议 - 客户", "状态", "模块", "标题", "信件树", "队列", "服务人员/所有者",
            "用户反馈时间", "解决方案提供时间", "实际处理时长", "首次响应时间", "首次响应时长（分钟）", "问题系统", "工单来源", "系统名称", "SLA豁免", "link的工单数", "link的工单号", "解决团队",
        }
        // headerNameArray1 = []string{
        // 	"工单 ID", "工单号", "所属经销商", "创建人（用户）", "创建人（支持）", "创建时间", "关闭时间", "挂起时长（分钟）", "SLA目标时间（分钟）", "总时长（分钟）",
        // 	"解决时间（分钟）", "解决时间差（分钟）", "类型", "紧急度 - 客户", "影响范围 - 客户", "服务水平协议 - 客户", "状态", "模块", "标题", "信件树", "队列", "服务人员/所有者",
        // 	"用户反馈时间", "解决方案提供时间", "实际处理时长", "首次响应时间", "首次响应时长（分钟）", "问题系统", "工单来源", "系统名称", "SLA豁免", "解决团队",
        // }
        headerNameArray2 = []string{"工单 ID", "工单号", "创建时间", "关闭时间", "服务水平协议", "工单来源", "类型", "模块", "L1第一次转L2的时间", "等待总时长L1（分钟）", "转移时长L1（分钟）", "L1转L2 SLA是否达标", "SLA豁免", "解决团队"}
        headerNameArray3 = []string{"工单 ID", "工单号", "创建时间", "关闭时间", "服务水平协议", "工单来源", "类型", "模块", "L1队列总时长", "L1队列总时长是否达标", "L2队列总时长", "L2队列总时长是否达标", "SLA豁免", "解决团队"}
    }
    if roles == "Fadada support team" {
        headerNameArray1 = []string{
            "工单 ID", "工单号", "工单标题", "所属经销商", "系统名称", "优先级", "创建时间", "状态", "指定处理人", "角色", "关闭时间", "总时长（分钟）",
        }
        CreateXlS(rows1, roles+" 第一张报表_"+searchType+"_"+starttime+"_"+strconv.FormatInt(time.Now().Unix(), 10), headerNameArray1)
    } else {
        // fmt.Println("rows3 ", len(rows3[0]), len(headerNameArray3))
        // fmt.Println("rows2 ", len(rows3[0]), len(headerNameArray3))
        CreateXlS(rows3, roles+" 第三张报表_"+searchType+"_"+starttime+"_"+strconv.FormatInt(time.Now().Unix(), 10), headerNameArray3)
        CreateXlS(rows2, roles+" 第二张报表_"+searchType+"_"+starttime+"_"+strconv.FormatInt(time.Now().Unix(), 10), headerNameArray2)
        CreateXlS(rows1, roles+" 第一张报表_"+searchType+"_"+starttime+"_"+strconv.FormatInt(time.Now().Unix(), 10), headerNameArray1)
    }
}

var slaTargetTime = map[int]int{}

var posFieldMapping = map[string]map[string]string{
    "POS": map[string]string{
        "jxs":       "POSjingxioashang", // 所属经销商
        "sfwkhyhtd": "sfwkhyhtd",
        "yhfksj":    "POSyonghufankuishijian", // POS用户反馈时间
        "jjtd":      "JJTD",                   // 解决团队
        "scyysj":    "IncidentStartTime",      // 首次响应时间
        "wtfl1":     "POSwentifenlei1",        //问题分类1
        "wtfl2":     "POSwentifenlei2",        //问题分类2
        "yxfw":      "POSyingxiangfanwei",     // pos影响范围
        "jjd":       "POSjinjidu",             // 紧急度
        "mk":        "POSmokuai",              //模块
        "mkfl":      "POSmokuaifenlei",        //模块分类
        "jjfksj":    "POSjiejueshijian",       // 解决方案提供时间
        "wtxt":      "POSwentixitong",         // 问题系统
        "wtyy":      "POSwentiyuanyin",        // 问题原因
        "gdly":      "POSgongdanlaiyuan",      // 工单来源
        "xtmc":      "dxcxitongmingcheng",     // 系统名称
        "czhj":      "POScaozuohuanjing",      // 操作环境
        "zdyx":      "POSzhongdayingxiang",    // 重大影响
        "slahm":     "SLAhm",                  // SLA豁免
    },
    "DSS": map[string]string{
        "yhfksj": "DSSyonghufankuishijian", // 用户反馈时间
        "jxs":    "POSjingxioashang",       // 所属经销商
        "jjtd":   "JJTD",                   // 解决团队
        "scyysj": "IncidentStartTime",      // 首次响应时间
        "wtfl1":  "DSSwentifenlei1",        //问题分类1
        "wtfl2":  "DSSwentifenlei2",        //问题分类2
        "jjd":    "DSSjinjidu",             // 紧急度
        "yxfw":   "DSSyingxiangfanwei",     // pos影响范围
        "mk":     "DSSmokuia",              //模块
        "mkfl":   "DSSmokuaifenlei",        //模块分类
        "jjfksj": "POSjiejueshijian",       // 解决方案提供时间
        "wtxt":   "DSSwentixitong",         // 问题系统
        "wtyy":   "DSSwentiyuanyin",        // 问题原因
        "gdly":   "gongdanlaiyuan",         // 工单来源
        "xtmc":   "dxcxitongmingcheng",     // 系统名称
        "czhj":   "DSScaozuohuanjing",      // 操作环境
        "zdyx":   "POSzhongdayingxiang",    // 重大影响
        "slahm":  "SLAhm",                  // SLA豁免
    },
    "D-Flow": map[string]string{
        "yhfksj": "DFlowyonghufankuishijian", // 用户反馈时间
        "jxs":    "POSjingxioashang",         // 所属经销商
        "jjd":    "DFlowjinjidu",             // 紧急度
        "yxfw":   "DFLOWYXFW",                // pos影响范围
        "wtxt":   "DSSwentixitong",
        "mk":     "DFlowmokuai",        //模块
        "jsfa":   "DFlowjiejueshijian", // 解决方案提供时间
        "jjtd":   "JJTD",               // 解决团队
        "xtmc":   "dxcxitongmingcheng", // 系统名称
        "czhj":   "DSScaozuohuanjing",  //操作环境
        "slahm":  "SLAhm",              // SLA豁免
    },
    "Fadada support team": map[string]string{
        "yhfksj": "DFlowyonghufankuishijian", // 用户反馈时间
        "jxs":    "POSjingxioashang",         // 所属经销商
        "jjd":    "DFlowjinjidu",             // 紧急度
        "yxfw":   "DFLOWYXFW",                // pos影响范围
        "wtxt":   "DSSwentixitong",
        "mk":     "DFlowmokuai",        //模块
        "jsfa":   "DFlowjiejueshijian", // 解决方案提供时间
        "jjtd":   "JJTD",               // 解决团队
        "xtmc":   "dxcxitongmingcheng", // 系统名称
        "czhj":   "DSScaozuohuanjing",  //操作环境
        "slahm":  "SLAhm",              // SLA豁免
    },
}

type TicketDatInfo struct {
    Jxs      string // 所属经销商
    Jjsjc    string // 解决时间差（分钟） SLA目标时间-解决时间（分钟）
    Jjfksj   string //
    Scyysj   string // 首次响应时间
    Jxdkh    string // 第一次选择的紧急度
    Jxdgx    string // 更新后最终的紧急度，如果未更新，则显示为空
    Jxdgxry  string // 更新人员姓名
    Yxfw     string // 第一次选择的影响范围
    Yxfwgx   string // 更新后最终的影响范围，如果未更新，则显示为空
    Yxfwgxry string // 影响范围更新人员姓名
    Slakh    string //第一次显示的优先级
    Slagx    string //服务水平协议 - 更新
    SlaGXRY  string // 服务水平协议 - 更新人员
    Xjs      string // 信件树
    Yhfksj   string //用户反馈时间
    Wtxt     string //问题系统
    Wtyy     string //问题原因
    Xtmc     string //系统名称
    Czhj     string //操作环境
    Wtfl1    string //问题分类1
    Wtfl2    string //问题分类2
    Mk       string //模块
    Mkfl     string //模块分类
    Gdly     string // 工单来源
    Zdyx     string // 重大影响
    Slahm    string //SLA豁免
    Jjtd     string // 解决团队

    LinkCount string
    LinkList  string
}

/*

 */
func TicketDetailData(ticketid int64, roles string) TicketDatInfo {
    var td TicketDatInfo
    articlelist := util.TicketArticleList(ticketid)
    dfvList := util.TicketynamicFieldValueGet(ticketid)
    linkList, lincount := util.TicketLinkList(ticketid)
    tempMapping := posFieldMapping[roles]

    td.Czhj = dfvList[tempMapping["czhj"]]
    td.Xtmc = dfvList[tempMapping["xtmc"]]
    td.Jjfksj = dfvList[tempMapping["jjfksj"]]

    td.Wtxt = dfvList[tempMapping["wtxt"]]
    td.Wtyy = dfvList[tempMapping["wtyy"]]
    td.Zdyx = dfvList[tempMapping["zdyx"]]
    td.Yhfksj = dfvList[tempMapping["yhfksj"]]
    td.Jxs = dfvList[tempMapping["jxs"]]
    td.Xjs = articlelist
    td.LinkList = linkList
    td.LinkCount = strconv.Itoa(lincount)

    // 获取 sla 的更新历史 39 :SLAUpdate
    slaHistory := util.TicketHistoryCustom(ticketid, 39, "")
    // sla
    if len(slaHistory) > 1 {
        arr := strings.Split(slaHistory[len(slaHistory)-1].Name, "%%")
        arr1 := strings.Split(slaHistory[0].Name, "%%")

        td.Slagx = arr[1]
        td.Slakh = arr1[1]
        td.SlaGXRY = slaHistory[len(slaHistory)-1].CreateUser
    } else if len(slaHistory) == 1 {
        arr := strings.Split(slaHistory[0].Name, "%%")
        td.Slakh = arr[1]
    }

    // 获取影响范围的历史
    if tempMapping["yxfw"] != "" {
        ImpactsHistory := util.TicketHistoryCustom(ticketid, 28, tempMapping["yxfw"])
        if len(ImpactsHistory) > 1 {
            arr := strings.Split(ImpactsHistory[len(ImpactsHistory)-1].Name, "%%")
            arr1 := strings.Split(ImpactsHistory[0].Name, "%%")
            td.Yxfwgx = arr[4]
            td.Yxfw = arr1[4]
            td.Yxfwgxry = ImpactsHistory[len(ImpactsHistory)-1].CreateUser
        } else {
            td.Yxfw = dfvList[tempMapping["yxfw"]]
        }
    }

    if tempMapping["jjd"] != "" {
        // 获取紧急度的历史
        jxdHistory := util.TicketHistoryCustom(ticketid, 28, tempMapping["jjd"])
        if len(jxdHistory) > 1 {
            arr := strings.Split(jxdHistory[len(jxdHistory)-1].Name, "%%")
            arr1 := strings.Split(jxdHistory[0].Name, "%%")
            td.Jxdgx = arr[4]
            td.Jxdkh = arr1[4]
            td.Jxdgxry = jxdHistory[len(jxdHistory)-1].CreateUser
            td.Jxdkh = dfvList[tempMapping["jjd"]]
        } else {
            td.Jxdkh = dfvList[tempMapping["jjd"]]
        }
    }

    return td
}

type TicketReprot struct {
    ID                int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:api路径"` // 工单的 ID
    TicketNumber      string    `json:"tn" gorm:"column:tn; comment:api中文描述"`             // 工单的单号
    SLAName           string    `json:"sla_name" gorm:"comment: ticket title"`
    SLAID             int       `json:"sla_id" `
    Title             string    `json:"title"`
    OwnerName         string    `json:"owner_name"`
    TypeName          string    `json:"type_name"`
    RoleName          string    `json:"role_name"`
    StateName         string    `json:"state_name"`
    PriorityName      string    `json:"priority_name" `
    PriorityID        int       `json:"priority_id" `                      // 工单的标题
    CloseTime         time.Time `json:"close_time" `                       // 工单角色
    TicketSource      string    `json:"ticket_source" gorm:"comment:api组"` // 工单角色 ID
    TicketClassify    string    `json:"ticket_classify"`                   // 工单是否锁定 id
    TwoTicketClassify string    `json:"two_ticket_classify"`
    SLAHM             string    `json:"sla_hm"`
    SolveTeam         string    `json:"solve_team"`
    Moudle            string    `json:"moudle"` // 类型 ID
    MoudleClassfiy    string    `json:"moudle_classfiy" `
    CreateCustomer    string    `json:"create_customer"`
    CreateTime        time.Time `json:"create_time" `
    CreateUserName    string    `json:"create_user_name" `
    CreateBy          string    `json:"create_by" gorm:"comment:api组"`      // 工单的创建人
    ChangeTime        string    `json:"change_time" gorm:"autoUpdateTime;"` // 工单的修改时间
    ChangeBy          string    `json:"change_by" gorm:"comment:工单的修改人"`    // 工单的修改人
}

// 查询数据
func selectTicket(roles, starttime, endtime, searchType string) []TicketReprot {
    searchTime := "th.create_time > '" + starttime + "' AND th.create_time < '" + endtime + "'"

    selectSQLDSS := `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                        s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                        q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name
                        , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                        , max(case when dfv.field_id=86 then dfv.value_text else null end) as ticket_source
                        , max(case when dfv.field_id=152 then dfv.value_text else null end) as ticket_classify
                        , max(case when dfv.field_id=153 then dfv.value_text else null end) as two_ticket_classify
                        , max(case when dfv.field_id=158 then dfv.value_text else null end) as moudle
                        , max(case when dfv.field_id=159 then dfv.value_text else null end) as moudle_classfiy
                        , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                        , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                    FROM ticket t
                        LEFT JOIN sla s ON s.id = t.sla_id
                        RIGHT JOIN queue q ON q.id = t.queue_id AND (q.name LIKE 'DataLake%' OR q.name LIKE '` + roles + `%') AND q.valid_id =1
                        LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(86,152,153,158,159,249, 240)
                        LEFT JOIN users u ON u.id = t.user_id
                        LEFT JOIN users u1 ON u1.id = t.create_by
                        LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id

                        LEFT JOIN ticket_history th ON t.id = th.ticket_id
                            AND (th.name LIKE '%closed successful%' OR th.name LIKE '%首次电话解决%' OR th.name LIKE '%用户无回应%' )
                    WHERE
                        t.sla_id IS NOT NULL AND ` + searchTime + `

                    GROUP BY t.tn`

    selectSQLPOS := `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                        s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                        q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name
                        , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                        , max(case when dfv.field_id=183 then dfv.value_text else null end) as ticket_source
                        , max(case when dfv.field_id=175 then dfv.value_text else null end) as ticket_classify
                        , max(case when dfv.field_id=176 then dfv.value_text else null end) as two_ticket_classify
                        , max(case when dfv.field_id=134 then dfv.value_text else null end) as moudle
                        , max(case when dfv.field_id=135 then dfv.value_text else null end) as moudle_classfiy
                        , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                        , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                    FROM ticket t
                        LEFT JOIN sla s ON s.id = t.sla_id
                        RIGHT JOIN queue q ON q.id = t.queue_id AND q.name LIKE '` + roles + `%' AND q.valid_id =1
                        LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(203, 183, 175, 176, 134, 135, 249,240)
                        LEFT JOIN users u ON u.id = t.user_id
                        LEFT JOIN users u1 ON u1.id = t.create_by
                        LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                        LEFT JOIN ticket_history th ON t.id = th.ticket_id
                            AND (th.name LIKE '%closed successful%' OR th.name LIKE '%首次电话解决%' OR th.name LIKE '%用户无回应%' )
                    WHERE
                        t.sla_id IS NOT NULL AND ` + searchTime + `

                    GROUP BY t.tn`
    selectSQLFadada := `SELECT t.id as id, t.title as title, t.tn as tn, t.create_time as create_time, u.full_name AS owner_name, ts.name AS state_name,
                            q.name AS role_name, u1.login AS createby, t.create_customer AS create_customer
                            , tp.name as priority_name, t.create_time as create_time
                            , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                        FROM ticket t
                            LEFT JOIN queue q ON q.id = t.queue_id
                            LEFT JOIN sla s ON s.id = t.sla_id
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN ticket_priority tp ON tp.id = t.ticket_priority_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th ON t.id = th.ticket_id
                            LEFT JOIN ticket_history th1 ON t.id = th1.ticket_id
                        WHERE
                            ` + searchTime + `
                            AND (th.name LIKE '%closed successful%' OR th.name LIKE '%首次电话解决%' OR th.name LIKE '%用户无回应%' )
                            AND th1.queue_id IN ( SELECT id FROM queue WHERE queue.name LIKE 'Fadada%' AND queue.valid_id =1 )
                        GROUP BY t.tn
                        `

    selectSQLDFlow := `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                            s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                            q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name, tt.name as type_name
                            , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                            , max(case when dfv.field_id=86 then dfv.value_text else null end) as ticket_source
                            , max(case when dfv.field_id=175 then dfv.value_text else null end) as ticket_classify
                            , max(case when dfv.field_id=176 then dfv.value_text else null end) as two_ticket_classify
                            , max(case when dfv.field_id=255 then dfv.value_text else null end) as moudle
                            , max(case when dfv.field_id=253 then dfv.value_text else null end) as moudle_classfiy
                            , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                            , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                        FROM ticket t
                            LEFT JOIN sla s ON s.id = t.sla_id
                            RIGHT JOIN queue q ON q.id = t.queue_id AND q.name LIKE '` + roles + `%' AND q.valid_id =1
                            LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(86,253,255,249,240)
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_type tt ON tt.id = t.type_id
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th ON t.id = th.ticket_id
                                AND (th.name LIKE '%closed successful%' OR th.name LIKE '%首次电话解决%' OR th.name LIKE '%用户无回应%' )
                        WHERE
                            t.sla_id IS NOT NULL AND ` + searchTime + `

                            GROUP BY t.tn`

    if searchType == "create" {
        searchTime = "t.create_time > '" + starttime + "' AND t.create_time < '" + endtime + "'"
        selectSQLDSS = `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                            s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                            q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name
                            , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                            , max(case when dfv.field_id=86 then dfv.value_text else null end) as ticket_source
                            , max(case when dfv.field_id=152 then dfv.value_text else null end) as ticket_classify
                            , max(case when dfv.field_id=153 then dfv.value_text else null end) as two_ticket_classify
                            , max(case when dfv.field_id=158 then dfv.value_text else null end) as moudle
                            , max(case when dfv.field_id=159 then dfv.value_text else null end) as moudle_classfiy
                            , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                            , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                        FROM ticket t
                            LEFT JOIN sla s ON s.id = t.sla_id
                            LEFT JOIN queue q ON q.id = t.queue_id
                            LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(86,152,153,158,159,249, 240)
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th ON t.id = th.ticket_id
                        WHERE

                            th.queue_id IN (SELECT id FROM queue WHERE queue.name LIKE "` + roles + `%") AND th.history_type_id = 1
                            AND ` + searchTime + `
                        GROUP BY t.tn`

        selectSQLPOS = `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                            s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                            q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name
                            , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                            , max(case when dfv.field_id=183 then dfv.value_text else null end) as ticket_source
                            , max(case when dfv.field_id=175 then dfv.value_text else null end) as ticket_classify
                            , max(case when dfv.field_id=176 then dfv.value_text else null end) as two_ticket_classify
                            , max(case when dfv.field_id=134 then dfv.value_text else null end) as moudle
                            , max(case when dfv.field_id=135 then dfv.value_text else null end) as moudle_classfiy
                            , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                            , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                        FROM ticket t
                            LEFT JOIN sla s ON s.id = t.sla_id
                            LEFT JOIN queue q ON q.id = t.queue_id
                            LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(203, 183, 175, 176, 134, 135, 249,240)
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th ON t.id = th.ticket_id
                        WHERE

                            th.queue_id IN (SELECT id FROM queue WHERE queue.name LIKE "` + roles + `%") AND th.history_type_id = 1
                            AND ` + searchTime + `
                        GROUP BY t.tn`

        selectSQLDFlow = `SELECT t.id as id, t.title as title, t.tn as tn, t.sla_id as sla_id, t.create_time as create_time,
                        s.name as sla_name, t.ticket_priority_id as priority_id, ts.name AS state_name,t.create_customer AS create_customer,
                        q.name AS role_name, u.full_name AS owner_name, u1.login AS create_user_name, tt.name as type_name
                        , max(case when IFNULL(th.id,0)<>0 then th.create_time ELSE NULL END) as close_time
                        , max(case when dfv.field_id=86 then dfv.value_text else null end) as ticket_source
                        , max(case when dfv.field_id=175 then dfv.value_text else null end) as ticket_classify
                        , max(case when dfv.field_id=176 then dfv.value_text else null end) as two_ticket_classify
                        , max(case when dfv.field_id=255 then dfv.value_text else null end) as moudle
                        , max(case when dfv.field_id=253 then dfv.value_text else null end) as moudle_classfiy
                        , max(case when dfv.field_id=249 then dfv.value_text else null end) as sla_hm
                        , max(case when dfv.field_id=240 then dfv.value_text else null end) as solve_team
                        FROM ticket t
                            LEFT JOIN sla s ON s.id = t.sla_id
                            LEFT JOIN queue q ON q.id = t.queue_id
                            LEFT JOIN dynamic_field_value dfv ON t.id = dfv.object_id AND dfv.field_id IN(86,253,255,249,240)
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_type tt ON tt.id = t.type_id
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th ON t.id = th.ticket_id
                        WHERE

                            th.queue_id IN (SELECT id FROM queue WHERE queue.name LIKE "` + roles + `%") AND th.history_type_id = 1
                            AND ` + searchTime + `
                            GROUP BY t.tn`
        selectSQLFadada = `SELECT t.id as id, t.title as title, t.tn as tn, t.create_time as create_time,  u.full_name AS owner_name, ts.name AS state_name,
                            q.name AS role_name, u1.login AS createby, t.create_customer AS create_customer
                            , tp.name as priority_name, t.create_time as create_time
                        FROM ticket t
                            LEFT JOIN queue q ON q.id = t.queue_id
                            LEFT JOIN sla s ON s.id = t.sla_id
                            LEFT JOIN users u ON u.id = t.user_id
                            LEFT JOIN users u1 ON u1.id = t.create_by
                            LEFT JOIN ticket_priority tp ON tp.id = t.ticket_priority_id
                            LEFT JOIN ticket_state ts ON t.ticket_state_id = ts.id
                            LEFT JOIN ticket_history th1 ON t.id = th1.ticket_id
                        WHERE
                            ` + searchTime + `
                            AND th1.queue_id IN ( SELECT id FROM queue WHERE queue.name LIKE 'Fadada%' AND queue.valid_id =1 )
                        GROUP BY t.tn
                        `
    }
    var selectSQL string
    if roles == "DSS" {
        selectSQL = selectSQLDSS
    } else if roles == "POS" {
        selectSQL = selectSQLPOS
    } else if roles == "D-Flow" {
        selectSQL = selectSQLDFlow
    } else if roles == "Fadada support team" {
        selectSQL = selectSQLFadada
    }
    var ticketList []TicketReprot

    global.GVA_DB_REPORT.Raw(selectSQL).Scan(&ticketList)
    fmt.Println("sql ----************ ", roles, selectSQL)
    return ticketList
}

// CreateXlS data为要写入的数据,fileName 文件名称, headerNameArray 表头数组
func CreateXlS(data [][]string, fileName string, headerNameArray []string) {
    f := excelize.NewFile()
    sheetName := "Sheet1"
    sheetWords := []string{
        "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
        "V", "W", "X", "Y", "Z",
        "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU",
        "AV", "AW", "AX", "AY", "AZ",
    }

    for k, v := range headerNameArray {
        //fmt.Println(" head line  ",sheetWords[k]+"1", " head line value :", v )
        f.SetCellValue(sheetName, sheetWords[k]+"1", v)
    }

    //设置列的宽度
    f.SetColWidth("Sheet1", "A", sheetWords[len(headerNameArray)-1], 18)
    headStyleID, _ := f.NewStyle(`{
        "font":{
            "color":"#333333",
            "bold":false,
            "family":"arial"
        },
        "alignment":{
            "vertical":"center",
            "horizontal":"center"
        }
        }`)
    //设置表头的样式
    f.SetCellStyle(sheetName, "A1", sheetWords[len(headerNameArray)-1]+"1", headStyleID)
    line := 1
    // 循环写入数据
    for _, v := range data {
        line++
        for kk, _ := range headerNameArray {
            //fmt.Println(" cell name ", sheetWords[kk]+strconv.Itoa(line), "value :", v[kk])
            f.SetCellValue(sheetName, sheetWords[kk]+strconv.Itoa(line), v[kk])
        }
    }

    // 保存文件
    if err := f.SaveAs("D:\\05-附件\\02-跑出来的报表\\" + fileName + ".xlsx"); err != nil {
        fmt.Println(err.Error())
    }
}
