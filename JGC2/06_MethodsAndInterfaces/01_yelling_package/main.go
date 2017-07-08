package yelling

import "strings"

type LoudString struct {
	s string
}

func NewLoudStirng() LoudString {
	return LoudString{}
}

/*Pass by value or Copy Receiver*/
func (ls LoudString) ValueChange(s string) {
	/*This method receives a copy of the loud string passed*/
	ls.s = strings.ToUpper(s)
}

/*Pass by reference or Pointer Receiver*/
func (ls *LoudString) PointerChange(s string) {
	/*This method receives a pointer to the loud string passed*/
	ls.s = strings.ToUpper(s)
}

func (ls *LoudString) String() string {
	return ls.s
}

func (ls *LoudString) Blank() {
	ls.s = ""
}
