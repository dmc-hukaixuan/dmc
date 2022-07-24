package cronmodule

import (
    model "dmc/kernel/model/admin"
    "dmc/kernel/service/stats"

    // "dmc/kernel/util/mail"
    "encoding/json"
)

type Report struct{}

func (*Report) Run(taskID int, taskData string) {
    // get report data
    stats.StatsRun(1)
    //  格式化搜索的日期
    var sendconfig model.ReportSendConfig
    json.Unmarshal([]byte(taskData), &sendconfig)
    // 解析发件的数据

    // 发送邮件
    //  mail.SendMail()
}
