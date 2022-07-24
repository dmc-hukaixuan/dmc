package stats

import (
    "dmc/global"
    model "dmc/kernel/model/report"
    "encoding/json"
    "fmt"
    "strconv"
    "time"
)

type DynamicStats interface {
    GenerateDynamicStatsRun(roles, searchType, startTime, stopTime string)
}

/*
   stats run
*/
func StatsRun(statsId int) {
    statsData := StatsGet(statsId)
    // StatsRestrictions := map[string]interface{}{}
    StatsRestrictions := model.StatsDynamicConfig{}
    json.Unmarshal([]byte(statsData.Config), &StatsRestrictions)
    fmt.Println("statsData.Config --------------- ,", statsData.Config)
    fmt.Println("timeFilter --------------- ,", StatsRestrictions.UseAsRestriction)
    searchType := StatsRestrictions.UseAsRestriction["searchType"]
    roles := StatsRestrictions.UseAsRestriction["role_name"]
    startTime := ""
    stopTime := ""
    restriction := StatsRestrictions.UseAsRestriction

    year, month, _ := time.Now().Date()
    thistime := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

    if timeFilter, ok := restriction["create_time"]; ok {
        // fmt.Println("time ", timeFilter)
        filter := timeFilter.(map[string]interface{})
        timeType, _ := filter["type"].(string)

        // timeFilter := map[string]string{}
        if timeType == "last_month" {
            startTime = thistime.AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
            stopTime = thistime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
        } else if timeType == "month_scale" {
            beginDay, _ := filter["TimeStart"].(string)
            stopDay, _ := filter["TimeStop"].(string)
            monthstring := fmt.Sprintf("%02d", int(month))
            startTime = strconv.Itoa(year) + "-" + monthstring + "-" + beginDay
            stopTime = strconv.Itoa(year) + "-" + monthstring + "-" + stopDay
        }
        // json.Unmarshal([]byte(restriction["create_time"].(string)), &timeFilter)
        fmt.Println("timeFilter --------------- ,", filter, startTime, stopTime, int(month))
    }
    fmt.Println("wwwwww --------------- ,", roles.(string), searchType, startTime, stopTime)

    Run(statsData.ScaleType).GenerateDynamicStatsRun(
        roles.(string),
        searchType.(string),
        "2022-03-01",
        "2022-03-15",
    )
}

// get data run a dynaymic stats
func Run(scaleType string) DynamicStats {
    switch scaleType {
    case "ticket_list":
        return &TicketList{}
    case "ticket_count":
        return &TicketList{}
    default:
        return &TicketList{}
    }
}

// get stats data
func StatsGet(statsId int) model.Stats {
    var stats model.Stats
    err := global.GVA_DB.Table("dmc_stats").Where("id = ?", statsId).First(&stats).Error
    if err != nil {

    }
    return stats
}

// stats restriction params get
func StatsRestrictionsGet(config string) {

}
