package process

import (
    model "dmc/kernel/model/admin"
    "dmc/kernel/service/admin/process"
    "dmc/kernel/service/admin/process/transitionaction"
    "dmc/kernel/service/ticket"
    "encoding/json"
)

//Check valid Transitions and Change Ticket's Activity
// if a Transition was positively checked
func ProcessTransition(UserID int, ticketID int64) (string, bool) {
    // ticket get
    // Get Ticket Data
    ticketData := ticket.TicketGet(ticketID, true)

    // %Transitions Hash for easier reading
    //contains all possible Transitions for the current Activity
    // Handle all possible TransitionEntityID's for the Process->Path's->ActivityEntityID down to
    // Transition.pm's TransitionCheck for validation
    // will return undef if nothing matched or the first matching TransitionEntityID
    ok, TransitionEntityID := process.TransitionCheck(ticketData, ticketData["ActivityID"].(string))
    // if we didn't get a TransitionEntityID
    // no check was successful -> return nothing
    if ok == false {
        // todo print debug info
        // if ( $Self->{Debug} > 0 ) {
        //     $Kernel::OM->Get('Kernel::System::Log')->Log(
        //         Priority => 'debug',
        //         Message  => 'No Transition matched for TicketID: '
        //             . "$Param{TicketID} ProcessEntityID: $Param{ProcessEntityID} "
        //             . "ActivityEntityID: $Param{ActivityEntityID}!",
        //     );
        // }
        return "", false
    }
    CheckOnly := 0
    // If we have a Transition without valid FutureActivitySet we have to complain
    if CheckOnly > 0 {
        // if ( $Self->{Debug} > 0 ) {
        //     $Kernel::OM->Get('Kernel::System::Log')->Log(
        //         Priority => 'debug',
        //         Message  => "Transition with ID $TransitionEntityID matched for "
        //             . "TicketID: $Param{TicketID} ProcessEntityID: $Param{ProcessEntityID} "
        //             . "ActivityEntityID: $Param{ActivityEntityID}!",
        //     );
        // }
        return TransitionEntityID, true
    }

    // If we should just check what Transition matched
    // return a hash containing

    //  Set the new ActivityEntityID on the Ticket
    Success := ProcessTicketActivitySet(ticketID, TransitionEntityID, UserID)
    if Success == false {
        // print log
        return "", false
    }

    // get this transition action
    TransitionActions := process.TransitionActionList(TransitionEntityID)
    for _, transitionAction := range TransitionActions {
        tad := model.TransitionActionData{}
        json.Unmarshal([]byte(transitionAction.Config), &tad)
        transitionaction.TransitionAction(tad.Module).Run(ticketID, tad.Configs)
    }
    // if we don't have Transition Actions on that transition,
    // if we have Transition Action and it isn't an array return,
    return "", false
}

//
func ProceeFirstNodeList() {

}

//    Set Ticket's ActivityEntityID
//     my $Success = $ProcessObject->ProcessTicketActivitySet(
//         ProcessEntityID  => 'P1',
//         ActivityEntityID => 'A1',
//         TicketID         => 123,
//         UserID           => 123,
//     );
//     Returns:
//     $Success = 1; # undef
//     1 if setting the Activity was executed
//     undef if setting failed
func ProcessTicketActivitySet(ticketID int64, ActivityEntityID string, userID int) bool {

    // If Ticket Update to the new ActivityEntityID was successful return 1
    return true
}
