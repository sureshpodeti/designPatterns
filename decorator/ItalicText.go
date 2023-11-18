package main

type ItalicText struct {
	Text IText
}

func (i *ItalicText) Make() string {
	return i.Text.Make() + ", " + "𝙩𝙚𝙭𝙩"
}
