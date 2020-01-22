package workflow

import "errors"

func NewAsyncWorkflow() asyncWorkflow {
	return asyncWorkflow{}
}

type asyncWorkflow struct {
	syncWorkflow
	missions []Workflow // 任务数组，任务之间按异步执行
}

func (me asyncWorkflow) AddMission(mt ...Workflow) asyncWorkflow {
	me.missions = append(me.missions, mt...)
	return me
}

func (me asyncWorkflow) Run() error {
	// 使用chan接收错误
	ch := make(chan error)
	count := me.MissionLength()
	errs := make([]error, 0)
	i := 0
	// 等待全部mission完成
	for err := range me.ProcessAsync(ch) {
		i += 1
		if err != nil {
			errs = append(errs, err)
		}
		if i >= count {
			close(ch)
			break
		}
	}
	if len(errs) != 0 {
		return errors.New(func() string {
			str := ""
			for _, err := range errs {
				str += err.Error()
			}
			return str
		}())
	} else {
		return nil
	}
}

func (me asyncWorkflow) ProcessAsync(ch chan error) <-chan error {
	go func() {
		for _, ms := range me.missions {
			go func() {
				if err := ms.Run(); err != nil {
					// 发送错误
					ch <- err
				} else {
					ch <- nil
				}
			}()
		}
	}()
	return ch
}
