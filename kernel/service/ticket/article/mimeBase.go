package article

import (
    "dmc/global"
    model "dmc/kernel/model/ticket"
    "dmc/kernel/util"
    date_time "dmc/kernel/util/time"
    "strconv"
    "time"

    //"kernel/service/ticket"
    "os"
)

// create a mime article
func ArticleCreate(articleData model.ArticleDataMimeCreate) int64 {
    incomingTime := date_time.SystemTime()
    // for the event handler, before any actions have taken place
    //ldTicketData := ticket.TicketGet()
    // add 'no body' if there is no body there!

    metaArticle := model.MetaArticle{
        TicketID:               articleData.TicketID,
        SenderType:             articleData.SenderTypeID,
        IsVisibleForCustomer:   articleData.IsVisibleForCustomer,
        CommunicationChannelID: ChannelIDGet(articleData.CommunicationChannel),
        UserID:                 articleData.UserID,
    }
    //  Create meta article
    articleID := metaArticleCreate(metaArticle)
    // article create failed
    if articleID == 0 {

    }
    // 写入 article 的富文本数据
    articleRichText := model.ArticleRichText{
        ArticleID:   articleID,
        TicketID:    articleData.TicketID,
        ContentSize: len(articleData.Body),
        Body:        articleData.Body,
        CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
        CreateBy:    articleData.UserID,
        ChangeTime:  time.Now().Format("2006-01-02 15:04:05"),
        ChangeBy:    articleData.UserID,
    }

    // rich text,
    err := global.GVA_DB.Table("article_rich_text").Create(&articleRichText).Error
    if err != nil {
        panic(err)
    }

    articleDataMime := model.ArticleDataMime{
        ArticleID:    articleID,
        TicketID:     articleData.TicketID,
        From:         articleData.From,
        ReplyTo:      articleData.ReplyTo,
        To:           articleData.To,
        Cc:           articleData.Cc,
        Bcc:          articleData.Bcc,
        Subject:      articleData.Subject,
        ContentType:  "text/plain; charset=utf-8",
        ContentSize:  len(articleData.Body),
        IncomingTime: incomingTime,
        Body:         articleData.Body,
        ContentPath:  time.Now().Format("2006/01/02"),
        CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
        CreateBy:     articleData.UserID,
        ChangeTime:   time.Now().Format("2006-01-02 15:04:05"),
        ChangeBy:     articleData.UserID,
    }
    //写入 article 的基础数据，去除富文本的内容
    articleDataMime.Body = articleDataMime.Body
    err1 := global.GVA_DB.Table("article_data_mime").Create(&articleDataMime).Error
    if err != nil {
        panic(err1)
    }
    // add attachments,if config is instore article in fiel system

    return articleID
}

// Create a new article.
func metaArticleCreate(metaArticle model.MetaArticle) int64 {
    metaArticle.InsertFingerprint = strconv.Itoa(os.Getppid()) + "-" + util.GenerateRandomString(32)
    err := global.GVA_DB.Table("article").Create(&metaArticle).Error
    if err != nil {
        panic(err)
    }
    return metaArticle.ID
}

func ChannelIDGet(ChannelName string) int {
    channel := map[string]int{
        "Email":             1,
        "Phone":             2,
        "Internal":          3,
        "Chat":              4,
        "Weixin":            5,
        "IsLand":            6,
        "WeixinMiniprogram": 7,
    }
    return channel[ChannelName]
}

func ArticleWriteAttachment(formID string, article_id int64, userID int, ticketID int64) {

}
