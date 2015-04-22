package tbox

type Char struct {
	Char        rune
	Attr        Attribute
	Front, Back Color
}

type Attribute uint64

type Color struct {
	R, G, B uint8
}

type Line []Char
