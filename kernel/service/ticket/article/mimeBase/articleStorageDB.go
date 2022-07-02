package mimeBase

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"io/ioutil"
	"time"
)

type DB struct{}

func (*DB) ArticleWriteAttachment(formID string, article_id int64, userID int, ticketID int64) {

	// get file content and file name
	files, err := ioutil.ReadDir("")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		att := model.ArticleDataMimeAttachment{
			TicketID:    ticketID,
			ArticleID:   article_id,
			Filename:    file.Name(),
			ContentType: "",
			ContentSize: file.Size(),
			Content:     "",
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			CreateBy:    userID,
			ChangeTime:  time.Now().Format("2006-01-02 15:04:05"),
			ChangeBy:    userID,
		}
		err := global.GVA_DB.Table("dynamic_field_value").Create(&att).Error
		if err != nil {
			panic(err)
		}
	}

}
