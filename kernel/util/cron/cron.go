package cron

import (
    "dmc/global"
    model "dmc/kernel/model/admin"
    "dmc/kernel/service/cronmodule"
    "fmt"

    "github.com/robfig/cron"
)

func CronStart() {

    // get corn list
    cronList := CronList(1)
    c := cron.New()
    for _, v := range cronList {
        entityID, err := c.AddFunc(v.CronTime, cronmodule.CornModule.Run(
            v.ID,
            v.TaskData,
        ))
        if err != nil {
            fmt.Println("err", err)
        }
        // update cron id
        CronUpdateIdent(v.ID, entityID)
    }

    c.Start()
}

func CronStop() {

}

func CronRestart() {

}

// get cron list
func CronList(validID int) []model.SchedulerTask {
    var scheduler []model.SchedulerTask
    err := global.GVA_DB.Table("dmc_scheduler_task").Scan(&scheduler).Error
    if err != nil {
        panic(err)
    }
    return scheduler
}

func CronAdd() {

}

func CronDelete() {

}

func CronUpdate() {

}

// update corn excution
func CronUpdateIdent(cornID int, Indet string) {
    global.GVA_DB.Raw("UPDATE dmc_scheduler_task SET ident = ? WHERE id = ?", Indet, cornID).Scan(&scheduler).Error
}
