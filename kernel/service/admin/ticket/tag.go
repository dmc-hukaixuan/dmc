package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
 add a ticket state
*/
func TagAdd(tag model.Tag) (stateID int, err error) {
	err = global.GVA_DB.Table("ticket_state").Create(&tag).Error
	if err != nil {
		return
	}
	return tag.ID, err
}

/*
	ticket state detial info
*/
func TagGet(tag_id int) (tag model.Tag) {
	err := global.GVA_DB.Table("ticket_state").Where("id = ?", tag_id).First(&tag).Error
	if err != nil {
		return
	}
	return tag
}

/*
	update ticket state
*/
func TagUpdate(tag model.Tag) (stateID int, err error) {
	err = global.GVA_DB.Table("ticket_state").Where("id = ?", tag.ID).Model(&tag).Omit("create_by", "create_time").Updates(tag).Error
	if err != nil {
		return
	}
	return tag.ID, err
}

/*
* get tag list as a hash of ID, Name pairs

useage

returns

	 %List = (
        1 => "new",
        2 => "closed successful",
        3 => "closed unsuccessful",
        4 => "open",
        5 => "removed",
        6 => "pending reminder",
        7 => "pending auto close+",
        8 => "pending auto close-",
        9 => "merged",
    );

*/
func TagList(validID int) map[string]string {
	var tp []model.TicketState
	selectSQL := `SELECT id, name FROM tag`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	tagList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return tagList
	}
	// do loop, build a json string
	for _, v := range tp {
		tagList[strconv.Itoa(v.ID)] = v.Name
	}
	return tagList
}
