package cron

import (
    "dmc/global"
    model "dmc/kernel/model/admin"
    "dmc/kernel/service/cronmodule"
    "fmt"

    "github.com/robfig/cron"
)

type CronJob struct {
    TaskType string
    TaskID   int
    TaskData string
}

/*
 cron 自动任务，支持无参函数作为回调，
 如果需要支持有参的参数作为到期的回调， 需要实现 job 的结构
*/
func CronStart() {
    // get corn list
    cronList := CronList(1)
    c := cron.New()
    for _, v := range cronList {
        err := c.AddJob(v.CronTime, CronJob{v.TaskType, v.ID, v.TaskData})
        if err != nil {
            fmt.Println("err", err)
            continue
        }

        // update cron id
        // CronUpdateIdent(v.ID, entityID)
    }

    c.Start()

}

// 实现 cron 中的 run 方法，用于自动任务中传递参数
func (c CronJob) Run() {
    cronmodule.CornRun(c.TaskType).Run(
        c.TaskID,
        c.TaskData,
    )
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

func CronGet(cronID int) model.SchedulerTask {
    var scheduler model.SchedulerTask
    err := global.GVA_DB.Table("dmc_scheduler_task").Scan(&scheduler).Error
    if err != nil {
        panic(err)
    }
    return scheduler
}

func CronDelete() {

}

func CronUpdate() {

}

// update corn excution
func CronUpdateIdent(cornID int, Indet string) {
    err := global.GVA_DB.Raw("UPDATE dmc_scheduler_task SET ident = ? WHERE id = ?", Indet, cornID).Error
    if err != nil {

    }
}
