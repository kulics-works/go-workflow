package workflow

type Workflow interface {
	Run() []error
	Length() int
	Add(...Workflow) Workflow
}
