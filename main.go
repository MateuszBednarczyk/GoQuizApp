package main

import (
	"fmt"
)

type Question struct {
	content    string
	answers    map[int]string
	correctAns int
}

type Asker interface {
	ask()
}

func main() {
	questions := [...]Question{
		{content: "How many fingers does human have?", answers: map[int]string{
			1: "two",
			2: "five",
			3: "twenty",
		}, correctAns: 3},
		{content: "Is go fun?", answers: map[int]string{
			1: "no",
			2: "yes",
		}, correctAns: 2},
	}

	for _, q := range questions {
		fmt.Println(q.content)
		for _, a := range q.answers {
			fmt.Println(a)
		}

		var userAns int
		_, err := fmt.Scan(&userAns)
		if err != nil {
			fmt.Println("Something went wrong")
		} else {
			if userAns == q.correctAns {
				fmt.Println("Nice")
			} else {
				fmt.Println("Nah")
			}
		}
	}
}
