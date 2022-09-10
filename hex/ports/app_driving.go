package ports

type APIPorts interface {
	EchoWords(words string) (string, error)
}
