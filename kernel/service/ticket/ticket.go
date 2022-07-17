package ticket

import (
    "dmc/global"
    model "dmc/kernel/model/ticket"
    "dmc/kernel/service/admin"
    "dmc/kernel/service/template"
    "dmc/kernel/service/ticket/number"
    dateTime "dmc/kernel/util/time"
    "encoding/json"
    "errors"
    "fmt"
    "sync"
    "time"

    "gorm.io/gorm"
)

/*
   创建工单函数
*/
func TicketCreate(TicketBaseData model.TicketBaseData) (TicketID int64, err error) {
    // get db object
    // db := pool.GetDB()
    numberbuild := number.TicketNumber()
    // check ticket number if not null, generate ticket number
    if TicketBaseData.TicketNumber == "" {
        TicketBaseData.TicketNumber = numberbuild.TicketNumberBuild()
    }
    fmt.Println("TicketBaseData ", TicketBaseData)

    if !errors.Is(global.GVA_DB.Table("sc_ticket").Create(&TicketBaseData).Error, gorm.ErrRecordNotFound) {
        return 0, errors.New("Ticket create failed")
    }
    // update dynamic_field

    // trgger event loop
    // trigger event

    return TicketBaseData.ID, err
}

/*
   获取g单所有动态字段的值
*/
// func TicketDynamicFieldValueGet(ticketid int64, ticketData *map[string]interface{}) *map[string]interface {
//     type DynamicFieldValue struct {
//         Label     string `json:"label"`
//         Name      string `json:"name"`
//         ValueText string `json:"value_text"`
//     }
//     var dfv []DynamicFieldValue
//     //db := pool.GetDB()
//     // sql string TicketHistoryGet ticket_history_2022_02
//     SelctSQL := `SELECT dfv.id, dfv.field_id, df.name as name, df.label as label, max(case when dfv.value_text IS NOT NULL then dfv.value_text ELSE dfv.value_date END) as value_text
//                 FROM dynamic_field_value dfv LEFT JOIN dynamic_field df ON df.id = dfv.field_id WHERE dfv.object_id = ?
//                 GROUP BY dfv.id`
//     // ask database and fetch result
//     global.GVA_DB.Raw(SelctSQL, ticketid).Scan(&dfv)
//     for _, v := range dfv {
//         ticketData[v.Name] = v.ValueText
//     }
//     return &ticketData
// }
var wg sync.WaitGroup

// add a history entry to an ticket
// we should bulk update ticket info
func TicketHistoryAdd(thfd model.TicketHistoryFieldData) {
    var historyLine []model.TicketHistory
    historyTime := dateTime.CurrentTimestamp()

    // get last time ticket history data
    existisHistory := TicketHistoryDistinctGet(thfd.TicketID)
    fieldData := thfd.TicketFiledList

    // 创建有5 个缓冲的通道，数据类型是  *TicketReprot
    ch := make(chan string, 10)
    wg.Add(8)
    for i := 0; i < 8; i++ {
        go func() {
            defer wg.Done()
            // build now ticket history data
            for field := range ch {
                startValue := ""
                leaveValue := fieldData[field]
                workingTime := 0
                calendarTime := 0
                // calcuate working time
                if fieldHistory, ok := existisHistory[field]; ok {
                    startValue = fieldHistory.LeaveValue
                    calendarID := 1
                    // working time are calculated according to the calendar set by the queue
                    if field == "Queue" || field == "QueueID" {
                        calendarID = 1
                    } else {
                        calendarID = 1
                    }
                    history, _ := dateTime.StringToTime(fieldHistory.CreateTime)
                    workingTime = dateTime.WorkingTime(
                        history,
                        time.Now(),
                        calendarID,
                    )
                }
                historyLine = append(historyLine, model.TicketHistory{
                    Object:       field,
                    TicketID:     thfd.TicketID,
                    StartValue:   startValue,
                    LeaveValue:   leaveValue.(string),
                    OwnerID:      thfd.OwnerID,
                    QueueID:      thfd.QueueID,
                    WorkingTime:  workingTime,
                    CalendarTime: calendarTime,
                    ArticleID:    thfd.ArticleID,
                    TemplateID:   thfd.TemplateID,
                    TemplateName: thfd.TemplateName,
                    Source:       thfd.Source,
                    CreateBy:     thfd.UserID,
                    CreateTime:   historyTime,
                })
            }
        }()
    }
    // 创建
    for _, field := range thfd.FieldOrder {
        ch <- field
    }
    close(ch)
    wg.Wait()
    // buld insert ticket history
    err := global.GVA_DB.Table("dmc_ticket_history").Create(&historyLine).Error
    if err != nil {
        panic(err)
    }
}

