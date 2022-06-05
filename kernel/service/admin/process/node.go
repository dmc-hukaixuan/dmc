package process

import (
	"dmc/global"
	"dmc/kernel/model/admin"
	"fmt"
)

func NodeAdd(processNode []admin.ProcessNode) (processNodeEnter admin.ProcessTransition, err error) {
	err = global.GVA_DB.Table("dmc_pm_activity").Create(&processNode).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processNodeEnter, err
}

// get process use node list
func NodeProcessListGet(processID string) (processNodeEnter map[string]int, err error) {
	var processNode []admin.ProcessNode
	err = global.GVA_DB.Raw("SELECT ticket_template_id FROM dmc_pm_activity WHERE process_id = ?", processID).Find(&processNode).Error
	for _, v := range processNode {
		processNodeEnter[v.NodeID] = v.ID
	}
	return processNodeEnter, err
}

// delete node
func NodeDelete(nodeID string) (err error) {
	//nodeListString := strings.Join(*nodeList, ",")
	deleteSQL := `DELETE FROM dmc_pm_activity WHERE node_id = ? `
	// ask database
	err = global.GVA_DB.Exec(deleteSQL, nodeID).Unscoped().Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	// delete node link ticket template
	NodeTemplateDelete(nodeID)
	return err
}

func NodeGet(processTransition *admin.ProcessTransition) (processTransitionEnter admin.ProcessTransition, err error) {
	err = global.GVA_DB.Table("dmc_pm_activity").Create(&processTransition).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTransitionEnter, err
}

// node update
func NodeUpdate(pt *admin.ProcessNode, templatelist []int) (ProcessNodeEnter admin.ProcessNode, err error) {
	updateSQL := `UPDATE dmc_pm_activity SET name = ?, config = ?, change_by = ?, change_time = ? WHERE id = ?`
	err = global.GVA_DB.Raw(updateSQL, pt.Name, pt.Config, pt.ChangeBy, pt.ChangeTime, pt.ID).Scan(&ProcessNodeEnter).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	// delete template list
	NodeTemplateDelete(pt.NodeID)
	var templateAdd []admin.ActivityTemplate
	// link ticket template
	for _, v := range templatelist {
		templateAdd = append(templateAdd, admin.ActivityTemplate{
			ProcessID:  "",
			NodeID:     pt.NodeID,
			TemplateID: v,
		},
		)
	}
	NodeTemplateAdd(templateAdd)
	return ProcessNodeEnter, err
}

// node update
func NodeTemplateAdd(template []admin.ActivityTemplate) (err error) {
	err = global.GVA_DB.Table("dmc_pm_activity_ticket_template").Create(&template).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return err
}

// get activity or node template list
func NodeTemplateGet(NodeID string) (TemplateList []admin.TemplateIDList) {
	global.GVA_DB.Raw("SELECT ticket_template_id FROM dmc_pm_activity_ticket_template WHERE node_id = ?", NodeID).Find(&TemplateList)
	return TemplateList
}

// delet node link template
func NodeTemplateDelete(NodeID string) (err error) {
	deleteSQL := `DELETE FROM dmc_pm_activity_ticket_template WHERE node_id = ? `
	// ask database
	err = global.GVA_DB.Exec(deleteSQL, NodeID).Unscoped().Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return err
}

// get node list by porcessid
func NodeListByProcessID(processID int) (ProcessNodeEnter []admin.ProcessNode, err error) {
	err = global.GVA_DB.Table("dmc_pm_activity").Where("process_id = ?", processID).Find(&ProcessNodeEnter).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return ProcessNodeEnter, err
}
