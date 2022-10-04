package ports

type EchoingPort interface {
	Echoing(words string) (string, error)
}

type AllocationViewerPort interface {
	ViewAllocation(date string) (string, error)
}
