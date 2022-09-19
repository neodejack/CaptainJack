package ports

type AllocationViewerDbPorts interface {
	CheckWhatDatesAreAvail() (string, error)
}
