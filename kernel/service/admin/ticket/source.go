package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
get type list

    my %List = $TypeObject->TypeList();

or

    my %List = $TypeObject->TypeList(
        Valid => 0,
    );
*/
func SourceList(validID int) map[string]string {
	var tp []model.TicketType
	selectSQL := `SELECT id, name FROM ticket_source`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	typeList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return typeList
	}
	// do loop, build a json string
	for _, v := range tp {
		typeList[strconv.Itoa(v.ID)] = v.Name
	}
	return typeList
}

/*
	get ticket ticket detail
*/
func SourceGet(sourceID int) (ts model.TicketSource) {
	err := global.GVA_DB.Table("ticket_source").Where("id = ?", sourceID).First(&ts).Error
	if err != nil {
		return
	}
	return ts
}

/*
	add a ticket source
*/
func SourceAdd(ts model.TicketSource) (sourceID int, err error) {
	err = global.GVA_DB.Table("ticket_source").Create(&ts).Error
	if err != nil {
		return
	}
	return ts.ID, err
}

/*
	ticket souce update
*/
func SourceUpdate(ts model.TicketSource) (sourceID int, err error) {
	err = global.GVA_DB.Table("ticket_source").Where("id = ?", ts.ID).Model(&ts).Omit("create_by", "create_time").Updates(ts).Error
	if err != nil {
		return
	}
	return ts.ID, err
}

/*
	type list get
*/
func SourceListGet(validID int) (tt []model.TicketType) {
	selectSQL := `SELECT tt.id, tt.name, tt.valid_id, tt.tnstart AS tnstart,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					tt.create_time AS create_time, tt.change_time AS change_time
					FROM ticket_type tt 
					LEFT JOIN users u ON u.id = tt.create_by 
					LEFT JOIN users u1 ON u1.id = tt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&tt).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return tt
}
