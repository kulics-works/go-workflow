package workflow

func NewAsyncWorkflow(wk ...Workflow) *asyncWorkflow {
	me := &asyncWorkflow{}
	for _, v := range wk {
		me.Add(v)
	}
	return me
}

type asyncWorkflow struct {
	syncWorkflow
}

func (me *asyncWorkflow) Add(wk ...Workflow) *asyncWorkflow {
	me.works = append(me.works, wk...)
	return me
}

func (me *asyncWorkflow) Run() []error {
	// 使用chan接收错误
	ch := make(chan []error)
	count := me.Length()
	errs := make([]error, 0)
	i := 0
	// 等待全部mission完成
	for err := range me.ProcessAsync(ch) {
		i++
		if len(err) != 0 {
			errs = append(errs, err...)
		}
		if i >= count {
			close(ch)
			break
		}
	}
	return errs
}

func (me *asyncWorkflow) ProcessAsync(ch chan []error) <-chan []error {
	go func() {
		for _, wk := range me.works {
			go func(wk Workflow) {
				if errs := wk.Run(); len(errs) != 0 {
					// 发送错误
					ch <- errs
				} else {
					ch <- nil
				}
			}(wk)
		}
	}()
	return ch
}
