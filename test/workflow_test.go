package workflow

import (
	"errors"
	"testing"

	"github.com/kulics-works/go-command"
	"github.com/kulics-works/go-workflow"
)

func TestWorkflow(t *testing.T) {
	list := workflow.NewSyncWorkflow()
	type localCommand struct {
		cli   *command.Local
		value int
	}
	ms := workflow.NewWork(localCommand{command.NewLocal(), 2020}, func(i interface{}) error {
		if params, ok := i.(localCommand); ok {
			t.Log(params.cli.GetURL(), params.value)
			return nil
		} else {
			return errors.New("params error")
		}
	})
	ms2 := workflow.NewWork(command.NewLocal(), func(i interface{}) error {
		cmd := i.(*command.Local)
		t.Log(cmd.GetIP())
		return nil
	})
	// add async work
	list.Add(workflow.NewAsyncWorkflow(ms, ms2, ms, workflow.NewSyncWorkflow(ms, ms, ms2)),
		ms2, ms2)
	list.Add(workflow.NewWorks(workflow.Parameters{command.NewLocal(), command.NewLocal(), command.NewLocal()},
		workflow.Functions{
			func(i interface{}) error {
				t.Log("multi work")
				return nil
			},
		}))
	err := list.Run()
	if len(err) != 0 {
		t.Fatal(err)
	}
}
