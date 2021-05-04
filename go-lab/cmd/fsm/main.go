package main

import (
	"fmt"
	"github.com/looplab/fsm"
	"time"
)

type Flow struct {
	To  string
	FSM *fsm.FSM
}

func NewFlow(to string) *Flow {
	d := &Flow{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"submitted",
		fsm.Events{
			{Name: "approvalLevel1", Src: []string{"submitted"}, Dst: "done"},
			{Name: "rejectLevel1", Src: []string{"submitted"}, Dst: "inititator"},
			{Name: "approvalLevel2Controller", Src: []string{"done", "submitted"}, Dst: "done"},
			{Name: "rejectLevel2Controller", Src: []string{"done", "submitted"}, Dst: "initiator"},
			{Name: "approvalLevel2Leader", Src: []string{"done", "submitted"}, Dst: "done"},
			{Name: "rejectLevel2Leader", Src: []string{"done", "submitted"}, Dst: "initiator"},
			{Name: "resubmit", Src: []string{"initiator"}, Dst: "submitted"},
		},
		fsm.Callbacks{
			"enter_state":              func(e *fsm.Event) { d.enterState(e) },
			"after_rejectLevel2Leader": func(e *fsm.Event) { d.rejection(e) },
			"after_resubmit":           func(e *fsm.Event) { d.resubmit(e) },
		},
	)

	return d
}

func (d *Flow) enterState(e *fsm.Event) {
	fmt.Printf("flow %s project is %s\n", d.To, e.Dst)
}

func (d *Flow) rejection(e *fsm.Event) {
	fmt.Printf("flow rejected %s\n", e.Event)
}

func (d *Flow) resubmit(e *fsm.Event) {
	fmt.Printf("flow resubmit %v %v\n", e.Args[0], e.Args[1])
}

func main() {
	flow := NewFlow("transformer")

	err := flow.FSM.Event("approvalLevel2Controller")
	if err != nil {
		fmt.Println(err)
	}

	err = flow.FSM.Event("rejectLevel2Leader")
	if err != nil {
		fmt.Println(err)
	}
	args1 := make([]interface{}, 2)
	args1[0] = "sqlTx"
	args1[1] = 1
	args2 := make([]interface{}, 1)
	args2[0] = time.Now()
	err = flow.FSM.Event("resubmit", args1, args2)
	if err != nil {
		fmt.Println(err)
	}

	err = flow.FSM.Event("approvalLevel2Leader")
	if err != nil {
		fmt.Println(err)
	}
}
