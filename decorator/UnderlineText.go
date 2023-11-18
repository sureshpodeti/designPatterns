package main

type UnderLineText struct {
	Text IText
}

func (u *UnderLineText) Make() string {
	return u.Text.Make() + ", " + "t̲h̲i̲s̲"

}
