package workflow

import (
	"errors"
	"testing"

	"fabric_deploy/remote_helper"
)

func TestWorkflow(t *testing.T) {
	list := NewSyncWorkflow()
	ms := NewMission().AddClient(remote_helper.NewLocal()).AddClient(remote_helper.NewLocal())
	t.Log(ms.ClientLength())
	ms = ms.AddClient(remote_helper.NewLocal())
	t.Log(ms.ClientLength())
	ms = ms.AddMission(2020, func(cmd remote_helper.Command, i parameter) error {
		if params, ok := i.(int); ok {
			t.Log(cmd.GetURL(), params)
			return nil
		} else {
			return errors.New("params error")
		}
	}).AddMission(nil, func(cmd remote_helper.Command, i parameter) error {
		t.Log(cmd.GetIP())
		return nil
	})
	err := list.AddMission(NewAsyncWorkflow().AddMission(ms).AddMission(ms)).Run()
	if err != nil {
		t.Fatal(err)
	}
}
