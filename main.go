package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
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
	//QUESTIONS SET
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

	//WINDOW CONFIG
	a := app.New()
	window := a.NewWindow("Quiz")
	window.Resize(fyne.NewSize(400, 200))
	content := container.NewGridWithRows(3)

	for _, q := range questions {

		correctAnsLabel := widget.NewLabel("Correct answer")
		wrongAnsLabel := widget.NewLabel("Bad answer")

		//LABEL
		queContent := widget.NewLabel(q.content)

		//INPUT
		input := widget.NewEntry()
		input.SetPlaceHolder("Here's place for your answer")

		//ANSWER BTN
		answerBtn := widget.Button{Text: "Answer", OnTapped: func() {
			if checkAnswer(q.correctAns, convertTextToInt(input.Text)) {
				content.Add(correctAnsLabel)
			} else {
				content.Add(wrongAnsLabel)
			}
		}}

		content.Add(queContent)
		content.Add(input)
		content.Add(&answerBtn)
	}

	window.SetContent(content)
	window.ShowAndRun()
}

func checkAnswer(correctAns int, givenAns int) bool {
	return givenAns == correctAns
}

func convertTextToInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		return 0
	} else {
		return result
	}
}
