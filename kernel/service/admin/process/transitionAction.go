package process

import (
	"dmc/global"
	"dmc/kernel/model/admin"
	"fmt"
)

// add new Trnsition=
// returns the id of the created Transition if success or undef otherwise
func TransitionActionAdd(processTransition *admin.ProcessTransitionAction) (err error) {
    err = global.GVA_DB.Table("dmc_pm_transition_action").Create(&processTransition).Error
    if err != nil {
        fmt.Println("err add:", err)
        return
    }
    return err
}

// delete transition
func TransitionActionDelete(transitionID string) (err error) {
    selectSQL := `DELETE FROM dmc_pm_transition_action WHERE transition_action_id = ? `
    // ask database
    err = global.GVA_DB.Exec(selectSQL, transitionID).Unscoped().Error
    if err != nil {
        fmt.Println("err add:", err)
        return
    }
    return err
}

// transition action data
func TransitionActionGet(transitionActionID *admin.ProcessTransitionAction) (processTransitionEnter admin.ProcessTransitionAction, err error) {
    err = global.GVA_DB.Table("dmc_pm_process").First(&processTransitionEnter, transitionActionID).Error
    if err != nil {
        fmt.Println("err add:", err)
        return
    }
    return processTransitionEnter, err
}

// process trsnsition update
func TransitionActionUpdate(pt *admin.ProcessTransitionAction) (processTransitionEnter admin.ProcessTransition, err error) {

    err = global.GVA_DB.Raw(`UPDATE dmc_pm_transition_action SET name = ?, config = ?, change_by = ?, change_time = ? WHERE id = ?`, pt.Name, pt.Config, pt.ChangeBy, pt.ChangeTime, pt.ID).Scan(&processTransitionEnter).Error
    if err != nil {
        fmt.Println("err add:", err)
        return
    }
    return processTransitionEnter, err
}

func TransitionActionList(transitionID string) (pta []admin.ProcessTransitionAction) {
    err := global.GVA_DB.Table("dmc_pm_transition_action").Where("transition_id = ?", transitionID).Find(&pta).Error
    if err != nil {
        fmt.Println("err add:", err)
        return
    }
    return pta
}
