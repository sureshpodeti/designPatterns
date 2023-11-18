package main

import "fmt"

func main() {

	/*
	 1. Pizza Customization
	 2. Text formatting {bold, italic, underline} to text
	*/
	text := &Text{}

	fmt.Println("Plain text: ", text.Make())

	bold := &BoldText{Text: text}

	fmt.Println("Bold text: ", bold.Make())

	italic := &ItalicText{Text: bold}

	fmt.Println("Bold and Italic text: ", italic.Make())

	underline := &UnderLineText{Text: italic}

	fmt.Println("Bold, Italic, Underline text: ", underline.Make())
}
