package mail

import (
    config "dmc/config/database"
    model "dmc/kernel/model/admin"
    "fmt"
    "io/ioutil"
    "regexp"

    "gopkg.in/gomail.v2"
)

// send mail
func SendMail(role string, sendconfig model.ReportSendConfig) {
    // get config
    // sendMailList := map[string][]string{
    // 	"POS":    {"Greta.Zhang@majorel.cn", "Jie.li@porsche.cn", "chengkai.pan@mhp.com", "rachel.geng@porsche.cn"},
    // 	"DSS":    {"Greta.Zhang@majorel.cn", "Jie.li@porsche.cn", "rachel.geng@porsche.cn", "szheng@gpstrategies.com", "jane@processsoftwareservice.com"},
    // 	"D-Flow": {"Greta.Zhang@majorel.cn", "rachel.geng@porsche.cn"},
    // 	// "Fadada support team": {"jiapanjpag@163.com", "hukaixuan2015@outlook.com", "hukaixuan26@163.com"},
    // 	// "POS":    {"hukaixuan26@163.com", "jiapanjpag@163.com"},
    // 	// "DSS":    {"hukaixuan26@163.com", "jiapanjpag@163.com"},
    // 	// "D-Flow": {"hukaixuan26@163.com", "jiapanjpag@163.com"},
    // }

    // ccList := map[string][]string{
    // 	"POS":                 {"diki.zhang@porsche.cn", "jay.chen@porsche.cn"},
    // 	"DSS":                 {"jay.chen@porsche.cn"},
    // 	"D-Flow":              {"jay.chen@porsche.cn"},
    // 	"Fadada support team": {"guiguipenny@sina.com"},
    // }
    // 跑出来的报表目录
    files, err := ioutil.ReadDir(`D:\05-附件\02-跑出来的报表`)
    if err != nil {
        panic(err)
    }
    // 循环规则
    //for k, v := range sendMailList {
    m := gomail.NewMessage()
    //发送人
    //m.SetHeader("From")
    m.SetHeaders(map[string][]string{
        "From":    {m.FormatAddress(config.GetConfig().Email.From, "Penny")},
        "To":      sendconfig.To, // v,
        "Cc":      sendconfig.Cc,
        "Subject": {sendconfig.Subject},
    })
    //内容
    m.SetBody("text/html", sendconfig.Body)

    // 发送给相应的人
    for _, file := range files {
        match, _ := regexp.MatchString(role, file.Name())
        if match {
            // 增加附件
            fmt.Println("file.Name()   ", file.Name())
            m.Attach("var/temp/report/" + file.Name())

        }
    }
    //拿到 token，并进行连接,第 4 个参数是填授权码
    d := gomail.NewDialer(config.GetConfig().Email.Host, config.GetConfig().Email.Port, config.GetConfig().Email.From, config.GetConfig().Email.Token)
    fmt.Println("-------1-----", d)
    // 发送邮件
    if err := d.DialAndSend(m); err != nil {
        fmt.Printf("--------DialAndSend err %v:", err)
        panic(err)
    }
}
