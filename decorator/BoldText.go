package main

type BoldText struct {
	Text IText
}

func (b *BoldText) Make() string {
	return b.Text.Make() + ", " + "𝘁𝗲𝘅𝘁"
}
