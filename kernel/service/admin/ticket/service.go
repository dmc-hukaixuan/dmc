/*
All service functions.
*/
package ticket

import (
	"dmc/global"
	"fmt"
	"strconv"

	model "dmc/kernel/model/ticket"
)

/*
	service list for drop down, or other
*/
func ServiceList(validID int) map[string]string {
	var sla []model.Service
	selectSQL := `SELECT id, name FROM service`
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

// service overview list
func ServiceListGet(service_id int) (ts []model.Service) {

	selectSQL := `SELECT s.id AS id, s.name AS name, s.valid_id AS valid_id, s.internal_note AS internal_note,
					s.external_note AS external_note,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					s.create_time AS create_time, s.change_time AS change_time
					FROM service s 
					LEFT JOIN users u ON u.id = s.create_by 
					LEFT JOIN users u1 ON u1.id = s.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&ts).Error
	if err != nil {
		panic(err)
	}
	return ts
}

/*
	service get
*/
func ServiceGet(service_id int) (ts model.Service) {
	err := global.GVA_DB.Table("service").Where("id = ?", service_id).First(&ts).Error
	if err != nil {
		return
	}
	slaList := serviceLinkSLAGet(service_id)
	ts.SLAList = slaList
	return ts
}

/*
	add service
*/
func ServiceAdd(service model.Service) (service_id int, err error) {
	err = global.GVA_DB.Table("sla").Create(&service).Error
	if err != nil {
		return
	}
	// link sla to service
	serviceLinkSLAAdd(service.ID, service.SLAList)
	return service.ID, err
}

func ServiceUpdate(service model.Service) (sla_id int, err error) {
	err = global.GVA_DB.Table("service").Where("id = ?", service.ID).Model(&service).Omit("create_by", "create_time").Updates(service).Error
	if err != nil {
		return
	}
	// delete service already link sla
	serviceLinkSLADelete(service.ID)
	// link sla to service
	serviceLinkSLAAdd(service.ID, service.SLAList)
	return service.ID, err
}

/*
	@param: sla id int
	@param: servcis list an int array
	@description : link service and sla
	@return: null
*/
func serviceLinkSLAAdd(service_id int, slas []int) {
	var sla []model.SLAService
	for _, v := range slas {
		sla = append(sla, model.SLAService{
			SLA:     v,
			Service: service_id,
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
func serviceLinkSLADelete(sla_id int) {
	deleteSQL := `DELETE FROM service_sla WHERE service_id = ? `
	// ask database
	err := global.GVA_DB.Exec(deleteSQL, sla_id).Unscoped().Error
	if err != nil {
		return
	}
}

/*
	get sla link service
*/
func serviceLinkSLAGet(service_id int) (services []int) {

	var slaService []model.SLAService
	err := global.GVA_DB.Table("service_sla").Where("service_id = ?", service_id).First(&slaService).Error
	if err != nil {
		return
	}
	// do loop
	for _, v := range slaService {
		services = append(services, v.SLA)
	}
	return services
}

// // service link tag
// func serviceLinkTagGet(service_id int) (services []int) {
// 	return []int
// }
