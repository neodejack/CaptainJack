package echoing

type Echo struct {
}

func NewAdapter() *Echo {
	return &Echo{}
}

func (echo Echo) Echoing(words string) (string, error) {
	annoyingWords := words + ", so what?"
	return annoyingWords, nil
}
