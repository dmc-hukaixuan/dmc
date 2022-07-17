package process

import (
    "dmc/global"
    "dmc/kernel/model/admin"
    "dmc/kernel/service/admin/process/transitionvalidation"
    "encoding/json"
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

/* TransitionCheck()

Checks if one or more Transition Conditions are true.

    my $TransitionCheck = $TransitionObject->TransitionCheck(
        TransitionEntityID => 'T1',
        or
        TransitionEntityID => ['T1', 'T2', 'T3'],
        Data       => {
            Queue         => 'Raw',
            DynamicField1 => 'Value',
            Subject       => 'Testsubject',
            ...
        },
    );

Returns:
    If called on a single TransitionEntityID

    $Checked = 1;       # 0

Returns:
    If called on an array of TransitionEntityIDs

    $Checked = 'T1'     # 0 if no Transition was true

*/

func TransitionCheck(ticketData map[string]interface{}, nodeID string) (Checked bool, transitionID string) {
    transitionNode := NodeTransitionGet(ticketData["ActivityID"].(string))
TRANSITIONENTITYID:
    for _, v := range transitionNode {
        // we have TransitionConditions
        ConditionSuccess := 0
        ConditionFail := 0
        condition := admin.TransitionCondition{}
        json.Unmarshal([]byte(v.Config), &condition)
        ConditionLinking := condition.ConditionLinking
        conditions := condition.Conditions
    CONDITIONNAME:
        for _, condition := range conditions {
            fields := condition.Fields
            CondType := condition.Condition
            FieldSuccess := 0
            FieldFail := 0
        FIELDLNAME:
            for _, field := range fields {
                fieldType := field.FieldType
                fieldName := field.FieldName
                fieldValue := field.FieldValue
                compare := field.Compare
                ticketFieldValue, _ := ticketData[fieldName]

                match := transitionvalidation.Validation(compare).Validate(fieldType, fieldValue, ticketFieldValue)

                if match {
                    // Successful check if we just need one matching Condition to make this
                    // Transition valid.
                    FieldSuccess++
                    if ConditionLinking == "or" && CondType == "or" {
                        // print debug info

                        // return
                        return true, v.NodeID
                    }
                    if ConditionLinking != "or" && CondType == "or" {
                        continue CONDITIONNAME
                    }
                } else {
                    FieldFail++
                    // print debug info

                    // Failed check if we have all 'and' conditions.
                    if ConditionLinking == "and" && CondType == "and" {
                        continue TRANSITIONENTITYID
                    }
                    // Try next Condition if all Condition Fields have to be true.
                    if CondType == "and" {
                        continue CONDITIONNAME
                    }
                }
                continue FIELDLNAME
            }
            if CondType == "and" {
                // if we had no failing check this condition matched
                if FieldFail == 0 {
                    // Successful check if we just need one matching Condition to make this Transition valid.
                    if ConditionLinking == "or" {
                        // $Self->DebugLog(
                        //     MessageType      => 'Success',
                        //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                        //     ConditionName    => $ConditionName,
                        //     ConditionType    => $CondType,
                        //     ConditionLinking => $ConditionLinking,
                        // );
                        return true, v.NodeID
                    }
                    ConditionSuccess++
                } else {
                    ConditionFail++
                    if ConditionLinking == "and" {
                        continue TRANSITIONENTITYID
                    }
                }
            } else if CondType == "or" {
                // If we had at least one successful check, this condition matched.
                if FieldSuccess > 0 {

                    // Successful check if we just need one matching Condition to make this Transition valid.
                    if ConditionLinking == "or" {

                        // $Self->DebugLog(
                        //     MessageType      => 'Success',
                        //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                        //     ConditionName    => $ConditionName,
                        //     ConditionType    => $CondType,
                        //     ConditionLinking => $ConditionLinking,
                        // );
                        return true, v.NodeID
                    }
                    ConditionSuccess++
                } else {
                    ConditionFail++
                    // Failed check if we have all 'and' conditions.
                    if ConditionLinking == "and" {
                        continue TRANSITIONENTITYID
                    }
                }
            } else if CondType == "xor" {
                // if we had exactly one successful check, this condition matched.
                if FieldSuccess == 1 {

                    // Successful check if we just need one matching Condition to make this Transition valid.
                    if ConditionLinking == "or" {

                        // $Self->DebugLog(
                        //     MessageType      => 'Success',
                        //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                        //     ConditionName    => $ConditionName,
                        //     ConditionType    => $CondType,
                        //     ConditionLinking => $ConditionLinking,
                        // );

                        return true, v.NodeID
                    }
                    ConditionSuccess++
                } else {
                    ConditionFail++
                }
            }
        }
        if ConditionLinking == "and" {
            // If we had no failing conditions this transition matched.
            if ConditionFail == 0 {

                // $Self->DebugLog(
                //     MessageType      => 'Success',
                //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                //     ConditionLinking => $ConditionLinking,
                // );
                return true, v.NodeID
            }
        } else if ConditionLinking == "or" {
            // If we had at least one successful condition, this transition matched.
            if ConditionSuccess > 0 {
                // $Self->DebugLog(
                //     MessageType      => 'Success',
                //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                //     ConditionLinking => $ConditionLinking,
                // );
                return true, v.NodeID
            }
        } else if ConditionLinking == "xor" {
            // If we had exactly one successful condition, this transition matched.
            if ConditionSuccess == 1 {
                // $Self->DebugLog(
                //     MessageType      => 'Success',
                //     TransitionName   => $Transitions->{$TransitionEntityID}->{Name},
                //     ConditionLinking => $ConditionLinking,
                // );
                return true, v.NodeID
            }
        }
    }
    return false, ""
}
