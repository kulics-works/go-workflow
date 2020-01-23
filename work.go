package workflow

type Parameters = []Parameter
type Parameter = interface{}
type Functions = []Function
type Function = func(i interface{}) error

type work struct {
	params Parameter
	exec   Function // 任务函数
}

func NewWork(params Parameter, exec Function) *work {
	return &work{
		params: params,
		exec:   exec,
	}
}

func (me *work) Run() []error {
	err := me.exec(me.params)
	if err != nil {
		return []error{err}
	}
	return nil
}

func (me *work) Length() int {
	return 1
}

// 最小单位，不支持继续添加
func (me *work) Add(wk ...Workflow) Workflow {
	return me
}

func NewWorks(params Parameters, execs Functions) *asyncWorkflow {
	asyncworks := NewAsyncWorkflow()
	for _, param := range params {
		syncworks := NewSyncWorkflow()
		for _, exec := range execs {
			syncworks.Add(NewWork(param, exec))
		}
		asyncworks.Add(syncworks)
	}
	return asyncworks
}
