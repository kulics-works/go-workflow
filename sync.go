package workflow

type syncWorkflow struct {
	works []Workflow // 里程碑队列，里程碑按同步顺序执行
	name  string
}

func NewSyncWorkflow(wk ...Workflow) *syncWorkflow {
	me := &syncWorkflow{}
	for _, v := range wk {
		me.Add(v)
	}
	return me
}

func (me *syncWorkflow) Name() string {
	return me.name
}

func (me *syncWorkflow) Length() int {
	return len(me.works)
}

func (me *syncWorkflow) Add(wk ...Workflow) *syncWorkflow {
	me.works = append(me.works, wk...)
	return me
}

func (me *syncWorkflow) Clear() {
	me.works = nil
}

func (me *syncWorkflow) Run() []error {
	for _, mt := range me.works {
		if errs := mt.Run(); len(errs) != 0 {
			return errs
		}
	}
	return nil
}
