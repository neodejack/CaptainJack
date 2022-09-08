package ports

type EchoingPort interface {
	Echoing(words string) (string, error)
}
