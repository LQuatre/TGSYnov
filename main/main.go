package main

import (
	"github.com/rivo/tview"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	YnovBox := tview.NewBox()
	YnovBox.SetBorder(true).SetTitle("TGS 2023")

	asciiLetter := ReadFile("assets/Ascii-letters/standard.txt")
	alphabetAsciiLetters := asciiLetter[298:(298 + (26 * 9))]

	ScrollingMessage := "Ynov est présent à la TGS 2023 !"
	asciiScrollingMessage := []string{"", "", "", "", "", "", "", "", ""}

	for _, letter := range ScrollingMessage {
		if letter == ' ' {
			for i := 0; i < 9; i++ {
				asciiScrollingMessage[i] += "         "
			}
		} else {
			letterIndex := int(letter - 'A')
			letterStart := letterIndex * 9
			letterEnd := letterStart + 9
			for i := 0; i < 9; i++ {
				asciiScrollingMessage[i] +=
			}
		}
	}

	ScrollingText := tview.NewTextView()
	ScrollingText.SetDynamicColors(true)
	ScrollingText.SetRegions(true)
	ScrollingText.SetWrap(false)
	ScrollingText.SetBorder(true).SetTitle("Scrolling Text")
	//ScrollingText.SetText(asciiScrollingMessage[0] + "\n" + asciiScrollingMessage[1] + "\n" + asciiScrollingMessage[2] + "\n" + asciiScrollingMessage[3] + "\n" + asciiScrollingMessage[4] + "\n" + asciiScrollingMessage[5] + "\n" + asciiScrollingMessage[6] + "\n" + asciiScrollingMessage[7] + "\n" + asciiScrollingMessage[8])
	ScrollingText.SetText(ScrollingMessage)
	if err := tview.NewApplication().SetRoot(YnovBox, true).Run(); err != nil {
		panic(err)
	}
}
