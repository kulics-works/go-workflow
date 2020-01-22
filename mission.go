package workflow

import (
	"fabric_deploy/remote_helper"
	"sync"

	"github.com/sirupsen/logrus"
)

type parameter = interface{}

func NewMission(params ...remote_helper.Command) mission {
	return mission{
		clients: params,
	}
}

type mission struct {
	syncWorkflow
	name    string
	clients []remote_helper.Command // 任务链接
}

func (me mission) Name() string {
	return me.name
}

func (me mission) ClientLength() int {
	return len(me.clients)
}

func (me mission) AddClient(it remote_helper.Command) mission {
	me.clients = append(me.clients, it)
	return me
}

func (me *mission) ClearClient() {
	me.clients = []remote_helper.Command{}
}

func (me mission) Lambda(fn func(it *mission)) mission {
	fn(&me)
	return me
}

func (me mission) AddMission(params parameter, exec func(remote_helper.Command, parameter) error) mission {
	me.missions = append(me.missions, submission{
		Params: params,
		Exec:   exec,
	})
	return me
}

func (me mission) Run() error {
	var wg sync.WaitGroup
	var err error
	wg.Add(me.ClientLength())

	for _, cli := range me.clients {
		go func(cmd remote_helper.Command) {
			for _, ms := range me.missions {
				logrus.Info("正在 ", cmd.GetIP(), " 执行")
				if e := ms.Exec(cmd, ms.Params); e != nil {
					err = e
				}
			}
			wg.Done()
		}(cli)
		if err != nil {
			return err
		}
	}
	wg.Wait()
	return nil
}

type submission struct {
	Params parameter
	Exec   func(remote_helper.Command, parameter) error // 任务函数
}

func (me submission) Run() error {
	return nil
}
