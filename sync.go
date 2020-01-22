package workflow

type syncWorkflow struct {
	missions []Workflow // 里程碑队列，里程碑按同步顺序执行
}

func NewSyncWorkflow() syncWorkflow {
	return syncWorkflow{}
}

func (me syncWorkflow) MissionLength() int {
	return len(me.missions)
}

func (me syncWorkflow) AddMission(mt ...Workflow) syncWorkflow {
	me.missions = append(me.missions, mt...)
	return me
}

func (me *syncWorkflow) ClearMission() {
	me.missions = nil
}

func (me syncWorkflow) Run() error {
	for _, mt := range me.missions {
		if err := mt.Run(); err != nil {
			return err
		}
	}
	return nil
}
