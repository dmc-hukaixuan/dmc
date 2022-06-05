package process

import (
	"dmc/global"
	"dmc/kernel/model/admin"
	"fmt"
)

// add new Trnsition=
// returns the id of the created Transition if success or undef otherwise
func TransitionAdd(processTransition []admin.ProcessTransition) (processTransitionEnter admin.ProcessTransition, err error) {
	err = global.GVA_DB.Table("dmc_pm_transition").Create(&processTransition).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTransitionEnter, err
}

// delete transition
func TransitionDelete(transitionID string) (err error) {
	deleteSQL := `DELETE FROM dmc_pm_transition WHERE transition_id = ? `
	// ask database
	err = global.GVA_DB.Exec(deleteSQL, transitionID).Unscoped().Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	// delete transition action
	TransitionActionDelete(transitionID)
	return err
}

// transition data
func TransitionGet(processTransition *admin.Process) (processTransitionEnter admin.Process, err error) {
	err = global.GVA_DB.Table("dmc_pm_transition").Create(&processTransition).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTransitionEnter, err
}

// process trsnsition update
func TransitionUpdate(pt *admin.ProcessTransition) (processTransitionEnter admin.ProcessTransition, err error) {
	err = global.GVA_DB.Raw(`UPDATE dmc_pm_transition SET name = ?, config = ?, change_by = ?, change_time = ? WHERE id = ?`, pt.Name, pt.Config, pt.ChangeBy, pt.ChangeTime, pt.ID).Scan(&processTransitionEnter).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTransitionEnter, err
}

// transition list
func TransitionListbyProceeID(proceeID int) (processTransitionEnter []admin.ProcessTransition, err error) {
	err = global.GVA_DB.Table("dmc_pm_transition").Where("process_id = ?", proceeID).Find(&processTransitionEnter).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTransitionEnter, err
}

func TransitionListbyProceeIDGet(proceeID string) (TransitionList map[string]int, err error) {
	var tl []admin.ProcessTransition
	err = global.GVA_DB.Table("dmc_pm_transition").Where("process_id = ?", proceeID).Find(&tl).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	for _, v := range tl {
		TransitionList[v.TransitionID] = v.ID
	}
	return TransitionList, err
}
