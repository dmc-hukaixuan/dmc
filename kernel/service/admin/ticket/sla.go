package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
	get sla data
*/
func SLAAdd(sla model.SLA) (sla_id int, err error) {
	err = global.GVA_DB.Table("sla").Create(&sla).Error
	if err != nil {
		return
	}
	return sla.ID, err
}

/*
	get sla data
*/
func SLAGet(slaID int) (sla model.SLA) {
	err := global.GVA_DB.Table("sla").Where("id = ?", slaID).First(&sla).Error
	if err != nil {
		return
	}
	return sla
}

/*
	return a map list of slas
*/
func SLAList(validID int) map[string]string {
	var sla []model.SLA
	selectSQL := `SELECT id, name FROM sla`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	slaList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&sla).Error
	if err != nil {
		return slaList
	}
	// do loop, build a json string
	for _, v := range sla {
		slaList[strconv.Itoa(v.ID)] = v.Name
	}
	return slaList
}

/*
	link SLA and services
*/
func SLAUpdate(sla model.SLA) (sla_id int, err error) {
	err = global.GVA_DB.Table("sla").Where("id = ?", sla.ID).Model(&sla).Omit("create_by", "create_time").Updates(sla).Error
	if err != nil {
		return
	}
	return sla.ID, err
}

/*
	type list getSLA
*/
func SLAListGet(validID int) (sla []model.SLA) {
	selectSQL := `SELECT s.id AS id, s.name AS name, s.valid_id AS valid_id, s.internal_note AS internal_note,
					s.external_note AS external_note,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					s.create_time AS create_time, s.change_time AS change_time
					FROM sla s 
					LEFT JOIN users u ON u.id = s.create_by 
					LEFT JOIN users u1 ON u1.id = s.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&sla).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return sla
}

/*
	@param: sla id int
	@param: servcis list an int array
	@description : link service and sla
	@return: null
*/
func slaLinkServiceAdd(sla_id int, services []int) {
	var sla []model.SLAService
	for _, v := range services {
		sla = append(sla, model.SLAService{
			SLA:     sla_id,
			Service: v,
		})
	}
	err := global.GVA_DB.Table("service_sla").Create(&sla).Error
	if err != nil {
		return
	}
}

/*
	delete SLA link services
*/
func slaLinkServiceDelete(sla_id int) {
	deleteSQL := `DELETE FROM service_sla WHERE sla = ? `
	// ask database
	err := global.GVA_DB.Exec(deleteSQL, sla_id).Unscoped().Error
	if err != nil {
		return
	}
}

/*
	get sla link service
*/
func slaLinkServiceGet(sla_id int) (services []int) {

	var slaService []model.SLAService
	err := global.GVA_DB.Table("service_sla").Where("sla_id = ?", sla_id).First(&slaService).Error
	if err != nil {
		return
	}
	// do loop
	for _, v := range slaService {
		services = append(services, v.Service)
	}
	return services
}
