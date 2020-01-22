package workflow

type Workflow interface {
	Run() error
	// MissionLength() int
	// AddMission(mt ...IMission) IMission
	// ClearMission()
}
