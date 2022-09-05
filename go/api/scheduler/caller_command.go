package scheduler

// Command  used by the call function
type CallCommand interface {
	Invoke() error
}

type WeeklySchedulerCommand struct {
	Data *SchedulerData
}
