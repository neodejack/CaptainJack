package ports

type APIPorts interface {
	EchoWords(words string) (string, error)
}

type AllocationViewerPorts interface {
	GetDateSelection(dateStr string) (string, error)
	CheckAllocationData(date string) (string, error)
	SendAllocationDate(data string) error
}