func TicketHistoryGet(ticketID int64) {

}

// get ticket distinct history
// every field last time value and upate time and so on
func TicketHistoryDistinctGet(ticketID int64) map[string]model.TicketHistory {
    var historyLine []model.TicketHistory
    sql := `SELECT id, object, leave_value, working_time, queue_id, working_time, create_time FROM dmc_history WHERE ticket_id = ? ORDER BY create_time`
    global.GVA_DB.Raw(sql, ticketID).Scan(&historyLine)

    historyValue := map[string]model.TicketHistory{}
    // do loop history line build a string
    for _, v := range historyLine {
        historyValue[v.Object] = v
    }
    return historyValue
}

func TicketHistoryDelete(ticketID int64) {

}

func TicketMerge(MainTicketID int64, MergeTicketID int64) {

}

func TicketMergeDynamicFields(MainTicketID int64, MergeTicketID int64) {

}

/*
   =head2 TicketMergeLinkedObjects()

   merge linked objects from one ticket into another, that is, move
   them from the merge ticket to the main ticket in the link_relation table.

       my $Success = $TicketObject->TicketMergeLinkedObjects(
           MainTicketID  => 123,
           MergeTicketID => 42,
           UserID        => 1,
       );

   =cut
*/
func TicketMergeLinkedObjects(MainTicketID int64, MergeTicketID int64) {

}

// get ticket data
//
func TicketGet(ticketID int64, dynamic_field bool) map[string]interface{} {
    var baseData model.TicketBaseData
    err := global.GVA_DB.Table("dmc_ticket").Where("id = ?", ticketID).Find(&baseData).Error
    if err != nil {
        panic(err)
    }

    ticketData := make(map[string]interface{})
    data, _ := json.Marshal(&baseData)
    json.Unmarshal(data, &ticketData)

    // fetch dynamicfield value
    if dynamic_field {
        // get ticket dynamicfield object list
        dynamicFieldList, _ := admin.DynamicFieldList("Ticket")
        for _, v := range dynamicFieldList {
            dynamicField := template.DynamicField(v.FieldType).ValueGet(v.ID, "Ticket", ticketID)
            ticketData["DynamicField_"+v.Name] = dynamicField
        }
    }

    return ticketData
}

func TicketWatchGet() {

}

/*
   to subscribe a ticket to watch it

   my $Success = $TicketObject->TicketWatchSubscribe(
       TicketID    => 111,
       WatchUserID => 123,
       UserID      => 123,
   );
*/
func TicketWatchSubscribe(TicketID int64, WatchUserID int64) {

}

/*
   to remove a subscription of a ticket

   my $Success = $TicketObject->TicketWatchUnsubscribe(
       TicketID    => 111,
       WatchUserID => 123,
       UserID      => 123,
   );

*/
func TicketWatchUnsubscribe(TicketID int64, WatchUserID int64) {

}

/*
   set ticket flags

   my $Success = $TicketObject->TicketFlagSet(
       TicketID => 123,
       Key      => "Seen",
       Value    => 1,
       UserID   => 123, # apply to this user
   );
*/
func TicketFlagSet() {

}
