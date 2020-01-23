package workflow

type Workflow interface {
	Run() []error
	Name() string
	Length() int
}
