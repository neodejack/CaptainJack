package funstufffortina

import "log"

type Echo struct {
}

func NewAdapter() *Echo {
	return &Echo{}
}

func (echo Echo) Echoing(words string) (string, error) {
	var annoyingWords string
	log.Println("what i get:", words)
	log.Println("what i want: i'm tina")

	if words=="i'm tina" {
		poem := `
			We are both poets
			So I ask you
			to write me into a poem
			and you say: here, this one
			is about shaping you
			into a wave
			or here
			you are a horse
			with lace reins
			and I look
			finding only music
			or what could be
			your mother so
			I ask again where am I
			and you say: who else
			could I mean
			when I write: sweet
			witch, write: teeth you
			say: can't you see it turn
			like you do	
		`
		annoyingWords = "don't be too sad girl. " +
			"you are gonna have some good food and take a solid night of sleep tonight. " +
			"things will get easier. " +
			"here's a cute poem for you if you are still felling sad" + poem

		return annoyingWords, nil

	} else {
		annoyingWords = words + ", so what?"
	}
	return annoyingWords, nil
}
